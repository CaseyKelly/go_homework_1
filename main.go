package main

import (
	"flag"
	"fmt"
	"strconv"
	"strings"
)

var (
	ex1 = flag.String("ex1", "", "String of operations without spaces using all the same operation")
	ex2 = flag.String("ex2", "", "String of operations with spaces using all the same operation")
)

func main() {

	flag.Parse()
	if strings.Contains(*ex1, "+") {
		trim := strings.Join(strings.Fields(*ex1), "")
		splt := strings.Split(trim, "+")
		sum, _ := strconv.Atoi(splt[0])
		for i := 1; i < (len(splt)); i++ {
			num, _ := strconv.Atoi(splt[i])
			diff := sum + num
			sum = diff
			if i == (len(splt) - 1) {
				fmt.Println(sum)
			}
		}
	} else if strings.Contains(*ex1, "-") {
		trim := strings.Join(strings.Fields(*ex1), "")
		splt := strings.Split(trim, "-")
		sum, _ := strconv.Atoi(splt[0])
		for i := 1; i < (len(splt)); i++ {
			num, _ := strconv.Atoi(splt[i])
			diff := sum - num
			sum = diff
			if i == (len(splt) - 1) {
				fmt.Println(diff)
			}
		}
	} else if strings.Contains(*ex1, "*") {
		trim := strings.Join(strings.Fields(*ex1), "")
		splt := strings.Split(trim, "*")
		sum, _ := strconv.Atoi(splt[0])
		for i := 1; i < (len(splt)); i++ {
			num, _ := strconv.Atoi(splt[i])
			diff := sum * num
			sum = diff
			if i == (len(splt) - 1) {
				fmt.Println(diff)
			}
		}
	} else if strings.Contains(*ex1, "/") {
		trim := strings.Join(strings.Fields(*ex1), "")
		splt := strings.Split(trim, "/")
		sum, _ := strconv.Atoi(splt[0])
		for i := 1; i < (len(splt)); i++ {
			num, _ := strconv.Atoi(splt[i])
			diff := sum / num
			sum = diff
			if i == (len(splt) - 1) {
				fmt.Println(diff)
			}
		}
	}
}
