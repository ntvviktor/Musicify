<script>
  import { isPlaying, audioPlayer } from "$lib/data";
  import { onMount } from "svelte";
  import Controls from "./Controls.svelte";

  let duration = 0;
  let currentTime = 0;
  // @ts-ignore
  let paused = true;
  let volume = 0.5;

  // @ts-ignore
  export let metadata;

  console.log(metadata);

  /**
   * @type {MediaSource}
   */
  var mediaSource;
  /**
   * @type {SourceBuffer}
   */
  var sourceBuffer_audio;
  let segNum = 0;
  var maxSegNum = metadata.SegmentURL.length;
  /**
   * @type {HTMLAudioElement}
   */
  var audio;
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
    // sourceBuffer_audio.addEventListener('updateend', loadSegment);
    sourceBuffer_audio.addEventListener("updateend", loadSegment);
  }

  function loadSegment() {
    if (segNum < maxSegNum) {
      getAudioSegment();
      segNum++;
    }
  }

  // @ts-ignore
  async function getAudioSegment() {
    //@ts-ignore
    let byterange = metadata.SegmentURL[segNum].MediaRange;
    const res = await fetch("http://127.0.0.1:3000/music/song1", {
      headers: {
        range: `bytes=${byterange}`,
      },
    });
    const buf = await res.arrayBuffer();
    // @ts-ignore
    sourceBuffer_audio.appendBuffer(buf);
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
  <Controls {currentTime} {duration} {audio}/>
  {/if}
</div>
