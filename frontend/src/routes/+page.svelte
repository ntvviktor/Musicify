<script>
  import { onMount } from "svelte";
  import { browser } from '$app/environment'
   
  import AudioPlayer from "../components/AudioPlayer.svelte";
  import ProgressBar from "../components/ProgressBar.svelte";
  import PlayButton from "../components/PlayButton.svelte";
  import Controls from "../components/Controls.svelte";

  /**
   * @type {any}
   */
  let metadata;

  onMount(async() => {
    // mediaSource = new MediaSource();
    // startup();
    const res = await fetch("http://127.0.0.1:3000/metadata")
    const content = await res.json();
    metadata = content;
  })
 


</script>

<main class="main">
  {#if metadata}
  <AudioPlayer metadata={metadata}/>
  {/if}
</main>

<style>
  :global(body) {
    box-sizing: border-box;
    margin: 0px;
    padding: 0px;
  }

  .main {
    width: 100%;
    height: 100%;
    display: flex;
    align-items: center;
    justify-content: center;
    flex-direction: column;
  }
</style>