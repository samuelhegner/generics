package main

import "fmt"

func main() {
	map1 := map[int]string{1: "sam", 2: "jeff", 3: "jess"}
	map2 := map[string]int{"sam": 25, "jeff": 45, "jess": 80}

	k1 := MapKeys(map1)
	k2 := MapKeys(map2)

	fmt.Println(k1)
	fmt.Println(k2)
}

// A generic function that returns a slice of the maps keys
func MapKeys[K comparable, V any](m map[K]V) []K {

	// Make an empty slice of the Key type and set its capacity to the length of the map
	keys := make([]K, 0, len(m))

	// Loop over the map and append every key to our keys slice
	for k := range m {
		keys = append(keys, k)
	}

	// return the slice
	return keys
}
