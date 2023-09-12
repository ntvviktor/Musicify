package main

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

/*
Require Ffmpeg multimedia framework with custom libfdk_aac library, Bento4 package installed
1. Ffmpeg command to convert flac to aac 256kbs and put into mp4 container
ffmpeg -i music.flac -c:a libfdk_aac -b:a 256k output.aac

2. Fragmented the audio file into each 10-second chunk of audio
mp4fragment --fragment-duration 10000 output.mp4 output_frag.mp4

3. Output the mpd file for metadata and byte range information
mp4dash --no-split --use-segment-list --no-media --mpd-name=[ID] output_frag.mp4
*/

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/music/song1", handleSong).
		Methods(http.MethodGet, http.MethodPost, http.MethodOptions)
	router.HandleFunc("/music/song1/init", handleInit).
		Methods(http.MethodGet, http.MethodPost, http.MethodOptions)
	router.HandleFunc("/metadata", handleMetadata).
		Methods(http.MethodGet, http.MethodPost, http.MethodOptions)
	router.Use(mux.CORSMethodMiddleware(router))
	server := &http.Server{
		Handler: router,
		// listen from any IP address (for WSL port forwarding purposes)
		Addr: "0.0.0.0:3000",
		// Good practice: enforce timeouts for servers you create!
		// WriteTimeout: 15 * time.Second,
		// ReadTimeout:  15 * time.Second,
	}
	log.Fatal(server.ListenAndServe())
}

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

type MPD struct {
	XMLName                   xml.Name `xml:"MPD"`
	Xmlns                     string   `xml:"xmlns,attr"`
	Profiles                  string   `xml:"profiles,attr"`
	MinBufferTime             string   `xml:"minBufferTime,attr"`
	MediaPresentationDuration string   `xml:"mediaPresentationDuration,attr"`
	Type                      string   `xml:"type,attr"`
	Period                    struct {
		AdaptationSet struct {
			MimeType         string `xml:"mimeType,attr"`
			StartWithSAP     string `xml:"startWithSAP,attr"`
			SegmentAlignment string `xml:"segmentAlignment,attr"`
			Representation   struct {
				ID                        string `xml:"id,attr"`
				Codecs                    string `xml:"codecs,attr"`
				Bandwidth                 string `xml:"bandwidth,attr"`
				AudioSamplingRate         string `xml:"audioSamplingRate,attr"`
				AudioChannelConfiguration struct {
					SchemeIdUri string `xml:"schemeIdUri,attr"`
					Value       string `xml:"value,attr"`
				} `xml:"AudioChannelConfiguration"`
				SegmentList SegmentList `xml:"SegmentList"`
			} `xml:"Representation"`
		} `xml:"AdaptationSet"`
	} `xml:"Period"`
}

type SegmentList struct {
	Timescale      string `xml:"timescale,attr"`
	Duration       string `xml:"duration,attr"`
	Initialization struct {
		SourceURL string `xml:"sourceURL,attr"`
		Range     string `xml:"range,attr"`
	} `xml:"Initialization"`
	SegmentURL []struct {
		MediaRange string `xml:"mediaRange,attr"`
	} `xml:"SegmentURL"`
}

func ParseInXML(dirname string) *SegmentList {
	xmlFile, err := os.Open(dirname)
	if err != nil {
		fmt.Println(err)
	}

	defer xmlFile.Close()

	var Metadata MPD
	err = xml.NewDecoder(xmlFile).Decode(&Metadata)
	if err != nil {
		log.Fatal(err)
	}
	SegmentList := Metadata.Period.AdaptationSet.Representation.SegmentList

	return &SegmentList
}
