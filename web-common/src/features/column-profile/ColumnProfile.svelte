<script lang="ts">
  import { COLUMN_PROFILE_CONFIG } from "@rilldata/web-common/layout/config";
  import {
    createQueryServiceTableColumns,
    createQueryServiceTableRows,
  } from "@rilldata/web-common/runtime-client";
  import { onMount } from "svelte";
  import { runtime } from "../../runtime-client/runtime-store";
  import { getColumnType } from "./column-types";
  import { getSummaries } from "./queries";
  import { defaultSort, sortByName, sortByNullity } from "./utils";
  import { keepPreviousData } from "@tanstack/svelte-query";

  export let connector: string;
  export let database: string;
  export let databaseSchema: string;
  export let objectName: string;
  export let containerWidth = 0;
  export let indentLevel = 0;

  let mode = "summaries";
  let container;

  $: ({ instanceId } = $runtime);

  onMount(() => {
    const observer = new ResizeObserver(() => {
      containerWidth = container?.clientWidth ?? 0;
    });
    observer.observe(container);
    return () => observer.unobserve(container);
  });

  $: profileColumns = createQueryServiceTableColumns(
    instanceId,
    objectName,
    {
      connector,
      database,
      databaseSchema,
    },
    { query: { placeholderData: keepPreviousData } },
  );

  /** get single example */
  $: exampleValue = createQueryServiceTableRows(instanceId, objectName, {
    connector,
    database,
    databaseSchema,
    limit: 1,
  });

  $: nestedColumnProfileQuery = getSummaries(
    instanceId,
    connector,
    database,
    databaseSchema,
    objectName,
    $profileColumns,
  );

  $: profile = $nestedColumnProfileQuery;
  let sortedProfile;
  const sortByOriginalOrder = null;

  let sortMethod = defaultSort;
  $: if (profile?.length && sortMethod !== sortByOriginalOrder) {
    sortedProfile = [...profile].sort(sortMethod);
  } else {
    sortedProfile = profile;
  }

  /** classes for the native select element */
  export const native_select =
    "rounded border border-6 border-transparent hover:bg-gray-200";
</script>

<!-- pl-16 -->
<div
  bind:this={container}
  class="pl-{indentLevel === 1
    ? '10'
    : '4'} pr-5 pb-2 flex justify-between text-gray-500 pt-1"
  class:flex-col={containerWidth < 325}
>
  <select bind:value={sortMethod} class={native_select} style:font-size="11px">
    <option value={sortByOriginalOrder}>show original order</option>
    <option value={defaultSort}>sort by type</option>
    <option value={sortByNullity}>sort by null %</option>
    <option value={sortByName}>sort by name</option>
  </select>
  <select
    bind:value={mode}
    class={native_select}
    class:hidden={containerWidth < 325}
    style:font-size="11px"
    style:transform="translateX(4px)"
  >
    <option value="summaries">show summary&nbsp;</option>
    <option value="example">show example</option>
    <option value="hide">hide reference</option>
  </select>
</div>

<div>
  {#if sortedProfile && exampleValue}
    {#each sortedProfile as column (column.name)}
      {@const hideRight = containerWidth < COLUMN_PROFILE_CONFIG.hideRight}
      {@const hideNullPercentage =
        containerWidth < COLUMN_PROFILE_CONFIG.hideNullPercentage}
      {@const compact =
        containerWidth < COLUMN_PROFILE_CONFIG.compactBreakpoint}
      <svelte:component
        this={getColumnType(column.type)}
        type={column.type}
        {connector}
        {database}
        {databaseSchema}
        {objectName}
        columnName={column.name}
        example={$exampleValue?.data?.data?.[0]?.[column.name]}
        {mode}
        {hideRight}
        {hideNullPercentage}
        {compact}
        enableProfiling={!$profileColumns?.isFetching}
      />
    {/each}
  {/if}
</div>
