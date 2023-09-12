# SpotifyClone Person Project

### Purpose
This project aims for simulate some feature of streaming in Spotify app, which is interesting and admireable
due to its speed. 

### How it works
This project use svelte for the frontend and Golang for backend
1. Create a AAC audio file using Ffmpeg command, use library libfdk_aac from Fraunhofer FDK AAC for higher quality encoder.
2. Use [Mp4dash](http://www.bento4.com/documentation/mp4dash/) to fragment the file into byte-range within the mp4 container.
3. Then in the frontend, fetch chunks of audio data in byte-range specified at the request headers.

### Inspiration
How Spotify streams most of their songs is extremely inspirational to me. Hence I want to learn and simulate this particular
function to discover.
