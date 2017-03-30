package main

import (
	"api"
	"github.com/gorilla/mux"
	"net/http"
	"web"
)

// TO DO
// Package for struct JSON Server
// Web part to view details
// Web part Update
// Import CSV
// Detail Select Column to show

// TO FIX
// Responsive details page view

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

	// Login
	r.HandleFunc("/login", web.Login)

	// API Part
	r.HandleFunc("/api/servers", api.GetAllItems).Methods("GET")
	r.HandleFunc("/api/servers", api.PostItem).Methods("POST")
	r.HandleFunc("/api/servers/{id}", api.DeleteItem).Methods("DELETE")
	r.HandleFunc("/api/servers/{id}", api.GetItem).Methods("GET")

	http.ListenAndServe(":9010", r)
}
