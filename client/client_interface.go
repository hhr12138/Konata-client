package client

import "time"

type Client interface {
	//Pipeline() Pipeliner
	//Pipelined(fn func(Pipeliner) error) ([]Cmder, error)

	ClientGetName() *StringCmd
	Echo(message interface{}) *StringCmd
	Ping() *StatusCmd
	Quit() *StatusCmd
	Del(keys ...string) *IntCmd
	Unlink(keys ...string) *IntCmd
	Dump(key string) *StringCmd
	Exists(keys ...string) *IntCmd
	Expire(key string, expiration time.Duration) *BoolCmd
	ExpireAt(key string, tm time.Time) *BoolCmd
	Keys(pattern string) *StringSliceCmd
	Migrate(host, port, key string, db int64, timeout time.Duration) *StatusCmd
	Move(key string, db int64) *BoolCmd
	ObjectRefCount(key string) *IntCmd
	ObjectEncoding(key string) *StringCmd
	ObjectIdleTime(key string) *DurationCmd
	Persist(key string) *BoolCmd
	PExpire(key string, expiration time.Duration) *BoolCmd
	PExpireAt(key string, tm time.Time) *BoolCmd
	PTTL(key string) *DurationCmd
	RandomKey() *StringCmd
	Rename(key, newkey string) *StatusCmd
	RenameNX(key, newkey string) *BoolCmd
	Restore(key string, ttl time.Duration, value string) *StatusCmd
	RestoreReplace(key string, ttl time.Duration, value string) *StatusCmd
	Sort(key string, sort Sort) *StringSliceCmd
	SortInterfaces(key string, sort Sort) *SliceCmd
	TTL(key string) *DurationCmd
	HTTL(key string) *DurationCmd
	ZTTL(key string) *DurationCmd
	LTTL(key string) *DurationCmd
	STTL(key string) *DurationCmd
	Type(key string) *StatusCmd
	GetShardCount() *IntCmd
	IScan(shard uint64, cursor string, match string, count int64) *IScanCmd
	Scan(cursor uint64, match string, count int64) *ScanCmd
	SScan(key string, cursor uint64, match string, count int64) *ScanCmd
	ScanSet(key string, cursor string, count int64) *ScanWithStringCursorCmd
	HScan(key string, cursor uint64, match string, count int64) *ScanCmd
	ScanHash(key string, cursor string, count int64) *ScanWithStringCursorCmd
	ZScan(key string, cursor uint64, match string, count int64) *ScanCmd
	Append(key, value string) *IntCmd
	BitCount(key string, bitCount *BitCount) *IntCmd
	BitOpAnd(destKey string, keys ...string) *IntCmd
	BitOpOr(destKey string, keys ...string) *IntCmd
	BitOpXor(destKey string, keys ...string) *IntCmd
	BitOpNot(destKey string, key string) *IntCmd
	BitPos(key string, bit int64, pos ...int64) *IntCmd
	Decr(key string) *IntCmd
	DecrExpire(key string, expiration time.Duration) *IntCmd
	DecrBy(key string, decrement int64) *IntCmd
	DecrByExpire(key string, decrement int64, expiration time.Duration) *IntCmd
	XDecrBy(key string, decrement int64) *IntCmd
	Get(key string) *StringCmd
	CGet(key string) *IntCmd
	DpsAddKey(key string) *StatusCmd
	GetBit(key string, offset int64) *IntCmd
	GetRange(key string, start, end int64) *StringCmd
	GetSet(key string, value interface{}) *StringCmd
	Incr(key string) *IntCmd
	IncrExpire(key string, expiration time.Duration) *IntCmd
	CIncr(key string) *IntCmd
	IncrBy(key string, value int64) *IntCmd
	IncrByExpire(key string, value int64, expiration time.Duration) *IntCmd
	CIncrBy(key string, value int64) *IntCmd
	IncrByFloat(key string, value float64) *FloatCmd
	AlIncrBy(key string, value int64) *IntCmd
	MGet(keys ...string) *SliceCmd
	CMGet(keys ...string) *SliceCmd
	MSet(pairs ...interface{}) *StatusCmd
	MSetNX(pairs ...interface{}) *BoolCmd
	Set(key string, value interface{}, expiration time.Duration) *StatusCmd
	CSet(key string, value int64, expiration time.Duration) *StatusCmd
	CasEx(key string, oldvalue interface{}, newvalue interface{}, expiration time.Duration) *IntCmd
	Cas(key string, oldvalue interface{}, newvalue interface{}) *IntCmd
	Cas2(key string, oldvalue interface{}, newvalue interface{}, expiration time.Duration) *IntCmd
	Cas2NX(key string, oldvalue interface{}, newvalue interface{}, expiration time.Duration) *IntCmd
	Cad(key string, value interface{}) *IntCmd
	SetBit(key string, offset int64, value int) *IntCmd
	SetNX(key string, value interface{}, expiration time.Duration) *BoolCmd
	SetXX(key string, value interface{}, expiration time.Duration) *BoolCmd
	SetRange(key string, offset int64, value string) *IntCmd
	StrLen(key string) *IntCmd
	HDel(key string, fields ...string) *IntCmd
	HExists(key, field string) *BoolCmd
	HGet(key, field string) *StringCmd
	HGetAll(key string) *StringStringMapCmd
	HIncrBy(key, field string, incr int64) *IntCmd
	HXDecrBy(key, field string, decrement int64) *IntCmd
	HIncrByFloat(key, field string, incr float64) *FloatCmd
	AlHIncrBy(key, field string, incr int64) *IntCmd
	HKeys(key string) *StringSliceCmd
	HLen(key string) *IntCmd
	HMGet(key string, fields ...string) *SliceCmd
	HMSet(key string, fields map[string]interface{}) *StatusCmd
	HMIncrBy(key string, fields map[string]int64) *IntSliceCmd
	HMXDecrBy(key string, fields map[string]int64) *IntSliceCmd
	HSet(key, field string, value interface{}) *BoolCmd
	HSetNX(key, field string, value interface{}) *BoolCmd
	HVals(key string) *StringSliceCmd
	HGetSet(key, field string, value interface{}) *StringCmd
	BLPop(timeout time.Duration, keys ...string) *StringSliceCmd
	BRPop(timeout time.Duration, keys ...string) *StringSliceCmd
	BRPopLPush(source, destination string, timeout time.Duration) *StringCmd
	LIndex(key string, index int64) *StringCmd
	LInsert(key, op string, pivot, value interface{}) *IntCmd
	LInsertBefore(key string, pivot, value interface{}) *IntCmd
	LInsertAfter(key string, pivot, value interface{}) *IntCmd
	LLen(key string) *IntCmd
	LPop(key string) *StringCmd
	LPush(key string, values ...interface{}) *IntCmd
	LPushX(key string, value interface{}) *IntCmd
	LRange(key string, start, stop int64) *StringSliceCmd
	LRem(key string, count int64, value interface{}) *IntCmd
	LSet(key string, index int64, value interface{}) *StatusCmd
	LTrim(key string, start, stop int64) *StatusCmd
	RPop(key string) *StringCmd
	RPopLPush(source, destination string) *StringCmd
	RPush(key string, values ...interface{}) *IntCmd
	RPushX(key string, value interface{}) *IntCmd
	SAdd(key string, members ...interface{}) *IntCmd
	SCard(key string) *IntCmd
	SDiff(keys ...string) *StringSliceCmd
	SDiffStore(destination string, keys ...string) *IntCmd
	SInter(keys ...string) *StringSliceCmd
	SInterStore(destination string, keys ...string) *IntCmd
	SIsMember(key string, member interface{}) *BoolCmd
	SMembers(key string) *StringSliceCmd
	SMove(source, destination string, member interface{}) *BoolCmd
	SPop(key string) *StringCmd
	SPopN(key string, count int64) *StringSliceCmd
	SRandMember(key string) *StringCmd
	SRandMemberN(key string, count int64) *StringSliceCmd
	SRem(key string, members ...interface{}) *IntCmd
	SUnion(keys ...string) *StringSliceCmd
	SUnionStore(destination string, keys ...string) *IntCmd
	XGet(key string) *XGetCmd
	XSet(key string, value interface{}, generation int64, expiration time.Duration) *XSetCmd
	ScanRow(key string, limit int, offset int, target string) *ScanRowCmd
	TtlQPush(key string, ttl time.Duration, item []byte) *IntCmd
	TtlQPopTo(key string, cursor int64) *StatusCmd
	TtlQDelete(key string) *StatusCmd
	TtlQLen(key string) *IntCmd
	TtlQScan(key string, startCursor, endCursor, limit int64) *TtlQScanCmd
	TtlQGet(key string, cursor int64) *StringCmd
	TtlQGetLatestCursor(key string) *IntCmd
	TtlQGetMeta(key string) *TtlQMetaCmd
	TtlQExists(key string) *BoolCmd
	HDrop(keys string) *IntCmd
	ZDrop(keys string) *IntCmd
	LDrop(keys string) *IntCmd
	SDrop(keys string) *IntCmd
	ZAdd(key string, members ...Z) *IntCmd
	ZAddNX(key string, members ...Z) *IntCmd
	ZAddXX(key string, members ...Z) *IntCmd
	ZAddCh(key string, members ...Z) *IntCmd
	ZAddNXCh(key string, members ...Z) *IntCmd
	ZAddXXCh(key string, members ...Z) *IntCmd
	ZIncr(key string, member Z) *FloatCmd
	ZIncrNX(key string, member Z) *FloatCmd
	ZIncrXX(key string, member Z) *FloatCmd
	ZAddWithLimit(key string, limit int64, members ...Z) *IntCmd
	ZIncrWithLimit(key string, limit int64, member Z) *FloatCmd
	ZCard(key string) *IntCmd
	ZCount(key, min, max string) *IntCmd
	ZLexCount(key, min, max string) *IntCmd
	ZIncrBy(key string, increment float64, member string) *FloatCmd
	ZInterStore(destination string, store ZStore, keys ...string) *IntCmd
	ZRange(key string, start, stop int64) *StringSliceCmd
	ZRangeWithScores(key string, start, stop int64) *ZSliceCmd
	ZRangeByScore(key string, opt ZRangeBy) *StringSliceCmd
	ZRangeByLex(key string, opt ZRangeBy) *StringSliceCmd
	ZRangeByScoreWithScores(key string, opt ZRangeBy) *ZSliceCmd
	ZRank(key, member string) *IntCmd
	ZRem(key string, members ...interface{}) *IntCmd
	ZRemRangeByRank(key string, start, stop int64) *IntCmd
	ZRemRangeByScore(key, min, max string) *IntCmd
	ZRemRangeByLex(key, min, max string) *IntCmd
	ZRevRange(key string, start, stop int64) *StringSliceCmd
	ZRevRangeWithScores(key string, start, stop int64) *ZSliceCmd
	ZRevRangeByScore(key string, opt ZRangeBy) *StringSliceCmd
	ZRevRangeByLex(key string, opt ZRangeBy) *StringSliceCmd
	ZRevRangeByScoreWithScores(key string, opt ZRangeBy) *ZSliceCmd
	ZRevRank(key, member string) *IntCmd
	ZScore(key, member string) *FloatCmd
	ZUnionStore(dest string, store ZStore, keys ...string) *IntCmd
	ZPopMax(key string, count uint64) *ZSliceCmd
	ZPopMin(key string, count uint64) *ZSliceCmd
	PFAdd(key string, els ...interface{}) *IntCmd
	PFCount(keys ...string) *IntCmd
	PFMerge(dest string, keys ...string) *StatusCmd
	BgRewriteAOF() *StatusCmd
	BgSave() *StatusCmd
	AlchemyBgSave() *SliceCmd
	ClientKill(ipPort string) *StatusCmd
	ClientList() *StringCmd
	ClientPause(dur time.Duration) *BoolCmd
	ConfigGet(parameter string) *SliceCmd
	ConfigResetStat() *StatusCmd
	ConfigSet(parameter, value string) *StatusCmd
	DBSize() *IntCmd
	FlushAll() *StatusCmd
	FlushAllAsync() *StatusCmd
	FlushDB() *StatusCmd
	FlushDBAsync() *StatusCmd
	Info(section ...string) *StringCmd
	LastSave() *IntCmd
	Save() *StatusCmd
	Shutdown() *StatusCmd
	ShutdownSave() *StatusCmd
	ShutdownNoSave() *StatusCmd
	SlaveOf(host, port string) *StatusCmd
	Time() *TimeCmd
	Eval(script string, keys []string, args ...interface{}) *Cmd
	EvalSha(sha1 string, keys []string, args ...interface{}) *Cmd
	ScriptExists(scripts ...string) *BoolSliceCmd
	ScriptFlush() *StatusCmd
	ScriptKill() *StatusCmd
	ScriptLoad(script string) *StringCmd
	DebugObject(key string) *StringCmd
	PubSubChannels(pattern string) *StringSliceCmd
	PubSubNumSub(channels ...string) *StringIntMapCmd
	PubSubNumPat() *IntCmd
	ClusterSlots() *ClusterSlotsCmd
	ClusterNodes() *StringCmd
	ClusterMeet(host, port string) *StatusCmd
	ClusterForget(nodeID string) *StatusCmd
	ClusterReplicate(nodeID string) *StatusCmd
	ClusterResetSoft() *StatusCmd
	ClusterResetHard() *StatusCmd
	ClusterInfo() *StringCmd
	ClusterKeySlot(key string) *IntCmd
	ClusterCountFailureReports(nodeID string) *IntCmd
	ClusterCountKeysInSlot(slot int) *IntCmd
	ClusterDelSlots(slots ...int) *StatusCmd
	ClusterDelSlotsRange(min, max int) *StatusCmd
	ClusterSaveConfig() *StatusCmd
	ClusterSlaves(nodeID string) *StringSliceCmd
	ClusterFailover() *StatusCmd
	ClusterAddSlots(slots ...int) *StatusCmd
	ClusterAddSlotsRange(min, max int) *StatusCmd
	GeoAdd(key string, geoLocation ...*GeoLocation) *IntCmd
	GeoPos(key string, members ...string) *GeoPosCmd
	GeoRadius(key string, longitude, latitude float64, query *GeoRadiusQuery) *GeoLocationCmd
	GeoRadiusRO(key string, longitude, latitude float64, query *GeoRadiusQuery) *GeoLocationCmd
	GeoRadiusByMember(key, member string, query *GeoRadiusQuery) *GeoLocationCmd
	GeoRadiusByMemberRO(key, member string, query *GeoRadiusQuery) *GeoLocationCmd
	GeoDist(key string, member1, member2, unit string) *FloatCmd
	GeoHash(key string, members ...string) *StringSliceCmd
	Command() *CommandsInfoCmd
	Raw(args ...string) *Cmd
	/* Abase special command */
	HashExists(key string) *BoolCmd
	SExists(key string) *BoolCmd
	LExists(key string) *BoolCmd
	ZExists(key string) *BoolCmd
	HSetExpire(key string, expiration time.Duration) *BoolCmd
	LSetExpire(key string, expiration time.Duration) *BoolCmd
	SSetExpire(key string, expiration time.Duration) *BoolCmd
	ZSetExpire(key string, expiration time.Duration) *BoolCmd
	HSetExpires(key string, expiration time.Duration) *BoolCmd
	LSetExpires(key string, expiration time.Duration) *BoolCmd
	SSetExpires(key string, expiration time.Duration) *BoolCmd
	ZSetExpires(key string, expiration time.Duration) *BoolCmd
}
