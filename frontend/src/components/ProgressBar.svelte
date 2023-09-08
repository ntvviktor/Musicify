<script>
  /**
   * @type {number}
   */
  export let value = 50;
  export let range = 0;
  /**
   * @type {number}
   */
   // @ts-ignore
   export let max;
   export let min = 0;
   export let step = 1;
   // @ts-ignore
   /**
   * @type {number}
   */
   let percentageValue;
  // @ts-ignore
  $: percentageValue = percentOf(value);

 // @ts-ignore
   $: percentOf = function (value) {
    value / max  * 100
  }

</script>

<div class="container" style="--buffered-width: {percentageValue}%; --seek-before-width: {range}%">
  <input type="range" id="seek-slider" min={min} max={max} step={step}/>
</div>

