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


```go
if (username != "") {
    // Check if username exists in database
    // Error: Username does not exists in database. Please try again.
    if (password != "") {
        // update accounts with NEW password provided
        if (email != "") {
            // Check if email exists in database
            // Error: Email already exists in database. Please try again.
            // Update Accounts with NEW email
            if (user_group) {
                // Append new user_group to accounts table (if any)
                // Add new usergroup and username to the `user_group` table (if any)
                // Update `accounts` table with NEW user_group
                if (status) {
                    // Update Accounts with NEW status
                } else {
                    // update accounts with OLD status, SELECT statement to get old status
                }
            } else {
                // update accounts with OLD user_group, SELECT statement to get old user_group
                if (status) {
                    // Update Accounts with NEW status
                } else {
                    // update accounts with OLD status, SELECT statement to get old status
                }
            }
        } else {
            // update accounts with OLD email, SELECT statement to get old email
            if (user_group) {
                // Append new user_group to accounts table (if any)
                // Add new usergroup and username to the `user_group` table (if any)
                // Update `accounts` table with NEW user_group
                if (status) {
                    // Update Accounts with NEW status
                } else {
                    // update accounts with OLD status, SELECT statement to get old status
                }
            } else {
                // update accounts with OLD user_group, SELECT statement to get old user_group
                if (status) {
                    // Update Accounts with NEW status
                } else {
                    // update accounts with OLD status, SELECT statement to get old status
                }
            }
        }
    } else {
        // update accounts with OLD password, SELECT statement to get old password
        if (email != "") {
            // Check if email exists in database
            // Error: Email already exists in database. Please try again.
            // Update Accounts with NEW email
            if (user_group) {
                // Append new user_group to accounts table (if any)
                // Add new usergroup and username to the `user_group` table (if any)
                // Update `accounts` table with NEW user_group
                if (status) {
                    // Update Accounts with NEW status
                } else {
                    // update accounts with OLD status, SELECT statement to get old status
                }
            } else {
                // update accounts with OLD user_group, SELECT statement to get old user_group
                if (status) {
                    // Update Accounts with NEW status
                } else {
                    // update accounts with OLD status, SELECT statement to get old status
                }
            }
        } else {
            // update accounts with OLD email, SELECT statement to get old email
            if (user_group) {
                // Append new user_group to accounts table (if any)
                // Add new usergroup and username to the `user_group` table (if any)
                // Update `accounts` table with NEW user_group
                if (status) {
                    // Update Accounts with NEW status
                } else {
                    // update accounts with OLD status, SELECT statement to get old status
                }
            } else {
                // update accounts with OLD user_group, SELECT statement to get old user_group
                if (status) {
                    // Update Accounts with NEW status
                } else {
                    // update accounts with OLD status, SELECT statement to get old status
                }
            }
        }
    }
} else {
    // Error: Please provide a username.
}

```