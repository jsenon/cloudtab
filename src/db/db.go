package db

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
	"time"
)

type Server struct {
	ID                 bson.ObjectId `json:"id" bson:"_id,omitempty"`
	CMDBName           string        `json:"CMDBName"`
	Function           string        `json:"Function"`
	SerialNumber       string        `json:"SerialNumber"`
	AssetCode          string        `json:"Assetcode" `
	HardwareDefinition `json:"Hardwarerows"`
	Localisation       `json:"Localisationrows"`
	Networking         []Networks `json:"Networksrows"`
	Remarks            string     `json:"Remarks"`
	Status             string     `json:"Status"`
	UpdateTime         time.Time  `json:"UpdateTime,omitempty" bson:"UpdateTime"`
	InsertTime         time.Time  `json:"InsertTime,omitempty" bson:"InsertTime"`
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
	Vlan       string `json:"Vlan"`
	MAC        string `json:"MAC"`
}

type Network struct {
	NCMDBName     string `json:"NCMDBName"`
	NFunction     string `json:"NFunction"`
	NSerialNumber string `json:"NSerialNumber"`
	NAssetCode    int    `json:"NAssetCode"`
	NRows         []Localisation
}

type Person struct {
	User     string
	Password string
	Mail     string
}

var db *mgo.Database

func init() {
	session, err := mgo.Dial("localhost")
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	db = session.DB("cloudtab")
}

func collection() *mgo.Collection {
	return db.C("items")
}

func GetAll() ([]Server, error) {
	res := []Server{}

	if err := collection().Find(nil).All(&res); err != nil {
		return nil, err
	}

	return res, nil
}

func Save(item Server) error {
	return collection().Insert(item)
}

func Update(key string, item string, values Networks) error {
	// Func ONLY available for []network{} collection
	// return collection().Remove(bson.M{"_id": bson.ObjectIdHex(id)})
	// db.items.updateOne({cmdbname:"ServerTest99"},{$addToSet:{"networking": [ { "ipaddr" : "abcd", "patchpanel" : "abcd", "serverport" : "abcd", "switch" : "abcd", "vlan" : "abcd", "mac" : "abcdefg" } ]}})
	_, err := collection().Upsert(bson.M{"_id": bson.ObjectIdHex(key)}, bson.M{"$addToSet": bson.M{item: values}})
	fmt.Println(err, key, bson.ObjectIdHex(key), item, values)
	return err
}

//Main Update func
func Updatemain(key string, item Server) (error, error) {
	// Func ONLY available for []network{} collection
	// return collection().Remove(bson.M{"_id": bson.ObjectIdHex(id)})
	// db.items.updateOne({cmdbname:"ServerTest99"},{$addToSet:{"networking": [ { "ipaddr" : "abcd", "patchpanel" : "abcd", "serverport" : "abcd", "switch" : "abcd", "vlan" : "abcd", "mac" : "abcdefg" } ]}})
	err := collection().Update(bson.M{"_id": bson.ObjectIdHex(key)}, item)

	UpdateTime := time.Now()
	erro := collection().Update(bson.M{"_id": bson.ObjectIdHex(key)}, bson.M{"$set": bson.M{"UpdateTime": UpdateTime}})
	fmt.Println(erro, UpdateTime)

	return err, erro
}

func GetOne(id string) (*Server, error) {
	res := Server{}

	if err := collection().Find(bson.M{"_id": bson.ObjectIdHex(id)}).One(&res); err != nil {
		return nil, err
	}
	return &res, nil
}

func Remove(id string) error {
	return collection().Remove(bson.M{"_id": bson.ObjectIdHex(id)})
}
