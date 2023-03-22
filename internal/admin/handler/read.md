接收前端参数，再将处理结果返回到前端。
handler 调用 Facade 层是：一对一接口调用，且 handler 层不做任何业务处理，目的是为了后续拓展直接替换 handler 为 RPC 框架而准备。
一个 handler 只能调用一个 Facade