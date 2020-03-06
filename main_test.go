package main

import (
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAdd(t *testing.T){
	tot := Add(1,2)
	assert.NotNil(t, tot, "total should not be il")
	assert.Equal(t,3,tot,"expecting 4")
}
func TestSubstract(t *testing.T){
	tot := Substract(0,0)
	assert.NotNil(t, tot, "total must be not ni")
	assert.Equal(t,0, tot,"expexting0")
}

func TestInts(t *testing.T){
	tt := []struct{
		name string
		numbers []int
		sum int
	}{
		{"one to five",[]int{4,2},6},
		{"no numbers",nil,0},
		{"one and minus one",[]int{1,-1},0},
	}
	for _,tc := range  tt {
		s:= Ints(tc.numbers...)
		if s != tc.sum{
			t.Errorf("sum of %v should be %v; got %v",tc.numbers, tc.sum,s)
		}
	}
}
func Router()*mux.Router{
		router := mux.NewRouter()
		router.HandleFunc("/",RootEndpoint).Methods("GET")
		return router

}

func TestRootEndpoint(t *testing.T) {
	request,_ := http.NewRequest("GET","/",nil)
	response := httptest.NewRecorder()
	Router().ServeHTTP(response,request)
	assert.Equal(t,200, response.Code,"OK response is expected")
	assert.Equal(t, "youyou",response.Body.String(),"Incorrect body found")
}