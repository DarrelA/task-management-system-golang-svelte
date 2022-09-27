<script>
  import axios from "axios";
  import { errorToast, successToast } from "../toast";
  import { Button, Form, FormGroup, Input, Label, Row, Col } from "sveltestrap";
  import { Modal, ModalBody, ModalHeader, ModalFooter, Table } from "sveltestrap";
  import Navbar from "../Navbar/IsLoggedInUser.svelte"

  let password = "" 
  let email = "";

  let message = ""
  let code = ""

  // For Modal
  let open = false;

  async function handleClick(e) {
    e.preventDefault();
    const json = {password, email};

    try {
      const response = await axios.post("http://localhost:4000/update-user", json, { withCredentials: true });
      if (response) {
        console.log(response)
        message = response.data.message;
        code = response.data.code;
        successToast(message);
        password = "";
        email = "";
      }
    } catch (error) {
      errorToast(error.response.data.message);
    }
  }

  const toggle = (e) => {
    e.preventDefault();
    open = !open;
  }

  function openModal() {
    open = !open
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
    <Form>
      <Row>
        <Col>    
          <FormGroup>
            <Label for="password">Password:</Label>
            <Input type="password" bind:value={password} placeholder="Password" />
          </FormGroup>
        
          <FormGroup>
            <Label for="email">Email:</Label>
            <Input type="email" bind:value={email} placeholder="Email" />
          </FormGroup>
        </Col>
      </Row>
    </Form>
  </ModalBody>

  <ModalFooter>
    <Button style="color: #fffbf0;" color="warning" on:click={handleClick}>Update User</Button>
    <Button class="back-button" color="danger" on:click={toggle}>Back</Button>
  </ModalFooter>
</Modal>