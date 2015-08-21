package data

//////////////////////////////////////////////////////////////////////////////
// Imports
//////////////////////////////////////////////////////////////////////////////

import (
	"fmt"
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

	collection := mdb.GetSession().DB("notes").C("user")

	err := collection.Insert(user)

	if err != nil {
		fmt.Printf("Can't insert user: %v\n", err)
		os.Exit(1)
	}

}

func UpdateUser(user model.User) {
    collection := mdb.GetSession().DB("notes").C("user")

    err := collection.Update(bson.M{"_id": user.Id}, user)

    if err != nil {
        fmt.Printf("Can't update user: %v\n", err)
        os.Exit(1)
    }
}

func FindOneBy(filter bson.M) model.User {

	result := model.User{}

    err := mdb.GetSession().DB("notes").C("user").Find(filter).One(&result)

    if err != nil {
        log.Fatal(err)
    }

    return result

}

func FindUsers(filter bson.M) []model.User {

	result := []model.User{}

    err := mdb.GetSession().DB("notes").C("user").Find(filter).All(&result)

    if err != nil {
        log.Fatal(err)
    }

    return result

}

func RemoveUser(filter bson.M) error {
    err := mdb.GetSession().DB("notes").C("user").Remove(filter)
    if err != nil {
        log.Fatal(err)
        return err
    }
    return nil
}