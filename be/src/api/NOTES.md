Team Google Docs:
https://docs.google.com/document/d/1gJ4HMwvIHg4xFzEETJy6QcPwH-7_qz8selOytcBvtTE/edit

Udemy Github:
https://github.com/GoesToEleven/golang-web-dev

Error Code middleware:
https://levelup.gitconnected.com/sending-http-error-codes-with-golang-and-gin-gonic-d915d1dd0166

Examples:

- Bank Application with gin:

  - https://dev.to/techschoolguru/implement-restful-http-api-in-go-using-gin-4ap1
  - https://github.com/techschool/simplebank
    https://gobyexample.com/interfaces
    https://github.com/gin-gonic/examples
    https://go.dev/doc/tutorial/web-service-gin
    https://8thlight.com/blog/emmanuel-byrd/2022/05/19/well-thought-project-layout-design-for-a-golang-backend.html

## Working Directory

```
📦BE
 ┣ 📂config
 ┃ ┣ 📜.env
 ┃ ┗ 📜schema.sql
 ┣ 📂src
 ┃ ┣ 📂api
 ┃ ┃ ┣ 📂middleware
 ┃ ┃ ┃ ┣ 📜email.go
 ┃ ┃ ┃ ┣ 📜errorhandler.go
 ┃ ┃ ┃ ┣ 📜get-user-group.go
 ┃ ┃ ┃ ┗ 📜password.go
 ┃ ┃ ┣ 📂route
 ┃ ┃ ┃ ┣ 📜admin-create-user.go
 ┃ ┃ ┃ ┗ 📜route.exe
 ┃ ┃ ┣ 📜NOTES.md
 ┃ ┃ ┣ 📜go.mod
 ┃ ┃ ┗ 📜go.sum
 ┃ ┗ 📜main.go
 ┗ 📜README.md
```

---

### Working with Gin
`Context.JSON` serialize given struct as JSON into the response body




---

### STEPS

### JWT

- login
  - [] username

  - [] password

  - [] assign new JWT token to user

  - [] encode JSON to JWT

  - [] decode JWT to JSON

`json.Marshal()` return JSON obj from data structure

`json.Unmarshal()` return data structure formed from JSON obj

````
curl http://localhost:4000/login
--include \
--header "Content-Type: application/json" \
--request "POST" \
--data '{"username": "user1","password": "password1"}'
```

- create user

  - [x] username

  - [x] password (hash in db)

  - bcrypt

  - [x] email (optional)

  - net/mail
  - regex

  - [x] active status

  - [x] composite key usergroup
  - [x] validate groupname

  - [x] admin privilege (optional. degault value in db = 0)

  - [x] timestamp

- work on API + postman
- connection to database (mySQL)
- data entry into database
- error handling middleware
- create composite key usergroup

  - [x] username exist

  - [x] password invalid format

  - [x] email invalid format

## To-Do

- [] Server struct with env variables
- [x] Get user group for dropdown
- [x] Refactor middlewares into separate files and import to route

---

- type User struct

  - dummy data in JSON format
  - fetch user data
  - `c.IndentedJSON` : serialized struct into JSON

- type Response struct

  - error codes for different response

```
router.POST("/admin-create-user", func(c *gin.Context) {
    username := c.PostForm("username")
    password := c.PostForm("password")
    email := c.PostForm("email")
    status := c.PostForm("status")

    // response 200
    c.JSON(http.StatusOK, gin.H{
        "message": "user successfully created",
        "username": username,
        "password": password,
        "email": email,
        "status": status,
    })
})
```

Using curl

```
curl http://localhost:4000/admin-create-user \
--include \
--header "Content-Type: application/json" \
--request "POST" \
--data '{"username": "team_member","password": "Admin_123","email": "team_member@tms.com","status": "Active"}'
```

`curl http://localhost:4000/get-users`

```
curl http://localhost:4000/get-user \
--include \
--header "Content-Type: application/json" \
--request "POST" \
--data '{"username": "alfred"}'
```

---

## Response middleware

https://sosedoff.com/2014/12/21/gin-middleware.html

```
type Response struct {
Status int
Error []string
}
```

Function to send response middleware

```
func ErrorHandler(c *gin.Context, code int, message interface{}) {
	c.IndentedJSON(code, gin.H{"error": message})
}
```

Example calling SendResponse

### Unauthorized Access

`SendResponse(c, helpers.Response{Status: http.StatusUnauthorized, Error: []string{"Username and password do not match"}})`

#### Wrong parameters

`SendResponse(c, helpers.Response{Status: http.StatusBadRequest, Error: []string{"One or more params are wrong"}})`

## Using gin framework to send JSON response

`c.IndentedJSON(400, gin.H{"code": 400, "error": "Password length should be between length 8 - 10 with numbers and special characters"})`

---
