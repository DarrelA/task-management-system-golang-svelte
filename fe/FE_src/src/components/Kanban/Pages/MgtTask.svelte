<script>
  import axios from "axios";
  import { errorToast, successToast } from "../../toast";
  import { Table, Row, Col, Button, Modal, ModalBody, ModalHeader, ModalFooter } from "sveltestrap";
  import CreateTask from "../Form/CreateTask.svelte";
  import Card from "./Card.svelte"
  import UpdateTask from "../Form/UpdateTask.svelte"

  let addTaskButton;

  export let task_name = ""
  export let task_description = ""
  export let task_notes = ""
  export let task_plan = ""
  export let appacronym;
  let tasksData = ""

  let size = "lg";
  let open = false;

  $: console.log(appacronym)

  function toggleAddTask(e) {
        e.preventDefault()
        open = !open;
        task_name = ""
        task_description = ""
        task_notes = ""
        task_plan = ""
        GetAllTasks()
    }
</script>   

<style>
</style>

<!-- This is dashboard where task(s) will be displayed after clicking into an application -->
<!-- 1. Open -->

<Modal isOpen={open} {toggleAddTask} {size}>
  <ModalHeader {toggleAddTask}>Create Task</ModalHeader>
  <ModalBody>
      <CreateTask bind:this={addTaskButton} {task_name} {task_description} {task_notes} {task_plan} {appacronym} />
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
</div>
