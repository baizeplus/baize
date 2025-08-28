package systemModels

type SseType struct {
	Key   string
	Value string
}

type Sse struct {
	UserIds      []string
	RedisPublish bool
	Sse          *SseType
}
