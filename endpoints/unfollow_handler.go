package endpoints

import (
  "encoding/json"
  "net/http"
  "postman-twitter/util"
  "postman-twitter/models"
)

/*
  UnFollowHandler removes the follow relationship from table
*/
func UnFollowHandler(r *http.Request) (interface{}, *util.HTTPError) {
  decoder := json.NewDecoder(r.Body)
  var follow models.Follow
  err := decoder.Decode(&follow)
  if err != nil {
      return nil, util.BadRequest(util.BAD_JSON_ERROR)
  }
  err = models.RemoveFollower(follow)
  if err != nil {
    return nil, util.InternalServerError(util.SQL_ERROR + " (" + err.Error() + ")")
  }
  return util.GENERIC_SUCCESS_RESPONSE, nil
}
