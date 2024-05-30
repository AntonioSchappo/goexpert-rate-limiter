package main

import (
	"github.com/AntonioSchappo/goexpert-rate-limiter/api"
	"github.com/AntonioSchappo/goexpert-rate-limiter/config"
	"github.com/AntonioSchappo/goexpert-rate-limiter/ratelimiter"
	"github.com/AntonioSchappo/goexpert-rate-limiter/repository"
)

func main() {
	configs, err := config.GetConfig(".")
	if err != nil {
		panic(err)
	}
	repo := repository.NewRedisRepository(configs.Host, "", 0)
	limiter := ratelimiter.NewRateLimiter(repo, configs.GetLimitConfig())
	server := api.NewAPIServer(configs.Port, limiter)
	server.Run()
}
