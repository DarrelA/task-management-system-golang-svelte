<script>
  import axios from "axios";
  import { errorToast, successToast } from "../../toast";
  import { Table, Row, Col, Button, Modal, ModalBody, ModalHeader, ModalFooter } from "sveltestrap";
  import CreateTask from "../Form/CreateTask.svelte";
  import Card from "./Card.svelte"

  let addTaskButton;

  export let task_name = ""
  export let task_description = ""
  export let task_notes = ""
  export let task_plan = ""
  let tasksData = ""

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

    async function GetAllTasks() {
    
    try {
      const response = await axios.get("http://localhost:4000/get-all-tasks?AppAcronym=durian", { withCredentials: true });

      if (response.data.error) {
        console.log(response.data.error);
      } else if (!response.data.error) {
        console.log(response.data)
        tasksData = response.data
      }
    } catch (error) {
      console.log(error);
    }
  }

  $: GetAllTasks()
</script>   

<style>
</style>

<Button style="float:right; font-weight: bold; color: black;" color="warning" on:click={toggleAddTask}>Create Task</Button>

<Row>
{#each tasksData as task}
{#if task.task_state === "Open"}
<Card>
  <span slot="task-name">{task.task_name}</span>
  <span slot="task-owner">Task Owner</span>
  <span slot="task-description">Task Description</span>
  <Button slot="move-left" on:left>&#8592;</Button>
  <Button slot="update-task" on:update-task>Update Task</Button>
  <Button slot="move-right" on:right>&#8594;</Button>
</Card>
{/if}
{/each}

{#each tasksData as task}
{#if task.task_state === "To Do"}
<Card>
  <span slot="task-name">{task.task_name}</span>
  <span slot="task-owner">Task Owner</span>
  <span slot="task-description">Task Description</span>
  <Button slot="move-left" on:left>&#8592;</Button>
  <Button slot="update-task" on:update-task>Update Task</Button>
  <Button slot="move-right" on:right>&#8594;</Button>
</Card>
{/if}
{/each}
</Row>

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