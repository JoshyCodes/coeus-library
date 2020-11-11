package files

import (
	"os"
	"fmt"
)

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