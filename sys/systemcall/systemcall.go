package systemcall

import (
	"os/exec"
	"fmt"
	
	"github.com/nickmancari/mvBot/pkg/media"
	"github.com/nickmancari/mvBot/sys/config"
)

var Settings = config.Read()

func Send(m *media.Media) error {

	if m.FinishedDownload == true {
		switch medium := m.Medium; {
		case medium == "Movie":
			
			for _, file := range m.Locations {

				fmt.Printf("Transferable Media File Found: %v\n", file)

			}

			err := remoteSendMovieFiles(m.Locations)
			if err == nil {
				err = Delete(m)
				return err
			}
			return err
		}
	}

	return nil
}

func Delete(m *media.Media) error {
	

	//logging
	fmt.Printf("Deleting Folder: %v\n", Settings.LocalFolder+m.FolderName)

	cmd := exec.Command("rm", "-rf", Settings.LocalFolder+m.FolderName)
	err := cmd.Run()

	if err == nil {
		fmt.Printf("Folder Deleted: %v\n", Settings.LocalFolder+m.FolderName)
		return err
	}
	
	fmt.Printf("Folder Error During Deletion: %v, Please Ensure Nomady is Run with Permissions\n", err)
	return err
}

func Migrate(m []*media.Media) error {

	var err error
	for _, files := range m {
		err = Send(files)
		if err != nil {
			return err
		}
	}
	
	return err
}

func remoteSendMovieFiles(locations []string) error {

	for _, file := range locations {
		cmd := exec.Command("scp", file, Settings.RemoteUser+"@"+Settings.RemoteServer+":"+Settings.RemoteFolder+"Movies/")
		err := cmd.Run()
		if err != nil {
			fmt.Println("File Failed to Send: "+file+" "+Settings.RemoteUser+"@"+Settings.RemoteServer+":"+Settings.RemoteFolder+"Movies/")
			return err
		}

		//logging help
		fmt.Printf("File Sent Successfully: %v\n", file)
	}

	return nil
}
