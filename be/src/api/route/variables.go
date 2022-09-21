package route

import (
	"backend/api/middleware"
	"database/sql"
)

var db *sql.DB = middleware.ConnectionToDatabase()
