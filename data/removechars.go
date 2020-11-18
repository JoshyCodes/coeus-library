package data

/**
 * Remove x chars from front of string
 * @params: data is string to be modified
 * @params: num of chars to remove from front of string
 * @return: new string
 */
func RemoveFrontChars(data string, num int) string {
	
	if len(data) > 0 {
		return data[num:]
	}

	return data
}

/**
 * Remove x chars from End of string
 * @params: data is string to be modified
 * @params: num of chars to remove from End of string
 * @return: new string
 */
func RemoveLastChars(data string, num int) string {
	
	if len(data) > 0 {
		return data[:len(data)-num]
	}
	
	return data
}

/**
 * Remove x chars from front/end of string
 * @params: data is string to be modified
 * @params: front is num of chars to remove from front of string
 * @params: last is num of chars to remove from end of string
 * @return: new string
 */
func RemoveBothEndChars(data string, front int, last int) string {
	
	if len(data) > 0 {
		return data[front:len(data)-last]
	}
	
	return data
}