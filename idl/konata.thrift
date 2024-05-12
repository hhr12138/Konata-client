// 即使定义结构体，在服务端也需要拼接成字符串给到日志模块主从同步，之后再反序列化成命令执行，所以直接规定协议，客户端传递字符串类型命令。
namespace go db.raft.konata_client

typedef string OpType

const OpType Write = "write"
const OpType Read = "read";
const OpType RemoveId = "remove_id";

typedef i32 ErrCode

//	Nil                     ErrorCode = 0
  //	ErrCodeCommon           ErrorCode = 50001
  //	ErrCodeIsNotLeader      ErrorCode = 50013 // 当前节点不是leader


const ErrCode Nil = 0;

const ErrCode ErrCodeRspParseFail = 40034;
const ErrCode ErrCodeCommandParseFail = 40035;

const ErrCode ErrCodeCommon = 50001
const ErrCode ErrCodeMasterReplace = 50012;
const ErrCode ErrCodeIsNotLeader = 50013;
 const 	 ErrCode ErrCodeNetworkPartition = 50014 // 网络分区异常
  const	ErrCode ErrCodeApplyTimeout      = 50015 // 写操作等待日志提交超时，需要重试
  const	ErrCode ErrCodeServiceDead       = 50016
 const 	ErrCode ErrCodeDiscardMsg        = 50017
 const 	ErrCode ErrCodeTypeErr           = 50018

struct BizErr {
    1: ErrCode code,
    2: string message,
    3: bool repeat,
}

struct Command {
    1: string req_id,
    2: string msg,
}

struct GetArgs {
    1: string req_id,
    2: string command,
    3: OpType op,
}

struct BaseReply {
    1: string Addr,
}

struct Reply {
    1: string value,
    2: BizErr error,
    255: BaseReply base,
}

struct PutAppendArgs {
    1: string req_id,
    2: string command,
    3: OpType op,
}

struct RequestVoteArgs {
    1: i64 term,
    2: i64 candidate_id,
    3: i64 last_log_index,
    4: i64 last_log_term,
}

struct RequestVoteReply {
    1: i64 term,
    2: bool vote_granted,
}

struct Log {
    1: i64 term,
    2: i64 index,
    3: string command,
}

struct RequestAppendArgs {
    1: i64 term,
    2: i64 leader_id,
    3: i64 prev_log_index,
    4: i64 prev_log_term,
    5: list<Log> entries,
    6: i64 leader_commit,
}

struct RequestAppendReply {
    1: i64 term,
    2: bool success,
    3: i64 next_index,
    4: i64 log_term
}

service KonataService {
    Reply Get(1: GetArgs args),
    Reply PutAppend(1: PutAppendArgs args),
    Reply RemoveReqId(1: GetArgs args),
    RequestVoteReply RequestVote(1: RequestVoteArgs args),
    RequestAppendReply AppendEntries(1: RequestAppendArgs args),
}

