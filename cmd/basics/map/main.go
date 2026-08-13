package main

import "fmt"

func main() {
	ages := make(map[string]int)
	ages["ann"] = 20
	ages["bob"] = 12
	fmt.Println(ages)
	fmt.Println(ages["ann"])
	fmt.Println(ages["no"])
	fmt.Println(ages)
}
