<script lang="ts">
  import IconButton from "@rilldata/web-common/components/button/IconButton.svelte";
  import * as DropdownMenu from "@rilldata/web-common/components/dropdown-menu";
  import ThreeDot from "@rilldata/web-common/components/icons/ThreeDot.svelte";
  import { Trash2Icon, Pencil } from "lucide-svelte";
  import DeleteUserGroupConfirmDialog from "./DeleteUserGroupConfirmDialog.svelte";
  import EditUserGroupDialog from "./EditUserGroupDialog.svelte";
  import type { V1OrganizationMemberUser } from "@rilldata/web-admin/client";

  export let groupName: string;
  export let currentUserEmail: string;
  export let searchUsersList: V1OrganizationMemberUser[];

  let isDropdownOpen = false;
  let isDeleteConfirmOpen = false;
  let isEditDialogOpen = false;
</script>

<!-- Managed groups cannot be deleted or edited -->
<DropdownMenu.Root bind:open={isDropdownOpen}>
  <DropdownMenu.Trigger class="flex-none">
    <IconButton rounded active={isDropdownOpen}>
      <ThreeDot size="16px" />
    </IconButton>
  </DropdownMenu.Trigger>
  <DropdownMenu.Content align="start">
    <DropdownMenu.Item
      class="font-normal flex items-center"
      on:click={() => {
        isEditDialogOpen = true;
      }}
    >
      <Pencil size="12px" />
      <span class="ml-2">Edit</span>
    </DropdownMenu.Item>
    <DropdownMenu.Item
      class="font-normal flex items-center"
      type="destructive"
      on:click={() => {
        isDeleteConfirmOpen = true;
      }}
    >
      <Trash2Icon size="12px" />
      <span class="ml-2">Delete</span>
    </DropdownMenu.Item>
  </DropdownMenu.Content>
</DropdownMenu.Root>

<DeleteUserGroupConfirmDialog bind:open={isDeleteConfirmOpen} {groupName} />

<EditUserGroupDialog
  bind:open={isEditDialogOpen}
  {groupName}
  {currentUserEmail}
  organizationUsers={searchUsersList}
/>
