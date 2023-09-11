<script>
  import { isPlaying, audioPlayer } from "$lib/data";
  import { onMount } from "svelte";
  import Controls from "./Controls.svelte";

  let duration = 0;
  let currentTime = 0;
  // @ts-ignore
  let paused = true;
  let volume = 0.5;

  /**
   * @type {{ SegmentURL: string | any[]; Initialization: { Range: any; }; }}
   */
  export let metadata;
  /**
   * @type {MediaSource}
   */
  let mediaSource;
  /**
   * @type {SourceBuffer}
   */
  let sourceBuffer_audio;
  let totalSegments = metadata.SegmentURL.length;
  /**
   * @type {boolean[]}
   */
  let requestedSegments = [];
  let segmentDuration = 0;

  // @ts-ignore
  let assetURL = "http://127.0.0.1:3000/music/song1";

  for (let i = 0; i < totalSegments; ++i) requestedSegments[i] = false;
  /**
   * @type {HTMLAudioElement}
   */
  let audio;
  // @ts-ignore
  function startup() {
    mediaSource.addEventListener("webkitsourceopen", sourceOpen, false);
    mediaSource.addEventListener("sourceopen", sourceOpen, false);
    audio.src = URL.createObjectURL(mediaSource);
  }

  // @ts-ignore
  async function sourceOpen(e) {
    sourceBuffer_audio = mediaSource.addSourceBuffer(
      "audio/mp4; codecs=mp4a.40.2"
    );
    // @ts-ignore
    let byterange = metadata.Initialization.Range;
    const res = await fetch("http://127.0.0.1:3000/music/song1/init", {
      headers: {
        range: `bytes=${byterange}`,
      },
    });
    // @ts-ignore
    const buf = await res.arrayBuffer();
    sourceBuffer_audio.appendBuffer(buf);

    // Add buffer on demand, so the key is no need "updateend" even listener? Mind-blowing
    // fetch the first 10s of the audio
    fetchAudioSegment(0);
    requestedSegments[0] = true;
    audio.addEventListener("timeupdate", checkBuffer);
    audio.addEventListener("canplay", function () {
      segmentDuration = audio.duration / totalSegments;
    });
    audio.addEventListener("seeking", seek);
  }

  // @ts-ignore
  async function fetchAudioSegment(segmentNumber) {
    //@ts-ignore
    let byterange = metadata.SegmentURL[segmentNumber].MediaRange;
    const res = await fetch("http://127.0.0.1:3000/music/song1", {
      headers: {
        range: `bytes=${byterange}`,
      },
    });
    const buf = await res.arrayBuffer();
    const curMode = sourceBuffer_audio.mode;
    console.log(curMode);


    // whatever normally would have called appendBuffer(buffer) can
    // now just call queue.push(buffer) instead
    sourceBuffer_audio.appendBuffer(buf);
  }

  // @ts-ignore
  function checkBuffer(e) {
    let currentSegment = getCurrentSegment();
    if (currentSegment === totalSegments && haveAllSegments()) {
      console.log("last segment", mediaSource.readyState);
      mediaSource.endOfStream();
      audio.removeEventListener("timeupdate", checkBuffer);
      // @ts-ignore
    } else if (shouldFetchNextSegment(currentSegment)) {
      requestedSegments[currentSegment] = true;
      fetchAudioSegment(currentSegment);
      console.log("time to fetch next chunk", audio.currentTime);
    }
  }

  // @ts-ignore
  function seek(e) {
    if (mediaSource.readyState === "open") {
      sourceBuffer_audio.abort();
      let currentSegment = getCurrentSegment();
      requestedSegments[currentSegment-1] = true;
      fetchAudioSegment(currentSegment-1);
      sourceBuffer_audio.addEventListener("updateend", () => {
        if(!requestedSegments[currentSegment]) {

          requestedSegments[currentSegment] = true;
          fetchAudioSegment(currentSegment);
          // audio.currentTime = sourceBuffer_audio.buffered.start(sourceBuffer_audio.buffered.length - 1);
          audio.play();
        }
      })
    } else {
      console.log("seek but not open");
    }
  }

  function getCurrentSegment() {
    return ((audio.currentTime / segmentDuration) | 0) + 1;
  }

  function haveAllSegments() {
    return requestedSegments.every(function (val) {
      return !!val;
    });
  }

  /**
   * @param {number} currentSegment
   */
  function shouldFetchNextSegment(currentSegment) {
    return (
      // @ts-ignore
      audio.currentTime > segmentDuration * currentSegment * 0.8 &&
      !requestedSegments[currentSegment]
    );
  }

  onMount(() => {
    mediaSource = new MediaSource();
    $audioPlayer = audio;
    startup();
  });
</script>

<div>
  <audio
    bind:this={audio}
    bind:duration
    bind:currentTime
    bind:paused
    bind:volume
    on:ended={() => {
      $isPlaying = false;
      currentTime = 0;
    }}
  />
  {#if audio}
    <Controls {currentTime} {duration} {audio} />
  {/if}
</div>

<style>
  audio {
    display: none;
  }
</style>
