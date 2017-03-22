package main

import (
	"api"

	"fmt"
	"github.com/gorilla/mux"

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

	// API Part
	r.HandleFunc("/servers", api.GetAllItems).Methods("GET")
	r.HandleFunc("/servers", api.PostItem).Methods("POST")
	r.HandleFunc("/servers/{id}", api.DeleItem).Methods("DELETE")
	r.HandleFunc("/servers/{id}", api.GetItems).Methods("GET")

	http.ListenAndServe(":9010", r)
}
