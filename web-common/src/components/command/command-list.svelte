<script lang="ts">
  import { cn } from "@rilldata/web-common/lib/shadcn";
  import { Command as CommandPrimitive } from "cmdk-sv";

  // type $$Props = CommandPrimitive.ListProps;

  let className: string | undefined | null = undefined;
  export { className as class };
</script>

<CommandPrimitive.List
  class={cn("max-h-[300px] overflow-y-auto overflow-x-hidden", className)}
  {...$$restProps}
>
  <slot />
</CommandPrimitive.List>
