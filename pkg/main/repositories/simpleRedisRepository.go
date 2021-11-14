package repositories

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"time"
)

type SimpleRedisRepository struct{}

var inMemoryData map[string]string
var synchronizer *time.Timer

const syncInterval = time.Second * 5000

const storePath = "./pkg/main/repositories/simpleRedisStore.json"

func (S SimpleRedisRepository) Connect() {
	inMemoryData = getStore()
	synchronizer = time.AfterFunc(syncInterval, func() {
		setStore(inMemoryData)
		synchronizer.Reset(syncInterval)
	})
}

func (S SimpleRedisRepository) Disconnect() {
	synchronizer.Stop()
	setStore(inMemoryData)
}

func (S SimpleRedisRepository) Get(key string) string {
	return inMemoryData[key]
}

func (S SimpleRedisRepository) Set(key string, value string) {
	inMemoryData[key] = value
}

func (S SimpleRedisRepository) Flush() {
	inMemoryData = make(map[string]string)
}

func getStore() map[string]string {
	file, err := os.Open(storePath)
	if err != nil {
		return make(map[string]string)
	}
	defer file.Close()
	data, err := ioutil.ReadAll(file)
	if err != nil {
		return make(map[string]string)
	}
	var store map[string]string
	if json.Unmarshal(data, store) != nil {
		return make(map[string]string)
	}
	return store
}

func setStore(store map[string]string) error {
	data, err := json.Marshal(store)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(storePath, data, 0777)
}
