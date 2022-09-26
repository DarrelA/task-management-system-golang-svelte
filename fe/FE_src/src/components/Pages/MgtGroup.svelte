<script>
  import Navbar from "../Navbar/IsLoggedInAdmin.svelte";
  import AddUserToGroup from "./AddUserToGroup.svelte";
  import TableGroup from "./TableGroup.svelte";
</script>

<Navbar />

<div>
  <br />

  <h3>Group Management</h3>
  <br />
  <!-- Table with Modal: Add Group -->
  <TableGroup />

  <!-- Modal: Add User To Group  -->
  <div class="add-user-to-group">
    <AddUserToGroup />
  </div>
</div>

<style>
  .add-user-to-group {
    margin-top: 5%;
  }
</style>
