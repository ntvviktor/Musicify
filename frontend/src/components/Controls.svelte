<script>
  import PlayButton from "./PlayButton.svelte";
  import { format } from "$lib/utilities";
  import { audioPlayer } from "$lib/data";
  /**
   * @type {HTMLAudioElement}
   */
  export let audio;
  export let duration = 0;
  export let currentTime = 0;
  // let raf = null;
  /**
   * @type {HTMLInputElement}
   */
  let seekSlider;
  /**
   * @type {HTMLDivElement}
   */
  let audioPlayerContainer;
  const whilePlaying = () => {
    seekSlider.value = String(Math.floor(audio.currentTime));
    audioPlayerContainer.style.setProperty(
      "--buffered-width",
      `${(Number(seekSlider.value) / Number(seekSlider.max)) * 100}%`
    );
    // raf = requestAnimationFrame(whilePlaying);
  };

  const showRangeProgress = (/** @type {HTMLInputElement} */ rangeInput) => {
    if (rangeInput === seekSlider) {
      audioPlayerContainer.style.setProperty('--buffered-width', Number(rangeInput.value) /Number(rangeInput.max) * 100 + '%')
    }
  };
  /**
   * @param {{ target: any; }} e
   */
  function movePosition(e) {
    showRangeProgress(e.target);
    // if (!audio.paused) {
    //   cancelAnimationFrame(raf);
    // }
  }

  function updatePosition() {
    // @ts-ignore
    $audioPlayer.currentTime = seekSlider.value;
    // if (!audio.paused) {
    //   requestAnimationFrame(whilePlaying);
    // }
  }

  audio.addEventListener("progress", whilePlaying);
</script>

<div class="control">
  <span>I Love You So</span>
  <span >The Walters</span>
  <PlayButton />
  <div class="progress">
    <span class="time">{format(currentTime)}</span>
    <div class="progress-slider">
      <div bind:this={audioPlayerContainer} class="container">
        <input
          bind:this={seekSlider}
          type="range"
          id="seek-slider"
          max={Math.floor(duration)}
          on:input={movePosition}
          on:change={updatePosition}
          step=0.01
        />
      </div>
    </div>
    <span class="duration">{format(duration)}</span>
  </div>
</div>

<style>
  .control{
    background-color: #2c303333;
    padding: 2.5rem;
    border-radius: 15px;
    display: flex;
    flex-direction: column;
    align-items: center;
    justify-content: center;
    gap: 11px;
    box-shadow: rgba(0, 0, 0, 0.16) 0px 3px 6px, rgba(0, 0, 0, 0.23) 0px 3px 6px;
  }
  .progress {
    display: flex;
    flex-direction: row;
    align-items: center;
    justify-content:center;
  }
  .container {
    width: 80%;
    margin: 0 12px;
    display: flex;
    flex-direction: column;
    align-items: center;
    justify-content: center;
  }
  input[type="range"] {
    position: relative;
    -webkit-appearance: none;
    width: 100%;
    margin: 0;
    padding: 0;
    height: 19px;
    margin: 0;
    float: left;
    outline: none;
    background-color: transparent;
  }
  input[type="range"]::-webkit-slider-runnable-track {
    width: 100%;
    height: 4px;
    cursor: pointer;
    background: linear-gradient(
      to right,
      rgba(0, 125, 181, 0.6) var(--buffered-width),
      rgba(0, 125, 181, 0.2) var(--buffered-width)
    );
    border-radius: 20%;
  }
  input[type="range"]::before {
    position: absolute;
    content: "";
    top: 8px;
    left: 0;
    width: 100%;
    height: 3px;
    background-color: #1a1a1a92;
    cursor: pointer;
  }
  input[type="range"]::-webkit-slider-thumb {
    position: relative;
    -webkit-appearance: none;
    box-sizing: content-box;
    border: 1px solid #1a1a1a92;
    height: 15px;
    width: 15px;
    border-radius: 50%;
    background-color: #fff;
    cursor: pointer;
    margin: -7px 0 0 0;
  }
  input[type="range"]:active::-webkit-slider-thumb {
    transform: scale(1.2);
    background: #1a1a1a92;
  }
  input[type="range"]::-moz-range-track {
    width: 100%;
    height: 4px;
    cursor: pointer;
    background: linear-gradient(
      to right,
      rgba(0, 125, 181, 0.6) var(--buffered-width),
      rgba(0, 125, 181, 0.2) var(--buffered-width)
    );
  }
  input[type="range"]::-moz-range-progress {
    background-color: #1a1a1a92;
  }
  input[type="range"]::-moz-focus-outer {
    border: 0;
  }
  input[type="range"]::-moz-range-thumb {
    box-sizing: content-box;
    border: 1px solid #1a1a1a92;
    height: 15px;
    width: 15px;
    border-radius: 50%;
    background-color: #fff;
    cursor: pointer;
  }
  input[type="range"]:active::-moz-range-thumb {
    transform: scale(1.2);
    background: #1a1a1a92;
  }
  input[type="range"]::-ms-track {
    width: 100%;
    height: 3px;
    cursor: pointer;
    background: transparent;
    border: solid transparent;
    color: transparent;
  }
  input[type="range"]::-ms-fill-lower {
    background-color: #1a1a1a92;
  }
  input[type="range"]::-ms-fill-upper {
    background: linear-gradient(
      to right,
      rgba(0, 125, 181, 0.6) var(--buffered-width),
      rgba(0, 125, 181, 0.2) var(--buffered-width)
    );
  }
  input[type="range"]::-ms-thumb {
    box-sizing: content-box;
    border: 1px solid #1a1a1a92;
    height: 15px;
    width: 15px;
    border-radius: 50%;
    background-color: #fff;
    cursor: pointer;
  }
  input[type="range"]:active::-ms-thumb {
    transform: scale(1.2);
    background: #1a1a1a92;
  }
</style>
