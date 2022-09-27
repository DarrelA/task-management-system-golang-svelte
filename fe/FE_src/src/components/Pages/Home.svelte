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
  <!-- {:else}
<Login/> -->
{/if}

<div class="masthead">
  <h2>Welcome {username} &#x1F642;</h2>
</div>

<style>
  .masthead {
    border-radius: 20px;
    position: absolute;
    background-color: grey;
    height: 30%;
    width: 75%;
    top: 50%;
    left: 50%;
    transform: translate(-50%, -50%);
    z-index: -1;
    display: flex;
    flex-direction: column;
    justify-content: center;
    text-align: center;
  }
</style>
