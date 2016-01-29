package main
import (
	"io/ioutil"
	"fmt"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

const basepath  string = "C:\\Users\\rajesh\\Downloads\\smaller"
var found = make(chan bool)

func main() {

	fileinfos, err := ioutil.ReadDir(basepath)
	check(err)
	fmt.Println(len(fileinfos))
	var i int = 0

	for i=0;i<len(fileinfos);i++ {
		fileName := basepath + "\\" + fileinfos[i].Name()
		go searchFile(fileName)
		fmt.Println(fileName , <-found)
	}


	fmt.Println("Done")
}

func searchFile(fileName string) {
	dat, err := ioutil.ReadFile(fileName)
	check(err)
	fileContents := string(dat)
	if strings.Contains(fileContents,"Dolokhov cut him short") {
		fmt.Println(fileName, " contains the search phrase" )
		found <- true
	}
}
