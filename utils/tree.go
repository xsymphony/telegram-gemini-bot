package utils

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
)

func PrintTree(dir string, indent string) error {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return err
	}

	for i, file := range files {
		if file.IsDir() {
			fmt.Printf("%s├── %s\n", indent, file.Name())
			if i == len(files)-1 {
				PrintTree(filepath.Join(dir, file.Name()), indent+"    ")
			} else {
				PrintTree(filepath.Join(dir, file.Name()), indent+"│   ")
			}
		} else {
			if i == len(files)-1 {
				fmt.Printf("%s└── %s\n", indent, file.Name())
			} else {
				fmt.Printf("%s├── %s\n", indent, file.Name())
			}
		}
	}

	return nil
}
