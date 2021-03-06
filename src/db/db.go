// Package db CloudTab.
//
// the purpose of this package is to provide DB Interface to mongodb
//
// Terms Of Service:
//
// there are no TOS at this moment, use at your own risk we take no responsibility
//
//     Schemes: http
//     Host: localhost
//     BasePath: /api
//     Version: 0.0.1
//     License: MIT http://opensource.org/licenses/MIT
//     Contact: Julien SENON <julien.senon@gmail.com>
//
package db

import (
	// "fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
	"time"
)

// Server represents the asset for this application
//
// A Server have multiple information to be stored.
//
// swagger:model server
type Server struct {
	// ID Server Generated
	ID bson.ObjectId `json:"id" bson:"_id,omitempty"`
	// Server Name
	CMDBName string `json:"CMDBName"`
	// Server Function
	Function string `json:"Function"`
	// Server Serial Number
	SerialNumber string `json:"SerialNumber"`
	// Server Asset Code
	AssetCode string `json:"Assetcode" `
	// Server Hardware Definition
	HardwareDefinition HardwareDefinition `json:"Hardwarerows"`
	// Server Localisation
	Localisation Localisation `json:"Localisationrows"`
	// Server Network Definition
	Networking []Networks `json:"Networksrows"`
	// Remark associate to server
	Remarks string `json:"Remarks"`
	// Server Status
	Status string `json:"Status"`
	// Server Update Time
	UpdateTime time.Time `json:"UpdateTime,omitempty" bson:"UpdateTime"`
	// Server Insert Time
	InsertTime time.Time `json:"InsertTime,omitempty" bson:"InsertTime"`
}

type HardwareDefinition struct {
	// Server Model
	Model string `json:"Model"`
	// Server CPU Speed
	CPU string `json:"CPU"`
	// Server Memory
	RAM string `json:"RAM"`
}

type Localisation struct {
	// Server Room
	Room string `json:"Room"`
	// Server Building
	Building string `json:"Building"`
	// Server Rack
	Rack string `json:"Rack"`
}

type Networks struct {
	// Server Ip Add associated
	IpAddr string `json:"Ipaddr"`
	// Server Patch Panel Number
	PatchPanel string `json:"Patchpanel"`
	// Server Port
	ServerPort string `json:"Serverport"`
	// Server Switch
	Switch string `json:"Switch"`
	// Server vlan
	Vlan string `json:"Vlan"`
	// Server VLAN
	MAC string `json:"MAC"`
}

var db *mgo.Database

// Init connection to MongoDB
func init() {
	session, err := mgo.Dial("localhost")
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	db = session.DB("cloudtab")
}

// Link connection to a collection
func collection() *mgo.Collection {
	return db.C("items")
}

// Func to retrieve element server from collection
func GetAll() ([]Server, error) {
	res := []Server{}

	if err := collection().Find(nil).All(&res); err != nil {
		return nil, err
	}

	return res, nil
}

// Func to add a server
func Save(item Server) error {
	return collection().Insert(item)
}

// Func used to Update Network Part as it s a slice
func Update(key string, item string, values Networks) error {
	// Func ONLY available for []network{} collection
	// return collection().Remove(bson.M{"_id": bson.ObjectIdHex(id)})
	// db.items.updateOne({cmdbname:"ServerTest99"},{$addToSet:{"networking": [ { "ipaddr" : "abcd", "patchpanel" : "abcd", "serverport" : "abcd", "switch" : "abcd", "vlan" : "abcd", "mac" : "abcdefg" } ]}})
	_, err := collection().Upsert(bson.M{"_id": bson.ObjectIdHex(key)}, bson.M{"$addToSet": bson.M{item: values}})
	// fmt.Println(err, key, bson.ObjectIdHex(key), item, values)
	return err
}

//Main Update func
func Updatemain(key string, item Server) (error, error) {
	err := collection().Update(bson.M{"_id": bson.ObjectIdHex(key)}, item)
	// fmt.Println(item)

	UpdateTime := time.Now()
	erro := collection().Update(bson.M{"_id": bson.ObjectIdHex(key)}, bson.M{"$set": bson.M{"UpdateTime": UpdateTime}})
	// fmt.Println(erro, UpdateTime)
	return err, erro
}

// Func to retrieve only one elements, used to get details about item
func GetOne(id string) (*Server, error) {
	res := Server{}

	if err := collection().Find(bson.M{"_id": bson.ObjectIdHex(id)}).One(&res); err != nil {
		return nil, err
	}
	return &res, nil
}

// Func to remve a server form collection
func Remove(id string) error {
	return collection().Remove(bson.M{"_id": bson.ObjectIdHex(id)})
}

// Func to remove network field from collection
// TODO multi criteria in $pull    { $pull: { Networksrows: {  $elemMatch: { Ipaddr: Ipaddr, MAC: MAC } } } },
func RemoveNetwork(id string, Ipaddr string) error {
	change := bson.M{"$pull": bson.M{"Networksrows": Ipaddr}}
	// fmt.Println(id)
	err := collection().UpdateId(bson.ObjectIdHex(id), change)
	return err
}

// Func to add multiple server
// JSON Struct is a slice of item JSON Server
func MultipleSave(item []Server) error {
	return collection().Insert(item)
}
