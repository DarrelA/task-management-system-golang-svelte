<script>
  import AdminNavbar from "../Navbar/IsLoggedInAdmin.svelte";
  import UserNavbar from "../Navbar/IsLoggedInUser.svelte";
  import CreateTask from "./CreateTask.svelte"

  const isAdmin = localStorage.getItem("isAdmin");
  let username = localStorage.getItem("username");
</script>

{#if isAdmin === "true"}
  <AdminNavbar />
{:else if isAdmin === "false"}
  <UserNavbar />
{/if}

<div class="masthead">
  <h2>Welcome {username} &#x1F642;</h2>
  <p>Do you have any tasks to complete today?</p>
</div>

<CreateTask />

<style>
  .masthead {
    border-radius: 20px;
    background-color: grey;
    height: 15%;
    width: 75%;
    margin: 20px auto;
    display: flex;
    flex-direction: column;
    justify-content: center;
    align-items: center;
    text-align: center;
    color: #fffbf0;
  }
</style>
