<script lang="ts">
  export let classes = "";
  export let isNull = false;
  export let dark = false;
  export let truncate = false;
  export let color = "text-gray-900";

  $: color = dark ? "" : color;
</script>

<span class:truncate class="{classes} {color}">
  {#if isNull}
    <span class="text-gray-400">-</span>
  {:else}
    <slot />
  {/if}
</span>
