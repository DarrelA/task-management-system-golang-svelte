<script>
  import axios from "axios";
  import { errorToast, successToast } from "../toast";
  import { Button, Form, FormGroup, Input, Label, Row, Col } from "sveltestrap";

  let loggedInUser = localStorage.getItem("username");
  let username = ""; //test w/o login
  let password = "" 
  let email = "";

  let message = ""
  let code = ""

  async function handleClick(e) {
    e.preventDefault();
    const json = {username, password, email};
    //console.log("sending: " + password + ":" + email)

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

</script>

<Form>
  <Row>
    <Col xs={3}>
      <FormGroup>
        <Label for="username">Username:</Label>
        <Input type="text" bind:value={username} placeholder="Username" />
        <Label>Leave input empty if not updating the required field</Label>
      </FormGroup>
    </Col>
    <Col xs={3}>
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
</Form>