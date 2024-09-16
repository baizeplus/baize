package systemModels

type SseType struct {
	Key   string
	Value string
}

type Sse struct {
	UserIds      []int64
	RedisPublish bool
	Sse          *SseType
}
