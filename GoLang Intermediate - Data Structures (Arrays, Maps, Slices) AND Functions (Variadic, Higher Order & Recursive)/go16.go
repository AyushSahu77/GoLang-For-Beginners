package main

import (
	"fmt"
)

func main() {
	//var map1 map[string]int
	//Syntax : map_name map[key datatype] value datatype

	// an empty map has nil values, as they are default map values
	// a nil map can't be assigned directly, to add values to a map it must be initialized at the time of creation

	map1 := map[string]int{"a": 1, "b": 2, "c": 3}
	fmt.Println(map1)
	fmt.Println(len(map1))
	fmt.Println(map1["b"]) //prints value associated with the key
	fmt.Print("\n")

	// map[] returns 2 things, "value" and "found"
	value1, found1 := map1["a"]
	fmt.Println(value1, found1)
	value2, found2 := map1["dd"]
	fmt.Println(value2, found2) // will return value 0 as it is default for int value
	fmt.Print("\n")

	map2 := map[string]string{"ddd": "444", "eee": "555", "fff": "666"}
	value3, found3 := map2["eee"]
	fmt.Println(value3, found3)
	value4, found4 := map2["a"]
	fmt.Println(value4, found4) // will return value "" empty string as it is default for string value
	fmt.Print("\n")

	map3 := map[string]string{"en": "english", "fr": "french", "gr": "german"}
	fmt.Println(map3)
	map3["hi"] = "hindi" // adding new value
	fmt.Println(map3)
	map3["en"] = "spanish" // updating existing value
	fmt.Println(map3)
	delete(map3, "en") //deleting a key-value pair, built in delete func
	fmt.Println(map3)
	fmt.Print("\n")

	// Looping over a map with range return 2 things, key and value
	map4 := map[string]string{"es": "spanish", "ca": "canadian", "eu": "british english"}
	for key, value := range map4 {
		fmt.Println(key, "==>", value)
	}
	map4 = make(map[string]string) // truncating a map i.e. clearing all key-value pairs
	fmt.Println(map4)
}
