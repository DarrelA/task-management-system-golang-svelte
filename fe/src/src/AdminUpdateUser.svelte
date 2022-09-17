<script>
	export let username = "";
	export let password = "";
	export let email = "";
	export let user_group = "";
	export let status = "";
	export let result = null

	async function handleClick() {
		await fetch("http://localhost:4000/admin-update-user", {
			mode: "no-cors",
			method: "POST",
			headers: {
				"Content-Type": "application/json"
			},
			body: JSON.stringify({
				username, password, email, user_group, status
			})
		})
		.then(response => {
			// response.json()
			console.log(response)
		})
		.catch(error => {
			console.error(error)
		})

		

		// const json = await response.json()
		// result = JSON.stringify(json)
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
<input type="text" bind:value={user_group} placeholder="User Group" > <br>
<input type="text" bind:value={status} placeholder="Active / Inactive" > <br>
<button on:click="{handleClick}">Update User</button>
</div>

<p>Result: {result}</p>

<ul>
	<li>Username: {username}</li>
	<li>Password: {password}</li>
	<li>Email: {email}</li>
	<li>UserGroup: {user_group}</li>
	<li>Status: {status}</li>
</ul>