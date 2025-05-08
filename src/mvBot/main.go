package main

import (
	"fmt"

	"github.com/nickmancari/mvBot/pkg/media"
	"github.com/nickmancari/mvBot/pkg/fshandler"
	"github.com/nickmancari/mvBot/sys/systemcall"
)

func main() {
  
	for {
	
		if fshandler.EmptyDir() == true {
			
			mediaFiles := media.Analyzer()
		
			err := systemcall.Migrate(mediaFiles)
			if err != nil {
				fmt.Println(err)
			}
		}

	}
}
