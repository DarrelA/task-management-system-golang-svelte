# Admin Update User

### Accounts Table

1. check if `username` is provided
1. retrieve username, password, email, usergroup, status (isactive) from frontend.
1. trim username and email
1. check if username exist (otherwise return error)
```
SELECT * FROM accounts WHERE username = ?
```
1. 