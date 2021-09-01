package main

import (
	"fmt"
	"math"
	"strings"
	"unicode/utf8"
)

func main() {
	reachItSomehow(9)
}

//---REGION: UTILS---
//Func: returns lenght of string
func getLenghtOfString(str string) int { return utf8.RuneCountInString(str) }

//Func: returns number of A's in a string
//Both "a" and "A" considered.
func numberOfCharA(str string) int { return (strings.Count(str, "A") + strings.Count(str, "a")) }

//---REGION: TASK FUNCTIONS---
//func: Q1) sorting string array inorder to number of spesific char in array elements.
func sortByNumberOfCharA(stringArray [6]string) ([6]string, error) {
	for i := 0; i < len(stringArray)-1; i++ {
		for j := 0; j < len(stringArray)-i-1; j++ {
			if numberOfCharA(stringArray[j]) < numberOfCharA(stringArray[j+1]) {
				//swap stringArray's values by indexes
				var temp string = stringArray[j]
				stringArray[j] = stringArray[j+1]
				stringArray[j+1] = temp
			} else if numberOfCharA(stringArray[j]) == numberOfCharA(stringArray[j+1]) {
				if getLenghtOfString(stringArray[j]) < getLenghtOfString(stringArray[j+1]) {
					//swap stringArray's values by indexes
					var temp string = stringArray[j]
					stringArray[j] = stringArray[j+1]
					stringArray[j+1] = temp
				}
			}
		}
	}

	return stringArray, nil
}

//func: Q2) recursive algorithm that reaches input somehow
func reachItSomehow(input int) int {
	var inpDouble float64 = float64(input)
	//runs until reaches 2
	if input == 2 {
		fmt.Println(input)
		return 2
	} else {
		fmt.Println(input)
		return reachItSomehow(int(math.Ceil((inpDouble - 1) / 2)))
	}

}

//func: Q3) find most repeated word
func findMostRepeatedCharacter(stringArray [7]string) string {
	var (
		maxKey    string = ""
		strMap    map[string]int
		maxNumber int = 0
	)
	//initialize map
	strMap = make(map[string]int)
	for i := 0; i < len(stringArray); i++ { //parse input array to fill map
		v, contains := strMap[stringArray[i]]
		if contains == false {
			strMap[stringArray[i]] = 1 //if key do not exist, add it with value of 1
		} else {
			strMap[stringArray[i]] = v + 1 //if key do exist, increase value by 1
		}
	}
	for k := range strMap { //iterate over map to find key with maximum value
		if strMap[k] > maxNumber {
			maxNumber = strMap[k]
			maxKey = k
		}
	}
	return maxKey
}
