-- DROPPING TABLES (TO AVOID MULTIPLE TABLE CREATING ERROR)
DROP TABLE IF EXISTS article;
DROP EXTENSION pgcrypto;

-- LOAD PRE COMPILED LIBRARY TO USE GET_RANDOM_UUID
CREATE EXTENSION pgcrypto;


-- CREATING TABLES
CREATE TABLE article (
  article_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  article_title VARCHAR(100) NULL,
  article_link VARCHAR(100) UNIQUE,
  article_date DATE
);

CREATE TABLE user_auth {
  username TEXT PRIMARY KEY,
  password TEXT
}


-- INSERTING DUMMY DATA
INSERT INTO article (article_title, article_link, article_date)
VALUES
  ('Wrappers for Go', 'TEST_LINK_1', NOW()),
  ('Introduction to LiveData', 'TEST_LINK_2', NOW());
