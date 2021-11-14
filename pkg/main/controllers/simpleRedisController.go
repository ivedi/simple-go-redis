package controllers

import (
	"fmt"
	"net/http"
	// "github.com/ivedi/simple-go-redis/pkg/main/repositories"
)

type SimpleRedisController struct{}

func (S SimpleRedisController) Get(w http.ResponseWriter, r *http.Request) {
	fmt.Println("hello darling")
	// repositories.SimpleRedisRepository.Get("test")
}

func (S SimpleRedisController) Set(w http.ResponseWriter, r *http.Request) {
	fmt.Println("hello papa")
	// repositories.SimpleRedisRepository.Set("test", "asdlfkj")
}
