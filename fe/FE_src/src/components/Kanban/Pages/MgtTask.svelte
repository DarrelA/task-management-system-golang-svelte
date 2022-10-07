<script>
  import axios from "axios";
  import { errorToast, successToast } from "../../toast";
  import { Table, Row, Col, Button, Modal, ModalBody, ModalHeader, ModalFooter, Container, Card, CardBody, CardHeader, CardSubtitle, CardText, CardTitle } from "sveltestrap";
  import CreateTask from "../Form/CreateTask.svelte";
  import UpdateTask from "../Form/UpdateTask.svelte";
  import Task from "../Card/Task.svelte";
  import Icon from '@iconify/svelte';

  export let appacronym = null;
  export let tasksData = [];

  let size = "xl";
  let openAddTask = false;
  let openUpdateTask = false;
  let createTaskButton;
  let updateTaskButton;

  let task_name = "";
  let task_description = "";
  let task_notes = "";
  let task_plan = "";
  let task_notes_existing;
  let task_state;
  let task_owner;
  let task_creator;

  // let IsPermitCreate = "Project Lead"
  // let IsPermitOpen = "Project Manager"
  // let IsPermitToDo = "Team Member"
  // let IsPermitDoing = "Team Member"
  // let IsPermitDone = "Project Lead"

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

  function toggleAddTask(e) {
    e.preventDefault();
    openAddTask = !openAddTask;
    GetAllTasks();
  }

  function toggleUpdateTask(e) {
    e.preventDefault();
    openUpdateTask = !openUpdateTask;
  }

  function editTask(taskname) {
    openUpdateTask = !openUpdateTask;
    task_name = taskname;
  }

  $: GetAllTasks();
  $: GetUserAppPermits();
</script>

