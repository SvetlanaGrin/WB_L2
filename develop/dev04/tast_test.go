package main

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestValid(t *testing.T) {
	arr := []string{"пятка", "пятак", "тяпка", "листок", "пятак", "слиток", "столик", "тортик"}
	m := make(map[string]struct{})
	m = map[string]struct{}{"листок": struct{}{}, "пятак": struct{}{}, "пятка": struct{}{}, "слиток": struct{}{}, "столик": struct{}{}, "тортик": struct{}{}, "тяпка": struct{}{}}
	require.Equal(t, Valid(arr), m)
}
func TestSet_Insert(t *testing.T) {
	set := NewSet()
	set.Insert("пятка", "пятак")
	m := make(map[string][]string)
	m = map[string][]string{"пятка": []string{"пятак"}}
	require.Equal(t, set.data, m)
}
func TestSet_CheckAnagram(t *testing.T) {
	set := NewSet()
	require.Equal(t, set.CheckAnagram("пятка", "пятак"), true)
}
func TestSet_FoundSet(t *testing.T) {
	arr := []string{"пятка", "пятак", "тяпка", "листок", "пятак", "слиток", "столик", "тортик"}
	mapSet := Valid(arr)
	set := NewSet()
	for elem, _ := range mapSet {
		set.FoundSet(elem)
	}
	m := make(map[string][]string)
	m = map[string][]string(map[string][]string{"листок": []string{"листок", "слиток", "столик"}, "пятка": []string{"пятак", "пятка", "тяпка"}, "тортик": []string{"тортик"}})
	require.Equal(t, set.data, m)

}
func TestSet_CheckLen(t *testing.T) {
	arr := []string{"пятка", "пятак", "тяпка", "листок", "пятак", "слиток", "столик", "тортик"}
	mapSet := Valid(arr)
	set := NewSet()
	for elem, _ := range mapSet {
		set.FoundSet(elem)
	}
	set.CheckLen()
	m := make(map[string][]string)
	m = map[string][]string{"листок": []string{"листок", "слиток", "столик"}, "пятак": []string{"пятак", "пятка", "тяпка"}}
	require.Equal(t, set.data, m)
}
