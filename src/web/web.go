package web

import (
	"html/template"
	"net/http"
	// "strconv"

	"db"

	// "encoding/json"
	"fmt"
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
