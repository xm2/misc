/*
sort hashmap by value
Input: map[string]int
*/
package main

import (
	"fmt"
	"sort"
)

/*
V1:
use another hashmap and array to help sorting
*/
func sortV1(m map[string]int) {
	n := map[int]string{}
	t := []int{}

	for k, v := range m {
		n[v] = k
		t = append(t, v)
	}

	sort.Ints(t)

	r := "{"
	for i, k := range t {
		r += fmt.Sprintf("\"%s\": %d", n[k], k)
		if i != len(t)-1 {
			r += ", "
		}
	}

	r += "}"

	fmt.Println("V1 output:")
	fmt.Println(r)
}

type kvPair struct {
	key   string
	value int
}

/*
V2:
wrapper hashmap in another struct array, then use sort.Slice to help sorting
*/
func sortV2(m map[string]int) {
	n := []kvPair{}
	for k, v := range m {
		n = append(n, kvPair{key: k, value: v})
	}

	sort.Slice(n, func(i, j int) bool {
		return n[i].value < n[j].value
	})

	r := "{"
	for i, v := range n {
		r += fmt.Sprintf("\"%s\": %d", v.key, v.value)
		if i != len(n)-1 {
			r += ", "
		}
	}

	r += "}"

	fmt.Println("V2 output:")
	fmt.Println(r)
}

/*
sort hashmap by value
Input: map[string]int
*/
func main() {
	m := map[string]int{
		"Bob":  120,
		"Mike": 3,
		"Eric": 1,
	}

	sortV1(m)
	sortV2(m)
}
