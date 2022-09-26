<script>
  import axios from "axios";
  import { Table, Button, Modal, ModalBody, ModalHeader, ModalFooter } from "sveltestrap";
  import FormAddGroup from "../Form/FormAddGroup.svelte"

  let usersGroupData = [];
  let groupname = "";
  let addButton;
  let openModal = false;
  const size = "lg";

  let loggedInUser = localStorage.getItem("username")

  $: getGroupInfo();

  export async function getGroupInfo() {
    try {
      const response = await axios.post("http://localhost:4000/get-users-in-group", {loggedInUser}, {withCredentials: true});
      if (response) {
        console.log(response.data)
        usersGroupData = response.data;
      }
    } catch (error) {
      console.log(error);
    }
  }

  const toggle = (e) => {
      e.preventDefault();
      openModal = !openModal;
      getGroupInfo();
      groupname = "";
    };
</script>

<Button color="primary" on:click={toggle}>Add Group</Button>

<br/><br/>

<Table bordered style="margin:0 auto;width:95%">
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

<Modal isOpen={openModal} {toggle} {size}>
  <ModalHeader {toggle}>Add Group</ModalHeader>
  <ModalBody>
    <FormAddGroup bind:this={addButton} {groupname}/>
  </ModalBody>
  <ModalFooter>
    <Button on:click={(e) => addButton.handleAddGroup(e)} style="background-color: #FCA311; border: none;">Add</Button>
    <Button class="back-button" color="danger" on:click={toggle}>Back</Button>
  </ModalFooter>
</Modal>