// Package api CloudTab.
//
// the purpose of this package is to provide Api Interface
//
// Terms Of Service:
//
// there are no TOS at this moment, use at your own risk we take no responsibility
//
//     Schemes: http
//     Host: localhost:9010
//     BasePath: /api
//     Version: 0.0.1
//     License: MIT http://opensource.org/licenses/MIT
//     Contact: Julien SENON <julien.senon@gmail.com>
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
// swagger:meta
package api

import (
	"net/http"

	"db"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2/bson"
	"io/ioutil"
	"time"
	// "strconv"
)

// Input failed validation.
// swagger:response validationError
type ValidationError struct {
	// The error message
	// in: body
	Body struct {
		Code    int32  `json:"code"`
		Message string `json:"message"`
		Field   string `json:"field"`
	} `json:"body"`
}

// Success Answer
// swagger:response successAnswer
type successAnswer struct {
	// The error message
	// in: body
	Body struct {
		Code    int32  `json:"code"`
		Message string `json:"message"`
		Field   string `json:"field"`
	} `json:"body"`
}

// Server id
// swagger:parameters deleteServer getServerById updateServer
type id struct {
	// id Server Generated
	//
	// in: path
	// required: true
	ID int64 `json:"id"`
}

// Server Param
// swagger:parameters createServer updateServer
type myServerBodyParams struct {
	// Server to submit
	//
	// in: body
	// required: true
	Server *db.Server `json:"server"`
}

// Server Multi
// swagger:parameters  importServer
type myMultipleServerBodyParams struct {
	// Server to submit
	//
	// in: body
	// required: true
	Server *[]db.Server `json:"multiserver"`
}

// swagger:route GET /servers servers listServer
//
// Lists servers
//
// This will show all available asset by default.
//
//
//     Responses:
//       default: validationError
//       200: []server
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

// swagger:route POST /servers servers createServer
//
// Add a server
//
// This will register asset.
//
//     Responses:
//       default: validationError
//       200: successAnswer
func PostItem(w http.ResponseWriter, req *http.Request) {
	var server db.Server

	decoder := json.NewDecoder(req.Body)

	// debug, err := json.MarshalIndent(decoder, "", "    ")
	// fmt.Println(" debug", debug)
	// fmt.Println(" err", err)

	errjson := decoder.Decode(&server)

	if errjson != nil {
		fmt.Println("Incorrect body")
		fmt.Println(errjson)
		return
	}
	// fmt.Println(server)

	id := bson.NewObjectId()
	ref := id.Hex()
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
	UpdateTime := time.Now()
	InsertTime := time.Now()

	fmt.Println(UpdateTime)
	fmt.Println(InsertTime)

	// IpAddr := server.Networking[0].IpAddr
	// PatchPanel := server.Networking[0].PatchPanel
	// ServerPort := server.Networking[0].ServerPort
	// Switch := server.Networking[0].Switch
	// Vlan := server.Networking[0].Vlan
	// MAC := server.Networking[0].MAC

	// item := db.Server{ID: id, CMDBName: Name, Networking: []db.Networks{{IpAddr: IpAddr, PatchPanel: PatchPanel, ServerPort: ServerPort, Switch: Switch, Vlan: Vlan, MAC: MAC}}}

	item := db.Server{ID: id, CMDBName: Name, Function: Function, SerialNumber: SerialNumber, AssetCode: AssetCode, HardwareDefinition: db.HardwareDefinition{Model: Model, CPU: CPU, RAM: RAM}, Localisation: db.Localisation{Room: Room, Building: Building, Rack: Rack}, Remarks: Remarks, Status: Status, UpdateTime: UpdateTime, InsertTime: InsertTime}
	if err := db.Save(item); err != nil {
		handleError(err, "Failed to save data: %v", w)
		return
	}

	// Idea is add static field and use loop for update array of networks
	// Networking: []db.Networks{{IpAddr: IpAddr, PatchPanel: PatchPanel, ServerPort: ServerPort, Switch: Switch, Vlan: Vlan, MAC: MAC}},
	var collec string
	collec = "networking"
	fmt.Println(server.Networking)

	for _, net := range server.Networking {
		fmt.Println(net)
		// Launch update with id = id
		if err := db.Update(ref, collec, net); err != nil {
			handleError(err, "Failed to save data: %v", w)
			return
		}
	}
	w.Write([]byte("OK"))
}

// swagger:route DELETE /servers/{id} servers deleteServer
//
// Delete servers
//
// This will delete asset.
//
//     Responses:
//       default: validationError
//       200: successAnswer
func DeleteItem(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id := vars["id"]

	if err := db.Remove(id); err != nil {
		handleError(err, "Failed to remove item: %v", w)
		return
	}
	w.Write([]byte("OK"))
}

// swagger:route GET /servers/{id} servers getServerById
//
// Lists specific server
//
// This will list details for specific server.
//
//     Responses:
//       default: validationError
//       200: server
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

// swagger:route PATCH /servers/{id} servers updateServer
//
// Update specific server
//
// This will update details for specific server.
//
//     Responses:
//       default: validationError
//       200: successAnswer
func UpdateItem(w http.ResponseWriter, req *http.Request) {
	// fmt.Println("Im in update api")
	vars := mux.Vars(req)
	id := vars["id"]
	// fmt.Println("id", id)

	var server db.Server

	decoder := json.NewDecoder(req.Body)

	errjson := decoder.Decode(&server)

	if errjson != nil {
		fmt.Println("Incorrect body")
		fmt.Println(errjson)
		return
	}

	if erro, err := db.Updatemain(id, server); err != nil && erro != nil {
		handleError(err, "Failed to save data: %v", w)
		return
	}
	w.Write([]byte("OK"))
}

// swagger:route POST /servers/import servers importServer
//
// Import / Add multiple Servers
//
// This will insert multiple servers.
//
//     Responses:
//       default: validationError
//       200: successAnswer
func PostMultipleItems(w http.ResponseWriter, req *http.Request) {
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		panic(err)
	}
	var server []db.Server
	err = json.Unmarshal(body, &server)
	// fmt.Println("Body:", body)
	if err != nil {
		//handle error
	}
	// Loop for multiple of server included in JSON
	for i := range server {
		fmt.Println(server[i].CMDBName)
		fmt.Println(server[i])
		if err := db.Save(server[i]); err != nil {
			handleError(err, "Failed to save data: %v", w)
			return
		}

	}
}
