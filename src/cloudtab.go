//go:generate swagger generate spec
// Package main CloudTab.
//
// the purpose of this application is to provide an CMDB application
// that will store information in mongodb backend
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
package main

import (
	"api"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"net/http"
	"web"
)

// TO FIX

// Struct JSON
// Server{
// 	CMDBName
// 	Function
// 	SerialNumber
// 	AssetCode
// 	HardwareDefinition {
// 		Model
// 		CPU
// 		RAM
// 	}
// 	Localisation {
// 		Room
// 		Building
// 		Rack
// 	}
// 	Networks {
// 		IpAddr
// 		PatchPanel
// 		ServerPort
// 		Switch
// 		Vlan
// 		MAC
// 	}
// 	Remarks
//  Status
// }

func main() {

	r := mux.NewRouter()

	// Remove CORS Header check to allow swagger and application on same host and port
	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type"})
	// To be changed
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS", "PATCH"})

	// Web Part
	r.HandleFunc("/index", web.Index)
	r.HandleFunc("/send", web.Send)
	r.HandleFunc("/sendupdate/{id}", web.SendUpdate)
	r.HandleFunc("/delete", web.Delete)
	r.HandleFunc("/details/{id}", web.Details)
	r.HandleFunc("/update/{id}", web.Update)
	r.HandleFunc("/netupdate/{id}", web.NetUpdate)
	r.HandleFunc("/sendnetupdate/{id}", web.SendNetUpdate)

	// Swagger
	r.HandleFunc("/api", web.ApiHelp)

	// Login
	r.HandleFunc("/login", web.Login)

	// API Part
	r.HandleFunc("/api/servers", api.GetAllItems).Methods("GET")
	r.HandleFunc("/api/servers", api.PostItem).Methods("POST")
	r.HandleFunc("/api/servers/{id}", api.DeleteItem).Methods("DELETE")
	r.HandleFunc("/api/servers/{id}", api.GetItem).Methods("GET")
	r.HandleFunc("/api/servers/{id}", api.UpdateItem).Methods("PATCH")
	r.HandleFunc("/api/servers/import", api.PostMultipleItems).Methods("POST")
	r.HandleFunc("/api/status/am-i-up", api.Statusamiup).Methods("GET")
	r.HandleFunc("/api/status/about", api.Statusabout).Methods("GET")

	// Health Check
	r.HandleFunc("/status/am-i-up", api.Statusamiup).Methods("GET")
	r.HandleFunc("/status/about", api.Statusabout).Methods("GET")

	// Static dir
	r.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir("templates/static/"))))

	http.ListenAndServe(":9010", handlers.CORS(originsOk, headersOk, methodsOk)(r))
}
