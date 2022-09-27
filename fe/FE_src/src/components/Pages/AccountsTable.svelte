<script>
  import axios from "axios";
  import { Button, Modal, ModalBody, ModalHeader, ModalFooter, Table } from "sveltestrap";
  import AdminUpdateUserForm from "./AdminUpdateUserForm.svelte";

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
</script>

<div>
  <h1>Users Database</h1>
  <Table bordered style="margin:0 auto;width:95%">
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

          <td
            ><Button color="primary" on:click={() => editUserData(userData.username, userData.email, userData.user_group, userData.status)}
              >Update User</Button
            ></td
          >
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
      <Button color="primary" on:click={(e) => updateButton.handleClick(e)}>Update User</Button>
      <Button class="back-button" color="danger" on:click={toggle}>Back</Button>
    </ModalFooter>
  </Modal>
</div>

<style>
  h1 {
    color: blueviolet;
    text-align: center;
    font-family: "Montserrat";
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
