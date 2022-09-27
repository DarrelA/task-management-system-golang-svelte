<script>
  import axios from "axios";
  import { errorToast, successToast } from "../toast";
  import { Button, Form, FormGroup, Input, Label, Row, Col } from "sveltestrap";
  import { Modal, ModalBody, ModalHeader, ModalFooter, Table } from "sveltestrap";

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

<h4>Update User Credentials</h4>
<Button color="primary" on:click={openModal}>Update User</Button>

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
    <Button color="primary" on:click={handleClick}>Update User</Button>
    <Button class="back-button" color="danger" on:click={toggle}>Back</Button>
  </ModalFooter>
</Modal>

<!-- <Form>
  <Row>
    <Col xs={3}>
      <FormGroup>
        <Label for="username">Username:</Label>
        <Input type="text" bind:value={username} placeholder="Username" readonly />
      </FormGroup>

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

  <Button color="primary" on:click={handleClick}>Update</Button>
</Form> -->