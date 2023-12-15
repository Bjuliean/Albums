package handler

import (
	"bytes"
	"encoding/json"
	logger "ful/RESTful/src/logs"
	"ful/RESTful/src/storage"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
)

func initGin() (*gin.Engine, *logger.Logger, *storage.Storage) {
	dbconfig := storage.DBConfig{
		Host:     "0.0.0.0",
		Port:     5432,
		User:     "user",
		Password: "qwerty",
		DBName:   "user",
		SSLMode:  "disable",
	}
	var dataStorage storage.Storage
	var reqHandler RequestHandler
	var logs logger.Logger

	logs = logger.NewLogrusLogger("../logs/tests_logs.txt")
	logs.CreateLogger()

	dataStorage = storage.NewPostgresStorage()
	dataStorage.ConnectToDatabase(&dbconfig)
	reqHandler = NewDefaultHandler(&dataStorage, &logs)

	return InitRouter(&reqHandler, &reqHandler), &logs, &dataStorage
}

func freeResources(logs *logger.Logger, db *storage.Storage) {
	(*logs).CloseLogger()
	(*db).CloseConnection()
}

func TestHandlerOK_GetAlbums(t *testing.T) {

	router, logs, db := initGin()
	defer freeResources(logs, db)

	req, _ := http.NewRequest("GET", "/albums", strings.NewReader(""))
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatal("status not ok")
	}
}

func TestHandlerOK_GetAlbum(t *testing.T) {

	router, logs, db := initGin()
	defer freeResources(logs, db)

	req, _ := http.NewRequest("GET", "/albums/1", strings.NewReader(""))
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatal("status not ok")
	}
}

func TestHandlerFATAL_GetAlbum(t *testing.T) {

	router, logs, db := initGin()
	defer freeResources(logs, db)

	req, _ := http.NewRequest("GET", "/albums/132", strings.NewReader(""))
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusNotFound {
		t.Fatal("status not ok")
	}
}

func TestHandlerOK_PostAlbum(t *testing.T) {

	router, logs, db := initGin()
	defer freeResources(logs, db)

	jbody := storage.Album{
		ID:     12,
		Title:  "ohoho",
		Artist: "aboba",
		Price:  123.45,
	}
	rbody, _ := json.Marshal(jbody)

	req, _ := http.NewRequest("POST", "/albums", bytes.NewReader(rbody))
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusCreated {
		t.Fatal("status not ok")
	}
}

func TestHandlerFATAL_DELETEAlbum(t *testing.T) {

	router, logs, db := initGin()
	defer freeResources(logs, db)

	req, _ := http.NewRequest("DELETE", "/albums/132", strings.NewReader(""))
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusNotFound {
		t.Fatal("status not ok")
	}
}

func TestHandlerOK_DELETEAlbum(t *testing.T) {

	router, logs, db := initGin()
	defer freeResources(logs, db)

	req, _ := http.NewRequest("DELETE", "/albums/2", strings.NewReader(""))
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusNoContent {
		t.Fatal("status not ok")
	}
}

func TestHandlerOK_UpdateAlbum(t *testing.T) {

	router, logs, db := initGin()
	defer freeResources(logs, db)

	jbody := storage.Album{
		ID:     12,
		Title:  "ohoho",
		Artist: "aboba",
		Price:  123.45,
	}
	rbody, _ := json.Marshal(jbody)

	req, _ := http.NewRequest("PUT", "/albums/3", bytes.NewReader(rbody))
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusCreated {
		t.Fatal("status not ok")
	}
}
