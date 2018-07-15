package store

import (
    "log"
    "net/http"
    "github.com/gorilla/mux"
)

var controller = &Controller{Repository: Repository{}}

// Route defines a route
type Route struct {
    Name        string
    Method      string
    Pattern     string
    HandlerFunc http.HandlerFunc
}

// Routes defines the list of routes of our API
type Routes []Route

var routes = Routes {

	Route {
        "/sign-up",
        "POST",
        "/sign-up",
        controller.Signup,
    },
	Route {
        "/sign-in",
        "GET",
        "/sign-in",
        controller.Signin,
    },
    Route {
        "/add-product",
        "POST",
        "/add-product",
        middleware(AuthenticationMiddleware(controller.AddProduct)),   // Added Rate Limit
    },
    Route {
        "/update-product",
        "PUT",
        "/update-product",
        AuthenticationMiddleware(controller.UpdateProduct),
    },
    // Get Product by {id}
    Route {
        "/products/{id}",
        "GET",
        "/products/{id}",
        middleware(AuthenticationMiddleware(controller.GetProduct)),   // Added Rate Limit
    },
    // Delete Product by {id}
    Route {
        "/delete-product/{id}",
        "DELETE",
        "/delete-product/{id}",
        AuthenticationMiddleware(controller.DeleteProduct),
    },
    // Search product with string
    Route {
        "/search/{query}",
        "GET",
        "/search/{query}",
        middleware(AuthenticationMiddleware(controller.SearchProduct)),  // Added Rate Limit
    }}

// NewRouter configures a new router to the API
func NewRouter() *mux.Router {
    
	limitRequest()
	
	router := mux.NewRouter().StrictSlash(true)
    for _, route := range routes { 
        var handler http.Handler
        log.Println(route.Name)
        handler = route.HandlerFunc
        
        router.
         Methods(route.Method).
         Path(route.Pattern).
         Name(route.Name).
         Handler(handler)
    }
    return router
}