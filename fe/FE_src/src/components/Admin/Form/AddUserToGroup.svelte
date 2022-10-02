<script>
  import axios from "axios";
  import { errorToast, successToast } from "../../toast";
  import MultiSelect from "svelte-multiselect";
  import { Form, FormGroup, Row, Col, Input, Label, Dropdown, DropdownToggle, DropdownItem, DropdownMenu } from "sveltestrap";

  export let username = [];
  let selected;
  let selectedUser;

  let groupsArray = [];
  let userArray = [];

  const loggedInUser = localStorage.getItem("username");

  export async function handleSubmit() {
    username = username.toString();
    const json = { loggedInUser, username: selectedUser, groupname: selected };

    try {
      const response = await axios.post("http://localhost:4000/add-user-to-group", json, { withCredentials: true });
      if (response) {
          successToast(response.data.message);
          selected = [];
          selectedUser = "";
      }
    } catch (error) {
        errorToast(error.response.data.message);
    }
  }

  async function GetUserData() {
    try {
      const response = await axios.get("http://localhost:4000/get-users", { withCredentials: true });
      console.log(response);
      if (response.data.error) {
        console.log(response.data.error);
      } else if (!response.data.error) {
        response.data.forEach((user) => {
          userArray.push(user.username);
        });

        userArray = userArray;
      }
    } catch (error) {
      console.error(error);
    }
  }
  $: GetUserData();

  async function GetUserGroups() {
    try {
      const response = await axios.get("http://localhost:4000/get-user-groups", { withCredentials: true });

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

<style>
</style>

<Form on:submit={handleSubmit}>
  <Row>
    <Col>
      <FormGroup>
        <Dropdown>
          <DropdownToggle style="width:100%" caret>Users</DropdownToggle>
           <DropdownMenu>
              {#each userArray as user}
                <DropdownItem active={user === selectedUser} on:click={() => (selectedUser = user)} placeholder={user}>
                  {user}
                </DropdownItem>
              {/each}
            </DropdownMenu>
        </Dropdown>
      </FormGroup>
    </Col>
    <Col class="col-md-6">
      <FormGroup>
        <Input placeholder={selectedUser} type="text" disabled />
      </FormGroup>
    </Col>
  </Row>
  <Row>
    <Col>
      <FormGroup>
        <Label for="usergroup">Group</Label>
        <MultiSelect bind:selected options={groupsArray} allowUserOptions={true} />
      </FormGroup>
    </Col>
  </Row>
</Form>
