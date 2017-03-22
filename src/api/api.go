package api

import (
	"net/http"

	"db"
	"encoding/json"
	"fmt"
	// "github.com/gorilla/mux"
	// "strconv"
)

func GetAllItems(w http.ResponseWriter, req *http.Request) {
	rs, err := db.GetAll()
	if err != nil {
		handleError(err, "Failed to load database items: %v", w)
		return
	}

	bs, err := json.Marshal(rs)
	if err != nil {
		handleError(err, "Failed to load marshal data: %v", w)
		return
	}

	w.Write(bs)
}

func handleError(err error, message string, w http.ResponseWriter) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte(fmt.Sprintf(message, err)))
}

// PostItem saves an item (form data) into the database.
func PostItem(w http.ResponseWriter, req *http.Request) {
	ID := req.FormValue("id")
	valueStr := req.FormValue("value")
	value, err := strconv.Atoi(valueStr)
	if err != nil {
		handleError(err, "Failed to parse input data: %v", w)
		return
	}

	item := db.Item{ID: ID, Value: value}

	if err = db.Save(item); err != nil {
		handleError(err, "Failed to save data: %v", w)
		return
	}

	w.Write([]byte("OK"))
}
