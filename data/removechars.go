package data

import()

func RemoveFrontChars(data string, num int) string {
	if len(data) > 0 {
	return data[num:]
	}
	return data
}