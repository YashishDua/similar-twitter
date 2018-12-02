package models

import (
  "github.com/google/uuid"
  "postman-twitter/database"
)
type Follow struct {
    FollowingUserID   uuid.UUID `db:"following_user_id"   json:"following_user_id"`
    FollowedByUserID  uuid.UUID `db:"followed_by_user_id" json:"followed_by_user_id"`
}

func AddFollower(follow Follow) error {
  sqlInsertQuery := "INSERT INTO follow (following_user_id, followed_by_user_id) VALUES " +
                    "(:following_user_id, :followed_by_user_id)"
  tx := database.DB.MustBegin()
  _, err := tx.NamedExec(sqlInsertQuery, follow)
  if err != nil {
    tx.Rollback()
    return err
  }
  tx.Commit()
  return nil
}
