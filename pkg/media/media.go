package media

import (
	
	"github.com/nickmancari/mvBot/pkg/fshandler"

)

type Media struct {
	FolderName		string
	FinishedDownload	bool
	FormatCounts		map[string]int
	Medium			string
	Locations		[]string
}
	

func (m *Media) FormatCheckCounter(format string) *Media {
	
	if m.FormatCounts == nil {
		m.FormatCounts = make(map[string]int)
	}

	count := fshandler.ContentCount(format, m.FolderName)
	
	m.FormatCounts[format] = count


	return m

}

func (m *Media) MediumCheck() *Media {
	

	for _, v := range m.FormatCounts {
		switch count := v; {
		case count > 1:
			m.Medium = "TV"
			return m
		}
	}
	
	m.Medium = "Movie"

	return m
}


func (m *Media) Formats(formats []string) *Media {
	

	for _, format := range formats {
		m.FormatCheckCounter(format)
	}

	return m
}

func (m *Media) FinishCheck(formats []string) *Media {

	var finishedStatus bool
	for _, format := range formats {
		finishedStatus = fshandler.DownloadFinished(format, m.FolderName)
		if finishedStatus == false {
			m.FinishedDownload = false
			return m
		} else {
			m.FinishedDownload = true
			return m
		}
	}

	return m

}

func (m *Media) Folder(folder string) *Media {

	m.FolderName = folder

	return m

}

func (m *Media) FileLocations() *Media {

	var locations []string
	for k, v := range m.FormatCounts {
		if v >= 1 {
			locations = fshandler.GetMediaFiles(m.FolderName, k)
			m.Locations = locations
			return m
		}
	}

	return m
}

func Analyzer() []*Media {

	formats := []string{".mp4", ".avi", ".mkv", ".mov", ".wmv"}

	folders := fshandler.GetFolders()

	var collection []*Media

	for _, folder := range folders {
		m := &Media{}
		media := m.Folder(folder).FinishCheck(formats)
		files := media.Formats(formats).MediumCheck().FileLocations()

		collection = append(collection, files)
	}

	return collection

}
