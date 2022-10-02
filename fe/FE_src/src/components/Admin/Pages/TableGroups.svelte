<script>
  import axios from "axios";
  import { Table, Row, Col, Button, Modal, ModalBody, ModalHeader, ModalFooter } from "sveltestrap";
  import FormAddGroup from "../Form/AddGroup.svelte"
  import AddUserToGroup from "../Form/AddUserToGroup.svelte"

  let usersGroupData = [];
  export let groupname = "";
  export let selected = [];
  export let selectedUser = "";

  let addGroupButton;
  let addUserButton;
  let openModalAddGroup = false;
  let openModalAddUser = false;
  const size = "lg";

  $: getGroupInfo();

  async function getGroupInfo() {
    try {
      const response = await axios.get("http://localhost:4000/get-users-in-group", {withCredentials: true});
      if (response) {
        console.log(response.data)
        usersGroupData = response.data;
      }
    } catch (error) {
      console.log(error);
    }
  }

  const toggleAddGroup = (e) => {
    e.preventDefault();
    openModalAddGroup = !openModalAddGroup;
    getGroupInfo();
    groupname = "";
  };

  const toggleAddUser = (e) => {
    e.preventDefault();
    openModalAddUser = !openModalAddUser;
    getGroupInfo();
    selected = [];
    selectedUser = "";
  };
</script>

<style>
    thead {
      background-color: #F4BB44;
    }

    tbody {
      background-color: #fffbf0;
    }

    th, tr {
      text-align: center;
    }
</style>


<div class="container-fluid">
  <Row>
    <Col>
        <h3>Group Management</h3>
    </Col>
    <Col>
        <Button style="float:right; font-weight: bold; margin-left: 10px; color: black;" color="warning" on:click={toggleAddUser}>Add User To Group</Button>
        <Button style="float:right; font-weight: bold; color: black;" color="warning" on:click={toggleAddGroup}>Add Group</Button>
    </Col>
  </Row>

  <br/>

  <Table bordered responsive>
    <thead>
      <tr>
        <th>Groupname</th>
        <th># of people</th>
      </tr>
    </thead>
    <tbody>
      {#each usersGroupData as userGroupData}
      <tr>
        <td>{userGroupData.user_group}</td>
        <td>{userGroupData.user_count}</td>
      </tr>
      {/each}
    </tbody>
  </Table>
</div>                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                           

<Modal isOpen={openModalAddGroup} {toggleAddGroup} {size}>
  <ModalHeader {toggleAddGroup}>Add Group</ModalHeader>
  <ModalBody>
    <FormAddGroup bind:this={addGroupButton} {groupname}/>
  </ModalBody>
  <ModalFooter>
    <Button style="color: #fffbf0;" color="warning" on:click={(e) => addGroupButton.handleAddGroup(e)}>Add</Button>
    <Button class="back-button" color="danger" on:click={toggleAddGroup}>Back</Button>
  </ModalFooter>
</Modal>

<Modal isOpen={openModalAddUser} {toggleAddUser} {size}>
  <ModalHeader {toggleAddUser}>Add User To Group</ModalHeader>
  <ModalBody>
    <AddUserToGroup bind:this={addUserButton} {selected} {selectedUser} />
  </ModalBody>
  <ModalFooter>
    <Button style="color: #fffbf0;" color="warning" on:click={(e) => addUserButton.handleSubmit(e)}>Add</Button>
    <Button class="back-button" color="danger" on:click={toggleAddUser}>Back</Button>
  </ModalFooter>
</Modal>