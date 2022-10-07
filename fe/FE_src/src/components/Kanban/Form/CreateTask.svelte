<script>
  import { errorToast, successToast } from "../../toast";
  import { Form, FormGroup, Input, Label, Col, Row, Dropdown, DropdownToggle, DropdownItem, DropdownMenu } from "sveltestrap";
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

  let planData = "";

  export async function handleSubmit(event) {
    event.preventDefault();
    const json = {
      task_app_acronym,
      task_name,
      task_description,
      task_notes,
      task_plan,
    };
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
        <Label>Task Name:</Label>
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
        <Dropdown>
          <DropdownToggle style="width:100%" caret>Plan Name</DropdownToggle>
          <DropdownMenu>
            {#each planData as plan}
              <DropdownItem
                on:click={() => (task_plan = plan.plan_name)}
                placeholder={plan}
              >
                {plan.plan_name}
              </DropdownItem>
            {/each}
          </DropdownMenu>
        </Dropdown>
        <FormGroup>
          <Input value={task_plan} type="text" readonly />
        </FormGroup>
      </FormGroup>
    </Col>
  </Row>
  <Row>
    <FormGroup>
      <Label>Task Description:</Label>
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
      <Label>Task Notes:</Label>
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
        <Label>Task State:</Label>
        <Input bind:value={task_state} readonly />
      </FormGroup>
    </Col>
    <Col>
      <FormGroup>
        <Label>Task Creator:</Label>
        <Input bind:value={username} readonly />
      </FormGroup>
    </Col>
    <Col>
      <FormGroup>
        <Label>Task Owner:</Label>
        <Input bind:value={username} readonly />
      </FormGroup>
    </Col>
  </Row>
</Form>
