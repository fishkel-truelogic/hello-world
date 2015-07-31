package rest

//////////////////////////////////////////////////////////////////////////////
// Imports
//////////////////////////////////////////////////////////////////////////////

import (
	"fmt"
	"gopkg.in/mgo.v2/bson"
    "os"
    "log"
    "net/http"
    "encoding/json"
    "strconv"
    "github.com/gorilla/mux"
    "hello-world/data/user"
    "hello-world/model"
)

//////////////////////////////////////////////////////////////////////////////
// Rest API Functions
//////////////////////////////////////////////////////////////////////////////

func GetUserById(w http.ResponseWriter, r *http.Request) {
	fmt.Println("------------------ GetUserById ------------------")
	
	vars := mux.Vars(r)
	if bson.IsObjectIdHex(vars["id"]) {

	    id := bson.ObjectIdHex(vars["id"])

		result := model.User{}

		result = data.FindOneBy(bson.M{"_id" : id})

	    j, err := json.Marshal(result)

	    if err != nil {
	    	log.Fatal(err)
		}

		fmt.Fprintf(w, string(j))
	} else {
		//TODO: manejar los distintos status
		fmt.Fprintf(w, "Bad Request")
	}

}

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Println("------------------ GetAllUsers ------------------")
	
	filter := userQueryParams(r)

	result := []model.User{}

	result = data.FindUsers(filter)

    j, err := json.Marshal(result)
    
    if err != nil {
    	log.Fatal(err)
	}

	fmt.Fprintf(w, string(j))
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("------------------ CreateUser ------------------")
	
	user := model.User{}

	decoder := json.NewDecoder(r.Body)
     
    err := decoder.Decode(&user)
    if err != nil {
        log.Fatal(err)
    }

	user.Id = bson.NewObjectId()

	data.InsertUser(user)

    j, err := json.Marshal(user)
    
    if err != nil {
    	log.Fatal(err)
	}

	fmt.Fprintf(w, string(j))
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("------------------ UpdateUser ------------------")
	

	vars := mux.Vars(r)
	if bson.IsObjectIdHex(vars["id"]) {

	    id := bson.ObjectIdHex(vars["id"])
		user := model.User{}

		decoder := json.NewDecoder(r.Body)
	     
	    err := decoder.Decode(&user)
	    if err != nil {
	        log.Fatal(err)
	    }

		user.Id = id

		data.UpdateUser(user)

	    j, err := json.Marshal(user)
	    
	    if err != nil {
	    	log.Fatal(err)
		}

		fmt.Fprintf(w, string(j))

	} else {
		fmt.Fprintf(w, "Bad request")
	}
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("------------------ DeleteUser ------------------")
	
	vars := mux.Vars(r)
	if bson.IsObjectIdHex(vars["id"]) {

	    id := bson.ObjectIdHex(vars["id"])

		err := data.RemoveUser(bson.M{"_id" : id})

	    if err != nil {
	    	log.Fatal(err)
	    	fmt.Fprintf(w, "error")
		} else {
			fmt.Fprintf(w, "deleted")
		}

	} else {
		fmt.Fprintf(w, "Bad request")
	}
}


//////////////////////////////////////////////////////////////////////////////
// Util Functions
//////////////////////////////////////////////////////////////////////////////

func strToInt(s string) int {
	i, err := strconv.Atoi(s)
    if err != nil {
        // handle error
        fmt.Println(err)
        os.Exit(2)
    }
    return i
}

func userQueryParams(r *http.Request) bson.M {
	
	filter := bson.M{}
	r.ParseForm()
	fmt.Println(r)
	
	firstname := r.Form.Get("firstname")
	if firstname != "" {
		filter["_first_name"] = firstname
		fmt.Println(firstname)
	}
	
	age := r.Form.Get("age")
	if age != "" {
		filter["_age"] = strToInt(age)
		fmt.Println(age)
	}
	
	lastname := r.Form.Get("lastname")
	if lastname != "" {
		filter["_last_name"] = lastname
		fmt.Println(lastname)
	}

	return filter
}


