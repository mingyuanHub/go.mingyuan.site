package common

import "strconv"

func IsContain(items []int, item int) bool {
	for _, eachItem := range items {
		if eachItem == item {
			return true
		}
	}
	return false
}

func IsContainStr(items []string, item string) bool {
	for _, eachItem := range items {
		if eachItem == item {
			return true
		}
	}
	return false
}

func StrArrToIntArr(strArr []string) []int {
	var intArr []int
	for _, item := range strArr {
		if i ,err := strconv.Atoi(item); err == nil {
			intArr = append(intArr, i)
		}
	}
	return intArr
}