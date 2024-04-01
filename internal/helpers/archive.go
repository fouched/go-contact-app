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

	c := make(chan int)
	go repo.CreateAllContactsArchive(fileName, c)
	for i := range c {
		archive.Progress = i
		ArchiveInstances[key] = archive
	}

	if archive.Progress == 100 {
		archive.Status = "Complete"
		archive.ArchiveFile = fileName
	}

	ArchiveInstances[key] = archive
	fmt.Println("Creating Archive 100%")

	return "Archive ready for download"
}

func Reset() {

}
