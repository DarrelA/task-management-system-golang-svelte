<script>
  import AdminNavbar from "../Navbar/IsLoggedInAdmin.svelte";
  import UserNavbar from "../Navbar/IsLoggedInUser.svelte";

  const isAdmin = localStorage.getItem("isAdmin");
  let username = localStorage.getItem("username");
</script>

{#if isAdmin === "true"}
  <AdminNavbar />
{:else if isAdmin === "false"}
  <UserNavbar />
  insert application component here for user
{/if}

<div class="masthead">
  <h2>Welcome {username} &#x1F642;</h2>
  <p>Do you have any tasks to complete today?</p>
</div>

<style>
  .masthead {
    border-radius: 20px;
    background-color: grey;
    height: 15%;
    width: 75%;
    margin: 0 auto;
    display: flex;
    flex-direction: column;
    justify-content: center;
    align-items: center;
    text-align: center;
    color: #fffbf0;
  }
</style>
