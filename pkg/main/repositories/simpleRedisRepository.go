package repositories

type SimpleRedisRepository struct{}

var inMemoryData map[string]string

func (S SimpleRedisRepository) Get(key string) string {
	return inMemoryData[key]
}

func (S SimpleRedisRepository) Set(key string, value string) {
	inMemoryData[key] := value
}
