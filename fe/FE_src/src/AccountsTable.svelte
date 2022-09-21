<script>
	import Axios from "axios"
    import { SvelteToast } from '@zerodevx/svelte-toast'
	import {Table} from "sveltestrap"
	import {Button,Modal,ModalBody,ModalFooter,ModalHeader} from 'sveltestrap';
	import AdminUpdateUser from "./AdminUpdateUser.svelte"

	let usersData = [];
	let open = false;
	let selectedUsername = ""
	let selectedUserData;

    $: getUsers()
	
	async function getUsers() {
		try {
			const response = await Axios.get("http://localhost:4000/get-users")
			if (response) {
				usersData = response.data
			}
		} catch (error) {
			console.log(error)
		}
	}

	async function getSelectedUserData() {
		try {
			const response = await Axios.post("http://localhost:4000/get-selected-users", selectedUsername)
			console.log(selectedUsername)
			if (response) {
				selectedUserData = response.data
				console.log(response.data)
			}
		} catch (error) {

		}
	}

	function editUserData(username) {
		open = !open
		selectedUsername = username
		getSelectedUserData()
	}
</script>

<main>
	<SvelteToast />
</main>

<style>
	h1 {
		color: purple;
	}
</style>

<h1>Accounts Table</h1>
<Table bordered>
	<thead>
		<tr>
		  <th>Username</th>
		  <th>Email</th>
		  <th>User Group(s)</th>
		  <th>Status</th>
		  <th>Edit</th>
		</tr>
	  </thead>
	  <tbody>
		{#each usersData as userData}
		<tr>
		  <th scope="row">{userData.username}</th>
		  <td>{userData.email}</td>
		  <td>{userData.user_group}</td>
		  <td>{userData.status}</td>
		  <td><Button color="primary" on:click={() => editUserData(userData.username)}>Edit User</Button></td>
		</tr>
		{/each}

		<div>
			<Modal isOpen={open} {editUserData}>
			  <ModalHeader {editUserData}>Update User</ModalHeader>
			  <ModalBody>
				<AdminUpdateUser selectedUsername={selectedUsername} />
			  </ModalBody>
			  <ModalFooter>
				<Button color="secondary" on:click={editUserData}>Back</Button>
				<Button color="primary" on:click={editUserData}>Update User</Button>
			  </ModalFooter>
			</Modal>
		  </div>
	  </tbody>
</Table>




