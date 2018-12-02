package middleware

import (
  "encoding/json"
  "strings"
  "net/http"
  "postman-twitter/auth"
  "postman-twitter/util"
  "postman-twitter/models"
  "postman-twitter/redis"
)

type CustomFunction = func(*http.Request) (interface{}, *util.HTTPError)

func Response(w http.ResponseWriter, payload interface{}) {
  w.Header().Set("Content-Type", "application/json")
  w.WriteHeader(http.StatusOK)
  if payload == util.GENERIC_SUCCESS_RESPONSE {
    payload = map[string]interface{}{
      "message": "Success",
      "status": http.StatusOK,
    }
  }
  json.NewEncoder(w).Encode(payload)
}

func Error(w http.ResponseWriter, err *util.HTTPError) {
  w.Header().Set("Content-Type", "application/json")
  w.WriteHeader(err.StatusCode)
  body := map[string]interface{}{
    "error": err.Message,
    "status": err.StatusCode,
  }
  json.NewEncoder(w).Encode(body)
}

func ResponseWrapper(httpFunction CustomFunction, authReq bool) http.HandlerFunc {
  return func(w http.ResponseWriter, r *http.Request) {
    if authReq {
      authErr := validateAuthorization(r)
      if authErr != nil {
        Error(w, authErr)
        return
      }
    }
    payload, err := httpFunction(r)
    if err != nil {
      Error(w, err)
      return
    }
    Response(w, payload)
  }
}

func validateAuthorization(r *http.Request) *util.HTTPError {
	authStrings := strings.SplitN(r.Header.Get("Authorization"), " ", 3)
	if len(authStrings) != 2 || authStrings[0] != "Bearer" {
		return util.Unauthorized(util.NO_BEARER_PRESENT)
	}
  isExpired := isInBlackList(authStrings[1])
  if isExpired {
    return util.Unauthorized(util.INVALID_JWT)
  }
	jwtAuthInfo, err := auth.DecodeJWTAuth(authStrings[1])
	if err != nil {
		return util.Unauthorized(err.Error())
	}
  var existingUserAuth models.UserAuth
  existingUserAuth, err = models.GetAuthDetails(jwtAuthInfo.Username)
  if err != nil {
    return util.Unauthorized(util.INVALID_JWT)
  }
  if existingUserAuth.Password != jwtAuthInfo.Password {
    return util.Unauthorized(util.INVALID_JWT)
	}
	return nil
}

// Check for token expiration in Redis Cache
func isInBlackList(key string) bool {
  value, err := redis.GetKeyValue(key)
  if err != nil {
    return false
  }
  if value == nil {
    return false
  }
  return true
}
