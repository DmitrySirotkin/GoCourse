package main

import (
	"encoding/json"
	"io/ioutil"
)

func main() {
	addressBook := make(map[string][]int)

	addressBook["Alex"] = []int{89996543210}
	addressBook["Bob"] = []int{89167243812}
	addressBook["Bob"] = append(addressBook["Bob"], 89155243627)

	addressBook["Sergey"] = []int{89996545542}
	addressBook["Dmitry"] = []int{89996546901}
	addressBook["Dmitry"] = append(addressBook["Alex"], 89123149999)
	addressBook["Evgenia"] = []int{89000045542}

	message, _ := json.Marshal(addressBook)

	ioutil.WriteFile("NumberDict", message, 777)
}
