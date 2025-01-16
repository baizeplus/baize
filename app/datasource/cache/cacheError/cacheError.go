package cacheError

const Nil CacheError = "cache: nil"

type CacheError string

func (e CacheError) Error() string { return string(e) }

func (CacheError) RedisError() {}
