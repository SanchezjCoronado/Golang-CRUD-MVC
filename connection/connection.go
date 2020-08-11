package connection

import (
	"gopkg.in/mgo.v2"
	"time"
)

var INFO= &mgo.DialInfo{
	Addrs: []string{},
	Timeout: 60 * time.Second,
	Database: "",
	Username: "",
	Password: "",
}

const DBNAME = ""

const DOCNAME = ""

var db *mgo.Database

const (
	COLLECTION = ""
	)

func Insert( shopping model.Shopping)  {

}