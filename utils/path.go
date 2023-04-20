package utils

import (
	"fmt"
	"os"
)

func GetFilesInDir(dir string) []string {
	var result = []string{}
	var index = 0
	var directories = []string{dir}
	for index < len(directories) {
		p := directories[index]
		read, _ := os.ReadDir(p)
		for _, file := range read {
			if file.IsDir() {
				directories = append(directories, fmt.Sprintf("%v/%v", p, file.Name()))
				continue
			}
			result = append(result, fmt.Sprintf("%v/%v", p, file.Name()))
		}
		index++
	}
	return result
}
