<script>
  import { errorToast, successToast } from "../../toast";
  import {
    Form,
    FormGroup,
    Input,
    Label,
    Col,
    Row
  } from "sveltestrap";
  import Select from 'svelte-select';
  import axios from "axios";

  export let appacronym; // url params

  export let task_name = "";
  export let task_description = "";
  export let task_notes = "";
  let task_state = "Open";
  export let task_plan = "";
  let task_app_acronym = appacronym;

  let username = localStorage.getItem("username");
  let message = "";

  let getPlansData = [];

  export async function handleSubmit(event) {
    event.preventDefault();
    const json = {
      task_app_acronym,
      task_name,
      task_description,
      task_notes,
      task_plan: task_plan.value,
    };
    console.log(json)
    try {
      const response = await axios.post(
        "http://localhost:4000/create-task",
        json,
        { withCredentials: true }
      );
      if (response) {
        message = response.data.message;
        successToast(message);
        task_name = "";
        task_description = "";
        task_notes = "";
        task_plan = "";
      }
    } catch (error) {
      errorToast(error.response.data.message);
    }
  }

  async function GetPlans() {
    try {
      const response = await axios.get(
        `http://localhost:4000/get-all-plans?AppAcronym=${appacronym}`,
        { withCredentials: true }
      );

      if (response.data) {
        getPlansData = response.data.plans;
      }
    } catch (error) {
      console.log(error);
    }
  }

  $: GetPlans();
  $: planItems = getPlansData.map(info => ({
    value: info.plan_name,
    label: info.plan_name
  }))

</script>

<Form>
  <Row>
    <Col>
      <FormGroup>
        <Label>Task Name</Label>
        <Input
          type="text"
          bind:value={task_name}
          placeholder="Enter a Task Name"
          autofocus
        />
      </FormGroup>
    </Col>
    <Col>
      <FormGroup>
        <Label>Plan Name</Label>
        <Select items={planItems} bind:value={task_plan}></Select>
      </FormGroup>
    </Col>
  </Row>
  <Row>
    <FormGroup>
      <Label>Task Description</Label>
      <Input
        type="textarea"
        bind:value={task_description}
        placeholder="Task Description"
        rows={2}
      />
    </FormGroup>
  </Row>
  <Row>
    <FormGroup>
      <Label>Task Notes</Label>
      <Input
        type="textarea"
        bind:value={task_notes}
        placeholder="Task Notes"
        rows={2}
      />
    </FormGroup>
  </Row>
  <Row>
    <Col>
      <FormGroup>
        <Label>Task State</Label>
        <Input bind:value={task_state} readonly />
      </FormGroup>
    </Col>
    <Col>
      <FormGroup>
        <Label>Task Creator</Label>
        <Input bind:value={username} readonly />
      </FormGroup>
    </Col>
    <Col>
      <FormGroup>
        <Label>Task Owner</Label>
        <Input bind:value={username} readonly />
      </FormGroup>
    </Col>
  </Row>
</Form>
