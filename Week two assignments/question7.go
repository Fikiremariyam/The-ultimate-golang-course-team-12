/*
Create a MultiKeyMap type that can store values based on multiple
keys (a combination of two or more strings). Implement functions to
Set, Get, and Delete values from this map. For example,
Set("first", "last", 25) should store a value 25 that can be
retrieved with Get("first", "last").
*/
package main

import (
	"fmt"
)

type Multikey struct {
	part1 string
	part2 string
}

type MultiKeyMap struct {
	count map[Multikey]interface{}
}

func NewMultiKeyMap() *MultiKeyMap {
	return &MultiKeyMap{
		count: make(map[Multikey]interface{}),
	}
}

func (p *MultiKeyMap) setvalue(key1 string, key2 string, value string) {
	p.count[Multikey{key1, key2}] = value

}
func (p *MultiKeyMap) get(key1, key2 string) (interface{}, bool) {

	value, exists := p.count[Multikey{key1, key2}]
	return value, exists
}
func (p *MultiKeyMap) Delete(key1, key2 string) {
	delete(p.count, Multikey{key1, key2})
}

func main() {

	object := NewMultiKeyMap()

	object.setvalue("first", "last", "25")
	object.setvalue("foo", "bar", "hello world")

	if value, found := object.get("first", "last"); found {
		fmt.Println("value for (first, last) is ", value)
	} else {
		fmt.Println("not found")
	}

	object.Delete("first", "last")
	if _, found := object.get("first", "last"); !found {
		fmt.Println("success fully deleted")
	}

}
