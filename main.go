package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

func Add(a int, b int)int{
	return (a+b)
}

func Substract(a int , b int) int {
	return a - b
}

func RootEndpoint(w http.ResponseWriter, r *http.Request)  {
	w.WriteHeader(200)
	w.Write([]byte("youyou"))

}
func Ints(vs ...int)int  {
	return ints(vs)
}
func ints(vs[]int)int{
	if len(vs) == 0 {
		return 0
	}
	return ints(vs[1:]) + vs[0]
}
func doubleHandler(w http.ResponseWriter, r *http.Request){
	text := r.FormValue("v")
	if text == "" {
		http.Error(w,"missing value", http.StatusBadRequest)
		return
	}
	// Atoi (string to int)
	v,err := strconv.Atoi(text)
	if err != nil {
		http.Error(w,"not a number :" + text, http.StatusBadRequest)
		return
	}
	fmt.Fprintln(w,v*2)

}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/",RootEndpoint).Methods("GET")
	router.HandleFunc("/double",doubleHandler)
	log.Fatal(http.ListenAndServe(":555",router))

}
