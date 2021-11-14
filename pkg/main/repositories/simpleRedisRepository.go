package repositories

type SimpleRedisRepository struct{}

var inMemoryData map[string]string

func (S SimpleRedisRepository) Initialize() {
	inMemoryData = make(map[string]string)
}

func (S SimpleRedisRepository) Get(key string) string {
	return "test" //string(inMemoryData[key])
}

func (S SimpleRedisRepository) Set(key string, value string) {
	inMemoryData[key] = value
}
