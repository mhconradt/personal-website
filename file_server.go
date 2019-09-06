package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"path/filepath"
	"strings"
)

type Info struct {
	Path string
	Size int
}

type FileServer struct {
	files map[string][]byte
}

func NewFileServer(dir string) (FileServer, error) {
	m, err := LoadDirectory(dir)
	if err != nil {
		return FileServer{}, err
	}
	return FileServer{m}, nil
}

func WalkDir(dir string) ([]Info, error) {
	infos := make([]Info, 0)
	dirInfo, err := ioutil.ReadDir(dir)
	if err != nil {
		return infos, err
	}
	for _, info := range dirInfo {
		n := info.Name()
		fullpath := filepath.Join(dir, n)
		if info.IsDir() {
			n := info.Name()
			ownInfos, err := WalkDir(filepath.Join(dir, n))
			if err != nil {
				fmt.Println("walk dir error")
				return infos, err
			}
			infos = append(infos, ownInfos...)
		} else {
			// recursion base case
			infos = append(infos, Info{fullpath, int(info.Size())})
		}
	}
	return infos, err
}

func LoadDirectory(dir string) (m map[string][]byte, err error) {
	infos, err := WalkDir(dir)
	if err != nil {
		return m, err
	}
	m = make(map[string][]byte)
	size := 0
	for _, info := range infos {
		b, err := ioutil.ReadFile(info.Path)
		if err != nil {
			fmt.Println("readfile error")
			return m, err
		}
		uri := "/" + info.Path
		m[uri] = b
		size += int(info.Size)
	}
	fmt.Println(fmt.Sprintf("serving %v MB of files from memory", size/1e6))
	return m, err
}

func (fs FileServer) mimeFromFilePath(fp string) (string, error) {
	chunks := strings.Split(fp, ".")
	c := len(chunks)
	if c == 0 {
		return "", fmt.Errorf("%v is not a valid filename. it is missing an extension", fp)
	}
	ext := chunks[c-1]
	mime, ok := map[string]string{
		"pdf": "application/pdf",
		"png": "image/png",
		"css": "text/css",
	}[ext]
	if !ok {
		return "", fmt.Errorf("extension %v for file %v not found", ext, fp)
	}
	return mime, nil
}

func (fs FileServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	mime, err := fs.mimeFromFilePath(r.RequestURI)
	if err != nil {
		w.WriteHeader(403)
		if _, err = w.Write([]byte(err.Error())); err != nil {
			fmt.Println("error writing response: ", err)
		}
		return
	}
	w.Header().Set("Content-Type", mime)
	data, ok := fs.files[r.RequestURI]
	if !ok {
		w.WriteHeader(404)
		fmt.Println(r.RequestURI)
		if _, err = w.Write([]byte("file not found")); err != nil {
			fmt.Println("error writing response: ", err)
		}
		return
	}
	w.WriteHeader(200)
	_, err = w.Write(data)
}
