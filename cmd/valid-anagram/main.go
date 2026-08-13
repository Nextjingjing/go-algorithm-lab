package main

import (
	"fmt"

	hashmap "github.com/Nextjingjing/go-algorithm-lab/algorithms/hash-map"
)

func main() {
	s := "anagram"
	t := "nagaram"

	fmt.Println(hashmap.IsAnagram(s, t))
}
