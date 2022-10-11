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
    let toggleAddTaskBtn;
    let toggleAddPlanBtn;
    
    let size = 'xl';
    let openAddTask = false;
    let openAddPlan = false;

    let task_name = '';
    let task_description = '';
    let task_notes = '';
    let task_plan = '';

    let checkPM = false;
    let IsPermitCreate = '';
    export let IsPermitOpen = '';
    export let IsPermitToDo = '';
    export let IsPermitDoing = '';
    export let IsPermitDone = '';

    function toggleTask(e) {
        toggleAddTaskBtn.toggleAddTask(e);
        openAddTask = !openAddTask;
    }

    function togglePlan(e) {
        toggleAddPlanBtn.toggleCreatePlan(e);
        openAddPlan = !openAddPlan;
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

  async function GetPMCheck() {
    try {
      const response = await axios.get(
        `http://localhost:4000/get-all-plans?AppAcronym=${appacronym}`,
        { withCredentials: true }
      );

      if (response.data) {
        checkPM = response.data.checkPM;
      }
    } catch (error) {
      console.log("error in get plans");
    }
  };


  $: GetUserAppPermits();

  const callBackfetch = (event) => {
    GetUserAppPermits()
  } 

  $: GetPMCheck();
</script>   
  
<style>
</style>

{#if isAdmin === "true"}
    <AdminNavbar />
{:else if isAdmin === "false"}
    <UserNavbar />
{/if}

<Col xs = "12">
    <MgtApp  appacronym={appacronym}  on:fetch={callBackfetch}/>
</Col>

<div class="container-fluid">
  {#if checkPM}
    <Button on:click={(e) => togglePlan(e)}>
      <Icon icon="bi:plus-lg" width="15" height="15" /> Plan
    </Button>
  {/if}

  {#if IsPermitCreate}
    <Button on:click={(e) => toggleTask(e)}>
      <Icon icon="bi:plus-lg" width="15" height="15" /> Task
    </Button>
  {/if}
</div>

<br/>

<div class="container-fluid">
    <Row>
        <Col xs="2">
            <MgtPlan bind:this={toggleAddPlanBtn} {appacronym} />
        </Col>
        <Col xs="10">
            <MgtTask bind:this={toggleAddTaskBtn} {appacronym} {IsPermitOpen} {IsPermitToDo} {IsPermitDoing} {IsPermitDone} />
        </Col>
    </Row>
</div>


<!-- Modal for Create Task -->
<Modal isOpen={openAddTask} {toggleTask} {size}>
    <ModalHeader {toggleTask}>Create Task</ModalHeader>
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
      <Button class="back-button" color="danger" on:click={toggleTask}>Back</Button>
    </ModalFooter>
  </Modal>