package controllers

import (
	"fmt"
	"net/http"
)

type SimpleRedisController struct{}

func (S SimpleRedisController) Get(w http.ResponseWriter, r *http.Request) {
	fmt.Println("hello darling")
}

func (S SimpleRedisController) Set(w http.ResponseWriter, r *http.Request) {
	fmt.Println("hello papa")
}
