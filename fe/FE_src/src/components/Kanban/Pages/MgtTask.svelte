<script>
  import Icon from '@iconify/svelte';
  import axios from 'axios';
  import {
    Button,
    Col,
    Modal,
    ModalBody,
    ModalFooter,
    ModalHeader,
    Row,
  } from 'sveltestrap';
  import { errorToast } from '../../toast';
  import Task from '../Card/Task.svelte';
  import TaskState from '../Card/TaskState.svelte';
  import CreateTask from '../Form/CreateTask.svelte';
  import UpdateTask from '../Form/UpdateTask.svelte';

  export let appacronym = null;
  
  let tasksData = [];
  let size = 'xl';
  let openUpdateTask = false;
  let updateTaskButton;

  let task_name = '';
  let task_description = '';
  let task_notes = '';
  let task_plan = '';
  let task_notes_existing;
  let task_state;
  let task_owner;
  let task_creator;
  let canUpdateTask = false;
  
  export let IsPermitOpen = '';
  export let IsPermitToDo = '';
  export let IsPermitDoing = '';
  export let IsPermitDone = '';

  export let appData;
  export let IsPermitCreate;
  // $:  console.log("fk",appData)
  // $:  console.log("dog", appData.app_permitCreate)
  

  async function GetAllTasks() {
    try {
      const response = await axios.get(
        `http://localhost:4000/get-all-tasks?AppAcronym=${appacronym}`,
        { withCredentials: true }
      );

      if (response.data) {
        tasksData = response.data;
      }
    } catch (error) {
      console.log('error');
    }
  }

  const demoteTask = async (task_name, task_state) => {
    const json = { task_app_acronym: appacronym, task_name, task_state };

    try {
      const response = await axios.put(
        'http://localhost:4000/task-state-transition',
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
        'http://localhost:4000/task-state-transition',
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

  export function toggleAddTask(e) {
    e.preventDefault();
    task_name = '';
    task_description = '';
    task_notes = '';
    task_plan = '';
    GetAllTasks();
  }

  function toggleUpdateTask(e) {
    e.preventDefault();
    openUpdateTask = !openUpdateTask;
    GetAllTasks();
  }

  function editTask(taskname) {
    openUpdateTask = !openUpdateTask;
    task_name = taskname;
  }

  $: GetAllTasks();
</script>

<div class="text-center" />
<Row>
  <Col>
    <TaskState title="Open">
      {#each tasksData as task}
        {#if task.task_state === 'Open'}
          <Task color={task.task_color}>
            <span slot="task-name">{task.task_name}</span>
            <span slot="task-owner">{task.task_owner}</span>
            <span slot="task-description" class="line-ellipsis"
              >{task.task_description}</span
            >
            <Row slot="task-actions">
              <Col>
                <Button
                  on:click={() => {
                    canUpdateTask = IsPermitOpen;
                    editTask(task.task_name);
                  }}>{IsPermitOpen ? 'Update Task' : 'Read Task'}</Button
                >
              </Col>
              {#if IsPermitOpen}
                <Col>
                  <Button
                    color="primary"
                    on:click={() => promoteTask(task.task_name, 'ToDo')}
                  >
                    &#8594;
                  </Button>
                </Col>
              {/if}
            </Row>
          </Task>
        {/if}
      {/each}
    </TaskState>
  </Col>

  <Col>
    <TaskState title="To Do">
      <br />
      {#each tasksData as task}
        {#if task.task_state === 'ToDo'}
          <Task color={task.task_color}>
            <span slot="task-name">{task.task_name}</span>
            <span slot="task-owner">{task.task_owner}</span>
            <span slot="task-description" class="line-ellipsis"
              >{task.task_description}</span
            >
            <Row slot="task-actions">
              <Col>
                <Button
                  on:click={() => {
                    canUpdateTask = IsPermitToDo;
                    editTask(task.task_name);
                  }}>{IsPermitToDo ? 'Update Task' : 'Read Task'}</Button
                >
              </Col>
              {#if IsPermitToDo}
                <Col>
                  <Button
                    color="primary"
                    on:click={() => promoteTask(task.task_name, 'Doing')}
                  >
                    &#8594;
                  </Button>
                </Col>
              {/if}
            </Row>
          </Task>
        {/if}
      {/each}
    </TaskState>
  </Col>

  <Col>
    <TaskState title="Doing">
      <br />
      {#each tasksData as task}
        {#if task.task_state === 'Doing'}
          <Task color={task.task_color}>
            <span slot="task-name">{task.task_name}</span>
            <span slot="task-owner">{task.task_owner}</span>
            <span slot="task-description" class="line-ellipsis"
              >{task.task_description}</span
            >
            <Row slot="task-actions">
              {#if IsPermitDoing}
                <Col>
                  <Button
                    color="primary"
                    on:click={() => demoteTask(task.task_name, 'ToDo')}
                  >
                    &#8592;
                  </Button>
                </Col>
              {/if}
              <Col>
                <Button
                  on:click={() => {
                    canUpdateTask = IsPermitDoing;
                    editTask(task.task_name);
                  }}>{IsPermitDoing ? 'Update Task' : 'Read Task'}</Button
                >
              </Col>
              {#if IsPermitDoing}
                <Col>
                  <Button
                    color="primary"
                    on:click={() => promoteTask(task.task_name, 'Done')}
                  >
                    &#8594;
                  </Button>
                </Col>
              {/if}
            </Row>
          </Task>
        {/if}
      {/each}
    </TaskState>
  </Col>

  <Col>
    <TaskState title="Done">
      <br />
      {#each tasksData as task}
        {#if task.task_state === 'Done'}
          <Task color={task.task_color}>
            <span slot="task-name">{task.task_name}</span>
            <span slot="task-owner">{task.task_owner}</span>
            <span slot="task-description" class="line-ellipsis"
              >{task.task_description}</span
            >
            <Row slot="task-actions">
              {#if IsPermitDone}
                <Col>
                  <Button
                    color="primary"
                    on:click={() => demoteTask(task.task_name, 'Doing')}
                  >
                    &#8592;
                  </Button>
                </Col>
              {/if}
              <Col>
                <Button
                  on:click={() => {
                    canUpdateTask = IsPermitDone;
                    editTask(task.task_name);
                  }}>{IsPermitDone ? 'Update Task' : 'Read Task'}</Button
                >
              </Col>
              {#if IsPermitDone}
                <Col>
                  <Button
                    color="primary"
                    on:click={() => promoteTask(task.task_name, 'Closed')}
                  >
                    &#8594;
                  </Button>
                </Col>
              {/if}
            </Row>
          </Task>
        {/if}
      {/each}
    </TaskState>
  </Col>

  <Col>
    <TaskState title="Close">
      <br />
      {#each tasksData as task}
        {#if task.task_state === 'Closed'}
          <Task color={task.task_color}>
            <span slot="task-name">{task.task_name}</span>
            <span slot="task-owner">{task.task_owner}</span>
            <span slot="task-description" class="line-ellipsis"
              >{task.task_description}</span
            >
            <Row slot="task-actions">
              <Col>
                <Button
                  on:click={() => {
                    canUpdateTask = false;
                    editTask(task.task_name);
                  }}>Read Task</Button
                >
              </Col>
            </Row>
          </Task>
        {/if}
      {/each}
    </TaskState>
  </Col>
</Row>

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
      {canUpdateTask}
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
  .line-ellipsis {
    display: -webkit-box;
    -webkit-line-clamp: 1;
    -webkit-box-orient: vertical;
    overflow: hidden;
    text-overflow: ellipsis;
  }
</style>
