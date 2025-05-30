<script lang="ts">
  import { page } from "$app/stores";
  import CanvasIcon from "@rilldata/web-common/components/icons/CanvasIcon.svelte";
  import ExploreIcon from "@rilldata/web-common/components/icons/ExploreIcon.svelte";
  import Tag from "@rilldata/web-common/components/tag/Tag.svelte";
  import Tooltip from "@rilldata/web-common/components/tooltip/Tooltip.svelte";
  import TooltipContent from "@rilldata/web-common/components/tooltip/TooltipContent.svelte";
  import { timeAgo } from "./utils";

  export let name: string;
  export let title: string;
  export let lastRefreshed: string;
  export let description: string;
  export let error: string;
  export let isMetricsExplorer: boolean;
  export let isEmbedded: boolean;

  $: organization = $page.params.organization;
  $: project = $page.params.project;

  $: lastRefreshedDate = lastRefreshed ? new Date(lastRefreshed) : null;

  $: href = isEmbedded
    ? undefined
    : isMetricsExplorer
      ? `/${organization}/${project}/explore/${name}`
      : `/${organization}/${project}/canvas/${name}`;
</script>

<svelte:element
  this={isEmbedded ? "button" : "a"}
  class="flex flex-col gap-y-0.5 group px-4 py-2 w-full"
  {href}
  role={isEmbedded ? "button" : "link"}
>
  <div class="flex gap-x-2 items-center">
    {#if isMetricsExplorer}
      <ExploreIcon size={"14px"} className="text-slate-500" />
    {:else}
      <CanvasIcon size={"14px"} className="text-slate-500" />
    {/if}
    <div
      class="text-gray-700 text-sm font-semibold group-hover:text-primary-600"
    >
      {title !== "" ? title : name}
    </div>
    {#if error !== ""}
      <Tag color="red">Error</Tag>
    {/if}
  </div>
  <div class="pl-[22px] flex gap-x-1 text-gray-500 text-xs font-normal">
    <span class="truncate">{name}</span>
    {#if lastRefreshedDate}
      <span>•</span>
      <Tooltip distance={8}>
        <span class="shrink-0">Last refreshed {timeAgo(lastRefreshedDate)}</span
        >
        <TooltipContent slot="tooltip-content">
          {lastRefreshedDate.toLocaleString()}
        </TooltipContent>
      </Tooltip>
    {/if}
    {#if description}
      <span>•</span>
      <span class="line-clamp-1">{description}</span>
    {/if}
  </div>
</svelte:element>
