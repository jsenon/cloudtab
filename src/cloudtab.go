package main

import (
	"api"
	"github.com/gorilla/mux"
	"web"

	"net/http"
)

// TO DO
// Package for struct JSON
// Webserver

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

	r.HandleFunc("/index", web.Index)
	r.HandleFunc("/send", web.Send)
	// API Part
	r.HandleFunc("/servers", api.GetAllItems).Methods("GET")
	r.HandleFunc("/servers", api.PostItem).Methods("POST")
	r.HandleFunc("/servers/{id}", api.DeleteItem).Methods("DELETE")
	r.HandleFunc("/servers/{id}", api.GetItem).Methods("GET")

	http.ListenAndServe(":9010", r)
}
