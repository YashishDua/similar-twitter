package util

const (
  //GERENIC SUCCESS
  GENERIC_SUCCESS_RESPONSE string = "Success"

  //GENERIC ERRORS
  BAD_JSON_ERROR string = "Bad JSON structure"
  DECODING_ERROR string = "Error in decoding the request body"
  SALTING_ERROR string = "Error while salting password"
  SQL_ERROR string = "Unable to perform SQL query"

  //AUTH
  AUTH_REQ bool = true
  AUTH_NOT_REQ bool = false
  USER_DOES_NOT_EXIST_ERROR string = "User doesn't exist"
  USER_ALREADY_EXIST_ERROR string = "User already exist"
  MISMATCH_PASSWORD_ERROR string = "Password didn't match"
  JWT_ERROR string = "Error while creating JWT token"
  INVALID_JWT string = "Invalid Auth Token"
  NO_BEARER_PRESENT string = "No Bearer in Authorization Header"



)
