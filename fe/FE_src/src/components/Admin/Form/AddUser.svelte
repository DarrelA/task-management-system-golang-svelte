<script>
  import axios from "axios";
  import { onMount } from "svelte";
  import { Form, FormGroup, Input, Label, Col, Row, Spinner, Styles } from "sveltestrap";
  import MultiSelect from "svelte-multiselect";
  import { errorToast, successToast } from "../../toast";

  let username = "";
  let email = "";
  let password = "";
  let status = "Active";

  let groupsArray = [];
  let selected = [];

  let loading = false;

  let loggedInUser = localStorage.getItem("username");

  async function getUsers() {
    try {
      const response = await axios.get("http://localhost:4000/get-users", { withCredentials: true });
      if (response) {
        usersData = response.data;
      }
    } catch (error) {
      console.log(error);
    }
  }

  export async function CreateUser(e) {
    e.preventDefault();
    const json = {
      loggedInUser,
      username,
      email,
      password,
      user_group: selected,
      status,
    };
    try {
      const response = await axios.post("http://localhost:4000/admin-create-user", json, { withCredentials: true });
      loading = true;

      setTimeout(() => {
        if (!response.data.error) {
          successToast(response.data.message);
          loading = false;

          username = "";
          password = "";
          email = "";
          selected = [];
          status = "Active";
          getUsers();
        }
      }, 500);
    } catch (error) {
      errorToast(error.response.data.message);
    }
  }

  onMount(() => {
    async function GetUserGroups() {
      try {
        const response = await axios.get("http://localhost:4000/get-user-groups", { withCredentials: true });

        if (response.data.error) {
          console.error(response.data.error);
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
    GetUserGroups();
  });

  let openModal = false;
  let size;
  const toggle = (e) => {
    e.preventDefault();
    getUsers();
    openModal = !openModal;
    size = "xl";

    username = "";
    password = "";
    email = "";
    selected = [];
    status = "Active";
  };
</script>

<Styles />
<div>
  {#if loading}
    <div class="loading-spinner">
      <Spinner size="lg"/>
    </div>
  {/if}
  <Form>
    <Row>
      <Col>
        <FormGroup>
          <Label>Username:</Label>
          <Input placeholder="username" bind:value={username} autofocus />
        </FormGroup>
      </Col>

      <Col>
        <FormGroup>
          <Label>Password:</Label>
          <Input type="password" placeholder="password" bind:value={password} />
        </FormGroup>
      </Col>
    </Row>

    <Row>
      <Col>
        <FormGroup>
          <Label>Email:</Label>
          <Input placeholder="example@email.com" bind:value={email} />
        </FormGroup>
      </Col>

      <Col>
        <FormGroup>
          <Label>User group(s)</Label>
          <MultiSelect bind:selected options={groupsArray} allowUserOptions={true} />
        </FormGroup>
      </Col>
    </Row>

    <Row>
      <Col>
        <FormGroup>
          <Input type="select" bind:value={status} placeholder="Status">
            <option>Inactive</option>
            <option>Active</option>
          </Input>
        </FormGroup>
      </Col>
    </Row>
  </Form>
</div>

<style>
  .loading-spinner {
    position: relative;
    left: 50%;
    top: 50%;
  }
</style>
