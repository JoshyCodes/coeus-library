package data

import "fmt"

/**
 * Converts map[string]interface{} object into single string object.
 * @params: data is the data object to be converted.
 * @return: string of converted data
 */
func JsonStringConvert(data map[string]interface{}) string {

	// set vars
	var jsn string
	
	// loop over data and use switch based on var type
	for k, v := range data {
		
		switch v := v.(type) {
		
		case string:
			jsn += fmt.Sprintf(`"%s" : "%s",`, k, v)
		
		case float64:
			jsn += fmt.Sprintf(`"%s" : %.2f,`, k, v)
		
		case []interface{}:
			objArray := ""

			for _, u := range v {
				objArray += fmt.Sprintf(`"%s",`, u)
			}
			
			jsn += fmt.Sprintf(`"%s" : [%s],`, k, RemoveLastChars(objArray, 1))
		
		case map[string]interface{}:
			jsn += fmt.Sprintf(`"%s" : %s,`, k, JsonStringConvert(v))
		
		default:
			fmt.Println(k, v, fmt.Sprintf("(MISMATCHED VARIABLE) %T", v))
		}
	}

	// return new json string value
	return fmt.Sprintf(`{%s}`, RemoveLastChars(jsn, 1))
}