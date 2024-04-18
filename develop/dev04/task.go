package main

import (
	"fmt"
	"slices"
	"sort"
	"strings"
)

type Set struct {
	data map[string][]string
}

func NewSet() *Set {
	data := make(map[string][]string)
	return &Set{data: data}
}

func (s *Set) CheckAnagram(key string, value string) bool {
	if len(value) != len(key) {
		return false
	}
	checkMap := make(map[rune]int, len(value))
	for _, elem := range value {
		checkMap[elem]++
	}
	for _, elem := range key {
		if _, found := checkMap[elem]; found {
			checkMap[elem]--
		} else {
			return false
		}
	}
	for _, elem := range value {
		if checkMap[elem] != 0 {
			return false
		}
	}
	return true
}

func (s *Set) Insert(key, value string) {
	i := sort.SearchStrings(s.data[key], value)
	s.data[key] = slices.Insert(s.data[key], i, value)
}

func (s *Set) FoundSet(value string) {
	value = strings.ToLower(value)
	for key := range s.data {
		if s.CheckAnagram(key, value) {
			s.Insert(key, value)
			return
		}
	}
	s.data[value] = []string{value}
}
func (s *Set) CheckLen() {
	for key := range s.data {
		if len(s.data[key]) == 1 {
			delete(s.data, key)
		}
	}
}

func Valid(arr []string) map[string]struct{} {
	m := make(map[string]struct{})
	for _, elem := range arr {
		m[elem] = struct{}{}
	}
	return m
}
func main() {
	arr := []string{"пятка", "пятак", "тяпка", "листок", "пятак", "слиток", "столик", "тортик"}
	m := Valid(arr)
	set := NewSet()
	for elem, _ := range m {
		set.FoundSet(elem)
	}
	set.CheckLen()
	fmt.Println(*set)
}
