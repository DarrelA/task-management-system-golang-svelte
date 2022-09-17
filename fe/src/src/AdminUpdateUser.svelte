<script>
	import MultiSelect from "./MultiSelect.svelte"

	export let username = "";
	export let password = "";
	export let email = "";
	export let user_group = "";
	export let status = "";
	export let result = null
	export let code = null
	let value;

	async function handleClick() {
		user_group = user_group.toString()
		await fetch("http://localhost:4000/admin-update-user", {
			mode: "cors",
			method: "POST",
			cache: "no-cache",
			headers: {
				"Content-Type": "application/json",
				"Access-Control-Allow-Origin": "*"
			},
			body: JSON.stringify({
				username, password, email, user_group, status
			})
		})
		.then(async response => {
			if (response) {
				// console.log(response)
				// result = response.status
				const json = await response.json()
				// result = JSON.stringify(json)
				result = json.message
				code = json.code
			}
		})
		.catch(error => {
			console.error(error)
		})
	}
</script>

<style>
	h1 {
		color: purple;
	}
</style>

<div style="text-align:center">
<h1>Admin Update User</h1>
<input type="text" bind:value={username} placeholder="Username" > <br>
<input type="text" bind:value={password} placeholder="Password" > <br>
<input type="text" bind:value={email} placeholder="Email" > <br>
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

<p>Message: {result}</p>
<p>Code: {code}</p>

<ul>
	<li>Username: {username}</li>
	<li>Password: {password}</li>
	<li>Email: {email}</li>
	<li>UserGroup: {user_group}</li>
	<li>Status: {status}</li>
</ul>