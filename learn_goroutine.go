package main
import (
	"time"
	"fmt"
)

func main() {
	a := Address{"line1", "line2"}
	c := CustomerInfo{"rajesh", a}
	payload := Payload{ c}
	go preparePayload(payload)
	time.Sleep(5000)
	fmt.Println("end of main")
}


func preparePayload( p Payload) {
	fmt.Println("in preparePayload")
	time.Sleep(1000)
	fmt.Println("preparePayload" ,p)
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