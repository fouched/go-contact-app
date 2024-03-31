package handlers

import (
	"fmt"
	"time"
)

var ArchiveInstance *Archive

type Archive struct {
	Status      string
	Progress    int
	ArchiveFile string
}

func NewArchive(archive *Archive) {
	ArchiveInstance = archive
}

func RunArchive(archive *Archive) string {
	fmt.Println("Creating Archive 0%")
	time.Sleep(2 * time.Second)
	archive.Progress = 20
	fmt.Println("Creating Archive 20%")
	time.Sleep(2 * time.Second)
	archive.Progress = 40
	fmt.Println("Creating Archive 40%")
	time.Sleep(2 * time.Second)
	archive.Progress = 60
	fmt.Println("Creating Archive 60%")
	time.Sleep(2 * time.Second)
	archive.Progress = 80
	fmt.Println("Creating Archive 80%")
	time.Sleep(2 * time.Second)
	archive.Progress = 100
	archive.Status = "Complete"
	fmt.Println("Creating Archive 100%")

	archive.ArchiveFile = "/some/path.csv"
	return "Archive ready for download"
}

func Reset() {

}
