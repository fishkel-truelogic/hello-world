package rest

//////////////////////////////////////////////
// Imports
//////////////////////////////////////////////
import (
    "fmt"
    "net/http"
    "github.com/gorilla/mux"
    "hello-world/rest/user"
    "hello-world/rest/note"
    "hello-world/utils"
)

//////////////////////////////////////////////
// Types and Vars
//////////////////////////////////////////////
type Route struct {
    Name        string
    Method      string
    Pattern     string
    HandlerFunc http.HandlerFunc
}

type Routes []Route

var paths map[string]string

var routes Routes 

//////////////////////////////////////////////
// Functions
//////////////////////////////////////////////
func InitRouter() {
    paths = make(map[string]string)
    paths["routes"]="/"
    paths["get_users"]="/users"
    paths["get_user"]="/user/{id}"
    paths["post_user"]="/user"
    paths["put_user"]="/user/{id}"
    paths["delete_user"]="/user/{id}"
    paths["get_notes_pub"]="/notes"
    paths["get_notes_user"]="/notes/{userid}"
    paths["get_notes_user_pub"]="/notes/pub/{userid}"
    paths["post_note"]="/note"
    paths["put_note"]="/note/{id}"
    paths["delete_note"]="/note/{id}"


    routes = Routes{
        //////////////////////
        // Index routes
        //////////////////////
        Route{
            "routes",
            "GET",
            paths["routes"],
            GetAllRoutes,
        },
        //////////////////////
        // User routes
        //////////////////////
        Route{
            "get_users",
            "GET",
            paths["get_users"],
            restUser.GetAllUsers,
        },
        
        Route{
            "get_user",
            "GET",
            paths["get_user"],
            restUser.GetUserById,
        }, 

       Route{
            "post_user",
            "POST",
            paths["post_user"],
            restUser.CreateUser,
        },

        Route{
            "user",
            "PUT",
            paths["put_user"],
            restUser.UpdateUser,
        },

        Route{
            "delete_user",
            "DELETE",
            paths["delete_user"],
            restUser.DeleteUser,
        },

        Route{
            "get_notes_pub",
            "GET",
            paths["get_notes_pub"],
            restNote.GetAllPublicNotes,
        },

        Route{
            "get_notes_user",
            "GET",
            paths["get_notes_user"],
            restNote.GetAllNotesByUserId,
        },

        Route{
            "get_notes_user_pub",
            "GET",
            paths["get_notes_user_pub"],
            restNote.GetPublicNotesByUserId,
        },

        Route{
            "post_note",
            "POST",
            paths["post_note"],
            restNote.CreateNote,
        },

        Route{
            "put_note",
            "PUT",
            paths["put_note"],
            restNote.UpdateNote,
        },

        Route{
            "delete_note",
            "DELETE",
            paths["delete_note"],
            restNote.DeleteNote,
        },
    }
}

func NewRouter() *mux.Router {

    InitRouter()

    router := mux.NewRouter().StrictSlash(true)
    for _, route := range routes {
        router.
            Methods(route.Method).
            Path(route.Pattern).
            Name(route.Name).
            Handler(route.HandlerFunc)
    }

    return router
}

func GetAllRoutes(w http.ResponseWriter, r *http.Request) {
    res := ""
    for k, v := range paths {

        res = utils.Concat(res, k)
        res = utils.Concat(res, ": ")
        res = utils.Concat(res, v)
        res = utils.Concat(res, "\n")

    }
    fmt.Println(res)
    fmt.Fprintf(w, res)
}

