package data

import ()

func IfArrayContainString(array []string, val string) bool {
	for _, data := range array {
		if data == val {
			return true
		}
	}
	return false
}