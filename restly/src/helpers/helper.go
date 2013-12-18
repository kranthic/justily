package helpers

import (
	"os"
	"fmt"
)

func GenUUID() (string, error){
	b := make([]byte, 16)
	f,err := os.Open("/dev/urandom")
	if err != nil{
		return "", err
	}
	defer f.Close()
	f.Read(b)
    return fmt.Sprintf("%x%x%x%x%x", b[0:4], b[4:6], b[6:8], b[8:10], b[10:]), nil
}

