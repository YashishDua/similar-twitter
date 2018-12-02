-- DROPPING TABLES (TO AVOID MULTIPLE TABLE CREATING ERROR)
DROP TABLE IF EXISTS user_auth;
DROP TABLE IF EXISTS follow;
DROP EXTENSION pgcrypto;


-- LOAD PRE COMPILED LIBRARY TO USE GET_RANDOM_UUID
CREATE EXTENSION pgcrypto;


-- CREATING TABLES
CREATE TABLE user_auth (
  user_id UUID UNIQUE DEFAULT gen_random_uuid(),
  username VARCHAR(50) PRIMARY KEY,
  password TEXT
);

CREATE TABLE follow (
  following_user_id UUID NOT NULL,
  followed_by_user_id UUID NOT NULL,
  PRIMARY KEY (following_user_id, followed_by_user_id)
);
