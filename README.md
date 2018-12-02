# postman-twitter
* Written in GoLang
* Uses Relational DB - Postgres (Pardon to not to use GraphQL)
* Uses JWT over web sessions
* Uses Redis to create blacklist (expired JWT Tokens) for session maintainence


## Libraries Used
Note: Just for reference. Go builds everything from vendor.
* https://github.com/go-chi/chi
* https://github.com/dgrijalva/jwt-go
* https://github.com/go-redis/redis
* https://github.com/lib/pq
* https://github.com/jmoiron/sqlx/reflectx
* https://golang.org/x/crypto/bcrypt

# How to build?

## Go Root
1. Install Go and set up ENVIRONMENT (https://golang.org/doc/install)
2. Clone repository to $HOME/go/src/

## Postgres
1. Install Postgres 
2. Create a database named, 'postman-twitter'
3. Run `db` bash script present in the project root to execute psql db schema on local

```bash
sh db
```

## Redis
1. Install Redis: brew install redis
2. Start Redis Server: redis-server /usr/local/etc/redis.conf
3. Default port: 6379

## Running Server
1. Run 'run' bash script present in the project root to build and run server
```bash
sh run
```
2. Current port: 8000
3. Change port configuration in /config/config_local.json

# Endpoints

## Signup

Endpoint: 
```
/api/v1/auth/signup
```
Method: POST <br>
Body: 
```json
{
	"username": "yashishdua",
	"password": "test"
}
```

## Signin

Endpoint: 
```
/api/v1/auth/signin
```
Method: POST <br>
Body: 
```json
{
	"username": "yashishdua",
	"password": "test"
}
``` 

Successful Response:
```json
{
  "jwt_token": "eyJhbGciOiJIUzI1NiIsInR5c...."
}
```

## Signout

Endpoint: 
```
/api/v1/auth/signout
```
Method: POST <br>
Header:
```
Key: Authorization
Value: Bearer <jwt_token>
```


## Follow
ASSUMPTION: You are providing valid user id <br>
Endpoint: 
```
/api/v1/user/follow 
```
Method: POST <br>
Header:
```
Key: Authorization
Value: Bearer <jwt_token>
```
Body: 

```json
{
	"following_user_id": "8559ab00-8a02-487e-8b82-3adbf4fbe69e",
	"followed_by_user_id": "8559ab00-8a02-487e-8b82-3adbf4fbe99e"
}
```

## UnFollow
ASSUMPTION: You are providing valid user id <br>
Endpoint: 
```
/api/v1/user/unfollow
```
Method: POST <br>
Header:
```
Key: Authorization
Value: Bearer <jwt_token>
```
Body: 

```json
{
	"following_user_id": "8559ab00-8a02-487e-8b82-3adbf4fbe69e",
	"followed_by_user_id": "8559ab00-8a02-487e-8b82-3adbf4fbe99e"
}
```

<br>

### These APIs are too naive, lot more to improve here!
Singing off.


