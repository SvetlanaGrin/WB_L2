package main

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestGrep_ReadFile(t *testing.T) {
	grep := newGrep()
	grep.ReadFile("test.txt")
	dataFile := make(map[string]string)
	dataFile["test.txt"] = "кошка cat имеет лапки, поэтому плохо пишет код|КОРОВА любит кодить, ее не берут на работу, потому что она не человек|бык|бык|БЫК|корова|волк|ВОЛК|"
	require.Equal(t, grep.dataFile, dataFile)
}
func TestGrep_FoundStr(t *testing.T) {
	grep := newGrep()
	grep.ReadFile("test.txt")
	grep.GetCheekString("код")
	grep.FoundStr()
	data := make(map[string][]string)
	data["test.txt"] = []string{"код", "кодить, ее не бер"}
	require.Equal(t, grep.data, data)
}
func TestGrep_FoundStrFlagA(t *testing.T) {
	grep := newGrep()
	grep.ReadFile("test.txt")
	grep.GetCheekString("код")
	grep.FoundStrFlagA(2)
	data := make(map[string][]string)
	data["test.txt"] = []string{"код", "кодить, ее "}
	require.Equal(t, grep.data, data)
}
func TestGrep_FoundStrFlagB(t *testing.T) {
	grep := newGrep()
	grep.ReadFile("test.txt")
	grep.GetCheekString("код")
	grep.FoundStrFlagB(2)
	data := make(map[string][]string)
	data["test.txt"] = []string{"пишет код ", "любит код "}
	require.Equal(t, grep.data, data)
}
func TestGrep_FoundStrFlagC(t *testing.T) {
	grep := newGrep()
	grep.ReadFile("test.txt")
	grep.ReadFile("file.txt")
	grep.GetCheekString("ко")
	grep.FoundStrFlagC(2)
	data := make(map[string][]string)
	data["file.txt"] = []string{"кошка cat", "корова", "котлетка с ", "с пюрешкой"}
	data["test.txt"] = []string{"кошка cat ", "пишет код", "любит кодить, ее ", "корова"}
	require.Equal(t, grep.data, data)
}

func TestGrep_CountStr(t *testing.T) {
	grep := newGrep()
	grep.ReadFile("test.txt")
	grep.ReadFile("file.txt")
	grep.GetCheekString("ко")
	grep.CountStr()
	count := make(map[string]int)
	count["file.txt"] = 4
	count["test.txt"] = 4
	require.Equal(t, grep.count, count)
}

func TestGrep_StrFlagn(t *testing.T) {
	grep := newGrep()
	grep.ReadFile("test.txt")
	grep.ReadFile("file.txt")
	grep.GetCheekString("ко")
	grep.StrFlagn()
	number := make(map[string][]int)
	number["file.txt"] = []int{1, 2, 6, 10}
	number["test.txt"] = []int{1, 1, 2, 7}
	require.Equal(t, grep.numberStr, number)
}
