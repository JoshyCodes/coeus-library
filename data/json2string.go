package data

import (
	"fmt"
)

func JsonStringConvert(data map[string]interface{}) string {
	var jsn string
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
			fmt.Println(k, v, "(unknown)")
			fmt.Printf("%T\n", v)
		}
	}
	return fmt.Sprintf(`{%s}`, RemoveLastChars(jsn, 1))
}