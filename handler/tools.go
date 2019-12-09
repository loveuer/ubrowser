package handler

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/h2non/filetype"
)

type FOLDERS struct {
	Name    string `json:"name"`
	Path    string `json:"path"`
	ModTime string `json:"modified"`
	Size    int64  `json:"size"`
	IsDir   bool   `json:"isDir"`
	Type    string `json:"type"`
}

func (f *FOLDERS) getType() {
	file, err := os.Open(f.Path)
	if err != nil {
		log.Printf("< handler > < tools > os open file err => %v\n", err)
		return
	}

	head := make([]byte, 261)
	file.Read(head)
	switch {
	case filetype.IsImage(head):
		f.Type = "image"
	case filetype.IsAudio(head):
		f.Type = "audio"
	case filetype.IsVideo(head):
		f.Type = "video"
	case filetype.IsArchive(head):
		f.Type = "archive"
	case filetype.IsDocument(head):
		f.Type = "document"
	default:
		f.Type = "other"
	}

	return
}

func getfolder(fp string) (*[]*FOLDERS, error) {
	var fds []*FOLDERS

	fs, err := ioutil.ReadDir(fp)
	if err != nil {
		log.Printf("< handler > < tools > open fullpath err => %v\n", err)
		return &fds, err
	}

	for _, val := range fs {
		one := &FOLDERS{
			Name:    val.Name(),
			Path:    filepath.Join(fp, val.Name()),
			ModTime: val.ModTime().Format("2006/01/02 15:04:05"),
			Size:    val.Size(),
			IsDir:   val.IsDir(),
			Type:    "folder",
		}
		if !one.IsDir {
			one.getType()
		}

		fds = append(fds, one)
	}

	return &fds, nil
}
