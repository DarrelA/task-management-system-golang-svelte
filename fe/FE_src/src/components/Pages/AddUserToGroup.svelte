<script>
   import axios from 'axios'
   import { toast } from '@zerodevx/svelte-toast'
   import { SvelteToast } from '@zerodevx/svelte-toast'
   import { Button, Badge, Form, FormGroup, FormText, Input, Label } from 'sveltestrap'
   import Navbar from '../Navbar/IsLoggedInAdmin.svelte'

   let message = ""
	let code = ""
   export let username;
	export let user_group = [];

	let groupsArray = [];
	let selected = []
	selected.push(...user_group)

   $: getUsers()
	
	async function getUsers() {
		try {
			const response = await axios.get("http://localhost:4000/get-users")
			if (response) {
				usersData = response.data
			}
		} catch (error) {
			console.log(error)
		}
	}

   async function handleSubmit() {
		let user_group = selected.join(",")
		const json = {username, user_group}

		try {
			const response = await axios.post("http://localhost:4000/admin-update-user", json)
			if (response) {
				console.log("POST BACK")
				message = response.data.message
				code = response.data.code
				toast.push(message)
				password = ""
			}
		} catch (error) {
			console.log(error)
		}
	}
 </script>

<form on:submit|preventDefault={handleSubmit}>
   <FormGroup>
      <label for="username">Username</label>
      <Input type="text" bind:value={username} />
   </FormGroup>
	<FormGroup>
		<Label for="usergroup">User Group(s):</Label>
		<MultiSelect bind:selected options={groupsArray} />
	</FormGroup>
   
</form>
