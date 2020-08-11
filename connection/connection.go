package connection

import (
	"errors"
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
func FindById(id string) (model.Shopping, error)  {
	var shopping model.Shopping
	if !bson.IsObjectIdHex(id){
		err := errors.New("Invalid ID")
		return shopping, err
	}

	session, err := mgo.DialWithInfo(INFO)
	if err != nil {
		log.Fatal(err)
		return shopping, err
	}
	defer session.Close()
	c := session.DB(DBNAME).C(DOCNAME)

	oid := bson.ObjectIdHex(id)
	err = c.FindId(oid).One(&shopping)
	return shopping, err
}

//Funcion Actualizar
func Update(shopping model.Shopping) error  {
	session, err := mgo.DialWithInfo(INFO)
	if err != nil {
		log.Fatal(err)
		return err
	}
	defer session.Close()
	c := session.DB(DBNAME).C(DOCNAME)
	err = c.UpdateId(shopping.ID, &shopping)
	return err
}

//Funcion encontrar por usuario
func FindByUser(idUser int) ([]model.Shopping, error)  {
	var shoppings []model.Shopping
	session, err := mgo.DialWithInfo(INFO)
	if err != nil {
		log.Fatal(err)
		return shoppings, err
	}
	defer session.Close()
	c := session.DB(DBNAME).C(DOCNAME)

	err = c.Find(bson.M{"user": idUser}).All(&shoppings)
	return shoppings, err
}

//Funcion Eliminar
func Delete(id string) error  {
	if !bson.IsObjectIdHex(id){
		err := errors.New("Invalid ID")
		return err
	}
	session, err := mgo.DialWithInfo(INFO)
	if err != nil {
		log.Fatal(err)
		return err
	}
	defer session.Close()
	c := session.DB(DBNAME).C(DOCNAME)

	oid := bson.ObjectIdHex(id)
	err = c.Remove(oid)
	return err
}