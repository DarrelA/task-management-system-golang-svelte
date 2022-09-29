package route

import (
	"backend/api/middleware"
	"backend/api/models"
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetOneApplication(w http.ResponseWriter, r *http.Request, c *gin.Context) {
	var application models.Application
	var data []models.Application

	result := middleware.SelectSingleApplication(application.AppAcronym)
	err := result.Scan(&application.AppAcronym)
	if (err != sql.ErrNoRows) {
		query := r.URL.Query()
		appAcronym := query.Get(application.AppAcronym)
		w.WriteHeader(200)
		w.Write([]byte(appAcronym))
		fmt.Println(application)
	} else if (err == sql.ErrNoRows) {
		middleware.ErrorHandler(c, 400, "AppAcronym does not exist")
	}

	// append response into slice
	data = append(data, application)

	// send data as array of JSON obj
	c.JSON(200, data)
}