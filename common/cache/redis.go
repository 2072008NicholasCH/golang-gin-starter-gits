package cache

const (
	// FiveMinutes is TTL for the cache
	FiveMinutes = 300
	// OneHour is TTL for the cache
	OneHour = 3600
	// OneMonth is TTL for the cache
	OneMonth = 2592000

)

const(
	prefix = "gin-starter-gits"
	// BookFindByUUID is a redis key dor finding book by uuid
	BookFindByUUID = prefix + ":book:find-book-by-uuid:%v"
)