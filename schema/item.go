package schema

import (
	"io/ioutil"
	"log"
	"encoding/json"
)

type Item struct{
	X string
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