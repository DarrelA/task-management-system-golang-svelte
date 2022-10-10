<script>
  import axios from "axios";
  import { errorToast, successToast } from "../../toast";
  import { Form, Row, Col, FormGroup, Input, Label } from "sveltestrap";
    
  let password;
  let email;

  export async function handleClick(e) {
    e.preventDefault();
    const json = {password, email};

    try {
      const response = await axios.post("http://localhost:4000/update-user", json, { withCredentials: true });
      if (response) {
        successToast(response.data.message);
        password = "";
        email = "";
      }
    } catch (error) {
      errorToast(error.response.data.message);
    }
  }
</script>
  
<style>
</style>

<Form>
    <Row>
      <Col>    
        <FormGroup>
          <Label for="password">Password</Label>
          <Input type="password" bind:value={password} placeholder="Password" />
        </FormGroup>
      
        <FormGroup>
          <Label for="email">Email</Label>
          <Input type="email" bind:value={email} placeholder="Email" />
        </FormGroup>
      </Col>
    </Row>
  </Form>
