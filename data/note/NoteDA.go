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

func InsertNote(note model.Note) error {

	collection := mdb.GetSession().DB("notes").C("note")

	err := collection.Insert(note)

	if err != nil {
		fmt.Printf("Can't insert note: %v\n", err)
		return err
	}
    return nil

}

func UpdateNote(note model.Note) {
    collection := mdb.GetSession().DB("notes").C("note")

    err := collection.Update(bson.M{"_id": note.Id}, note)

    if err != nil {
        fmt.Printf("Can't update note: %v\n", err)
        os.Exit(1)
    }
}


func FindNotes(filter bson.M) []model.Note {

	result := []model.Note{}

    err := mdb.GetSession().DB("notes").C("note").Find(filter).All(&result)

    if err != nil {
        log.Fatal(err)
    }

    return result

}

func RemoveNote(filter bson.M) error {
    err := mdb.GetSession().DB("notes").C("note").Remove(filter)
    if err != nil {
        log.Fatal(err)
        return err
    }
    return nil
}