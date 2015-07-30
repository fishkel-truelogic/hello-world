package main


import (
	"fmt"
	"gopkg.in/mgo.v2"
    "gopkg.in/mgo.v2/bson"
    "os"
    "log"
)


type User struct {
	Id    bson.ObjectId `bson:"_id"`
	FirstName string    `bson:"_first_name"`
	LastName string     `bson:"_last_name"`
	Age int             `bson:"_age"`
	Details []string    `bson:"_details"` 
}

func InsertUser(user User, sess *mgo.Session) {

	collection := sess.DB("test").C("foo")

	err := collection.Insert(user)

	if err != nil {
		fmt.Printf("Can't insert user: %v\n", err)
		os.Exit(1)
	}

}

func FindUserBy(filter bson.M, sess *mgo.Session) User {

	result := User{}

    err := sess.DB("test").C("foo").Find(filter).One(&result)

    if err != nil {
        log.Fatal(err)
    }

    return result

}

func ConnectToMongoDB() *mgo.Session {
	uri := os.Getenv("MONGOHQ_URL")

	if uri == "" {
		fmt.Println("no connection string provided")
		os.Exit(1)
	}

	sess, err := mgo.Dial(uri)

	if err != nil {
		fmt.Printf("Can't connect to mongo, go error %v\n", err)
		os.Exit(1)
	}

	return sess
}

func main() {
		
	sess := ConnectToMongoDB()
	defer sess.Close()
	sess.SetSafe(&mgo.Safe{})
	

	user := &User{}
	user.Id = bson.NewObjectId()
	user.FirstName = "Leonel"
	user.LastName = "Fishkel"
	user.Age = 22
	user.Details = []string{"Hola", "Chau"}

	//InsertUser(user, sess)

    result := User{}

    result = FindUserBy(bson.M{"_first_name" : "Leonel"}, sess)

    fmt.Println("User age:", result.Age)

}