package restNote

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
    "hello-world/data/note"
    "hello-world/model"
)

//////////////////////////////////////////////////////////////////////////////
// Rest API Functions
//////////////////////////////////////////////////////////////////////////////

func GetPublicNotesByUserId(w http.ResponseWriter, r *http.Request) {
	fmt.Println("------------------ GetPublicNotesByUserId ------------------")
	
	vars := mux.Vars(r)
	if bson.IsObjectIdHex(vars["userid"]) {

	    id := bson.ObjectIdHex(vars["userid"])

		result := []model.Note{}

		result = data.FindNotes(bson.M{"_user_id" : id, "_pri" : false})

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

func GetAllNotesByUserId(w http.ResponseWriter, r *http.Request) {
	fmt.Println("------------------ GetAllNotesByUserId ------------------")
	
	vars := mux.Vars(r)
	if bson.IsObjectIdHex(vars["userid"]) {

	    id := bson.ObjectIdHex(vars["userid"])

		result := []model.Note{}

		result = data.FindNotes(bson.M{"_user_id" : id})

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

func GetAllPublicNotes(w http.ResponseWriter, r *http.Request) {
	fmt.Println("------------------ GetAllPublicNotes ------------------")

	result := []model.Note{}

	filter := bson.M{"_pri" : "false"}

	result = data.FindNotes(filter)

    j, err := json.Marshal(result)
    
    if err != nil {
    	log.Fatal(err)
    	w.WriteHeader(http.StatusInternalServerError)
    	return
	}

	fmt.Fprintf(w, string(j))
}

func CreateNote(w http.ResponseWriter, r *http.Request) {
	fmt.Println("------------------ CreateNote ------------------")
	
	note := model.Note{}

	fmt.Println(r);
	
	decoder := json.NewDecoder(r.Body)
     
    err := decoder.Decode(&note)

    if err != nil {
        log.Fatal(err)
        w.WriteHeader(http.StatusInternalServerError)
    	return
    }

	note.Id = bson.NewObjectId()

	err = data.InsertNote(note)
	
	if err != nil {
    	w.WriteHeader(http.StatusInternalServerError)
    	fmt.Fprintf(w, err.Error());
    	return
	}

    j, err := json.Marshal(note)
    
    if err != nil {
    	log.Fatal(err)
    	w.WriteHeader(http.StatusInternalServerError)
    	return
	}
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, string(j))
}

func UpdateNote(w http.ResponseWriter, r *http.Request) {
	fmt.Println("------------------ UpdateNote ------------------")
	

	vars := mux.Vars(r)
	if bson.IsObjectIdHex(vars["id"]) {

	    id := bson.ObjectIdHex(vars["id"])
		note := model.Note{}

		decoder := json.NewDecoder(r.Body)
	     
	    err := decoder.Decode(&note)
	    if err != nil {
	        log.Fatal(err)
	        w.WriteHeader(http.StatusInternalServerError)
    		return
	    }
		note.Id = id

		data.UpdateNote(note)

	    j, err := json.Marshal(note)
	    
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

func DeleteNote(w http.ResponseWriter, r *http.Request) {
	fmt.Println("------------------ DeleteNote ------------------")
	
	vars := mux.Vars(r)
	if bson.IsObjectIdHex(vars["id"]) {

	    id := bson.ObjectIdHex(vars["id"])

		err := data.RemoveNote(bson.M{"_id" : id})

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
/*
func noteQueryParams(r *http.Request) bson.M {
	
	filter := bson.M{}
	r.ParseForm()
	fmt.Println(r)
	
	userId := r.Form.Get("userid")
	if userId != "" {
		filter["_user_id"] = userId
		fmt.Println(userId)
	}
	
	private := r.Form.Get("private")
	if private != "" {
		filter["_private"] = private == "true"
		fmt.Println(private)
	}
	
	return filter
}
*/


