package data

import()

func RemoveFrontChars(data string, num int) string {
	if len(data) > 0 {
	return data[num:]
	}
	return data
}

func RemoveLastChars(data string, num int) string {
	if len(data) > 0 {
		return data[:len(data)-num]
	}
	return data
}

func RemoveBothEndChars(data string, front int, last int) string {
	if len(data) > 0 {
		return data[front:len(data)-last]
	}
	return data
}