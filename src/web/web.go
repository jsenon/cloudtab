package web

import (
	"html/template"
	"net/http"
	// "strconv"
	"db"
	// "encoding/json"
	"fmt"
)

type Server struct {
	ID               string               `json:"id" bson:"_id,omitempty"`
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

func Render(w http.ResponseWriter, filename string, data interface{}) {
	tmpl, err := template.ParseFiles(filename)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	if err := tmpl.Execute(w, data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func Index(res http.ResponseWriter, req *http.Request) {
	// var rs Server
	rs, err := db.GetAll()

	data := struct {
		Title        string
		CMDBName     string
		Function     string
		SerialNumber string
	}{
		Title:        "cloudtab",
		CMDBName:     rs[0].CMDBName,
		Function:     rs[0].Function,
		SerialNumber: rs[0].SerialNumber,
	}

	Render(res, "templates/index.html", data)

	fmt.Println("err", err)
	fmt.Println("rs", rs)

	// for i := range rs {

	// 	data := struct {
	// 		Title        string
	// 		CMDBName     string
	// 		Function     string
	// 		SerialNumber string
	// 	}{
	// 		Title:        "cloudtab",
	// 		CMDBName:     rs[i].CMDBName,
	// 		Function:     rs[i].Function,
	// 		SerialNumber: rs[i].SerialNumber,
	// 	}

	// 	Render(res, "templates/index.html", data)
	// }
}

func handleError(err error, message string, w http.ResponseWriter) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte(fmt.Sprintf(message, err)))
}
