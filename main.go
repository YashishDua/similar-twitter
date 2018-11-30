package main

import (
  "fmt"
  "log"
  "net/http"
  "github.com/go-chi/chi"
  "github.com/go-chi/chi/middleware"
  "postman-twitter/config"
  "postman-twitter/database"
  "postman-twitter/util"
  "postman-twitter/endpoints"
)

func main() {
  config.Init()
  database.Init()
  defer database.DB.Close()

  router := chi.NewRouter()
  router.Use(middleware.Logger)
  router.Get("/", util.ResponseWrapper(Hello))

  router.Post("/signup", util.ResponseWrapper(endpoints.SignUpHandler))
  router.Post("/signin", util.ResponseWrapper(endpoints.SignInHandler))

  fmt.Println("Running on " + config.ServerConfig.Port)
  log.Fatal(http.ListenAndServe(":" + config.ServerConfig.Port, router))
}

//HealhCheck API
func Hello(r *http.Request) (interface{}, *util.HTTPError) {
  return "Under Construction", nil
}
