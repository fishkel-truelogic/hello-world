package data

//////////////////////////////////////////////////////////////////////////////
// Imports
//////////////////////////////////////////////////////////////////////////////

import (
	"fmt"
	//"gopkg.in/mgo.v2"
    "gopkg.in/mgo.v2/bson"
    "os"
    "log"
    "hello-world/mdb"
    "hello-world/model"
)

//////////////////////////////////////////////////////////////////////////////
// Data Access Functions
//////////////////////////////////////////////////////////////////////////////

func InsertUser(user model.User) {

	collection := mdb.GetSession().DB("test").C("foo")

	err := collection.Insert(user)

	if err != nil {
		fmt.Printf("Can't insert user: %v\n", err)
		os.Exit(1)
	}

}

func FindOneBy(filter bson.M) model.User {

	result := model.User{}

    err := mdb.GetSession().DB("test").C("foo").Find(filter).One(&result)

    if err != nil {
        log.Fatal(err)
    }

    return result

}

func FindUsers(filter bson.M) []model.User {

	result := []model.User{}

    err := mdb.GetSession().DB("test").C("foo").Find(filter).All(&result)

    if err != nil {
        log.Fatal(err)
    }

    return result

}