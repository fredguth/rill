<script lang="ts">
  import { formatDataType } from "../../lib/formatters";
  import Base from "./Base.svelte";
  export let isNull = false;
  export let inTable = false;
  export const textAlign = "text-right";
  export let customStyle = "";
  export let dark: boolean;
  export let type: string;
  export let value;
  export let truncate: boolean;
  export let color = "";
</script>

<Base
  {truncate}
  {isNull}
  classes="ui-copy-number font-semibold {customStyle} {dark
    ? 'text-slate-200'
    : 'text-slate-800'} {inTable && 'text-left'}"
  {dark}
  {color}
>
  <slot name="value">
    {formatDataType(value, type)}
  </slot>
</Base>
