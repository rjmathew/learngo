package main
import (
	"net/http"
	"encoding/json"
	"fmt"
)

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe("localhost:8081", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	response, err := publishJson()
	if (err != nil) {
		panic(err)
	}
	fmt.Fprintf(w, string(response))
}

func publishJson() ([]byte, error) {
	a := Address{"line1", "line2"}
	c := CustomerInfo{"rajesh", a}
	payload := Payload{ c}
	result, err := json.MarshalIndent(payload, "", "")
	if (err!=nil) {
		panic(err)
	}
	strval := string(result)
	fmt.Println("#",payload,"#")
	fmt.Println("*" ,strval, "*")
	return json.MarshalIndent(payload, "", "  ")
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