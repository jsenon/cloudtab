// Package web CloudTab.
//
// the purpose of this package is to provide Web HTML Interface
//
// Terms Of Service:
//
// there are no TOS at this moment, use at your own risk we take no responsibility
//
//     Schemes: http
//     Host: localhost
//     BasePath: /
//     Version: 0.0.1
//     License: MIT http://opensource.org/licenses/MIT
//     Contact: Julien SENON <julien.senon@gmail.com>
//
package web

import (
	"html/template"
	"net/http"
	// "strconv"

	"db"

	// "encoding/json"
	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2/bson"
)

// Present Information on Dedicated WebPortal

func Index(res http.ResponseWriter, req *http.Request) {
	// var rs Server
	rs, err := db.GetAll()
	t, _ := template.ParseFiles("templates/index.html")

	t.Execute(res, rs)
	if err != nil {
		return
	}
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

	rs, err := db.GetOne(id)
	if err != nil {
		return
	}
	// fmt.Println("rs", rs)
	t, _ := template.ParseFiles("templates/details.html")
	t.Execute(res, rs)
}

func Login(res http.ResponseWriter, req *http.Request) {
	t, _ := template.ParseFiles("templates/login.html")

	t.Execute(res, req)

}

func Update(res http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id := vars["id"]
	rs, err := db.GetOne(id)
	if err != nil {
		return
	}
	t, _ := template.ParseFiles("templates/update.html")
	t.Execute(res, rs)
}
