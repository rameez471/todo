package schema

import (
	"io/ioutil"
	"log"
	"encoding/json"
	"strconv"
)

type Item struct{
	X string
	Position int
	Done bool
}

func (i *Item) Pretty() string {
	return i.MarkDone()+" "+strconv.Itoa(i.Position) + ". " + i.X
}

func (i *Item) MarkDone() string {
	if i.Done {
		return "[X]"
	}
	return "[ ]"
}

func ReadItems(filename string) ([]Item, error) {
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
		return []Item{}, err
	}
	var items []Item
	err = json.Unmarshal(b, &items)
	if err != nil {
		log.Fatal(err)
		return []Item{}, err
	}
	for i,_ := range items{
		items[i].Position = i+1
	}
	return items, err
}

func WriteItems(filename string, items []Item) error {
	b, err := json.Marshal(items)
	err = ioutil.WriteFile(filename, b, 0644)
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}