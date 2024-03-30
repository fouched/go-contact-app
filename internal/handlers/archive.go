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

func NewArchive(ai *Archive) {
	ArchiveInstance = ai
}

func RunArchive() string {
	fmt.Println("Creating Archive 0%")
	ArchiveInstance.Status = "Running"
	time.Sleep(1 * time.Second)
	ArchiveInstance.Progress = 20
	fmt.Println("Creating Archive 20%")
	time.Sleep(1 * time.Second)
	ArchiveInstance.Progress = 40
	fmt.Println("Creating Archive 40%")
	time.Sleep(1 * time.Second)
	ArchiveInstance.Progress = 60
	fmt.Println("Creating Archive 60%")
	time.Sleep(1 * time.Second)
	ArchiveInstance.Progress = 80
	fmt.Println("Creating Archive 80%")
	time.Sleep(1 * time.Second)
	ArchiveInstance.Progress = 100
	ArchiveInstance.Status = "Complete"
	fmt.Println("Creating Archive 100%")

	ArchiveInstance.ArchiveFile = "/some/path.csv"
	return "Archive ready for download"
}

func Reset() {

}
