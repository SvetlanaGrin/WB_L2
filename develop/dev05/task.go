package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Grep struct {
	cheekStr  string
	dataFile  map[string]string
	data      map[string][]string
	count     map[string]int
	numberStr map[string][]int
	builder   strings.Builder
}

func (g *Grep) GetCheekString(str string) {
	g.cheekStr = str
}
func (g *Grep) ReadFile(nameFile string) {
	file, err := os.Open(nameFile)
	if err != nil {
		log.Fatal(err.Error())
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		g.builder.WriteString(scanner.Text())
		g.builder.WriteString("|")

	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err.Error())
		os.Exit(1)
	}
	g.dataFile[nameFile] = g.builder.String()
	g.builder.Reset()
}

func newGrep() *Grep {
	return &Grep{data: make(map[string][]string),
		dataFile:  make(map[string]string),
		count:     make(map[string]int),
		numberStr: make(map[string][]int),
		builder:   strings.Builder{}}
}
func (g *Grep) CountStr() {
	for key, _ := range g.dataFile {
		allIndex := g.FindIndex(key)
		g.count[key] = len(allIndex)
		fmt.Println(key, ":", len(allIndex))
	}
}
func (g *Grep) StrFlagn() {
	for key, _ := range g.dataFile {
		allIndex := g.FindIndex(key)
		count := 1
		for i := 0; i < len(allIndex); i++ {
			for j := 0; j < allIndex[i][0]; j++ {
				if string([]byte(g.dataFile[key])[j]) == "|" {
					count++
				}
			}
			fmt.Println(key, ":", count)
			g.numberStr[key] = append(g.numberStr[key], count)

		}
	}
}

func (g *Grep) FindIndex(key string) [][]int {
	re := regexp.MustCompile(g.cheekStr)
	return re.FindAllIndex([]byte(g.dataFile[key]), -1)
}
func (g *Grep) FoundStr() {
	for key, _ := range g.dataFile {
		allIndex := g.FindIndex(key)
		for i := 0; i < len(allIndex); i++ {
			for j := allIndex[i][0]; j < allIndex[i][0]+30 && j < len([]byte(g.dataFile[key])); j++ {
				if string([]byte(g.dataFile[key])[j]) == "|" {
					break
				}
				g.builder.WriteByte([]byte(g.dataFile[key])[j])

			}
			g.data[key] = append(g.data[key], g.builder.String())
			g.builder.Reset()
			fmt.Println(key, ":", g.data[key][i])

		}
	}
}
func (g *Grep) FoundStrFlagA(n int) {
	for key, _ := range g.dataFile {
		allIndex := g.FindIndex(key)
		for i := 0; i < len(allIndex); i++ {
			strSpl := strings.Split(string([]byte(g.dataFile[key])[allIndex[i][0]:]), " ")
			for j := 0; j < len(strSpl) && j < n; j++ {
				if strings.Contains(strSpl[j], "|") {
					g.builder.WriteString(strings.Split(strSpl[j], "|")[0])
					break
				}
				g.builder.WriteString(strSpl[j])
				g.builder.WriteString(" ")

			}
			g.data[key] = append(g.data[key], g.builder.String())
			g.builder.Reset()
		}
	}
}
func (g *Grep) Print() {
	for key, elems := range g.data {
		for _, elem := range elems {
			println(key, ":", elem)
		}
	}
}
func (g *Grep) FoundStrFlagB(n int) {
	for key, _ := range g.dataFile {
		allIndex := g.FindIndex(key)
		for i := 0; i < len(allIndex); i++ {
			strSpl := strings.Split(string([]byte(g.dataFile[key])[:allIndex[i][1]]), " ")
			for j := len(strSpl) - n; j < len(strSpl); j++ {
				if j < 0 {
					continue
				}
				if strings.Contains(strSpl[j], "|") {
					g.builder.Reset()
					g.builder.WriteString(strings.Split(strSpl[j], "|")[len(strings.Split(strSpl[j], "|"))-1])
					g.builder.WriteString(" ")
					continue
				}
				g.builder.WriteString(strSpl[j])
				g.builder.WriteString(" ")

			}
			g.data[key] = append(g.data[key], g.builder.String())
			g.builder.Reset()
		}
	}
}

func (g *Grep) FoundStrFlagC(n int) {
	g.FoundStrFlagB(n)
	g.FoundStrFlagA(n)
	data := make(map[string][]string)
	for key, elems := range g.data {
		for i := 0; i < len(elems)/2; i++ {
			g.data[key][i] = string([]byte(g.data[key][i])[:len(g.data[key][i])-len(g.cheekStr)-1])
			g.builder.WriteString(g.data[key][i])
			g.builder.WriteString(g.data[key][i+len(elems)/2])

			fmt.Println(key, ":", g.builder.String())
			data[key] = append(data[key], g.builder.String())
			g.builder.Reset()
		}
	}
	g.data = data
}

func main() {
	grepCommand := flag.NewFlagSet("grep", flag.ExitOnError)

	// Count subcommand flag pointers
	// Adding a new choice for --metric of 'substring' and a new --substring flag
	grepA := grepCommand.String("A", "0", " печатать +N строк после совпадения")
	grepB := grepCommand.String("B", "0", "печатать +N строк до совпадения")
	grepC := grepCommand.String("C", "0", "(A+B) печатать ±N строк вокруг совпадения")
	grepc := grepCommand.Bool("c", false, "(количество строк)")
	grepn := grepCommand.Bool("n", false, "напечатать номер строки")
	sortNotFlag := grepCommand.Bool("", true, "default")
	grep := newGrep()
	var cheekFlag int
	switch os.Args[1] {
	case "grep":
		grep.GetCheekString(os.Args[2])
		if _, err := os.Stat(os.Args[3]); err == nil {
			for i := 3; i < len(os.Args); i++ {
				if _, err := os.Stat(os.Args[i]); errors.Is(err, os.ErrNotExist) {
					cheekFlag = i
					fmt.Println(cheekFlag)
					break
				} else {
					grep.ReadFile(os.Args[i])
				}
			}
		}
		grepCommand.Parse(os.Args[cheekFlag:])
	default:
		flag.PrintDefaults()
		os.Exit(1)
	}

	if grepCommand.Parsed() {

	}
	// Required Flags
	if *grepA != "0" {
		n, err := strconv.Atoi(os.Args[cheekFlag+1])
		if err != nil {
			return
		}
		fmt.Println(n)
		grep.FoundStrFlagA(n)
		grep.Print()
		*sortNotFlag = false
	}
	if *grepB != "0" {
		n, err := strconv.Atoi(os.Args[cheekFlag+1])
		if err != nil {
			return
		}
		fmt.Println(n)
		grep.FoundStrFlagB(n)
		grep.Print()
		*sortNotFlag = false

	}
	if *grepC != "0" {
		n, err := strconv.Atoi(os.Args[cheekFlag+1])
		if err != nil {
			return
		}
		fmt.Println(n)
		grep.FoundStrFlagC(n)
		*sortNotFlag = false
	}
	if *grepc {
		*sortNotFlag = false
		grep.CountStr()
	}
	if *grepn {
		*sortNotFlag = false
		grep.StrFlagn()
	}
	if *sortNotFlag {
		grep.FoundStr()
	}
	// If the metric flag is substring, the substring or substringList flag is required
	fmt.Printf("-n: %v, -k: %v, -r: %v, -u: %v\n",
		*grepA,
		*grepB,
		*grepC,
		*grepc,
	)
}
