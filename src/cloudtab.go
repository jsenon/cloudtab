package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
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

func main() {

}
