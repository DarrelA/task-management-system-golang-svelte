<script>
    import axios from "axios";
    import { errorToast, successToast } from "../../toast";
    import { Table, Row, Col, Button, Modal, ModalBody, ModalHeader, ModalFooter, Card, CardBody, CardSubtitle, CardText } from "sveltestrap";
    import { navigate } from "svelte-routing";
    import UpdateApplication from "../Form/UpdateApplication.svelte";
    import Icon from '@iconify/svelte';

    import {createEventDispatcher} from "svelte"
  import { onMount } from "svelte";

    const dispatch = createEventDispatcher()

    let updateAppButton;

    export let app_description = "";
    export let app_startDate = "";
    export let app_endDate = "";
    export let app_permitCreate = "";
    export let app_permitOpen = "";
    export let app_permitTodo = "";
    export let app_permitDoing = "";
    export let app_permitDone = "";
    export let appacronym;
    let appData = "";
    let isProjectLead = false;

    let size = "lg";
    let open = false;

    function toggle(e) {
        GetApplicationData()
          e.preventDefault()
          open = !open;
          size = "xl";
          app_description = ""
          app_startDate = ""
          app_endDate = ""
          app_permitCreate = ""
          app_permitOpen = ""
          app_permitTodo = ""
          app_permitDoing = ""
          app_permitDone = ""
    }

    function handleBack() {
        navigate("/home")
    }

   export async function GetApplicationData() {
    try {
      const response = await axios.get(`http://localhost:4000/get-application?AppAcronym=${appacronym}`, { withCredentials: true });
      if (response.data.error) {
        console.log(response.data.error);
      } else if (!response.data.error) {
        appData = response.data.applications
        isProjectLead = response.data.isLead
        dispatch("fetch")
      }
    } catch (error) {
      console.log(error)
    }
  }
  
//   $: GetApplicationData()

  onMount(() => {
    GetApplicationData()
  })
</script>   

<!-- TO BE DONE BY AMOS -->
<div class="container-fluid">
  <br/>

  <Card>
    <CardBody style="text-align: center;">
        <Row>
            <Col>
                <CardSubtitle>Application</CardSubtitle>
                <CardText>
                    {appData.app_acronym}
                </CardText>
            </Col>
            <Col>
                <CardSubtitle>Permit Create</CardSubtitle>
                <CardText>
                    {appData.app_permitCreate}
                </CardText>
            </Col>
            <Col>
                <CardSubtitle>Permit Open</CardSubtitle>
                <CardText>
                    {appData.app_permitOpen}
                </CardText>
            </Col>
            <Col>
                <CardSubtitle>Permit To Do</CardSubtitle>
                <CardText>
                    {appData.app_permitTodo}
                </CardText>
            </Col>
            <Col>
                <CardSubtitle>Permit Doing</CardSubtitle>
                <CardText>
                    {appData.app_permitDoing}
                </CardText>
            </Col>
            <Col>
                <CardSubtitle>Permit Done</CardSubtitle>
                <CardText>
                    {appData.app_permitDone}
                </CardText>
            </Col>
            <Col>
                <CardText>
                    <Row>
                        <Col>
                        {#if isProjectLead}
                          <Button style="font-weight: bold; color: black;" color="warning" on:click={toggle}>
                            <Icon icon="bi:pencil-square" width="25" height="25" />
                          </Button>
                          {/if}
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
<br />

<Modal isOpen={open} {toggle} {size}>
  <ModalHeader {toggle}>Update Application</ModalHeader>
  <ModalBody>
      <UpdateApplication bind:this={updateAppButton} {app_startDate} {app_endDate} {app_permitCreate} {app_permitOpen} {app_permitTodo} {app_permitDoing} {app_permitDone} {appacronym} />
  </ModalBody>
  <ModalFooter>
    <Button style="color: #fffbf0;" color="warning" on:click={(e) => updateAppButton.handleSubmit(e)}>Update Application</Button>
    <Button class="back-button" color="danger" on:click={toggle}>Back</Button>
  </ModalFooter>
</Modal>

<style>
</style>