package main

import (
	"bytes"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strconv"
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

func TestDoubleHandler(t *testing.T){
	tt :=[] struct{
		name string
		value string
		double int
		err string
	}{
		{name:"double of two",value:"2",double:4},
		{name:"missing value",value:"",err:"missing value"},
		{name:"not a number",value:"x",err:"not a number :x"},
	}

	for _,tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			req, err := http.NewRequest("GET","localhost:555/double?v="+tc.value, nil)
			if err != nil {
				t.Fatalf("could not created request: %v", err)
			}
			rec := httptest.NewRecorder()
			doubleHandler(rec,req)

			res := rec.Result()
			defer res.Body.Close()

			b, err := ioutil.ReadAll(res.Body)
			if err != nil {
				t.Fatalf("could not read response: %v",err )
			}
			if tc.err != ""{
				if res.StatusCode != http.StatusBadRequest{
					t.Errorf("expected status Bad Request; got %v", res.StatusCode)
				}
				if msg := string(bytes.TrimSpace(b)); msg != tc.err{
					t.Errorf("expected message %q; got %q",tc.err,msg)
				}
				return
			}

			d, err := strconv.Atoi(string(bytes.TrimSpace(b)))
			if err != nil {
				t.Fatalf("expected an integr ; got %s",b)
			}
			if d!= tc.double {
				t.Fatalf("expected double to be %v; got %v",tc.double,d)
			}
		})

	}



}

func TestRouting(t *testing.T){
	srv := httptest.NewServer(handler())
	defer srv.Close()

	res, err := http.Get(fmt.Sprintf("%s/double?v=2", srv.URL))
	if err != nil {
		t.Fatalf("could not send GET request: %v",err)
	}

}