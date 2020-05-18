package main

import "fmt"

func main() {

}

// keep the value from optimizaion.
func keep(value interface{}) {
	switch v := value.(type) {
	case struct{}:
		fmt.Printf("%v", v)
	default:
		fmt.Printf("")
	}
}
