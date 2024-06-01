package storage

import (
	"chatgo/internal/chatgo/domain"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

type Filesystem struct {
}

func NewFilesystem() *Filesystem {
	return &Filesystem{}
}

func (s *Filesystem) SaveFiles(files map[string]*domain.File) {
	now := time.Now()
	for _, file := range files {
		log.Printf("File: %v Content: %v", file.Path, file.Content)
		bytes := []byte(file.Content)
		path := "out/" + strconv.FormatInt(now.Unix(), 10) + "/" + file.Path

		// Ensure the directory exists
		dir := filepath.Dir(path)
		if err := os.MkdirAll(dir, os.ModePerm); err != nil {
			log.Printf("error creating directory %v: %v", dir, err)
			continue
		}

		// Write the file
		err := os.WriteFile(path, bytes, 0644)
		if err != nil {
			log.Printf("error write file %v: %v", file.Path, err)
		}
	}
}
