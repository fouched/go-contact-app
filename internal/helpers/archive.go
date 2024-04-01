package helpers

import (
	"fmt"
	"time"
)

var ArchiveInstances = make(map[int]Archive)

type Archive struct {
	Status      string
	Progress    int
	ArchiveFile string
}

func RunArchive(key int) string {
	archive := ArchiveInstances[key]
	fmt.Println("Creating Archive 0%")
	time.Sleep(2 * time.Second)
	archive.Progress = 20
	ArchiveInstances[key] = archive
	fmt.Println("Creating Archive 20%")
	time.Sleep(2 * time.Second)
	archive.Progress = 40
	ArchiveInstances[key] = archive
	fmt.Println("Creating Archive 40%")
	time.Sleep(2 * time.Second)
	archive.Progress = 60
	ArchiveInstances[key] = archive
	fmt.Println("Creating Archive 60%")
	time.Sleep(2 * time.Second)
	archive.Progress = 80
	ArchiveInstances[key] = archive
	fmt.Println("Creating Archive 80%")
	time.Sleep(2 * time.Second)
	archive.Progress = 100
	archive.Status = "Complete"
	archive.ArchiveFile = "/some/path.csv"
	ArchiveInstances[key] = archive
	fmt.Println("Creating Archive 100%")

	return "Archive ready for download"
}

func Reset() {

}
