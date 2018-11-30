package util

import (
  "encoding/json"
  "net/http"
  "reflect"
)

type CustomFunction = func(*http.Request) (interface{}, *HTTPError)

func Response(w http.ResponseWriter, payload interface{}) {
  w.Header().Set("Content-Type", "application/json")
  w.WriteHeader(http.StatusOK)
  if reflect.TypeOf(payload) == reflect.TypeOf("") {
    payload = map[string]interface{}{
      "message": "Success",
      "status": http.StatusOK,
    }
  }
  json.NewEncoder(w).Encode(payload)
}

func Error(w http.ResponseWriter, err *HTTPError) {
  w.Header().Set("Content-Type", "application/json")
  w.WriteHeader(err.StatusCode)
  body := map[string]interface{}{
    "message": err.Message,
    "status": err.StatusCode,
  }
  json.NewEncoder(w).Encode(body)
}

func ResponseWrapper(httpFunction CustomFunction) http.HandlerFunc {
  return func(w http.ResponseWriter, r *http.Request) {
    payload, err := httpFunction(r)
    if err != nil {
      Error(w, err)
      return
    }
    Response(w, payload)
  }
}
