package restUser

//////////////////////////////////////////////////////////////////////////////
// Imports
//////////////////////////////////////////////////////////////////////////////

import (
	"fmt"
	"gopkg.in/mgo.v2/bson"
    "log"
    "net/http"
    "encoding/json"
    "github.com/gorilla/mux"
    "hello-world/data/user"
    "hello-world/model"
    "hello-world/utils"
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
		w.WriteHeader(http.StatusBadRequest)
		return
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
    	w.WriteHeader(http.StatusInternalServerError)
    	return
	}

	fmt.Fprintf(w, string(j))
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("------------------ CreateUser ------------------")
	fmt.Println(r)
	fmt.Println("-----------------------------")
	fmt.Println(r.Body)
	user := model.User{}

	decoder := json.NewDecoder(r.Body)
     
    err := decoder.Decode(&user)
    if err != nil {
        log.Fatal(err)
        w.WriteHeader(http.StatusInternalServerError)
    	return
    }

	user.Id = bson.NewObjectId()

	data.InsertUser(user)

    j, err := json.Marshal(user)
    
    if err != nil {
    	log.Fatal(err)
    	w.WriteHeader(http.StatusInternalServerError)
    	return
	}
	w.WriteHeader(http.StatusCreated)
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
	        w.WriteHeader(http.StatusInternalServerError)
    		return
	    }

		user.Id = id

		data.UpdateUser(user)

	    j, err := json.Marshal(user)
	    
	    if err != nil {
	    	log.Fatal(err)
	    	w.WriteHeader(http.StatusInternalServerError)
    	    return
		}

		fmt.Fprintf(w, string(j))

	} else {
		w.WriteHeader(http.StatusBadRequest)
    	return
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
	    	w.WriteHeader(http.StatusInternalServerError)
    		return
		} else {
			fmt.Fprintf(w, "deleted")
		}

	} else {
		w.WriteHeader(http.StatusBadRequest)
    	return
	}
}


//////////////////////////////////////////////////////////////////////////////
// Util Functions
//////////////////////////////////////////////////////////////////////////////

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
		filter["_age"] = utils.StrToInt(age)
		fmt.Println(age)
	}
	
	lastname := r.Form.Get("lastname")
	if lastname != "" {
		filter["_last_name"] = lastname
		fmt.Println(lastname)
	}

	return filter
}


