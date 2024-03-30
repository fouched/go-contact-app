package handlers

import (
	"fmt"
	"time"
)

type Archive struct {
	Status string
}

func NewArchive() *Archive {
	return &Archive{
		Status: "Waiting",
	}
}

func (a *Archive) Progress() int {
	return 0
}

func (a *Archive) Run() string {
	fmt.Println("Creating Archive")
	time.Sleep(5 * time.Second)
	fmt.Println("Archive ready for download")
	return "Archive ready for download"
}

func (a *Archive) Reset() {

}

func (a *Archive) ArchiveFile() string {
	return "/some/path"
}
