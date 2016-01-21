package main
import "fmt"

func main() {
	aCar := newCar()
	bCar := car{}
	aTruck := truck{}
	fmt.Println(aCar)
	aCar.mpg = 4
	fmt.Println(aCar)
	vinProviders :=[...]vinprovider {bCar, aTruck}
	for k,v := range vinProviders {
		fmt.Println(k, v, v.getVin())
	}
}

	type vehicle struct {
		mpg int
		numDoors int
	}

type car struct {
	vehicle

}

func (t truck) getVin() int {
return 2
}


func (c car) getVin() int {
    return 1
}

type vinprovider interface {
	getVin() int
}

type truck struct {
	vehicle
}

func newCar() *car {
	result := car{}
	result.mpg = 5
	return &result
}

func (v *vehicle) getMpg() {
	fmt.Println("mpg is ", &v.mpg)
}
