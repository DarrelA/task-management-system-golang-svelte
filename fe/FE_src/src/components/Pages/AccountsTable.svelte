<script>
  import axios from "axios";
  import { Button, Row, Col, Modal, ModalBody, ModalHeader, ModalFooter, Table } from "sveltestrap";
  import AdminUpdateUserForm from "./AdminUpdateUserForm.svelte";
  import AdminCreateUserForm from "./AdminCreateUserForm.svelte";

  let usersData = [];
  let open = false;

  let username = "";
  let email = "";
  let user_group = "";
  let status = "";

  let size = "lg";

  let updateButton;

  $: getUsers();

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

  function editUserData(selectedUsername, selectedEmail, selectedUserGroup, selectedStatus) {
    open = !open;
    username = selectedUsername;
    email = selectedEmail;
    if (selectedUserGroup != "") {
      user_group = selectedUserGroup.split(",");
    } else {
      user_group = selectedUserGroup;
    }
    status = selectedStatus;
  }

  export const toggle = (e) => {
    e.preventDefault();
    open = !open;
    getUsers();
  };

  let openModal = false;
  let addButton;
  const toggleAdd = (e) => {
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

<style>
  thead { 
      background-color: #F4BB44;
      /* color: #fffbf0; */
  }

  tbody {
      background-color: #fffbf0;
  }

  th, tr {
      text-align: center;
  }

  .inactive {
    color: red;
    font-weight: bold;
  }

  .active {
    color: mediumseagreen;
    font-weight: bold;
  }
</style>

<div class="container-fluid">
  <Row>
    <Col>
        <h3>User Management</h3>
    </Col>
    <Col>
        <Button style="float:right; font-weight: bold; margin-left: 10px; color: black;" color="warning" on:click={toggleAdd}>Add User</Button>
    </Col>
  </Row>

  <br/>

  <Table bordered responsive>
    <thead>
      <tr>
        <th>Username</th>
        <th>Email</th>
        <th style="width:35%">User Group(s)</th>
        <th>Status</th>
        <th>Edit</th>
      </tr>
    </thead>
    <tbody>
      {#each usersData as userData}
        <tr>
          <td>{userData.username}</td>
          <td>{userData.email}</td>
          <td style="width:35%">{userData.user_group}</td>
          <td class:active={userData.status === "Active"} class:inactive={userData.status === "Inactive"}>{userData.status}</td>

          <td>
            <Button style="font-weight: bold; color: black;" color="warning" on:click={() => editUserData(userData.username, userData.email, userData.user_group, userData.status)}>Update User</Button>
          </td>
        </tr>
      {/each}
    </tbody>
  </Table>

  <Modal isOpen={open} {toggle} {size}>
    <ModalHeader {toggle}>Update User</ModalHeader>
    <ModalBody>
      <AdminUpdateUserForm bind:this={updateButton} {username} {email} {user_group} {status} />
    </ModalBody>

    <ModalFooter>
      <Button style="color: #fffbf0;" color="warning" on:click={(e) => updateButton.handleClick(e)}>Update User</Button>
      <Button class="back-button" color="danger" on:click={toggle}>Back</Button>
    </ModalFooter>
  </Modal>

  <Modal isOpen={openModal} {toggleAdd} {size}>
    <ModalHeader {toggleAdd}>Add User</ModalHeader>
    <ModalBody>
      <AdminCreateUserForm bind:this={addButton} />
    </ModalBody>

    <ModalFooter>
      <Button style="color: #fffbf0;" color="warning" on:click={(e) => addButton.CreateUser(e)}>Add User</Button>
      <Button class="back-button" color="danger" on:click={toggleAdd}>Back</Button>
    </ModalFooter>
  </Modal>
</div>
