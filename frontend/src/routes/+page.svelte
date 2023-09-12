<script>
  import { onMount } from "svelte";
  import AudioPlayer from "../components/AudioPlayer.svelte";
  /**
   * @type {any}
   */
  let metadata;

  onMount(async() => {
    const res = await fetch("http://127.0.0.1:3000/metadata");
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
    display: flex;
    align-items: center;
    justify-content: center;
    flex-direction: column;
    position: absolute;               
    top: 50%;
    left: 50%;
    transform: translate(-50%, -50%);
  }
</style>