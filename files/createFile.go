package files

import (
	"os"
	"fmt"
)

/**
 * Creates file based on filename provided
 * @params: filename to be created
 * @return: error if applicable
 */
func CreateLogFile(fileName string) (err error) {

	file, err := os.Create(fileName)
    if err != nil {
        fmt.Println(err)
        return
    }
    
    err = file.Close()
    if err != nil {
        fmt.Println(err)
        return
	}
    
    return
}