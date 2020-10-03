package controller

import (
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"os"
	"sort"
	"time"
)

type FileInfo struct {
	Name      string    `json:"name"`
	UpdatedAt time.Time `json:"updated_at"`
}

func GetFileList() ([]FileInfo, error) {
	var base = os.Getenv("STATIC_DIR")
	files, err := ioutil.ReadDir(base) // base is defined in upload.go file
	if err != nil {
		return []FileInfo{}, err
	}
	fileList := []FileInfo{}
	for _, file := range files {
		fileItem := FileInfo{
			Name:      file.Name(),
			UpdatedAt: file.ModTime(),
		}
		fileList = append(fileList, fileItem)
	}
	// TODO: sort by modtime.
	sort.SliceStable(fileList, func(i, j int) bool {
		return fileList[i].UpdatedAt.After(fileList[j].UpdatedAt)
	})
	return fileList, nil
}

func Upload(file *multipart.FileHeader) (string, error) {
	var base = os.Getenv("STATIC_DIR")
	src, err := file.Open()
	if err != nil {
		return "", err
	}
	defer src.Close()
	n := fmt.Sprintf("%d-%s", time.Now().UTC().Unix(), file.Filename)
	dst := fmt.Sprintf("%s/%s", base, n)

	out, err := os.Create(dst)
	if err != nil {
		return "", err
	}
	defer out.Close()

	_, err = io.Copy(out, src)
	return n, err
}
