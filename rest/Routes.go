package rest

import (
    "net/http"
    "github.com/gorilla/mux"
    "hello-world/rest/user"
)

type Route struct {
    Name        string
    Method      string
    Pattern     string
    HandlerFunc http.HandlerFunc
}

type Routes []Route

func NewRouter() *mux.Router {

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

var routes = Routes{
    //////////////////////
    // User routes
    //////////////////////
    Route{
        "Users",
        "GET",
        "/users",
        rest.GetAllUsers,
    },
    
    Route{
        "user",
        "GET",
        "/user/{id}",
        rest.GetUserById,
    }, 

    Route{
        "user",
        "POST",
        "/user",
        rest.CreateUser,
    },
    
}