package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/alecthomas/assert"
	"github.com/gin-gonic/gin/binding"
	"google.golang.org/protobuf/proto"
)

var (
	addr = "http://127.0.0.1:10099"
)

func RequestResult(t *testing.T, url string, reqData proto.Message) []byte {
	body, _ := proto.Marshal(reqData)

	req, _ := http.NewRequest(http.MethodPost, url, bytes.NewReader(body))
	req.Header.Set("Content-Type", binding.MIMEPROTOBUF)

	resp, err := http.DefaultClient.Do(req)

	//打印响应所有信息
	fmt.Printf("%+#v", resp)

	assert.NoError(t, err, "请求错误")
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	res, _ := ioutil.ReadAll(resp.Body)

	t.Log("请求成功")
	return res
}
