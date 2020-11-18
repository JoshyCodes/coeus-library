package data

/**
 * Check if string value exists within an array
 * @params: array of strings
 * @params: val is string value to search for
 * @return: bool response of string status
 */func IfArrayContainString(array []string, val string) bool {

	for _, data := range array {
	
		if data == val {
			return true
		}
	}
	
	return false
}