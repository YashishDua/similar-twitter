package main

import (
  "fmt"
  "log"
  "net/http"
  "github.com/go-chi/chi"
  chiMiddleWare "github.com/go-chi/chi/middleware"
  "postman-twitter/config"
  "postman-twitter/database"
  "postman-twitter/util"
  "postman-twitter/middleware"
  "postman-twitter/endpoints"
  "postman-twitter/redis"
)

func main() {
  config.Init()
  database.Init()
  redis.Init()
  defer database.DB.Close()

  router := chi.NewRouter()
  router.Use(chiMiddleWare.Logger)
  router.Get("/", middleware.ResponseWrapper(Hello, util.AUTH_NOT_REQ))

  router.Post("/signup", middleware.ResponseWrapper(endpoints.SignUpHandler, util.AUTH_NOT_REQ))
  router.Post("/signin", middleware.ResponseWrapper(endpoints.SignInHandler, util.AUTH_NOT_REQ))

  router.Post("/signout", middleware.ResponseWrapper(endpoints.LogOutHandler, util.AUTH_REQ))
  router.Post("/follow", middleware.ResponseWrapper(endpoints.FollowHandler, util.AUTH_REQ))

  fmt.Println("Running on " + config.ServerConfig.Port)
  log.Fatal(http.ListenAndServe(":" + config.ServerConfig.Port, router))
}

//HealhCheck API
func Hello(r *http.Request) (interface{}, *util.HTTPError) {
  return "Under Construction", nil
}
