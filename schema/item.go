package schema

import (
	"io/ioutil"
	"log"
	"encoding/json"
	"strconv"
	"sort"
)

type Item struct{
	X string
	Position int
	Done bool
	Priority int
}

func (i *Item) Pretty() string {
	return i.MarkDone()+" "+strconv.Itoa(i.Position) + ". " + i.X + " ("+i.PrettyPriority()+") "
}

func (i *Item) PrettyPriority() string{
	switch i.Priority {
	case 1:
		return "High"
	case 2:
		return "Medium"
	case 3:
		return "Low"
	default:
		return ""
	}
}

func DeleteItem(items []Item, i int) []Item {
	items[i] = items[len(items)-1]
	items[len(items)-1] = Item{}
	items = items[:len(items)-1]
	return items
}

func (i *Item) MarkDone() string {
	if i.Done {
		return "[X]"
	}
	return "[ ]"
}

func (i *Item) SetPriority(pri int) {
	switch pri {
	case 1:
		i.Priority = 1
	case 3:
		i.Priority = 3
	default:
		i.Priority = 2
	}
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
	sort.Slice(items[:], func(i, j int) bool {
		if items[i].Priority == items[j].Priority {
			return items[i].Position < items[j].Position
		}
		return items[i].Priority < items[j].Priority
	})
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