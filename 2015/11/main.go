package main

import (
	"fmt"
	"log"
	"strings"
	"time"
)

func check(e error) {
	if e != nil {
		log.Panic(e)
	}
}
func NotForbidden(password string) bool {
	forbid := []string{"i", "o", "l"}
	for _, char := range forbid {
		if strings.Contains(password, char) {
			return false
		}
	}
	return true
}

func HasPairs(password string) bool {
	s := strings.Split(password, "")
	pairMap := map[string]int{}
	for i := 0; i < len(password)-1; i++ {
		if s[i] == s[i+1] {
			key := fmt.Sprintf("%s%s", s[i], s[i+1])
			if value, ok := pairMap[key]; ok {
				if i-value > 1 {
					return true
				}
			} else {
				pairMap[key] = i
			}
		}
	}
	return len(pairMap) >= 2
}

func HasStraight(password string) bool {
	aToZ := strings.Split("abcdefghijklmnopqrstuvwxyz", "")

	for i := 0; i < len(aToZ)-2; i++ {
		straight := fmt.Sprintf("%s%s%s", aToZ[i], aToZ[i+1], aToZ[i+2])
		// if NotForbidden(straight) {
		if strings.Contains(password, straight) {
			return true
		}
		// }
	}
	return false
}

func Increment(char string) string {
	if char == "z" {
		return "a"
	}
	r := []rune(char)[0]
	r++
	return string(r)
}

func IncrementPass(password string) string {
	passwordArr := strings.Split(password, "")
	for i := len(password) - 1; i >= 0; i-- {
		passwordArr[i] = Increment(passwordArr[i])
		if passwordArr[i] != "a" {
			return strings.Join(passwordArr, "")
		}
	}
	return "hello"
}

func Part1(data string) string {
	newPass := IncrementPass(data)
	for {
		if !NotForbidden(newPass) {
			newPass = IncrementPass(newPass)
			continue
		}
		if !HasPairs(newPass) {
			newPass = IncrementPass(newPass)
			continue
		}
		if !HasStraight(newPass) {
			newPass = IncrementPass(newPass)
			continue
		}
		return newPass
	}
}

func main() {
	startTime := time.Now()
	fmt.Printf("Part 1: %v\n", Part1("hepxcrrq"))
	fmt.Println(time.Since(startTime))
	fmt.Printf("Part 2: %v\n", Part1("hepxxyzz"))
	fmt.Println(time.Since(startTime))
}
