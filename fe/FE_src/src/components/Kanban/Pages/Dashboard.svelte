<script>
    import { Table, Row, Col, Button, Modal, ModalBody, ModalHeader, ModalFooter, Card, CardBody, CardSubtitle, CardText } from "sveltestrap";
    import { navigate } from "svelte-routing";
    import AdminNavbar from "../../Admin/NavBar/IsLoggedInAdmin.svelte";
    import UserNavbar from "../../User/NavBar/IsLoggedInUser.svelte";
    import CreateTask from "../Form/CreateTask.svelte";
    import MgtApp from "../Pages/MgtApp.svelte";
    import MgtPlan from "../Pages/MgtPlan.svelte";
    import MgtTask from "../Pages/MgtTask.svelte";
    import Icon from '@iconify/svelte';


    const isAdmin = localStorage.getItem("isAdmin");
    export let appacronym;


</script>   
  
<style>
</style>

{#if isAdmin === "true"}
    <AdminNavbar />
{:else if isAdmin === "false"}
    <UserNavbar />
{/if}

<Col xs = "15">
    <MgtApp appacronym={appacronym} />
</Col>

<br/>

<div class="container-fluid">
    <Row>
        <Col xs = "2">
            <MgtPlan appacronym={appacronym} />
        </Col>
        <Col xs = "10">
            <MgtTask appacronym={appacronym} />
        </Col>
    </Row>
</div>

