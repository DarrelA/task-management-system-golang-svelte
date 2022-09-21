<script>
	import MultiSelect from "./MultiSelect.svelte"
	import Axios from "axios"
	import { toast } from '@zerodevx/svelte-toast'
	import { SvelteToast } from '@zerodevx/svelte-toast'


	export let username = "";
	export let password = "";
	export let email = "";
	export let user_group = "";
	export let status = "";
	const color = "danger"
	// export let result = null
	// export let code = null
	let message = "";
	// let value;

	async function handleClick() {
		user_group = user_group.toString()
		const json = {username, password, email, user_group, status}

		try {
			const response = await Axios.post("http://localhost:4000/admin-update-user", json)
			if (response) {
				message = response.data.message
				toast.push(message)
			}
		} catch (error) {
			console.log(error)
		}
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

<div style="text-align:center">
<h1>Admin Update User</h1>
<input type="text" bind:value={username} placeholder="Username" > <br>
<input type="password" bind:value={password} placeholder="Password" > <br>
<input type="email" bind:value={email} placeholder="Email" > <br>
<!-- <input type="text" bind:value={user_group} placeholder="User Group" > <br> -->
<MultiSelect bind:value={user_group} placeholder="User Group(s)">
	<option value="Project Lead" >Project Lead</option>
	<option value="Project Manager">Project Manager</option>
	<option value="Team Member">Team Member</option>
</MultiSelect> <br>
<select bind:value={status} placeholder="Status">
	<option value="Active" >Active</option>
	<option value="Inactive">Inactive</option>
</select> <br>

<button on:click="{handleClick}">Update User</button>
</div>

<p>Message: {message}</p>
<!-- <p>Code: {code}</p> -->

<ul>
	<li>Username: {username}</li>
	<li>Password: {password}</li>
	<li>Email: {email}</li>
	<li>UserGroup: {user_group}</li>
	<li>Status: {status}</li>
</ul>