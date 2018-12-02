package endpoints

import (
  "encoding/json"
  "net/http"
  "github.com/go-chi/chi"
  "github.com/google/uuid"
  "postman-twitter/util"
  "postman-twitter/models"
)

func FollowHandler(r *http.Request) (interface{}, *util.HTTPError) {
  userID := chi.URLParam(r, "userID")
  // Not checking for valid userID
  if len(userID) == 0 {
    return nil, util.BadRequest(util.URL_QUERY_ERROR)
  }
  decoder := json.NewDecoder(r.Body)
  var follow models.Follow
  err := decoder.Decode(&follow)
  if err != nil{
      return nil, util.BadRequest(util.BAD_JSON_ERROR)
  }
  userUUID, _ := uuid.Parse(userID)
  follow.FollowingUserID = userUUID
  err = models.AddFollower(follow)
  if err != nil {
    return nil, util.InternalServerError(util.SQL_ERROR)
  }
  return util.GENERIC_SUCCESS_RESPONSE, nil
}
