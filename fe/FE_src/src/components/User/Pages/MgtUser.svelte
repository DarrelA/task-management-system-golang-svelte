<script>
  import { Modal, ModalBody, ModalHeader, ModalFooter, Button, Row, Col } from "sveltestrap";
  import Navbar from "../NavBar/IsLoggedInUser.svelte";
  import FormUser from "../Form/AddUser.svelte";

  export let password = "" 
  export let email = "";

  let updateUserButton;

  // For Modal
  let open = false;

  const toggle = (e) => {
    e.preventDefault();
    open = !open;
    password = ""
    email = ""
  }
</script>

<Navbar />

<div class="container-fluid">
  <br />
  <div class="container-fluid">
    <Row>
      <Col>
          <h3>User Management</h3>
      </Col>
      <Col>
          <Button style="float:right; font-weight: bold; color: black;" color="warning" on:click={toggle}>Update User</Button>
      </Col>
    </Row>
  </div>
</div>

<Modal isOpen={open} {toggle} >
  <ModalHeader>Update User</ModalHeader>
  <ModalBody>
    <FormUser bind:this={updateUserButton} {password} {email} />
  </ModalBody>

  <ModalFooter>
    <Button style="color: #fffbf0;" color="warning" on:click={(e) => updateUserButton.handleClick(e)}>Update</Button>
    <Button class="back-button" color="danger" on:click={toggle}>Back</Button>
  </ModalFooter>
</Modal>