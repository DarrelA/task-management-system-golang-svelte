<script>
  import { errorToast, successToast } from "../../toast";
  import {
    Form,
    FormGroup,
    Input,
    Label,
    Button,
    Modal,
    ModalHeader,
    ModalFooter,
    Col,
    Row,
    Spinner,
    ModalBody,
    Styles,
  } from "sveltestrap";
  import axios from "axios";

  export let appacronym; // from app params
  let plan_name = "";
  let plan_acronym = appacronym;
  export let plan_color;
  let plan_start = "";
  let plan_end = "";

  let username = localStorage.getItem("username");
  let message = "";

  let planData = "";

  export async function handleSubmit(event) {
    event.preventDefault();
    const json = { plan_name, plan_acronym, plan_color, plan_start, plan_end };
    console.log(json);
    try {
      const response = await axios.post(
        `http://localhost:4000/create-plan?AppAcronym=${appacronym}`,
        json,
        { withCredentials: true }
      );
      if (response) {
        message = response.data.message;
        successToast(message);
        plan_name = "";
        plan_start = "";
        plan_end = "";
      }
    } catch (error) {
      console.log(error);
      errorToast(error.response.data.message);
    }
  }

  async function GetPlans() {
    try {
      const response = await axios.get(
        `http://localhost:4000/get-all-plans?AppAcronym=${appacronym}`,
        { withCredentials: true }
      );

      if (response.data.error) {
        console.log(response.data.error);
      } else if (!response.data.error) {
        planData = response.data;
      }
    } catch (error) {
      console.log(error);
    }
  }
  $: GetPlans();
</script>

<Form>
  <Row>
    <Col>
      <FormGroup>
        <Label>Plan Name:</Label>
        <Input
          type="text"
          bind:value={plan_name}
          placeholder="Enter Plan Name"
          autofocus
        />
      </FormGroup>
    </Col>
    <Col>
      <FormGroup>
        <Label>Plan Color:</Label>
        <Input type="color" bind:value={plan_color} style="width:100%" />
      </FormGroup>
    </Col>
  </Row>
  <Row>
    <Col>
      <FormGroup>
        <Label>Start Date:</Label>
        <Input type="date" bind:value={plan_start} />
      </FormGroup>
    </Col>
    <Col>
      <FormGroup>
        <Label>End Date:</Label>
        <Input type="date" bind:value={plan_end} />
      </FormGroup>
    </Col>
  </Row>
</Form>
