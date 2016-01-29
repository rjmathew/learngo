package main
import (
	"io/ioutil"
	"fmt"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	dat, err := ioutil.ReadFile("C:\\Users\\rajesh\\Downloads\\big.txt")
	check(err)
	fmt.Println(len(dat))
	//fmt.Print(string(dat))
	var start, end int = 0,0

	threshhold := 1000000
	counter := 0
	for counter=0;counter <= len(dat)/threshhold; counter++ {
		start = counter * threshhold
		end = ((counter+1) * threshhold) - 1
		if (end >= len(dat)) {
			end = len(dat) - 1
		}
		subarray := dat[start:end]
		fmt.Println(counter, start, end ," - " ,len(subarray))
		err := ioutil.WriteFile("C:\\Users\\rajesh\\Downloads\\smaller\\small" + strconv.Itoa(counter) + ".txt" ,subarray, 0644)
		check(err)
	}

	fmt.Println(threshhold)
}
