<script>
  import axios from "axios";

  import Navbar from "../Navbar/IsLoggedInAdmin.svelte";
  import AccountsTable from "./AccountsTable.svelte";
  import AdminCreateUserForm from "./AdminCreateUserForm.svelte";
</script>

<Navbar />
<br />
<AdminCreateUserForm />
<AccountsTable />

<style>
</style>
