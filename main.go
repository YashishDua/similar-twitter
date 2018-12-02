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
const (
  AUTH_REQ bool = true
  AUTH_NOT_REQ bool = false
)

func routes() chi.Router {
	router := chi.NewRouter()
  router.Get("/ping", middleware.ResponseWrapper(Ping, AUTH_NOT_REQ))

	router.Route("/user", func(r chi.Router) {
    r.Post("/{userID}/follow", middleware.ResponseWrapper(endpoints.FollowHandler, AUTH_REQ))
    r.Post("/{userID}/unfollow", middleware.ResponseWrapper(endpoints.UnFollowHandler, AUTH_REQ))
	})

  router.Route("/auth", func(r chi.Router) {
    r.Post("/signup", middleware.ResponseWrapper(endpoints.SignUpHandler, AUTH_NOT_REQ))
    r.Post("/signin", middleware.ResponseWrapper(endpoints.SignInHandler, AUTH_NOT_REQ))
    r.Post("/signout", middleware.ResponseWrapper(endpoints.LogOutHandler, AUTH_REQ))
	})

	return router
}

func main() {
  config.Init()
  database.Init()
  redis.Init()
  defer database.DB.Close()

  router := chi.NewRouter()
  router.Use(chiMiddleWare.Logger)
  router.Mount("/api/v1", routes())

  fmt.Println("Running on " + config.ServerConfig.Port)
  log.Fatal(http.ListenAndServe(":" + config.ServerConfig.Port, router))
}

//HealhCheck API
func Ping(r *http.Request) (interface{}, *util.HTTPError) {
  return "Pong", nil
}
