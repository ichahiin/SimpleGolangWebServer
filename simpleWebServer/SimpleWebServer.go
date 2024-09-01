package main 
import (
	"net/http"
)

func main () {
	// Creating a new Route handler:
	// The first part is the route, while the second argument is the handlerfunction that has the signature described bellow
	http.HandleFunc("/Hello-world", func (w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World!")); // we cast our string to a byte array 
	});
	// We initialiaze the http server:
	http.ListenAndServe(":8080", nil)
}
