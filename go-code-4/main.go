package main

import (
	"fmt"
	"log"
	"net/http"

	redis "github.com/go-redis/redis/v7"
)

func connectRedis() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     "redislink:6379", // use default Addr
		Password: "",               // no password set
		DB:       0,                // use default DB
	})
}

func main() {
	http.HandleFunc("/ping", pingDb)
	http.HandleFunc("/v1/nome", saveQueryDb)
	http.HandleFunc("/v1/endereco", saveHeaderDb)
	http.HandleFunc("/get", getInfo)
	log.Println("Run Server port:8080")
	log.Fatal(http.ListenAndServe("0.0.0.0:8080", nil))
}

func pingDb(w http.ResponseWriter, r *http.Request) {
	redisDb := connectRedis()
	pong, err := redisDb.Ping().Result()
	if err != nil {
		log.Fatal(err)
	}
	w.Write([]byte(pong))
}

func saveQueryDb(w http.ResponseWriter, r *http.Request) {
	redisDb := connectRedis()
	err := redisDb.Set("nome", r.URL.Query().Get("nome"), 0).Err()
	if err != nil {
		panic(err)
	}
	w.Write([]byte("nome salvo com sucesso!"))
	w.WriteHeader(http.StatusCreated)
}

func saveHeaderDb(w http.ResponseWriter, r *http.Request) {
	redisDb := connectRedis()
	err := redisDb.Set("endereco", r.Header.Get("X-Endereco"), 0).Err()
	if err != nil {
		panic(err)
	}
	w.Write([]byte("endereco salvo com sucesso!"))
	w.WriteHeader(http.StatusCreated)
}

func getInfo(w http.ResponseWriter, r *http.Request) {
	redisDb := connectRedis()
	val, err := redisDb.Get(r.URL.Query().Get("item")).Result()
	if err == redis.Nil {
		w.WriteHeader(http.StatusNotFound)
		msg := fmt.Sprintf("nao foi possivel encontrar o item %s", r.URL.Query().Get("item"))
		w.Write([]byte(msg))
	} else if err != nil {
		panic(err)
	} else {
		w.WriteHeader(http.StatusFound)
		w.Write([]byte(val))
	}
}
