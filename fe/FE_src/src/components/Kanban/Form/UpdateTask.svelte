<script>
    import axios from "axios";
    import { errorToast, successToast } from "../../toast";
    import { Form, FormGroup, Input, Label, Row, Col, Dropdown, DropdownToggle, DropdownItem, DropdownMenu } from "sveltestrap";

    // url params
    export let appacronym;

    export let task_name;
    export let task_description;
    export let task_notes_existing;
    export let task_state;
    export let task_plan;
    export let task_owner;
    export let task_creator;
    export let canUpdateTask;
    
    let task_notes = "";
    let task_app_acronym = appacronym;

    let username = localStorage.getItem("username")
    let planData = []

  export async function handleSubmit(event) {
      event.preventDefault()
      task_owner = username
      const json = {task_notes, task_plan, task_owner}
      try {
        const response = await axios.post(`http://localhost:4000/update-task?AppAcronym=${task_app_acronym}&TaskName=${task_name}`, 
          json, 
          {withCredentials: true});

        if (response) {
          successToast(response.data.message);
          task_notes = ""
          GetTask();
        }
      } catch(error) {
        errorToast(error.response.data.message);
      }
    }
  
  async function GetTask() {
    try {
      const response = await axios.get(
        `http://localhost:4000/get-one-task?AppAcronym=${appacronym}&TaskName=${task_name}`,
        { withCredentials: true }
      );

      if (response.data) {
        task_plan = response.data.task_plan;
        task_description = response.data.task_description;
        task_notes_existing = response.data.task_notes;
        task_creator = response.data.task_creator;
        task_owner = response.data.task_owner;
        task_state = response.data.task_state;
      }
    } catch (error) {
      console.log(error);
    }
  }

  async function GetPlans() {
    try {
      const response = await axios.get(
        `http://localhost:4000/get-all-plans?AppAcronym=${task_app_acronym}`,
        { withCredentials: true }
      );

      if (response.data) {
        planData = response.data;
      }
    } catch (error) {
      console.log(error);
    }
  }

  console.log(canUpdateTask)

  $: GetPlans();
  $: GetTask();
</script>
  
<style>
</style>

<Form>
  <Row>
    <Col>
      <FormGroup>
              <Label>Task Name</Label>
              <Input type="text" value={task_name} placeholder="Task Name" readonly />
            </FormGroup>
          </Col>
          <Col>
            <FormGroup>
              <Dropdown>
                <DropdownToggle style="width:100%" caret disabled={!canUpdateTask}>Plan Name</DropdownToggle>
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
          <Col>
            <FormGroup>
              <Label>Task Description</Label>
              <Input type="textarea" bind:value={task_description} placeholder="No task description" rows={3} readonly />
            </FormGroup>
            
            <Row>
              <Col>
                <FormGroup>
                  <Label>Task Creator</Label>
                  <Input type="text" bind:value={task_creator} placeholder="Task Creator" readonly />
                </FormGroup>
              </Col>
              <Col>
                <FormGroup>
                  <Label>Task Owner</Label>
                  <Input type="text" bind:value={task_owner} placeholder="Task Owner" readonly />
                </FormGroup>
              </Col>
              <Col>
                <FormGroup>
                  <Label>Task State</Label>
                  <Input type="text" bind:value={task_state} placeholder="Task State" readonly />
                </FormGroup>
              </Col>
            </Row>
            
            <Row>
              <FormGroup>
                <Label>Task Notes</Label>
                <Input type="textarea" bind:value={task_notes} placeholder="Enter task notes" rows={6} readonly={!canUpdateTask}/>
              </FormGroup>
            </Row>
          </Col>
  
          <Col>
            <FormGroup>
              <Label>Task Notes Log</Label>
              <Input type="textarea" bind:value={task_notes_existing} placeholder="No task notes" rows={15} readonly />
            </FormGroup>
          </Col>
        </Row>
      </Form>