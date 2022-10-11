<script>
  import axios from "axios"
  import { Form, FormGroup, Input, Label, Row, Col } from "sveltestrap"
  import Select from "svelte-select"

  // url params
  export let appacronym

  export let task_name
  export let task_description
  export let task_notes_existing
  export let task_state
  export let task_plan
  export let task_owner
  export let task_creator

  let task_notes = ""

  async function GetTask() {
    try {
      const response = await axios.get(`http://localhost:4000/get-one-task?AppAcronym=${appacronym}&TaskName=${task_name}`, { withCredentials: true })

      if (response.data) {
        task_plan = response.data.task_plan
        task_description = response.data.task_description
        task_notes_existing = response.data.task_notes
        task_creator = response.data.task_creator
        task_owner = response.data.task_owner
        task_state = response.data.task_state
        console.log(response.data)
      }
    } catch (error) {
      console.log(error)
    }
  }

  $: GetTask()
</script>

<Form>
  <Row>
    <Col>
      <FormGroup>
        <Label>Task Name</Label>
        <Input type="text" value={task_name} placeholder="Task Name" readonly />
      </FormGroup>
      <FormGroup>
        <Label>Plan</Label>
        <Input type="text" value={task_plan} placeholder="NIL" readonly />
      </FormGroup>
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
    </Col>
    <Col>
      <FormGroup>
        <Label>Task Notes Log</Label>
        <Input type="textarea" bind:value={task_notes_existing} placeholder="No task notes" rows={14} readonly />
      </FormGroup>
    </Col>
  </Row>
</Form>

<style>
</style>
