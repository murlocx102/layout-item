package etcdkit

import (
	"context"
	"time"

	"go.etcd.io/etcd/api/v3/mvccpb"
	clientv3 "go.etcd.io/etcd/client/v3"
	"go.uber.org/zap"
	"google.golang.org/grpc/resolver"
)

// 客户端发现服务
type Discovery struct {
	schema       string
	EtcdAddrs    []string
	DialTimeout  int
	closeCh      chan struct{}
	watchCh      clientv3.WatchChan
	cli          *clientv3.Client
	keyPrefix    string
	srvAddrsList []resolver.Address
	cc           resolver.ClientConn
	logger       *zap.Logger
}

// NewServiceDiscovery  新建发现服务
func NewServiceDiscovery(etcdAddrs []string, logger *zap.Logger) *Discovery {
	return &Discovery{
		schema:      "etcd",
		EtcdAddrs:   etcdAddrs,
		DialTimeout: 3,
		logger:      logger,
	}
}

// 实现 resolver.Builder Scheme接口
func (r *Discovery) Scheme() string {
	return r.schema
}

// 实现 resolver.Builder Build接口
func (r *Discovery) Build(target resolver.Target, cc resolver.ClientConn, opts resolver.BuildOptions) (resolver.Resolver, error) {
	r.cc = cc

	r.keyPrefix = BuildPrefix(Server{Name: target.Endpoint, Version: target.Authority})
	if _, err := r.start(); err != nil {
		return nil, err
	}
	return r, nil
}

// 实现 resolver.Resolver ResolveNow接口
func (r *Discovery) ResolveNow(o resolver.ResolveNowOptions) {}

// 实现 resolver.Resolver Close接口
func (r *Discovery) Close() {
	r.closeCh <- struct{}{}
}

// start 开启
func (r *Discovery) start() (chan<- struct{}, error) {
	var err error
	r.cli, err = clientv3.New(clientv3.Config{
		Endpoints:   r.EtcdAddrs,
		DialTimeout: time.Duration(r.DialTimeout) * time.Second,
	})
	if err != nil {
		return nil, err
	}
	resolver.Register(r)

	r.closeCh = make(chan struct{})

	if err = r.sync(); err != nil {
		return nil, err
	}

	go r.watch()

	return r.closeCh, nil
}

// watch 监视和更新服务
func (r *Discovery) watch() {
	ticker := time.NewTicker(time.Minute)
	r.watchCh = r.cli.Watch(context.Background(), r.keyPrefix, clientv3.WithPrefix())

	for {
		select {
		case <-r.closeCh:
			return
		case res, ok := <-r.watchCh:
			if ok {
				r.update(res.Events)
			}
		case <-ticker.C:
			if err := r.sync(); err != nil {
				r.logger.Error("同步失败", zap.Error(err))
			}
		}
	}
}

// update 更新服务
func (r *Discovery) update(events []*clientv3.Event) {
	for _, ev := range events {
		var info Server
		var err error

		switch ev.Type {
		case mvccpb.PUT:
			info, err = ParseValue(ev.Kv.Value)
			if err != nil {
				continue
			}
			addr := resolver.Address{Addr: info.Addr, Metadata: info.Metadata}
			if !Exist(r.srvAddrsList, addr) {
				r.srvAddrsList = append(r.srvAddrsList, addr)
				r.cc.UpdateState(resolver.State{Addresses: r.srvAddrsList})
			}
		case mvccpb.DELETE:
			info, err = SplitPath(string(ev.Kv.Key))
			if err != nil {
				continue
			}
			addr := resolver.Address{Addr: info.Addr}
			if s, ok := Remove(r.srvAddrsList, addr); ok {
				r.srvAddrsList = s
				r.cc.UpdateState(resolver.State{Addresses: r.srvAddrsList})
			}
		}
	}
}

// sync 同步获取所有地址信息
func (r *Discovery) sync() error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	res, err := r.cli.Get(ctx, r.keyPrefix, clientv3.WithPrefix())
	if err != nil {
		return err
	}
	r.srvAddrsList = []resolver.Address{}

	for _, v := range res.Kvs {
		info, err := ParseValue(v.Value)
		if err != nil {
			continue
		}
		addr := resolver.Address{Addr: info.Addr, Metadata: info.Metadata}
		r.srvAddrsList = append(r.srvAddrsList, addr)
	}
	r.cc.UpdateState(resolver.State{Addresses: r.srvAddrsList})
	return nil
}
