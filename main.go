package main

import (
	"fmt"
	"log"
	"net/http"
)

func homehandle(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Welcome to the HomePage!")
}

func posthandle(w http.ResponseWriter, r *http.Request){

	if r.Method == http.MethodPost{
		fmt.Fprintf(w,"Post request received ")
	}else{
		http.Error(w,"Method not allowed",http.StatusMethodNotAllowed)
	}    
}


func handlerequest() {

	http.HandleFunc("/",homehandle)
	http.HandleFunc("/post",posthandle)

	fmt.Println("server is running on port 8000")
	log.Fatal(http.ListenAndServe(":8000",nil))

}

func main() {

	handlerequest()
}