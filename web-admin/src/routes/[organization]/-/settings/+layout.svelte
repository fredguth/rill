<!-- ORG SETTINGS -->

<script lang="ts">
  import { page } from "$app/stores";
  import LeftNav from "@rilldata/web-admin/components/nav/LeftNav.svelte";
  import { isEnterprisePlan } from "@rilldata/web-admin/features/billing/plans/utils";
  import type { PageData } from "./$types";
  import ContentContainer from "@rilldata/web-admin/components/layout/ContentContainer.svelte";

  export let data: PageData;

  $: ({ subscription, neverSubscribed, billingPortalUrl } = data);

  $: organization = $page.params.organization;
  $: basePage = `/${organization}/-/settings`;
  $: onEnterprisePlan =
    subscription?.plan?.name && isEnterprisePlan(subscription.plan.name);
  $: hideBillingSettings = neverSubscribed;
  $: hideUsageSettings = onEnterprisePlan || !billingPortalUrl;

  $: navItems = [
    { label: "General", route: "", hasPermission: true },
    {
      label: "Billing",
      route: "/billing",
      hasPermission: !hideBillingSettings,
    },
    {
      label: "Usage",
      route: "/usage",
      hasPermission: !hideBillingSettings && !hideUsageSettings,
    },
  ];
</script>

<ContentContainer title="Organization settings" maxWidth={1100}>
  <div class="container flex-col md:flex-row">
    <LeftNav
      {basePage}
      baseRoute="/[organization]/-/settings"
      {navItems}
      minWidth="180px"
    />
    <div class="flex flex-col gap-y-6 w-full">
      <slot />
    </div>
  </div>
</ContentContainer>

<style lang="postcss">
  .container {
    @apply flex pt-6 gap-6 max-w-full overflow-hidden;
  }
</style>
