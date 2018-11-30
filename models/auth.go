package models

import (
  "postman-twitter/database"
)

type UserAuth struct {
  Username string `db:"username", json:"username"`
  Password string `db:"password", json:"password"`
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
func SignIn(userAuth UserAuth) (UserAuth, error) {
  sqlSelectQuery := "SELECT * FROM user_auth WHERE username = $1"
	existingUserAuth := UserAuth{}
	err := database.DB.Get(&existingUserAuth, sqlSelectQuery, userAuth.Username)
	return existingUserAuth, err
}
