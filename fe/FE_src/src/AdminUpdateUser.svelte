<script>
	import Axios from "axios"
	import { toast } from '@zerodevx/svelte-toast'
	import { SvelteToast } from '@zerodevx/svelte-toast'
	import MultiSelect from "svelte-multiselect"
	import { Button, Form, FormGroup, FormText, Input, Label } from 'sveltestrap';

	let username = "";
	let password = "";
	let email = "";
	let status = "";

	let code = ""
	let message = "";

	const user_groups = ["Project Lead", "Project Manager", "Team Member"]
	let selected = []

	async function handleSubmit() {
		let user_group = selected.toString()
		console.log(user_group)
		const json = {username, password, email, user_group, status}

		try {
			const response = await Axios.post("http://localhost:4000/admin-update-user", json)
			if (response) {
				message = response.data.message
				code = response.data.code
				toast.push(message)
				username = "", password = "", email = "", status = "", selected = []
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
	.disabled {
		background-color: lightgrey;
	}
</style>

<Form>
<FormGroup>
	<Label for="username">Username:</Label>
	<Input type="text" bind:value={username} placeholder="Username" class="disabled" disabled />
</FormGroup>
<FormGroup>
	<Label for="password">Password:</Label>
	<Input type="password" bind:value={password} placeholder="Password" />
</FormGroup>
<FormGroup>
	<Label for="email">Email:</Label>
	<Input type="email" bind:value={email} placeholder="Email" />
</FormGroup>
<FormGroup>
	<Label for="usergroup">User Group(s):</Label>
	<MultiSelect bind:selected options={user_groups} />
</FormGroup>
<FormGroup>
    <Label for="status">Status</Label>
    <Input type="select" bind:value={status} placeholder="Status">
      <option>Inactive</option>
      <option>Active</option>
    </Input>
</FormGroup>
<Button on:submit="{handleSubmit}">Update User</Button>
</Form>