{#if IsPermitCreate}
  <Button color="primary" on:click={toggleAddTask}>Create Task</Button>
{/if}

<div class="text-center">
  
</div>
<Row>
  <Col>
    <Card class="mb-3">
      <CardHeader>
        <CardTitle>Open</CardTitle>
      </CardHeader>
      <CardBody>
        <CardSubtitle>
          <Button>
            <Icon icon="bi:plus-lg" width="25" height="25"/>
          </Button>
        </CardSubtitle>
        <CardText>
            <!-- All Open task will be displayed here -->
            <br/>
            {#each tasksData as task}
            {#if task.task_state === "Open"}
              <Task color={task.task_color}>
                <span slot="task-name">{task.task_name}</span>
                <span slot="task-owner">{task.task_owner}</span>
                <span slot="task-description">{task.task_description}</span>
                <Row slot="task-actions">
                  {#if IsPermitOpen}
                    <Col>
                      <Button on:click={() =>editTask(task.task_name)}>Update Task</Button>
                    </Col>
                    <Col>
                      <Button color="primary" on:click={() => promoteTask(task.task_name, "ToDo")}>
                        &#8594;
                      </Button>
                    </Col>
                  {/if}
                </Row>
              </Task>
            {/if}
          {/each}
        </CardText>
      </CardBody>
    </Card>
  </Col>

  <Col>
    <Card class="mb-3">
      <CardHeader>
        <CardTitle>To Do</CardTitle>
      </CardHeader>
      <CardBody>
        <CardText>
            <!-- All To Do task will be displayed here -->
            <br/>
            {#each tasksData as task}
            {#if task.task_state === "ToDo"}
              <Task color={task.task_color}>
                <span slot="task-name">{task.task_name}</span>
                <span slot="task-owner">{task.task_owner}</span>
                <span slot="task-description">{task.task_description}</span>
                <Row slot="task-actions">
                  {#if IsPermitToDo}
                    <Col>
                      <Button on:click={() =>editTask(task.task_name)}>Update Task</Button>
                    </Col>
                    <Col>
                      <Button color="primary" on:click={() => promoteTask(task.task_name, "Doing")}>
                        &#8594;
                      </Button>
                    </Col>
                  {/if}
                </Row>
              </Task>
            {/if}
          {/each}
        </CardText>
      </CardBody>
    </Card>
  </Col>

  <!-- @TODO: Fix delay in ui because of sending email -->
  <Col>
    <Card class="mb-3">
      <CardHeader>
        <CardTitle>Doing</CardTitle>
      </CardHeader>
      <CardBody>
        <CardText>
            <!-- All Doing task will be displayed here -->
            <br/>
            {#each tasksData as task}
            {#if task.task_state === "Doing"}
              <Task color={task.task_color}>
                <span slot="task-name">{task.task_name}</span>
                <span slot="task-owner">{task.task_owner}</span>
                <span slot="task-description">{task.task_description}</span>
                <Row slot="task-actions">
                  {#if IsPermitDoing}
                    <Col>
                      <Button
                        color="primary"
                        on:click={() => demoteTask(task.task_name, "ToDo")}>
                        &#8592;
                      </Button>
                    </Col>
                    <Col>
                      <Button on:click={() =>editTask(task.task_name)}>Update Task</Button>
                    </Col>
                    <Col>
                      <Button
                      color="primary"
                      on:click={() => promoteTask(task.task_name, "Done")}>
                      &#8594;
                    </Button>
                    </Col>
                  {/if}
                </Row>
              </Task>
            {/if}
          {/each}
        </CardText>
      </CardBody>
    </Card>  
  </Col>

  <Col>
    <Card class="mb-3">
      <CardHeader>
        <CardTitle>Done</CardTitle>
      </CardHeader>
      <CardBody>
        <CardText>
            <!-- All Done task will be displayed here -->
            <br/>
            {#each tasksData as task}
            {#if task.task_state === "Done"}
              <Task color={task.task_color}>
                <span slot="task-name">{task.task_name}</span>
                <span slot="task-owner">{task.task_owner}</span>
                <span slot="task-description">{task.task_description}</span>
                <Row slot="task-actions">
                  {#if IsPermitDone}
                    <Col>
                      <Button
                        color="primary"
                        on:click={() => demoteTask(task.task_name, "Doing")}>
                        &#8592;
                      </Button>
                    </Col>
                    <Col>
                      <Button on:click={() =>editTask(task.task_name)}>Update Task</Button>
                    </Col>
                    <Col>
                      <Button
                        color="primary"
                        on:click={() => promoteTask(task.task_name, "Closed")}>
                        &#8594;
                      </Button>
                    </Col>
                  {/if}
                </Row>
              </Task>
            {/if}
          {/each}
        </CardText>
      </CardBody>
    </Card> 
  </Col>

  <Col>
    <Card class="mb-3">
      <CardHeader>
        <CardTitle>Close</CardTitle>
      </CardHeader>
      <CardBody>
        <CardText>
            <!-- All Close task will be displayed here -->
            <br/>
            {#each tasksData as task}
            {#if task.task_state === "Closed"}
              <Task color={task.task_color}>
                <span slot="task-name">{task.task_name}</span>
                <span slot="task-owner">{task.task_owner}</span>
                <span slot="task-description">{task.task_description}</span>
              </Task>
            {/if}
          {/each}
        </CardText>
      </CardBody>
    </Card>
  </Col>
</Row>

<!-- Modal for Create Task -->
<Modal isOpen={openAddTask} {toggleAddTask} {size}>
  <ModalHeader {toggleAddTask}>Create Task</ModalHeader>
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
    <Button class="back-button" color="danger" on:click={toggleAddTask}>Back</Button>
  </ModalFooter>
</Modal>

<!-- Modal for Update Task -->
<Modal isOpen={openUpdateTask} {toggleUpdateTask} {size}>
  <ModalHeader {toggleUpdateTask}>Update Task</ModalHeader>
  <ModalBody>
    <UpdateTask
      bind:this={updateTaskButton}
      {task_name}
      {task_description}
      {task_notes_existing}
      {task_plan}
      {task_state}
      {task_creator}
      {task_owner}
      {appacronym}
    />
  </ModalBody>
  <ModalFooter>
    <Button
      style="color: #fffbf0;"
      color="warning"
      on:click={(e) => updateTaskButton.handleSubmit(e)}>Update Task</Button
    >
    <Button class="back-button" color="danger" on:click={toggleUpdateTask}>Back</Button>
  </ModalFooter>
</Modal>

<style>
</style>
