// 即使定义结构体，在服务端也需要拼接成字符串给到日志模块主从同步，之后再反序列化成命令执行，所以直接规定协议，客户端传递字符串类型命令。
namespace go db.raft.Konata_clent

struct GetArgs {
    1: string key,
}

struct GetReply {
    1: string value,
}

struct PutAppendArgs{
    1: string key,
    2: string value,
    3: string op,
}

struct PutAppendReply{
    1: bool success,
    2: string message,
}

service KonataService {
    GetReply Get(1: GetArgs args),
    PutAppendReply PutAppend(1: PutAppendArgs args),
}

