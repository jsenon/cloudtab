package main

import (
	"api"

	"fmt"
	"github.com/gorilla/mux"

	"net/http"
)

// TO DO
// Websocket Gorilla
// Backend DB MongoDB

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

// Mongo Help
// mgo.Dial("server1.example.com,server2.example.com")

func ErrorWithJSON(w http.ResponseWriter, message string, code int) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)
	fmt.Fprintf(w, "{message: %q}", message)
}

func ResponseWithJSON(w http.ResponseWriter, json []byte, code int) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)
	w.Write(json)
}

func main() {

	r := mux.NewRouter()
	r.HandleFunc("/servers", api.GetAllItems).Methods("GET")
	r.HandleFunc("/servers", api.PostItems).Methods("POST")

	http.ListenAndServe(":9010", r)
}

// Function Get Server
// Function Add Server
// Function Remove Server
// Function Update Server
