package main

import (
	"encoding/json"
	"fmt"
)

// Item types with at least standard and fragile variants, with a volume property.

type Item struct {
	ID        string `json:"id"`
	IsFragile bool   `json:"isFragile"`
	Volume    int    `Volume:"volume"`
}

// Container types with capacity constraints, possibly locking and fragile-safe behavior.

type Container struct {
	CapacityVolume int              `json:"capacityVolume"`
	IsLocked       bool             `json:"isLocked"`
	FragileSafe    bool             `json:"fragileSafe"`
	Items          map[string]*Item `json:"items"`
}

// Methods to add, remove, and move items between containers, enforcing rules (capacity, locking, grouping, fragility).

func (C *Container) add(I *Item) bool {
	if C.IsLocked {
		fmt.Println("Container is Locked Can not Add")
		return false
	}

	if C.FragileSafe && !I.IsFragile {
		fmt.Println("Container is for Fragile objects can not add non Fragile objects in here")
		return false
	}

	if _, ok := C.Items[I.ID]; ok == true {
		fmt.Println("Item already in Container")
		return false
	}

	if C.CapacityVolume >= I.Volume {
		C.Items[I.ID] = I
		C.CapacityVolume -= I.Volume
		fmt.Println("Item addded")
		return true
	}

	fmt.Println("Not enough space for the item")
	return false
}

func (C *Container) remove(I *Item) bool {
	if C.IsLocked {
		fmt.Println("Container is Locked Can not Add")
		return false
	}

	_, ok := C.Items[I.ID]
	if !ok {
		fmt.Println("Item not in Container")
		return false
	} else {
		delete(C.Items, I.ID)
		C.CapacityVolume += I.Volume
		fmt.Println("Deleted I")
		return true
	}
}

func (C *Container) move(I *Item, otherContainer *Container) bool {
	if C.FragileSafe && !otherContainer.FragileSafe {
		fmt.Println("Can not Move item from a fragile safe to a non fragile safe")
		return false
	}

	_, ok := C.Items[I.ID]
	if !ok {
		fmt.Println("Item is not in Container")
		return false
	}

	C.remove(I)
	if !otherContainer.add(I) {
		C.add(I)
		return false
	}
	return true
}

func (I *Item) ToJSON() (string, error) {
	js, err := json.MarshalIndent(I, "", "	")
	if err != nil {
		return "", err
	}
	return string(js), err
}

func (C *Container) ToJSON() (string, error) {
	js, err := json.MarshalIndent(C, "", "	")
	if err != nil {
		return "", err
	}
	return string(js), err
}

func ItemFromJSON(data string) (*Item, error) {
	var I Item
	err := json.Unmarshal([]byte(data), &I)
	if err != nil {
		return nil, err
	}
	return &I, err
}

func ContainerFromJSON(data string) (*Container, error) {
	var C Container
	err := json.Unmarshal([]byte(data), &C)
	if err != nil {
		return nil, err
	}
	return &C, err
}

// func main() {
// 	i1 := Item{ID: "1", Volume: 20, IsFragile: false}
// 	i2 := Item{ID: "2", Volume: 20, IsFragile: false}
// 	i3 := Item{ID: "3", Volume: 30, IsFragile: true}

// 	c1 := Container{CapacityVolume: 30, FragileSafe: true, Items: map[string]*Item{}}
// 	c2 := Container{CapacityVolume: 30, FragileSafe: false, Items: map[string]*Item{}}
// 	c3 := Container{CapacityVolume: 40, FragileSafe: false, Items: map[string]*Item{}}

// 	c1.add(&i1)
// 	c1.add(&i2)
// 	c1.add(&i3)

// 	c2.add(&i1)
// 	c2.add(&i2)
// 	c2.add(&i3)

// 	c2.remove(&i1)
// 	c2.remove(&i1)
// 	c2.add(&i1)

// 	c2.move(&i1, &c3)
// 	fmt.Println(c3.CapacityVolume)

// 	c2JSON, _ := c1.ToJSON()
// 	fmt.Println(c2JSON)
// }
