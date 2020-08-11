package connection

import (
	"github.com/SanchezjCoronado/Golang-CRUD-MVC/model"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
	"time"
)

//Conexion a la base de datos mongoDB
var INFO= &mgo.DialInfo{
	Addrs: []string{},//Ingresar la ip y el puerto
	Timeout: 60 * time.Second,//Tiempo de salida
	Database: "",//Nombre de la base de datos
	Username: "",//Usuario del la db
	Password: "",//Contraseña de la db
}
//Nombre de la base de datos
const DBNAME = ""
//Nombre del documento de mogo
const DOCNAME = ""

var db *mgo.Database

const (
	COLLECTION = ""
	)


//Funcion Insertar
func Insert(shopping model.Shopping) error {
	session, err := mgo.DialWithInfo(INFO)
	defer session.Close()

	shopping.ID = bson.NewObjectId()
	session.DB(DBNAME).C(DOCNAME).Insert(shopping)

	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}

//Funcion encontrar por ID
