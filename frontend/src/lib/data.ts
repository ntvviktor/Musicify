import { writable } from "svelte/store"

export const audioPlayer = writable();
export const status = writable('default');
export const isPlaying = writable(false);

export const trackList = writable([
    {
        title:"I Love You So",
        artist:"The Walters",
        urlMetadata:"http://127.0.0.1:3000/metadata"
    }
])