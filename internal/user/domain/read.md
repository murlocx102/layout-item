领域划分与边界

加入是领域模型,形成高内聚的领域组件

领域驱动开发中不仅进行了水平分层，还进行了垂直切片，将应用层以下划分成了不同领域（Domain），每个领域责任明确且高度内聚。

领域的划分应该满足单一职责原则，每个领域应当只对同一类行为者负责，每次系统的修改都应该分析属于哪个领域，如果某些领域总是同时被修改，他们应当被合并为一个领域。一旦领域划分后，不同领域之间需要制定严格的边界，领域暴露的接口，事件，领域之间的依赖关系都该被严格把控.


基础设施为外部 pkg目录