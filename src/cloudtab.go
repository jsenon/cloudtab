package main

import (
	"api"
	"github.com/gorilla/mux"
	"web"

	"net/http"
)

// TO DO
// Package for struct JSON Server
// Web part to view details
// Web part Update
// Import CSV
// Detail Select Column to show

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
	r.HandleFunc("/delete", web.Delete)
	r.HandleFunc("/details/{id}", web.Details)

	// API Part
	r.HandleFunc("/servers", api.GetAllItems).Methods("GET")
	r.HandleFunc("/servers", api.PostItem).Methods("POST")
	r.HandleFunc("/servers/{id}", api.DeleteItem).Methods("DELETE")
	r.HandleFunc("/servers/{id}", api.GetItem).Methods("GET")

	http.ListenAndServe(":9010", r)
}
