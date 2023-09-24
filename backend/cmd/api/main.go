package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
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
