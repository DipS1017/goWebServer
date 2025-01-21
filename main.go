package main

import (
	"fmt"
	"net/http"
  "github.com/DipS1017/goWebServer/internal/db"
)



func main() {
  db.InitDB()
  mux:=http.NewServeMux()
  mux.HandleFunc("/",handleRoot)
  fmt.Println("Server is running on port 8080")
  http.ListenAndServe(":8080",mux) 

}
func handleRoot(w http.ResponseWriter, r *http.Request){
  fmt.Fprintf(w,"Hello world")

}
