package api

import (
	"net/http"

	"db"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	// "strconv"
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

func GetAllItems(w http.ResponseWriter, req *http.Request) {
	rs, err := db.GetAll()
	if err != nil {
		handleError(err, "Failed to load database items: %v", w)
		return
	}

	bs, err := json.MarshalIndent(rs, "", "    ")
	if err != nil {
		handleError(err, "Failed to load marshal data: %v", w)
		return
	}

	w.Write(bs)
}

func handleError(err error, message string, w http.ResponseWriter) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte(fmt.Sprintf(message, err)))
}

// PostItem saves an item (form data) into the database.
func PostItem(w http.ResponseWriter, req *http.Request) {
	var server Server
	decoder := json.NewDecoder(req.Body)
	errjson := decoder.Decode(&server)
	if errjson != nil {
		fmt.Println("Incorrect body")
		return
	}
	fmt.Println("server", server)

	ID := server.ID
	fmt.Println("ID", ID)
	Name := server.CMDBName
	Function := server.Function
	SerialNumber := server.SerialNumber

	item := db.Server{ID: ID, CMDBName: Name, Function: Function, SerialNumber: SerialNumber}

	if err := db.Save(item); err != nil {
		handleError(err, "Failed to save data: %v", w)
		return
	}
	w.Write([]byte("OK"))
}

func DeleteItem(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id := vars["id"]

	if err := db.Remove(id); err != nil {
		handleError(err, "Failed to remove item: %v", w)
		return
	}

	w.Write([]byte("OK"))
}

func GetItem(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id := vars["id"]

	rs, err := db.GetOne(id)
	if err != nil {
		handleError(err, "Failed to read database: %v", w)
		return
	}

	bs, err := json.Marshal(rs)
	if err != nil {
		handleError(err, "Failed to marshal data: %v", w)
		return
	}

	w.Write(bs)
}
