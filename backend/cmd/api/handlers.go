package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func handleInit(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "range")
	if r.Method == http.MethodOptions {
		return
	}
	list := ParseInXML("output/stream.mpd")
	init := list.Initialization.Range
	file, err := os.Open("music_frag.mp4")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	defer file.Close()
	fileInfo, _ := file.Stat()
	byteRange := strings.Split(init, "-")
	end, _ := strconv.Atoi(byteRange[1])
	start, _ := strconv.Atoi(byteRange[0])
	fmt.Println(start)
	fmt.Println(end)
	buffer := make([]byte, end-start+1)
	n, err := file.ReadAt(buffer, int64(start))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/octet-stream")
	contentRange := fmt.Sprintf("bytes %d-%d/%d", start, end, fileInfo.Size())
	w.Header().Set("Content-Range", contentRange)
	io.CopyN(w, bytes.NewReader(buffer), int64(n))
}

func handleMetadata(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "range")
	if r.Method == http.MethodOptions {
		return
	}
	list := ParseInXML("output/stream.mpd")
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(*list)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// TODO: Before use segmented audio file, need to transcode AAC audio into MP4 container

func handleSong(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "range")

	if r.Method == http.MethodOptions {
		return
	}
	rangeHeader := r.Header.Get("range")
	file, err := os.Open("music_frag.mp4")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	defer file.Close()
	fileInfo, _ := file.Stat()
	rangeHeader = strings.Trim(rangeHeader, "bytes=")
	byteRange := strings.Split(rangeHeader, "-")
	start, _ := strconv.Atoi(byteRange[0])
	end, _ := strconv.Atoi(byteRange[1])
	fmt.Println(start)
	fmt.Println(end)
	buffer := make([]byte, end-start+1)
	n, err := file.ReadAt(buffer, int64(start))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/octet-stream")
	contentRange := fmt.Sprintf("bytes %d-%d/%d", start, end, fileInfo.Size())
	w.Header().Set("Content-Range", contentRange)
	io.CopyN(w, bytes.NewReader(buffer), int64(n))
}
