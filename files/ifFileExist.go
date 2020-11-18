package files

import (
	"os"
)

/**
 * Check if file exists
 * @params: filename of file to check if existing
 * @return: bool to represent file status
 */
func FileExists(filename string) bool {

    info, err := os.Stat(filename)
    if os.IsNotExist(err) {
        return false
    }
    
    return !info.IsDir()
}