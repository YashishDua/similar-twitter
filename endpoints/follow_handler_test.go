package endpoints

import (
  "testing"
  "postman-twitter/database"
)

func setupFollowTest() {
	database.InitTestDB()
	database.RefreshTableInDB("follow")
}

func TestFollow(t *testing.T) {
  //TODO
}
