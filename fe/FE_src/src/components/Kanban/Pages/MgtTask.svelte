<script>
  import axios from "axios";
  import { errorToast, successToast } from "../../toast";
  import {
    Table,
    Row,
    Col,
    Button,
    Modal,
    ModalBody,
    ModalHeader,
    ModalFooter,
    Container,
  } from "sveltestrap";
  import CreateTask from "../Form/CreateTask.svelte";
  import Card from "./Card.svelte";
  import UpdateTask from "../Form/UpdateTask.svelte";

  export let appacronym = null;
  export let tasksData = [];

  let size = "lg";
  let open = false;
  export let createTaskButton;

  let task_name = "";
  let task_description = "";
  let task_notes = "";
  let task_plan = "";

  let userAppPermits = JSON.parse(localStorage.getItem("userAppPermits"));
  let {
    IsPermitCreate,
    IsPermitOpen,
    IsPermitToDo,
    IsPermitDoing,
    IsPermitDone,
  } = userAppPermits;

  export async function GetAllTasks() {
    try {
      const response = await axios.get(
        `http://localhost:4000/get-all-tasks?AppAcronym=${appacronym}`,
        { withCredentials: true }
      );

      if (response.data) {
        tasksData = response.data;
      }
    } catch (error) {
      console.log("error");
    }
  }

  const GetUserAppPermits = async () => {
    try {
      const response = await axios.get(
        `http://localhost:4000/get-user-app-permits?appacronym=${appacronym}`,
        {
          withCredentials: true,
        }
      );
      if (response) {
        localStorage.setItem("userAppPermits", JSON.stringify(response.data));
      }
    } catch (e) {
      e.response && e.response.data.message
        ? errorToast(e.response.data.message)
        : errorToast(e.message);
    }
  };

  const demoteTask = async (task_name, task_state) => {
    const json = { task_app_acronym: appacronym, task_name, task_state };

    try {
      const response = await axios.put(
        "http://localhost:4000/task-state-transition",
        json,
        {
          withCredentials: true,
        }
      );
      if (response) {
        GetAllTasks();
      }
    } catch (e) {
      e.response && e.response.data.message
        ? errorToast(e.response.data.message)
        : errorToast(e.message);
    }
  };

  const promoteTask = async (task_name, task_state) => {
    const json = { task_app_acronym: appacronym, task_name, task_state };

    try {
      const response = await axios.put(
        "http://localhost:4000/task-state-transition",
        json,
        {
          withCredentials: true,
        }
      );
      if (response) {
        GetAllTasks();
      }
    } catch (e) {
      e.response && e.response.data.message
        ? errorToast(e.response.data.message)
        : errorToast(e.message);
    }
  };

  function toggle(e) {
    e.preventDefault();
    open = !open;
    task_name = "";
    task_description = "";
    task_notes = "";
    task_plan = "";
    GetAllTasks();
  }

  $: GetAllTasks();
  $: GetUserAppPermits();
</script>

{#if IsPermitCreate}
  <Button color="primary" on:click={toggle}>Create Task</Button>
{/if}

<Row>
  <Col>
    Open
    {#each tasksData as task}
      {#if task.task_state === "Open"}
        <Card>
          <span slot="task-name">{task.task_name}</span>
          <span slot="task-owner">{task.task_owner}</span>
          <span slot="task-description">{task.task_description}</span>
          <Row slot="task-actions">
            {#if IsPermitOpen}
              <Col><Button on:update-task><UpdateTask /></Button></Col>
              <Col
                ><Button
                  color="primary"
                  on:click={() => promoteTask(task.task_name, "ToDo")}
                  >&#8594;</Button
                ></Col
              >
            {/if}
          </Row>
        </Card>
      {/if}
    {/each}
  </Col>

  <Col>
    To Do
    {#each tasksData as task}
      {#if task.task_state === "ToDo"}
        <Card>
          <span slot="task-name">{task.task_name}</span>
          <span slot="task-owner">{task.task_owner}</span>
          <span slot="task-description">{task.task_description}</span>
          <Row slot="task-actions">
            {#if IsPermitToDo}
              <Col><Button on:update-task><UpdateTask /></Button></Col>
              <Col
                ><Button
                  color="primary"
                  on:click={() => promoteTask(task.task_name, "Doing")}
                  >&#8594;</Button
                ></Col
              >
            {/if}
          </Row>
        </Card>
      {/if}
    {/each}
  </Col>

  <!-- @TODO: Fix delay in ui because of sending email -->
  <Col>
    Doing
    {#each tasksData as task}
      {#if task.task_state === "Doing"}
        <Card>
          <span slot="task-name">{task.task_name}</span>
          <span slot="task-owner">{task.task_owner}</span>
          <span slot="task-description">{task.task_description}</span>
          <Row slot="task-actions">
            {#if IsPermitDoing}
              <Col
                ><Button
                  color="primary"
                  on:click={() => demoteTask(task.task_name, "ToDo")}
                  >&#8592;</Button
                ></Col
              >
              <Col><Button on:update-task><UpdateTask /></Button></Col>
              <Col
                ><Button
                  color="primary"
                  on:click={() => promoteTask(task.task_name, "Done")}
                  >&#8594;</Button
                ></Col
              >
            {/if}
          </Row>
        </Card>
      {/if}
    {/each}
  </Col>

  <Col>
    Done
    {#each tasksData as task}
      {#if task.task_state === "Done"}
        <Card>
          <span slot="task-name">{task.task_name}</span>
          <span slot="task-owner">{task.task_owner}</span>
          <span slot="task-description">{task.task_description}</span>
          <Row slot="task-actions">
            {#if IsPermitDone}
              <Col
                ><Button
                  color="primary"
                  on:click={() => demoteTask(task.task_name, "Doing")}
                  >&#8592;</Button
                ></Col
              >
              <Col><Button on:update-task><UpdateTask /></Button></Col>
              <Col
                ><Button
                  color="primary"
                  on:click={() => promoteTask(task.task_name, "Closed")}
                  >&#8594;</Button
                ></Col
              >
            {/if}
          </Row>
        </Card>
      {/if}
    {/each}
  </Col>

  <Col>
    Close
    {#each tasksData as task}
      {#if task.task_state === "Closed"}
        <Card>
          <span slot="task-name">{task.task_name}</span>
          <span slot="task-owner">{task.task_owner}</span>
          <span slot="task-description">{task.task_description}</span>
        </Card>
      {/if}
    {/each}
  </Col>
</Row>

<!-- Modal for Create task -->
<Modal isOpen={open} {toggle} {size}>
  <ModalHeader {toggle}>Create Task</ModalHeader>
  <ModalBody>
    <CreateTask
      bind:this={createTaskButton}
      {task_name}
      {task_description}
      {task_notes}
      {task_plan}
      {appacronym}
    />
  </ModalBody>
  <ModalFooter>
    <Button
      style="color: #fffbf0;"
      color="warning"
      on:click={(e) => createTaskButton.handleSubmit(e)}>Create Task</Button
    >
    <Button class="back-button" color="danger" on:click={toggle}>Back</Button>
  </ModalFooter>
</Modal>

<style>
</style>
