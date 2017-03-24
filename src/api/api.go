package api

import (
	"net/http"

	"db"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2/bson"
	// "strconv"
)

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
	var server db.Server

	decoder := json.NewDecoder(req.Body)

	// debug, err := json.MarshalIndent(decoder, "", "    ")
	// fmt.Println(" debug", debug)
	// fmt.Println(" err", err)

	errjson := decoder.Decode(&server)

	if errjson != nil {
		fmt.Println("Incorrect body")
		return
	}
	fmt.Println(server)

	id := bson.NewObjectId()
	Name := server.CMDBName
	Function := server.Function
	SerialNumber := server.SerialNumber
	AssetCode := server.AssetCode
	Model := server.HardwareDefinition.Model
	CPU := server.HardwareDefinition.CPU
	RAM := server.HardwareDefinition.RAM
	Room := server.Localisation.Room
	Building := server.Localisation.Building
	Rack := server.Localisation.Rack
	Remarks := server.Remarks
	Status := server.Status

	IpAddr := server.Networking[0].IpAddr
	PatchPanel := server.Networking[0].PatchPanel
	ServerPort := server.Networking[0].ServerPort
	Switch := server.Networking[0].Switch
	Vlan := server.Networking[0].Vlan
	MAC := server.Networking[0].MAC

	item := db.Server{ID: id, CMDBName: Name, Function: Function, SerialNumber: SerialNumber, AssetCode: AssetCode, HardwareDefinition: db.HardwareDefinition{Model: Model, CPU: CPU, RAM: RAM}, Localisation: db.Localisation{Room: Room, Building: Building, Rack: Rack}, Networking: []db.Networks{{IpAddr: IpAddr, PatchPanel: PatchPanel, ServerPort: ServerPort, Switch: Switch, Vlan: Vlan, MAC: MAC}}, Remarks: Remarks, Status: Status}
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
	fmt.Println("id", id)

	rs, err := db.GetOne(id)
	// fmt.Println("rs", rs)
	// fmt.Println("err", err)

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
