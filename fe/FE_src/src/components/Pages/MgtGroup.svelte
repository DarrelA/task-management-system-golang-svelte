<script>
  import axios from 'axios'
  import { toast } from '@zerodevx/svelte-toast'
  import Navbar from '../Navbar/IsLoggedInAdmin.svelte'
  import AdminUpdateUser from './AdminUpdateUser.svelte';
  import AddUserToGroup from './AddUserToGroup.svelte';
  import AdminUpdateUserForm from './AdminUpdateUserForm.svelte';

	let groupname = ""
  let message = ""

  async function CreateGroup(){
    const json = {user_group: groupname}
    console.log(json)
    try{
      const response = await axios.post("http://localhost:4000/admin-create-group", json)
      console.log(response.data)
      if (response.data)
      {
        message = response.data.message
        toast.push(message)
      }
    } catch (e){
      console.log(e)
    }
  }

</script>

<style>
	input {
		color: purple;
	}
</style>

<Navbar/>
<h3>To be edited</h3>
<input type="text" bind:value={groupname}>

<button on:click={CreateGroup}>Calculate</button>