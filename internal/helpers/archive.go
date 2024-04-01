package helpers

import (
	"fmt"
	"github.com/fouched/go-contact-app/internal/repo"
	"strconv"
)

var ArchiveInstances = make(map[int]Archive)

type Archive struct {
	Status      string
	Progress    int
	ArchiveFile string
}

func RunArchive(key int) string {
	archive := ArchiveInstances[key]
	fileName := "./archive/" + strconv.Itoa(key) + ".csv"
	err := repo.CreateAllContactsArchive(fileName)
	if err != nil {
		fmt.Println("Error creating archive:" + err.Error())
		archive.Progress = 100
		archive.Status = "Error"
		ArchiveInstances[key] = archive
		return "Error creating archive"
	}

	archive.Progress = 100
	archive.Status = "Complete"
	archive.ArchiveFile = fileName
	ArchiveInstances[key] = archive
	fmt.Println("Creating Archive 100%")

	return "Archive ready for download"
}

func Reset() {

}
