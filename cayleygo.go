//Copyright 2014 Anthony Alaribe. All rights researved.
//Use this source code as governed by the MIT license
//that can be found in the LICENSE file
/*
  This cayleygo package makes accessing your cayley graph data accessible over the REST API.
  With this package, you can write triples to the database, update data already in the database, and delete data from the database.


*/
package cayleygo

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

var HAddress string = "http://localhost:64210"

type Quad struct {
	Subject   string `json:"subject"`
	Predicate string `json:"predicate"`
	Object    string `json:"object"`
	Label     string `json:"label"`
}

type Triads []Quad

func StringList(preds []string) string {
	var items string
	items = "["
	for i, _ := range preds {
		items = items + `"` + preds[i] + `",`
	}
	items = items + "]"

	return items
}

func Write(add string, q Triads) error {
	address := add + "/api/v1/write"
	t, err := json.Marshal(q)
	triad := string(t)
	if err != nil {
		return err
	}
	fmt.Println(string(triad))
	var x bytes.Buffer
	x.Write([]byte(triad))
	resp, err := http.Post(address, "text/json", &x)
	if err != nil {
		return err
	}
	_, readerr := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if readerr != nil {
		return readerr
	}
	return nil

}

func Delete(add string, q Triads) error {
	address := add + "/api/v1/delete"
	t, err := json.Marshal(q)
	triad := string(t)
	if err != nil {
		return err
	}
	var x bytes.Buffer
	x.Write([]byte(triad))
	resp, err := http.Post(address, "text/json", &x)
	if err != nil {
		return err
	}
	_, readerr := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if readerr != nil {
		return readerr
	}
	return nil

}

func Gremlin(add string, q string) ([]byte, error) {
	address := add + "/api/v1/query/gremlin"
	var x bytes.Buffer
	x.Write([]byte(q))
	resp, err := http.Post(address, "text/plain", &x)
	if err != nil {
		return nil, err
	}
	data, readerr := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	return data, readerr

}
