<script>
    import axios from "axios";
    import { errorToast, successToast } from "../../toast";
    import { Table, Row, Col, Button, Modal, ModalBody, ModalHeader, ModalFooter, Card, CardBody, CardSubtitle, CardText } from "sveltestrap";
    import { navigate } from "svelte-routing";
    import AdminNavbar from "../../Admin/NavBar/IsLoggedInAdmin.svelte";
    import UserNavbar from "../../User/NavBar/IsLoggedInUser.svelte";
    import CreateTask from "../Form/CreateTask.svelte";
    import MgtTask from "./MgtTask.svelte"
    import MgtPlan from "./MgtPlan.svelte"
    import Icon from "@iconify/svelte"

    const isAdmin = localStorage.getItem("isAdmin");
    let addTaskButton;

    export let task_name = ""
    export let task_description = ""
    export let task_notes = ""
    export let task_plan = ""
    export let appacronym;

    let size = "xl";
    let open = false;

    function handleBack() {
        navigate("/home")
    }

    function toggleAddTask(e) {
        e.preventDefault()
        open = !open;
        task_name = ""
        task_description = ""
        task_notes = ""
        task_plan = ""
    }

    let isOpen = false;
    const toggle = () => (isOpen = !isOpen);
</script>   
  
<style>
</style>

{#if isAdmin === "true"}
    <AdminNavbar />
{:else if isAdmin === "false"}
    <UserNavbar />
{/if}

<div class="container-fluid">
    <br/>

    <Card>
        <CardBody style="text-align: center;">
            <Row>
                <Col>
                    <CardSubtitle>Application</CardSubtitle>
                    <CardText>
                        Application Name
                    </CardText>
                </Col>
                <Col>
                    <CardSubtitle>Permit Create</CardSubtitle>
                    <CardText>
                        Permit Create Username
                    </CardText>
                </Col>
                <Col>
                    <CardSubtitle>Permit Open</CardSubtitle>
                    <CardText>
                        Permit Open Username
                    </CardText>
                </Col>
                <Col>
                    <CardSubtitle>Permit To Do</CardSubtitle>
                    <CardText>
                        Permit To Do Username
                    </CardText>
                </Col>
                <Col>
                    <CardSubtitle>Permit Doing</CardSubtitle>
                    <CardText>
                        Permit Doing Username
                    </CardText>
                </Col>
                <Col>
                    <CardSubtitle>Permit Done</CardSubtitle>
                    <CardText>
                        Permit Done Username
                    </CardText>
                </Col>
                <Col>
                    <CardText>
                        <Row>
                            <Col>
                                <Button style="font-weight: bold; color: black;" color="warning">
                                    <Icon icon="bi:pencil-square" width="25" height="25" />
                                </Button>
                            </Col>
                            <Col>
                                <Button style="font-weight: bold; color: black;" color="warning" on:click={handleBack}>
                                    <Icon icon="bi:arrow-left-square" width="25" height="25" />
                                </Button>
                            </Col>
                        </Row>
                    </CardText>
                </Col>
            </Row>
        </CardBody>
    </Card>
</div>

<br/>

<div class="container-fluid">
    <Row>
        <Col xs = "2">
            <MgtPlan appacronym={appacronym} />
        </Col>
<<<<<<< HEAD
        <Col xs = "10">
<<<<<<< HEAD
            <MgtTask />
=======
        <Col>
            <Button style="float:right; font-weight: bold; color: black;  margin-left: 10px;" color="warning" on:click={handleBack}>Back</Button> 
            <!-- <Button style="float:right; font-weight: bold; color: black;" color="warning" on:click={toggleAddTask}>Add Task</Button> -->
            <!-- <Button style="float:right; font-weight: bold; color: black;" color="warning" on:click={toggleUpdateTask}>Update Task</Button> -->
>>>>>>> 1c2ec3ec2d766e7af79147c1ce9bc26807a73b82
=======
            <MgtTask appacronym={appacronym} />
>>>>>>> main
        </Col>
    </Row>
</div>

<<<<<<< HEAD
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
=======
<MgtTask />

>>>>>>> 1c2ec3ec2d766e7af79147c1ce9bc26807a73b82
