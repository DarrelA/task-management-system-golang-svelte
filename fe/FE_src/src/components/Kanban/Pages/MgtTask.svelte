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
  import UpdateTask from '../Form/UpdateTask.svelte';
  import ViewTask from "../Form/ViewTask.svelte";

  export let appacronym = null;
  
  let tasksData = [];
  let size = 'xl';
  let openUpdateTask = false;
  let openViewTask = false;
  let updateTaskButton;

  let task_name = '';
  let task_description = '';
  let task_notes = '';
  let task_plan = '';
  let task_notes_existing;
  let task_state;
  let task_owner;
  let task_creator;
  
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

  function toggleViewTask(e) {
    e.preventDefault();
    openViewTask = !openViewTask;
  }

  function viewTask(taskname) {
    openViewTask = !openViewTask;
    task_name = taskname;
  }

  function editTask(taskname) {
    openUpdateTask = !openUpdateTask;
    task_name = taskname;
  }

  $: GetAllTasks();
</script>

<Row>
  <Col>
    <TaskState title="Open">
      {#each tasksData as task}
        {#if task.task_state === 'Open'}
          <Task color={task.task_color} style="text-align: center;">
            <span slot="button" style="float: right;">
              {#if IsPermitOpen}
              <Button size="sm" on:click={() => {editTask(task.task_name);}}>
                <Icon icon="bi:pencil-square" />
              </Button>
              {/if}
              <Button size="sm" on:click={() => {viewTask(task.task_name);}}>
                <Icon icon="fluent:eye-24-regular" />
              </Button>
            </span>
            <span slot="task-name">{task.task_name}</span>
            <span slot="task-owner" style="font-size: 15px;">
              <Icon icon="carbon:user-avatar-filled-alt" /> {task.task_owner}
            </span>
            <span slot="task-description" class="line-ellipsis" style="font-size: 15px;">
              <Icon icon="material-symbols:description-outline-rounded" /> {task.task_description}
            </span>
            <span slot="task-actions" style="float: right;">
              {#if IsPermitOpen}
              <Button size="sm" color="warning" on:click={() => promoteTask(task.task_name, 'ToDo')}>
                <Icon icon="akar-icons:arrow-right" />
              </Button>
              {/if}
              {#if !IsPermitOpen}
              <div class="invisible">
                <Button size="sm" color="warning">
                  <Icon icon="akar-icons:arrow-left" />
                </Button>
                <Button size="sm" color="warning">
                  <Icon icon="akar-icons:arrow-right" />
                </Button>
              </div>
              {/if}
            </span>
          </Task>
        {/if}
      {/each}
    </TaskState>
  </Col>

  <Col>
    <TaskState title="To Do">
      {#each tasksData as task}
        {#if task.task_state === 'ToDo'}
          <Task color={task.task_color}>
            <span slot="button" style="float: right;">
              {#if IsPermitToDo}
              <Button size="sm" on:click={() => {editTask(task.task_name);}}>
                <Icon icon="bi:pencil-square" />
              </Button>
              {/if}
              <Button size="sm" on:click={() => {viewTask(task.task_name);}}>
                <Icon icon="fluent:eye-24-regular" />
              </Button>
            </span>
            <span slot="task-name">{task.task_name}</span>
            <span slot="task-owner" style="font-size: 15px;">
              <Icon icon="carbon:user-avatar-filled-alt" /> {task.task_owner}
            </span>
            <span slot="task-description" class="line-ellipsis" style="font-size: 15px;">
              <Icon icon="material-symbols:description-outline-rounded" /> {task.task_description}
            </span>
            <span slot="task-actions" style="float: right;">
              {#if IsPermitToDo}
              <Button size="sm" color="warning" on:click={() => promoteTask(task.task_name, 'Doing')}>
                <Icon icon="akar-icons:arrow-right" />
              </Button>
              {/if}
              {#if !IsPermitToDo}
              <div class="invisible">
                <Button size="sm" color="warning">
                  <Icon icon="akar-icons:arrow-left" />
                </Button>
                <Button size="sm" color="warning">
                  <Icon icon="akar-icons:arrow-right" />
                </Button>
              </div>
              {/if}
            </span>
          </Task>
        {/if}
      {/each}
    </TaskState>
  </Col>

  <Col>
    <TaskState title="Doing">
      {#each tasksData as task}
        {#if task.task_state === 'Doing'}
          <Task color={task.task_color}>
            <span slot="button" style="float: right;">
              {#if IsPermitDoing}
              <Button size="sm" on:click={() => {editTask(task.task_name);}}>
                <Icon icon="bi:pencil-square" />
              </Button>
              {/if}
              <Button size="sm" on:click={() => {viewTask(task.task_name);}}>
                <Icon icon="fluent:eye-24-regular" />
              </Button>
            </span>
            <span slot="task-name">{task.task_name}</span>
            <span slot="task-owner" style="font-size: 15px;">
              <Icon icon="carbon:user-avatar-filled-alt" /> {task.task_owner}
            </span>
            <span slot="task-description" class="line-ellipsis" style="font-size: 15px;">
              <Icon icon="material-symbols:description-outline-rounded" /> {task.task_description}
            </span>
            <span slot="task-actions" style="float: right;">
              {#if IsPermitDoing}
                <Button size="sm" color="warning" on:click={() => demoteTask(task.task_name, 'ToDo')}>
                  <Icon icon="akar-icons:arrow-left" />
                </Button>
                <Button size="sm" color="warning" on:click={() => promoteTask(task.task_name, 'Done')}>
                  <Icon icon="akar-icons:arrow-right" />
                </Button>
              {/if}
              {#if !IsPermitDoing}
              <div class="invisible">
                <Button size="sm" color="warning">
                  <Icon icon="akar-icons:arrow-left" />
                </Button>
                <Button size="sm" color="warning">
                  <Icon icon="akar-icons:arrow-right" />
                </Button>
              </div>
              {/if}
            </span>
          </Task>
        {/if}
      {/each}
    </TaskState>
  </Col>

  <Col>
    <TaskState title="Done">
      {#each tasksData as task}
        {#if task.task_state === 'Done'}
          <Task color={task.task_color}>
            <span slot="button" style="float: right;">
              {#if IsPermitDone}
              <Button size="sm" on:click={() => {editTask(task.task_name);}}>
                <Icon icon="bi:pencil-square" />
              </Button>
              {/if}
              <Button size="sm" on:click={() => {viewTask(task.task_name);}}>
                <Icon icon="fluent:eye-24-regular" />
              </Button>
            </span>
            <span slot="task-name">{task.task_name}</span>
            <span slot="task-owner" style="font-size: 15px;">
              <Icon icon="carbon:user-avatar-filled-alt" /> {task.task_owner}
            </span>
            <span slot="task-description" class="line-ellipsis" style="font-size: 15px;">
              <Icon icon="material-symbols:description-outline-rounded" /> {task.task_description}
            </span>
            <span slot="task-actions" style="float: right;">
              {#if IsPermitDone}
                <Button size="sm" color="warning" on:click={() => demoteTask(task.task_name, 'Doing')}>
                  <Icon icon="akar-icons:arrow-left" />
                </Button>
                <Button size="sm" color="warning" on:click={() => promoteTask(task.task_name, 'Closed')}>
                  <Icon icon="akar-icons:arrow-right" />
                </Button>
              {/if}
              {#if !IsPermitDone}
              <div class="invisible">
                <Button size="sm" color="warning">
                  <Icon icon="akar-icons:arrow-left" />
                </Button>
                <Button size="sm" color="warning">
                  <Icon icon="akar-icons:arrow-right" />
                </Button>
              </div>
              {/if}
            </span>
          </Task>
        {/if}
      {/each}
    </TaskState>
  </Col>

  <Col>
    <TaskState title="Close">
      {#each tasksData as task}
        {#if task.task_state === 'Closed'}
          <Task color={task.task_color}>
            <span slot="button" style="float: right;">
              <Button size="sm" on:click={() => {viewTask(task.task_name);}}>
                <Icon icon="fluent:eye-24-regular" />
              </Button>
            </span>
            <span slot="task-name">{task.task_name}</span>
            <span slot="task-owner" style="font-size: 15px;">
              <Icon icon="carbon:user-avatar-filled-alt" /> {task.task_owner}
            </span>
            <span slot="task-description" class="line-ellipsis" style="font-size: 15px;">
              <Icon icon="material-symbols:description-outline-rounded" /> {task.task_description}
            </span>
            <span slot="task-actions" style="float: right;">
              <div class="invisible">
                <Button size="sm" color="warning">
                  <Icon icon="akar-icons:arrow-left" />
                </Button>
                <Button size="sm" color="warning">
                  <Icon icon="akar-icons:arrow-right" />
                </Button>
              </div>
            </span>
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

<!-- Modal for View Task -->
<Modal isOpen={openViewTask} {toggleViewTask} {size}>
  <ModalHeader {toggleViewTask}>View Task</ModalHeader>
  <ModalBody>
    <ViewTask 
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
    <Button class="back-button" color="danger" on:click={toggleViewTask}>Back</Button>
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

  .invisible {
    visibility: hidden;
  }
</style>
