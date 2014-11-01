package cayleygo

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

var HAddress string = "http://localhost:64210"

type quad struct {
	Subject   string `json:"subject"`
	Predicate string `json:"predicate"`
	Object    string `json:"object"`
	Label     string `json:"label"`
}

type triads []quad

func Write(add string, q triads) {
	address := add + "/api/v1/write"
	t, err := json.Marshal(q)
	triad := string(t)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println(string(triad))
	var x bytes.Buffer
	x.Write([]byte(triad))
	resp, err := http.Post(address, "text/json", &x)
	if err != nil {
		fmt.Println(err)
	}
	a, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(a))
	resp.Body.Close()

}

func Delete(add string, q triads) {
	address := add + "/api/v1/delete"
	t, err := json.Marshal(q)
	triad := string(t)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println(triad)
	var x bytes.Buffer
	x.Write([]byte(triad))
	resp, err := http.Post(address, "text/json", &x)
	if err != nil {
		fmt.Println(err)
	}
	a, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(a))
	resp.Body.Close()

}

func Gremlin(add string, q string) []byte {
	address := add + "/api/v1/query/gremlin"
	var x bytes.Buffer
	x.Write([]byte(q))
	resp, err := http.Post(address, "text/plain", &x)
	if err != nil {
		fmt.Println(err)
	}
	a, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(a))
	resp.Body.Close()
	return a

}
