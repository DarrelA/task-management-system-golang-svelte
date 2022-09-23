<script>
  import axios from "axios";
  import { errorToast, successToast } from "../toast";
  import { Form, FormGroup, FormText, Input, Label } from "sveltestrap";

  import MultiSelect from "svelte-multiselect";

  let message = "";
  let code = "";
  export let username;
  export let user_group = [];

  let groupsArray = [];
  let selected = [];
  selected.push(...user_group);

  $: getUsers();

  async function getUsers() {
    try {
      const response = await axios.get("http://localhost:4000/get-users");
      if (response) {
        let usersData = response.data;
        console.log(usersData);
      }
    } catch (error) {
      console.log(error);
    }
  }

  async function handleSubmit(e) {
    e.preventDefault();
    let user_group = selected.join(",");
    const json = { username, user_group };

    try {
      const response = await axios.post("http://localhost:4000/admin-update-user", json);
      if (response) {
        console.log("POST BACK");
        message = response.data.message;
        code = response.data.code;
        successToast(message);
        password = "";
      }
    } catch (error) {
      errorToast(error);
      console.log(error);
    }
  }
</script>

<Form on:submit={handleSubmit}>
  <FormGroup>
    <label for="username">Username</label>

    <Input type="text" bind:value={username} />
  </FormGroup>
  <FormGroup>
    <Label for="usergroup">User Group(s):</Label>
    <MultiSelect bind:selected options={groupsArray} />
  </FormGroup>
</Form>
