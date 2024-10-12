package main

import "fmt"

var a = 10

func main() {
	z := 20
	qualquerCoisa(z)

}

// chaves define o scope
func qualquerCoisa(x int) {
	fmt.Println(a)
	fmt.Println(x)
}
