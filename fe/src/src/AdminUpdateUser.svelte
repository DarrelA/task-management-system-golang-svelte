<script>
	export let username = "";
	export let password = "";
	export let email = "";
	export let user_group = "";
	export let status = "";

	function usernameInput(event) {
		const enteredUsername = event.target.value
		username = enteredUsername
		console.log("Username", enteredUsername)
	}

	function passwordInput(event) {
		const enteredPassword = event.target.value
		console.log("Password", enteredPassword)
	}

	function emailInput(event) {
		const enteredEmail = event.target.value
		console.log("Email", enteredEmail)
	}

	function userGroupInput(event) {
		const enteredUserGroup = event.target.value
		console.log("User Group", enteredUserGroup)
	}

	function statusInput(event) {
		const enteredStatus = event.target.value
		console.log("Status", enteredStatus)
	}

	function handleClick() {
		alert("User updated successfully!")
	}
</script>

<style>
	h1 {
		color: purple;
	}
</style>

<div style="text-align:center">
<h1>Admin Update User</h1>
<input type="text" value={username} placeholder="Username" on:input="{usernameInput}"> <br>
<input type="text" value={password} placeholder="Password" on:input="{passwordInput}"> <br>
<input type="text" value={email} placeholder="Email" on:input="{emailInput}"> <br>
<input type="text" value={user_group} placeholder="User Group" on:input="{userGroupInput}"> <br>
<input type="text" value={status} placeholder="Active / Inactive" on:input="{statusInput}"> <br>
<button on:click="{handleClick}">Update User</button>
</div>