# Making a webserver in go
This is a guided exercise on creating a simple `REST API` in `Golang`.
This project is divided in to two part, 
* **part1** Where we create a basic web server.
* **part2** Where We create a shopping cart list service `CRUD REST API`.
## Basic Web Server
Frist we need to create initialiaze a new go modules to hold all of our package.
```sh
go mod init go.mod
```
Then we create out go Program boilerPlate

```go
package main

func main (){
}
```

inside the `func main` function we initialize the http server by first importing the `"net/http"` package, and then calling the `http.ServeAndListen` function on port `8080`.
```go 
package main

import "net/http"

func main (){
    http.ListenAndServe(":8080", nil);
}
```
if we add our main.go 
