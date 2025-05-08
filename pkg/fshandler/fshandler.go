package fshandler

import (
	"fmt"
	"os"
	"strings"

	"github.com/nickmancari/mvBot/sys/config"
)

var Settings = config.Read()

func EmptyDir() bool {
	
	entries, err := os.ReadDir(Settings.LocalFolder)
	if err != nil {
		fmt.Printf("Directory Empty Check Error: %v\n", err)
	}
	
	var i int
	for _, content := range entries {
		if content != nil {
			i++
		}
	}

	if i > 0 {
		return true
	}

	return false

}

func DownloadFinished(mediaType string, folder string) bool {
	
	entries, err := os.ReadDir(Settings.LocalFolder+folder)
	if err != nil {
		fmt.Printf("Reading Directory for Content Finished Check Error: %v\n", err)
	}

	//Debug helper
	_ = mediaType

	for _, file := range entries {
		if strings.Contains(file.Name(), ".part") {
			return false
		}
	}

	return true

}

func ContentCount(mediaType string, folder string) int {

	entries, err := os.ReadDir(Settings.LocalFolder+folder)
	if err != nil {
		fmt.Printf("Reading Directory for Content Count Error: %v\n", err)
	}

	var i int
	for _, file := range entries {
		if strings.Contains(file.Name(), mediaType) {
			i++
		}
	}

	return i
}

func GetFolders() []string {
	
	folders, err := os.ReadDir(Settings.LocalFolder)
	if err != nil {
		fmt.Printf("Reading Folder to Get Files Error: %v", err)
	}

	var content []string

	for _, folder := range folders {

		content = append(content, folder.Name())
	}

	return content

}

func GetMediaFiles(folder string, format string) []string {

	entries, err := os.ReadDir(Settings.LocalFolder+folder)
	if err != nil {
		fmt.Printf("Readin Folder to get Files Error: %v", err)
	}

	var path string
	var fullPathCollection []string
	for _, file := range entries {
		if strings.Contains(file.Name(), format) {
			path = Settings.LocalFolder+folder+"/"+file.Name()
			fullPathCollection = append(fullPathCollection, path)
		}
	}

	return fullPathCollection
}
