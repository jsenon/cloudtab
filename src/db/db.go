package db

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
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

func GetOne(id string) (*Server, error) {
	res := Server{}

	if err := collection().Find(bson.M{"_id": bson.ObjectIdHex(id)}).One(&res); err != nil {
		return nil, err
	}
	fmt.Println("res", res)
	return &res, nil
}

func Remove(id string) error {
	return collection().Remove(bson.M{"_id": bson.ObjectIdHex(id)})
}
