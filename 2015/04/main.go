package main

import (
	"crypto/md5"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"
)

func check(e error) {
	if e != nil {
		log.Panic(e)
	}
}

func Md5(in string) string {
	data := []byte(in)
	return fmt.Sprintf("%x", md5.Sum(data))
}

func Part1(input string) int {
	i := 0
	for {
		iString := strconv.Itoa(i)
		newString := fmt.Sprintf("%s%s", input, iString)
		md5 := Md5(newString)
		if strings.HasPrefix(md5, "00000") {
			return i
		} else {
			i++
		}
	}
}

func Part2(input string) int {
	i := 0
	for {
		iString := strconv.Itoa(i)
		newString := fmt.Sprintf("%s%s", input, iString)
		md5 := Md5(newString)
		if strings.HasPrefix(md5, "000000") {
			return i
		} else {
			i++
		}
	}
}

func main() {
	startTime := time.Now()
	data := "yzbqklnj"
	fmt.Printf("Part 1: %v\n", Part1(data))
	fmt.Println(time.Since(startTime))
	fmt.Printf("Part 2: %v\n", Part2(data))
	fmt.Println(time.Since(startTime))
}
