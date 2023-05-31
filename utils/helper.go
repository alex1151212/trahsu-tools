package utils

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func ChDir(path string) {
	err := os.Chdir(path)
	if err != nil {
		log.Fatal(err)
	}
}

func FindDirFileName(path string, exts []string) (int, []string) {

	var fileNames []string

	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// 如果是檔案且附檔名屬於要檢查的附檔名之一，就將檔案名稱存儲到切片中
		if !info.IsDir() {
			for _, ext := range exts {
				if strings.HasSuffix(info.Name(), string(ext)) {
					fileNames = append(fileNames, info.Name())
					break
				}
			}
		}

		return nil
	})

	if err != nil {
		fmt.Println(err)

	}

	// 輸出符合條件的檔案名稱
	for _, filename := range fileNames {
		fmt.Println(filename)
	}
	return len(fileNames), fileNames
}
