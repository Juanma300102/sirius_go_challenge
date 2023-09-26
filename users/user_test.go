package users_test

import (
	"bytes"
	"challenge/db"
	"challenge/models"
	"challenge/router"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestUserList(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := router.SetupRouter()
	db.Migrate()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/v1/users/", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
}

func TestUserCreate(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := router.SetupRouter()
	db.Migrate()
	user := models.User{
		Email: "admin@gmail.com",
		Firstname: "admin",
		Lastname: "admin",
		Dni: "43944733",
		Adress: "Home",
		Password: "1234!",
	}

	w := httptest.NewRecorder()

	body, _ := json.Marshal(user)

	req, _ := http.NewRequest("POST", "/api/v1/users/", bytes.NewReader(body))
	router.ServeHTTP(w, req)
	assert.Equal(t, 201, w.Code)
}

func TestUserUpdate(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := router.SetupRouter()
	db.Migrate()
	user := models.User{
		Email: "admin@gmail.com",
		Firstname: "admin",
		Lastname: "admin",
		Dni: "43944733",
		Adress: "Home",
		Password: "1234!",
	}

	w := httptest.NewRecorder()

	body, _ := json.Marshal(user)

	req, _ := http.NewRequest("POST", "/api/v1/users/", bytes.NewReader(body))

	router.ServeHTTP(w, req)
	assert.Equal(t, 201, w.Code)

	user = models.User{
		Dni: "00000000",
	}

	w = httptest.NewRecorder()

	body, _ = json.Marshal(user)

	req, _ = http.NewRequest("PUT", "/api/v1/users/1", bytes.NewReader(body))

	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
}

func TestUserDelete(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := router.SetupRouter()
	db.Migrate()
	user := models.User{
		Email: "admin@gmail.com",
		Firstname: "admin",
		Lastname: "admin",
		Dni: "43944733",
		Adress: "Home",
		Password: "1234!",
	}

	w := httptest.NewRecorder()

	body, _ := json.Marshal(user)

	req, _ := http.NewRequest("POST", "/api/v1/users/", bytes.NewReader(body))

	router.ServeHTTP(w, req)
	assert.Equal(t, 201, w.Code)

	w = httptest.NewRecorder()

	req, _ = http.NewRequest("DELETE", "/api/v1/users/1", nil)

	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
}