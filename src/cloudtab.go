package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"net/http"
	"webserver"
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

type Server struct {
	CMDBName         string               `json:"CMDBName"`
	Function         string               `json:"Function"`
	SerialNumber     string               `json:"SerialNumber"`
	AssetCode        int                  `json:"Assetcode"`
	HardwareRows     []HardwareDefinition `json:"Hardwarerows"`
	LocalisationRows []Localisation       `json:"Localisationrows"`
	NetworksRows     []Networks           `json:"Networksrows"`
	Remarks          string               `json:"Remarks"`
	Status           string               `json:"Status"`
}

type HardwareDefinition struct {
	Model string `json:"Model"`
	CPU   string `json:"CPU"`
	RAM   string `json:"RAM"`
}

type Localisation struct {
	Room     string `json:"Room"`
	Building string `json:"Building"`
	Rack     string `json:"Rack"`
}

type Networks struct {
	IpAddr     string `json:"Ipaddr"`
	PatchPanel string `json:"Patchpanel"`
	ServerPort string `json:"Serverport"`
	Switch     string `json:"Switch"`
	Vlan       int    `json:"Vlan"`
	MAC        string `json:"MAC"`
}

type Network struct {
	CMDBName     string `json:"NCMDBName"`
	Function     string `json:"NFunction"`
	SerialNumber string `json:"NSerialNumber"`
	AssetCode    int    `json:"NAssetCode"`
	Rows         []Localisation
}

type Person struct {
	User     string
	Password string
	Mail     string
}

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
	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	defer session.Close()
	r := mux.NewRouter()
	r.HandleFunc("/servers", getserversHandler(session)).Methods("GET")
	r.HandleFunc("/servers", postserversHandler(session)).Methods("POST")

	http.Handle("/", r)
	http.ListenAndServe(":9010", r)
}

// Function Get Server
func (s *mgo.Session) getserversHandler(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		session := s.Copy()
		defer session.Close()

		c := session.DB("store").C("servers")

		var books []Book
		err := c.Find(bson.M{}).All(&books)
		if err != nil {
			ErrorWithJSON(w, "Database error", http.StatusInternalServerError)
			log.Println("Failed get all servers: ", err)
			return
		}

		respBody, err := json.MarshalIndent(servers, "", "  ")
		if err != nil {
			log.Fatal(err)
		}

		ResponseWithJSON(w, respBody, http.StatusOK)
	}
}

// Function Add Server
func (s *mgo.Session) postserversHandler(w http.ResponseWriter, r *http.Request) {

}

// Function Remove Server
// Function Update Server
