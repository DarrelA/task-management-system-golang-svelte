<script>
    import axios from "axios";
    import { errorToast, successToast } from "../../toast";
    import { Form, FormGroup, Input, Label, Row, Col, Button, Modal, ModalHeader, ModalBody, ModalFooter } from "sveltestrap";
    
    let task_name = ""
    let task_description = ""
    let task_notes = ""
    let task_state = ""
    let task_plan = ""
    let task_app_acronym = ""
    let task_owner = ""
    let task_creator = ""

    let username = localStorage.getItem("username")

    let size = "xl";
    let open = false;

    async function handleSubmit(event) {
        event.preventDefault()
        const json = {task_app_acronym, task_name, task_description, task_notes, task_plan}
        try {
            const response = await axios.post("http://localhost:4000/update-task", json, {withCredentials: true});
            if (response) {
                successToast(response.data.message);
                task_name = ""
                task_description = ""
                task_notes = ""
                task_plan = ""
            }
        } catch(error) {
            errorToast(error.response.data.message);
        }
    }

    function handleClick() {
        open = !open
    }

    function toggle(e) {
        e.preventDefault()
        open = !open;
    }
</script>
  
<style>
</style>

<!-- TO BE DONE BY BEATRICE -->
<Button style="font-weight: bold; color: black;" color="warning" on:click={handleClick} >Update Task</Button>

<Modal isOpen={open} {toggle} {size}>
  <ModalHeader {toggle}>Update Task</ModalHeader>
  <ModalBody>
    <Form>
      <Row>
        <Col>
          <FormGroup>
            <Label>Task Name</Label>
            <Input type="text" bind:value={task_name} placeholder="Task Name" />
          </FormGroup>
        </Col>
        <Col>
          <FormGroup>
              <Label>Plan Name</Label>
              <Input type="select" bind:value={task_plan} placeholder="Select a Plan">
                  <option value="Sprint 1">Sprint 1</option>
                  <option value="Sprint 2">Sprint 2</option>
              </Input>
          </FormGroup>
        </Col>
      </Row>
      <Row>
        <Col>
          <FormGroup>
            <Label>Task Description</Label>
            <Input type="textarea" bind:value={task_description} placeholder="Task Description" rows={4} />
          </FormGroup>

          <FormGroup>
            <Label>Task Creator</Label>
            <Input type="text" bind:value={task_creator} placeholder="Task Creator" />
          </FormGroup>

          <FormGroup>
            <Label>Task Owner</Label>
            <Input type="text" bind:value={task_owner} placeholder="Task Owner" />
          </FormGroup>
        </Col>
        
        <Col>
          <FormGroup>
            <Label>Task Notes</Label>
            <Input type="textarea" bind:value={task_description} placeholder="Task Description" rows={11} disabled />
          </FormGroup>
        </Col>

        <FormGroup>
          <Label>Task Notes</Label>
          <Input type="textarea" bind:value={task_description} placeholder="Task Notes" rows={5} />
        </FormGroup>
      </Row>
    </Form>
  </ModalBody>

  <ModalFooter>
    <Button style="color: #fffbf0;" color="warning" on:click={(e) => handleSubmit(e)}>Update Task</Button>
    <Button class="back-button" color="danger" on:click={toggle}>Back</Button>
  </ModalFooter>
</Modal>
