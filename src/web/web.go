package web

import (
	"html/template"
	"net/http"
	// "strconv"

	"db"

	// "encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2/bson"
)

// Present Information on Dedicated WebPortal

func Index(res http.ResponseWriter, req *http.Request) {
	// var rs Server
	rs, err := db.GetAll()

	// t := template.Must(template.New("tmpl").Parse(tmpl))
	t, _ := template.ParseFiles("templates/index.html")

	t.Execute(res, rs)

	fmt.Println("err", err)
	fmt.Println("rs", rs)
}

func Send(res http.ResponseWriter, req *http.Request) {
	var CMDBName string
	var Function string
	var SerialNumber string

	req.ParseForm()

	id := bson.NewObjectId()
	CMDBName = req.FormValue("CMDBName")
	Function = req.FormValue("Function")
	SerialNumber = req.FormValue("SerialNumber")

	item := db.Server{ID: id, CMDBName: CMDBName, Function: Function, SerialNumber: SerialNumber}
	if err := db.Save(item); err != nil {
		// handleError(err, "Failed to save data: %v", res)
		return
	}

	// http redirect to index
	http.Redirect(res, req, "/index", http.StatusSeeOther)
}

func Delete(res http.ResponseWriter, req *http.Request) {
	var id string
	req.ParseForm()

	id = req.FormValue("id")

	if err := db.Remove(id); err != nil {
		// handleError(err, "Failed to remove item: %v", w)
		return
	}
	// http redirect to index
	http.Redirect(res, req, "/index", http.StatusSeeOther)
}

func Details(res http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id := vars["id"]
	fmt.Println("id", id)

	rs, err := db.GetOne(id)
	fmt.Println("rs", rs)
	fmt.Println("err", err)
	t, _ := template.ParseFiles("templates/details.html")
	t.Execute(res, rs)
}
