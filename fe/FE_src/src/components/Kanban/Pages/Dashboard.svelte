<script>
    import { Table, Row, Col, Button, Modal, ModalBody, ModalHeader, ModalFooter, Card, CardBody, CardSubtitle, CardText } from "sveltestrap";
    import AdminNavbar from "../../Admin/NavBar/IsLoggedInAdmin.svelte";
    import UserNavbar from "../../User/NavBar/IsLoggedInUser.svelte";
    import CreateTask from "../Form/CreateTask.svelte";
    import MgtApp from "../Pages/MgtApp.svelte";
    import MgtPlan from "../Pages/MgtPlan.svelte";
    import MgtTask from "../Pages/MgtTask.svelte";
    import Icon from '@iconify/svelte';
    import axios from "axios";
    import { errorToast } from '../../toast';

    const isAdmin = localStorage.getItem("isAdmin");
    
    export let appacronym;
    let createTaskButton;
    let toggleButton;
    let size = 'xl';
    let openAddTask = false;

    let task_name = '';
    let task_description = '';
    let task_notes = '';
    let task_plan = '';

    let IsPermitCreate = '';
    export let IsPermitOpen = '';
    export let IsPermitToDo = '';
    export let IsPermitDoing = '';
    export let IsPermitDone = '';

    function toggle(e) {
        toggleButton.toggleAddTask(e);
        openAddTask = !openAddTask;
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
        IsPermitCreate = response.data.IsPermitCreate;
        IsPermitOpen = response.data.IsPermitOpen;
        IsPermitToDo = response.data.IsPermitToDo;
        IsPermitDoing = response.data.IsPermitDoing;
        IsPermitDone = response.data.IsPermitDone;
      }
    } catch (e) {
      e.response && e.response.data.message
        ? errorToast(e.response.data.message)
        : errorToast(e.message);
    }
  };

  $: GetUserAppPermits();
</script>   
  
<style>
    .btn {
        text-align: center;
    }
</style>

{#if isAdmin === "true"}
    <AdminNavbar />
{:else if isAdmin === "false"}
    <UserNavbar />
{/if}

<Col xs = "12">
    <MgtApp appacronym={appacronym} />
</Col>

<div class="btn">
    <Icon icon="bi:plus-lg" width="25" height="25" /> Plan
    {#if IsPermitCreate}
        <Button on:click={(e) => toggle(e)}>
            <Icon icon="bi:plus-lg" width="25" height="25" /> Task
        </Button>
    {/if}
</div>

<br/>

<div class="container-fluid">
    <Row>
        <Col xs="2">
            <MgtPlan appacronym={appacronym} />
        </Col>
        <Col xs="10">
            <MgtTask bind:this={toggleButton} appacronym={appacronym} {IsPermitOpen} {IsPermitToDo} {IsPermitDoing} {IsPermitDone} />
        </Col>
    </Row>
</div>


<!-- Modal for Create Task -->
<Modal isOpen={openAddTask} {toggle} {size}>
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