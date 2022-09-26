<script>
  import axios from "axios";
  import { toast } from "@zerodevx/svelte-toast";
  import { SvelteToast } from "@zerodevx/svelte-toast";
  import MultiSelect from "svelte-multiselect";
  import { Button, Badge, Form, FormGroup, FormText, Input, Label, Dropdown, DropdownToggle, DropdownItem, DropdownMenu } from "sveltestrap";
  import Navbar from "../Navbar/IsLoggedInAdmin.svelte";
  import ProtectedRoute from "../ProtectedRoute.svelte";

  let message = "";
  let code = "";
  export let username = [];
  //   export let user_group = [];

  let groupsArray = [];
  let userArray = [];
  let selected = [];
  let selectedUser = "";
  //   selected.push(...user_group);

  async function handleSubmit() {
    // let user_group = selected.join(",");
    console.log(username);
    username = username.toString();
    const json = { username: selectedUser, groupname: selected };
    console.log(json);

    try {
      const response = await axios.post("http://localhost:4000/add-user-to-group", json, { withCredentials: true });
      if (response) {
        message = response.data.message;
        console.log(response.data);
        code = response.data.code;
        toast.push(message);
      }
    } catch (error) {
      console.log(error);
    }
  }

  async function GetUserData() {
    try {
      const response = await axios.get("http://localhost:4000/get-users");
      console.log(response);
      if (response.data.error) {
        console.log(response.data.error);
      } else if (!response.data.error) {
        response.data.forEach((user) => {
          userArray.push(user.username);
        });

        userArray = userArray;
      }
    } catch (e) {
      console.error(e);
    }
  }
  $: GetUserData();

  async function GetUserGroups() {
    try {
      const response = await axios.get("http://localhost:4000/get-user-groups");

      console.log(response);
      if (response.data.error) {
        console.log(response.data.error);
      } else if (!response.data.error) {
        const usergroups = response.data;

        usergroups.forEach((group) => {
          groupsArray.push(group);
        });

        groupsArray = groupsArray;
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

<!-- <Navbar /> -->

<form on:submit|preventDefault={handleSubmit}>
  <FormGroup>
    <Dropdown>
      <DropdownToggle caret>Users</DropdownToggle>
      <DropdownMenu>
        {#each userArray as user}
          <DropdownItem active={user === selectedUser} on:click={() => (selectedUser = user)} placeholder={user}>
            {user}</DropdownItem
          >
        {/each}
      </DropdownMenu>
    </Dropdown>
    <input disabled placeholder={selectedUser} />
  </FormGroup>
  <FormGroup>
    <Label for="usergroup">User Group(s):</Label>
    <MultiSelect bind:selected options={groupsArray} allowUserOptions={true} />
  </FormGroup>
  <Button color="primary">Add User</Button>
</form>

<style>
</style>
