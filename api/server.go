// The api directory contains the code of the REST API server.
// It serves the purspose od encapsulation, testability and reusability.
// As well as decoupling the server part form themain package.
----------------------------------------------------------------------------

package api 

import (
	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

// We create a struct that holds the list of our items.
type Item struct {
	ID uuid.UUID `json: "id"`;
	Name string `json: "name"`;
}

// Create another struct that will compose the gorilla mux router
type Server struct {
	*mux.Router;

	shoppingItems []Item;
}

func main (){
	f
}
