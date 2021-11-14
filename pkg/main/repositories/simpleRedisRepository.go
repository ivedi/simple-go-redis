package repositories

type SimpleRedisRepository struct{}

var InMemoryData map[string]string

func (S SimpleRedisRepository) Get(key string) string {
	return InMemoryData[key]
}

func (S SimpleRedisRepository) Set(key string, value string) string {
	repositories.InMemoryData[key] := value
}
