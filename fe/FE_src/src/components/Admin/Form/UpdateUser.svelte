<script>
  import axios from "axios";
  import { errorToast, successToast } from "../../toast";
  import MultiSelect from "svelte-multiselect";
  import { Form, FormGroup, Input, Label, Row, Col } from "sveltestrap";

  export let username;
  let password;
  export let email;
  export let user_group = [];
  export let status;

  let groupsArray = [];
  let selected = [];
  selected.push(...user_group);

  const loggedInUser = localStorage.getItem("username");

  export async function handleClick(e) {
    e.preventDefault();
    let user_group = selected.join(",");
    const json = { loggedInUser, username, password, email, user_group, status };

    try {
      const response = await axios.post("http://localhost:4000/admin-update-user", json, { withCredentials: true });
      if (response) {
        successToast(response.data.message);
        password = "";
      }
    } catch (error) {
      errorToast(error.response.data.message);
    }
  }

  async function GetUserGroups() {
    
    try {
      const response = await axios.get("http://localhost:4000/get-user-groups", { withCredentials: true });

      if (response.data.error) {
        console.log(response.data.error);
      } else if (!response.data.error) {
        const usergroups = response.data;

        usergroups.forEach((group) => {
          groupsArray.push(group);
        });

        groupsArray = groupsArray;
      }
    } catch (error) {
      console.error(error);
    }
  }
  $: GetUserGroups();
</script>

<Form>
  <Row>
    <Col>
      <FormGroup>
        <Label for="username">Username:</Label>
        <Input type="text" bind:value={username} readonly />
      </FormGroup>
    </Col>

    <Col>
      <FormGroup>
        <Label for="password">Password:</Label>
        <Input type="password" bind:value={password} placeholder="Password" />
      </FormGroup>
    </Col>
  </Row>

  <Row>
    <Col>
      <FormGroup>
        <Label for="email">Email:</Label>
        <Input type="email" bind:value={email} placeholder="Email" />
      </FormGroup>
    </Col>
    
    <Col>
      <FormGroup>
        <Label for="status">Status</Label>
        <Input type="select" bind:value={status} placeholder="Status">
          <option>Inactive</option>
          <option>Active</option>
        </Input>
      </FormGroup>
    </Col>
  </Row>

  <FormGroup>
    <Label for="usergroup">User Group(s):</Label>
    <MultiSelect bind:selected options={groupsArray} />
  </FormGroup>
  
</Form>

<style>
</style>
