package main

import (
	. "fmt"
	"time"
)

type Entry struct {
	key   string
	milis int64
}

type EntryList []Entry

func (e EntryList) Swap(i, j int) {
	e[i], e[j] = e[j], e[i]
}

func (e EntryList) Len() int {
	return len(e)
}

func (e EntryList) Less(i, j int) bool {
	return e[i].milis < e[j].milis
}

func sortMapByValue(m map[string]int64) EntryList {
	e := make(EntryList, len(m))
	i := 0
	for k, v := range m {
		e[i] = Entry{k, v}
		i++
	}

	// sort.Sort(e)
	sorted := true

	for sorted {
		sorted = false
		for i := 0; i < e.Len()-1; i++ {

			if e.Less(i+1, i) {
				sorted = true
				e.Swap(i, i+1)

			}
		}
	}

	return e
}

func main() {

	m := make(map[string]int64)

	m["a"] = time.Now().UTC().UnixNano()
	// time.Sleep(1 * time.Second)
	m["c"] = time.Now().UTC().UnixNano()
	// time.Sleep(1 * time.Second)
	m["b"] = time.Now().UTC().UnixNano()

	Println(sortMapByValue(m))
	Println(m["y"])
}
