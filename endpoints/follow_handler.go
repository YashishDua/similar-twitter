package endpoints

import (
  "encoding/json"
  "net/http"
  "log"
  "postman-twitter/util"
  "postman-twitter/models"
)

func FollowHandler(r *http.Request) (interface{}, *util.HTTPError) {
  decoder := json.NewDecoder(r.Body)
  var follow models.Follow
  err := decoder.Decode(&follow)
  log.Println(follow)
  if err != nil {
      return nil, util.BadRequest(util.BAD_JSON_ERROR)
  }
  err = models.AddFollower(follow)
  if err != nil {
    return nil, util.InternalServerError(util.SQL_ERROR)
  }
  return util.GENERIC_SUCCESS_RESPONSE, nil
}
