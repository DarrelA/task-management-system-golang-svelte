<script>
  import axios from 'axios'
  import { toast } from '@zerodevx/svelte-toast'
  import { SvelteToast } from '@zerodevx/svelte-toast'
	import MultiSelect from "svelte-multiselect"
  import {
	  Button, 
		FormGroup,   
		Label,
		DropdownItem,
		DropdownMenu,
		DropdownToggle,
		Dropdown,
		Input
	} from 'sveltestrap'

	let isOpen = false;

  import Navbar from '../Navbar/IsLoggedInAdmin.svelte'

  let message = ""
	let code = ""
  export let username = "";
	export let user_group = [];
  
	let groupsArray = [];
	let userArray = [];
	let selected = []
	selected.push(...user_group)

	async function handleSubmit() {
		let user_group = selected.join(",")
		console.log(username)
		username = username.toString()
		const json = {username, user_group}
		console.log(json)

		try {
			const response = await axios.post("http://localhost:4000/add-user-to-group", json, { withCredentials: true })
			if (response) {
				message = response.data.message
				code = response.data.code
				toast.push(message)
			}
		} catch (error) {
			console.log(error)
		}
	}

	async function GetUserData() {
		try {
			const response = await axios.get(
				"http://localhost:4000/get-users"
			);
			console.log(response)
			if (response.data.error) {
				console.log(response.data.error);
			} else if (!response.data.error) {
				response.data.forEach((user) => {
					userArray.push(user.username)
				});

			userArray = userArray
			}
		}	catch (e) {
			console.error(e);
		}
	}
	$: GetUserData();

   async function GetUserGroups() {
      try {
        const response = await axios.get(
          "http://localhost:4000/get-user-groups"
        );

		  console.log(response)
        if (response.data.error) {
          console.log(response.data.error);
        } else if (!response.data.error) {
          const usergroups = response.data;

          usergroups.forEach((group) => {
            groupsArray.push(group);
          });

		  groupsArray = groupsArray
        }
      } catch (e) {
        console.error(e);
      }
    }
    $: GetUserGroups();
 </script>

<main>
	<SvelteToast />
</main>

<style>
</style>

<Navbar />

<form on:submit|preventDefault={handleSubmit}>
	<Label for="username">Username</Label>
	<select name="username" bind:value={username}>
		{#each userArray as username}
		<option>{username}</option>
		{/each}
	</select>

	<!-- <Dropdown {isOpen} toggle={() => (isOpen = !isOpen)}>
		<DropdownToggle tag="div" class="d-inline-block">
			<Input />
		</DropdownToggle>
		<DropdownMenu>
			{#each userArray as username}
			<DropdownItem bind:value={username}>{username}</DropdownItem>
			{/each}
		</DropdownMenu>
	</Dropdown> -->

	<FormGroup>
		<br />
		<Label for="usergroup">User Group(s):</Label>
		<MultiSelect bind:selected options={groupsArray} />
	</FormGroup>
   <Button color="primary" >Add User</Button>
</form>
