package models

import (
  "github.com/google/uuid"
  "postman-twitter/database"
)

type UserAuth struct {
  ID        *uuid.UUID  `db:"user_id" json:"user_id"`
  Username  string      `db:"username", json:"username"`
  Password  string      `db:"password", json:"password"`
}

func SignUp(userAuth UserAuth) error {
  sqlInsertQuery := "INSERT INTO user_auth (username, password) VALUES " +
                    "(:username, :password)"
  tx := database.DB.MustBegin()
  _, err := tx.NamedExec(sqlInsertQuery, userAuth)
  if err != nil {
    tx.Rollback()
    return err
  }
  tx.Commit()
  return nil
}

// TODO: TRANSACTIONS
func GetCredentials(username string) (UserAuth, error) {
  sqlSelectQuery := "SELECT * FROM user_auth WHERE username = $1"
	existingUserAuth := UserAuth{}
	err := database.DB.Get(&existingUserAuth, sqlSelectQuery, username)
	return existingUserAuth, err
}
