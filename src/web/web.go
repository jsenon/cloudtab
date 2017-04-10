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
package web

import (
	"html/template"
	"net/http"
	"time"
	// "strconv"

	"db"
	"fmt"

	// "encoding/json"
	"github.com/gorilla/mux"
	"github.com/gorilla/schema"

	"gopkg.in/mgo.v2/bson"
)

// Present Information on Dedicated WebPortal

// Func to display all server on table view
func Index(res http.ResponseWriter, req *http.Request) {
	// var rs Server
	rs, err := db.GetAll()
	t, _ := template.ParseFiles("templates/index.html")

	t.Execute(res, rs)
	if err != nil {
		return
	}
}

// Func used to add server from index page
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

// Func to delete {id} Server
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

// Func to display full web details for {id} Servers
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

// Func to display login web form
// Not Used
func Login(res http.ResponseWriter, req *http.Request) {
	t, _ := template.ParseFiles("templates/login.html")

	t.Execute(res, req)

}

// Func to display update web form
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

// Func call by update forms
func SendUpdate(res http.ResponseWriter, req *http.Request) {

	var serverdecode db.Server

	vars := mux.Vars(req)
	id := vars["id"]

	// Use gorilla schema packages that fills a struct with form values
	decoder := schema.NewDecoder()
	fmt.Println(id)
	req.ParseForm()

	err := decoder.Decode(&serverdecode, req.PostForm)
	if err != nil {
		// Handle error
	}

	serverdecode.UpdateTime = time.Now()

	//
	// We need to pass Updatemain(key string , item db.Server)
	//
	if erro, err := db.Updatemain(id, serverdecode); err != nil && erro != nil {
		return
	}

	fmt.Println("err:", err)

	http.Redirect(res, req, "/update/"+id, http.StatusSeeOther)
}

// Func redirected to Swagger
func ApiHelp(res http.ResponseWriter, req *http.Request) {
	// URL To be changed
	http.Redirect(res, req, "/swagger.html?url=http://localhost:9010/swaggermain.json", http.StatusSeeOther)
}

func NetUpdate(res http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id := vars["id"]
	rs, err := db.GetOne(id)
	if err != nil {
		return
	}
	t, _ := template.ParseFiles("templates/netupdate.html")
	t.Execute(res, rs)
}

func SendNetUpdate(res http.ResponseWriter, req *http.Request) {

	var serverdecode db.Server

	vars := mux.Vars(req)
	id := vars["id"]

	// Use gorilla schema packages that fills a struct with form values
	decoder := schema.NewDecoder()
	req.ParseForm()

	fmt.Println(id)

	valuesip := req.Form["ipdyn"]
	valuespp := req.Form["pp"]
	valuessw := req.Form["sw"]
	valuesvlan := req.Form["vlan"]
	valuesmac := req.Form["mac"]
	serverport := "12"

	// Take into account that ip is mandatory and always set
	for i := range valuesip {
		// fmt.Println("x:", x)
		// fmt.Println("ipdynamique:", valuesip[i])
		// fmt.Println("pp:", valuespp[i])
		// fmt.Println("switch:", valuessw[i])
		// fmt.Println("vlan:", valuesvlan[i])
		// fmt.Println("mac:", valuesmac[i])
		net := db.Networks{IpAddr: valuesip[i], PatchPanel: valuespp[i], ServerPort: serverport, Switch: valuessw[i], Vlan: valuesvlan[i], MAC: valuesmac[i]}
		fmt.Println("net:", net)
		// item := db.Server{ID: bson.ObjectIdHex(id), Networking: []db.Networks{{IpAddr: valuesip[i], PatchPanel: valuespp[i], ServerPort: serverport, Switch: valuessw[i], Vlan: valuesvlan[i], MAC: valuesmac[i]}}}
		// if erro, err := db.Updatemain(id, item); err != nil && erro != nil {
		// 	return
		// }
		if err := db.Update(id, "networking", net); err != nil {
			return
		}
	}

	// valuespp := req.Form["pp"]
	// for i := range valuespp {
	// 	fmt.Println("pp loop", valuespp[i])
	// }

	// We can write form to html
	// fmt.Fprintln(res, req.Form)

	err := decoder.Decode(&serverdecode, req.PostForm)
	if err != nil {
		// Handle error
	}
	http.Redirect(res, req, "/netupdate/"+id, http.StatusSeeOther)

}
