package main
import "fmt"

func main() {
	var x int = 123;
	fmt.Println(x);
	var y *int;
	y = &x;
	*y = 234;
	fmt.Println(x)
}
