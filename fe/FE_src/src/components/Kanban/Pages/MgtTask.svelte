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
<<<<<<< HEAD

    async function GetAllTasks() {
    
    try {
      const response = await axios.get(`http://localhost:4000/get-all-tasks?AppAcronym=${appacronym}`, { withCredentials: true });

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
=======
>>>>>>> 1c2ec3ec2d766e7af79147c1ce9bc26807a73b82
</script>   

<style>
</style>

<<<<<<< HEAD
<Button style="float:right; font-weight: bold; color: black;" color="warning" on:click={toggleAddTask}>Create Task</Button>


<Row>
  Open
{#each tasksData as task}
{#if task.task_state === "Open"}
<Card>
  <span slot="task-name">{task.task_name}</span>
  <span slot="task-owner">{task.task_owner}</span>
  <span slot="task-description">{task.task_description}</span>
  <Button color="primary" slot="move-left" on:left>&#8592;</Button>
  <span slot="update-task" on:update-task><UpdateTask /></span>
  <Button color="primary" slot="move-right" on:right>&#8594;</Button>
</Card>
{/if}
{/each}

To Do
{#each tasksData as task}
{#if task.task_state === "ToDo"}
<Card>
  <span slot="task-name">{task.task_name}</span>
  <span slot="task-owner">{task.task_owner}</span>
  <span slot="task-description">{task.task_description}</span>
  <Button slot="move-left" on:left>&#8592;</Button>
  <span slot="update-task" on:update-task><UpdateTask /></span>
  <Button slot="move-right" on:right>&#8594;</Button>
</Card>
{/if}
{/each}

Doing
  {#each tasksData as task}
  {#if task.task_state === "Doing"}
  <Card>
    <span slot="task-name">{task.task_name}</span>
    <span slot="task-owner">{task.task_owner}</span>
    <span slot="task-description">{task.task_description}</span>
    <Button slot="move-left" on:left>&#8592;</Button>
    <span slot="update-task" on:update-task><UpdateTask /></span>
    <Button slot="move-right" on:right>&#8594;</Button>
  </Card>
  {/if}
  {/each}

  Done
  {#each tasksData as task}
  {#if task.task_state === "Done"}
  <Card>
    <span slot="task-name">{task.task_name}</span>
    <span slot="task-owner">{task.task_owner}</span>
    <span slot="task-description">{task.task_description}</span>
    <Button slot="move-left" on:left>&#8592;</Button>
    <span slot="update-task" on:update-task><UpdateTask /></span>
    <Button slot="move-right" on:right>&#8594;</Button>
  </Card>
  {/if}
  {/each}

Close
  {#each tasksData as task}
  {#if task.task_state === "Closed"}
  <Card>
    <span slot="task-name">{task.task_name}</span>
    <span slot="task-owner">{task.task_owner}</span>
    <span slot="task-description">{task.task_description}</span>
    <Button slot="move-left" on:left>&#8592;</Button>
    <span slot="update-task" on:update-task><UpdateTask /></span>
    <Button slot="move-right" on:right>&#8594;</Button>
  </Card>
  {/if}
  {/each}
</Row>

=======
<!-- This is dashboard where task(s) will be displayed after clicking into an application -->
<!-- 1. Open -->
>>>>>>> 1c2ec3ec2d766e7af79147c1ce9bc26807a73b82

<Modal isOpen={open} {toggleAddTask} {size}>
  <ModalHeader {toggleAddTask}>Create Task</ModalHeader>
  <ModalBody>
      <CreateTask bind:this={addTaskButton} {task_name} {task_description} {task_notes} {task_plan} {appacronym} />
  </ModalBody>
  <ModalFooter>
      <Button style="color: #fffbf0;" color="warning" on:click={(e) => addTaskButton.handleSubmit(e)}>Create Task</Button>
      <Button class="back-button" color="danger" on:click={toggleAddTask}>Back</Button>
  </ModalFooter>
<<<<<<< HEAD
</Modal>
=======
</Modal>

<!-- 2. To Do -->
<!-- 3. Doing -->
<!-- 4. Done -->
<!-- 5. Closed -->
<div class="container-fluid">
  <br/>
  <Button style="float:right; font-weight: bold; color: black;" color="warning" on:click={toggleAddTask}>Create Task</Button>
</div>
>>>>>>> 1c2ec3ec2d766e7af79147c1ce9bc26807a73b82
