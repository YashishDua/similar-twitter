package main

import (
  "fmt"
  "log"
  "net/http"
  "github.com/go-chi/chi"
  "github.com/go-chi/chi/middleware"
  "go-yashish/config"
  "go-yashish/database"
  "go-yashish/util"
)

func main() {
  config.Init()
  database.Init()
  defer database.DB.Close()

  router := chi.NewRouter()
  router.Use(middleware.Logger)
  router.Get("/", util.ResponseWrapper(Hello))

  fmt.Println("Running on " + config.ServerConfig.Port)
  log.Fatal(http.ListenAndServe(":" + config.ServerConfig.Port, router))
}

//HealhCheck API
func Hello(r *http.Request) (interface{}, *util.HTTPError) {
  return "Under Construction", nil
}
