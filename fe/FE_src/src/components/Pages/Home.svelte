<script>
  import { navigate } from 'svelte-routing';
  import AdminNavbar from '../Navbar/IsLoggedInAdmin.svelte'
  import UserNavbar from '../Navbar/IsLoggedInUser.svelte'
  import Login from './Login.svelte';

  const isAdmin = localStorage.getItem("isAdmin");
  // if (isAdmin === null) {
  //   navigate('/')
  // }
</script>

{#if isAdmin==="true"} 
<AdminNavbar/>
{:else if isAdmin==="false"}
<UserNavbar/>
<!-- {:else}
<Login/> -->
{/if}

<h3>Welcome to TMS</h3>
