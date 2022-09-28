<script>
  import { navigate } from "svelte-routing";
  import AdminNavbar from "../Navbar/IsLoggedInAdmin.svelte";
  import UserNavbar from "../Navbar/IsLoggedInUser.svelte";
  import Login from "./Login.svelte";

  const isAdmin = localStorage.getItem("isAdmin");
  let username = localStorage.getItem("username");
  // if (isAdmin === null) {
  //   navigate('/')
  // }
</script>

{#if isAdmin === "true"}
  <AdminNavbar />
{:else if isAdmin === "false"}
  <UserNavbar />
  insert application component here for user
  <!-- {:else}
<Login/> -->
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
