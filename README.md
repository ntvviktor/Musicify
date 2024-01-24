# Musicify Person Project

### Purpose
This project aims for simulate some feature of streaming in Spotify app: stream music with http byte-range request.
### How it works
The project uses Svelte for the frontend and Golang for backend:
1. Transcode audio using Ffmpeg command, use library [libfdk_aac](https://en.wikipedia.org/wiki/Fraunhofer_FDK_AAC) for higher quality encoder.
2. Use [Mp4dash](http://www.bento4.com/documentation/mp4dash/) to fragment the file into byte-range within the mp4 container.
3. Then in the frontend, fetch chunks of audio data in byte-range specified at the request headers.

### Inspiration
Credit inspiration to this [github repo](https://github.com/nickdesaulniers/netfix/blob/gh-pages/demo/).

### Demo
Clone this project
```
cd backend
go run cmd/api/main.go

cd frontend
npm install
npm run dev
go to localhost
```

Notes TODO: seeking audio feature can becomes buggy and might not work properly.

