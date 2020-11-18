package files
 
import (
	"log"
	"os"
)
 
/**
 * Create dir if not already existing
 * @params: path is name/location of dir
 */
func CreateDir(path string) {

	_, err := os.Stat(path)
 
	if os.IsNotExist(err) {
	
		errDir := os.MkdirAll(path, 0755)
		if errDir != nil {
			log.Fatal(err)
		}
	}
}