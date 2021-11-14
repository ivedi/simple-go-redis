package controllers

import (
	"io/ioutil"
	"net/http"

	// "github.com/ivedi/simple-go-redis/pkg/main/repositories"
	"cmd/main/main.go/pkg/main/repositories"

	"github.com/gorilla/mux"
)

type SimpleRedisController struct{}

var repository = repositories.SimpleRedisRepository{}

func (S SimpleRedisController) Initialize() {
	repository.Connect()
}

func (S SimpleRedisController) Get(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["key"]
	value := repository.Get(key)
	w.Write([]byte(value))
}

func (S SimpleRedisController) Set(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["key"]
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	value := string(body)
	repository.Set(key, value)
}

func (S SimpleRedisController) Flush(w http.ResponseWriter, r *http.Request) {
	repository.Flush()
}

func (S SimpleRedisController) Close() {
	repository.Disconnect()
}
