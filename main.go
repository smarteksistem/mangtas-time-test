/*
02/02/2022
metafiliana@gmail.com
metafiliana@smarteksistem.com

Smartek - Mangtas golang client test
*/

package main

import (
	"fmt"
)

/*
Task: Given a list of 4 integers, find all possible valid 24 hour times (eg: 12:34) that the given integers can be assembled into and return the total number of valid times.
You can not use the same number twice.
Times such as 34:12 and 12:60 are not valid.

Notes: Input integers can not be negative.
Input integers can yeald 0 possible valid combinations.

Example:
	Input: [1, 2, 3, 4]
	Valid times: ["12:34", "12:43", "13:24", "13:42", "14:23", "14:32", "23:14", "23:41"]
	Return: 8
*/

// remove func
// remove the slice component by the slice key (s)
// the slice is still ordered
func remove(slice []int, s int) []int {
	return append(slice[:s], slice[s+1:]...)
}

func possibleTimes(digits []int) int {
	// copy digits value
	// availableDigits used in 1 iteration
	availableDigits := make([]int, len(digits))
	copy(availableDigits, digits)

	// possible valid 24 hour time
	// stored in this map uniquely
	possibleTimesMap := make(map[string]bool)

	for kh1, hour1 := range availableDigits {
		// copy the value
		// so it would be possible to iterate with
		// decrement value of digits
		tempDigits1 := make([]int, len(availableDigits))
		copy(tempDigits1, availableDigits)

		// possible hour is between 0 and 2
		if hour1 > 2 {
			continue
		}

		firstHour := hour1

		// remove the digit, so it wont be doubled
		tempDigits1 = remove(tempDigits1, kh1)

		for kh2, hour2 := range tempDigits1 {
			// if first digit of hour is 2
			// possible 2nd digit of hour is between 0 and 3
			if firstHour == 2 && hour2 > 3 {
				continue
			}

			tempDigits2 := make([]int, len(tempDigits1))
			copy(tempDigits2, tempDigits1)

			// remove the digit, so it wont be doubled
			tempDigits2 = remove(tempDigits2, kh2)

			for km1, minute1 := range tempDigits2 {
				// possible first digit of minute value
				// is <= 5
				if minute1 > 5 {
					continue
				}

				tempDigits3 := make([]int, len(tempDigits2))
				copy(tempDigits3, tempDigits2)

				// remove the digit, so it wont be doubled
				tempDigits3 = remove(tempDigits3, km1)

				for _, minute2 := range tempDigits3 {
					// possible 2nd digit of minute value
					// is <= 9
					if minute2 > 9 {
						continue
					}

					// set string key of valid 24hr time
					timesString := fmt.Sprintf("%d%d:%d%d", hour1, hour2, minute1, minute2)

					// check if possible time isn't available in maps
					// store it in maps
					if _, ok := possibleTimesMap[timesString]; !ok {
						possibleTimesMap[timesString] = true

						// break the last minutes loop
						break
					}
				}
			}
		}
	}

	return len(possibleTimesMap)
}

func main() {
	// Example test cases.
	fmt.Println(possibleTimes([]int{1, 2, 3, 4}))
	fmt.Println(possibleTimes([]int{2, 2, 1, 9}))

	// added test case with 0 possible valid time
	fmt.Println(possibleTimes([]int{2, 2, 10, 9}))
}
