package main
import (
	"net/http"
	"fmt"
	"io/ioutil"
	"encoding/json"
)

func main() {
	url := "http://localhost:8081"
	res, err := http.Get(url)
	if (err != nil) {
		panic(err)
	}
	body, err := ioutil.ReadAll(res.Body)
	if (err != nil) {
		panic(err)
	}
	var p Payload
	err = json.Unmarshal(body, &p)
	if (err != nil) {
		panic(err)
	}
	fmt.Println(p)
}




type Payload struct {
	Envelope CustomerInfo
}

type CustomerInfo struct {
	Name string
	Address Address
}

type Address struct {
	Line1 string
	Line2 string
}