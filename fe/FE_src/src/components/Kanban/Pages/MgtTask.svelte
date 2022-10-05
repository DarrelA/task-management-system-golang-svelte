<script>
  import axios from "axios";
  import { errorToast, successToast } from "../../toast";
<<<<<<< Updated upstream
  import {  Button, Card, CardBody, CardFooter, CardHeader, CardSubtitle, CardText, CardTitle, Row, Col } from "sveltestrap";
=======
  import { Table, Row, Col, Button, Modal, ModalBody, ModalHeader, ModalFooter } from "sveltestrap";
  import CreateTask from "../Form/CreateTask.svelte";

  let addTaskButton;

  export let task_name = ""
  export let task_description = ""
  export let task_notes = ""
  export let task_plan = ""

  let size = "lg";
  let open = false;

  function toggleAddTask(e) {
        e.preventDefault()
        open = !open;
        task_name = ""
        task_description = ""
        task_notes = ""
        task_plan = ""
    }
>>>>>>> Stashed changes
</script>   

<style>
</style>

<<<<<<< Updated upstream

<div class="text-center">
  <Row>
    <!-- Task: Open -->
    <Col>
      <Card class="mb-3">
        <CardHeader>
          <CardTitle>Open</CardTitle>
        </CardHeader>
        <CardBody>
          <CardSubtitle>Card subtitle</CardSubtitle>
          <CardText>
            Some quick example text to build on the card title and make up the bulk of
            the card's content.
          </CardText>
          <Button>Button</Button>
        </CardBody>
      </Card>
    </Col>
    <!-- Task: To Do -->
    <Col>
      <Card class="mb-3">
        <CardHeader>
          <CardTitle>To Do</CardTitle>
        </CardHeader>
        <CardBody>
          <CardSubtitle>Card subtitle</CardSubtitle>
          <CardText>
            Some quick example text to build on the card title and make up the bulk of
            the card's content.
          </CardText>
          <Button>Button</Button>
        </CardBody>
      </Card>
    </Col>
    <!-- Task: Doing -->
    <Col>
      <Card class="mb-3">
        <CardHeader>
          <CardTitle>Doing</CardTitle>
        </CardHeader>
        <CardBody>
          <CardSubtitle>Card subtitle</CardSubtitle>
          <CardText>
            Some quick example text to build on the card title and make up the bulk of
            the card's content.
          </CardText>
          <Button>Button</Button>
        </CardBody>
      </Card>
    </Col>
    <!-- Task: Done -->
    <Col>
      <Card class="mb-3">
        <CardHeader>
          <CardTitle>Done</CardTitle>
        </CardHeader>
        <CardBody>
          <CardSubtitle>Card subtitle</CardSubtitle>
          <CardText>
            Some quick example text to build on the card title and make up the bulk of
            the card's content.
          </CardText>
          <Button>Button</Button>
        </CardBody>
      </Card>
    </Col>
    <!-- Task: Closed -->
    <Col>
      <Card class="mb-3">
        <CardHeader>
          <CardTitle>Closed</CardTitle>
        </CardHeader>
        <CardBody>
          <CardSubtitle>Card subtitle</CardSubtitle>
          <CardText>
            Some quick example text to build on the card title and make up the bulk of
            the card's content.
          </CardText>
          <Button>Button</Button>
        </CardBody>
      </Card>
    </Col>
  </Row>
=======
<!-- This is dashboard where task(s) will be displayed after clicking into an application -->
<!-- 1. Open -->

<Modal isOpen={open} {toggleAddTask} {size}>
  <ModalHeader {toggleAddTask}>Create Task</ModalHeader>
  <ModalBody>
      <CreateTask bind:this={addTaskButton} {task_name} {task_description} {task_notes} {task_plan} />
  </ModalBody>
  <ModalFooter>
      <Button style="color: #fffbf0;" color="warning" on:click={(e) => addTaskButton.handleSubmit(e)}>Create Task</Button>
      <Button class="back-button" color="danger" on:click={toggleAddTask}>Back</Button>
  </ModalFooter>
</Modal>

<!-- 2. To Do -->
<!-- 3. Doing -->
<!-- 4. Done -->
<!-- 5. Closed -->
<div class="container-fluid">
  <br/>
  <Button style="float:right; font-weight: bold; color: black;" color="warning" on:click={toggleAddTask}>Create Task</Button>
>>>>>>> Stashed changes
</div>