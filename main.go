package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
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
func main() {
	router := mux.NewRouter()
	router.HandleFunc("/",RootEndpoint).Methods("GET")
	log.Fatal(http.ListenAndServe(":555",router))

}
