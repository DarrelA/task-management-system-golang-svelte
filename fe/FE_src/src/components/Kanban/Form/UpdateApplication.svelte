<script>
  import axios from "axios";
  import { onMount } from "svelte";
  import { Form, FormGroup, Input, Label, Col, Row, Styles, Icon, Accordion, AccordionItem, Button, Modal, ModalHeader, ModalBody, Dropdown, DropdownToggle, DropdownMenu, DropdownItem } from "sveltestrap";
  import AddGroup from "../../Admin/Form/AddGroup.svelte";
  import { errorToast, successToast } from "../../toast";
  import AddApplication from "./AddApplication.svelte";

  let app_startDate = ""
  let app_endDate = ""
  let app_permitCreate = ""
  let app_permitOpen = ""
  let app_permitToDo = ""
  let app_permitDoing = ""
  let app_permitDone = ""
  let app_acronym = ""
  let permitCreate = ""
  let permitOpen = ""
  let permitToDo = ""
  let permitDoing = ""
  let permitDone = ""
  export let appacronym;

  let appData = "";

  let size = "xl";
  let open = false; 

  export async function handleSubmit(event) {
    event.preventDefault()
    const json = {app_acronym, app_startDate, app_endDate, permitCreate, permitOpen, permitToDo, permitDoing, permitDone}
    try {
      const response = await axios.post("http://localhost:4000/update-application", json, {withCredentials: true});
      setTimeout(() => {
        if (response) {
         successToast(response.data.message);
         app_startDate = ""
         app_endDate = ""
         permitCreate = ""
         permitOpen = ""
         permitToDo = ""
         permitDoing = ""
         permitDone = ""
        }
      }, 500) 
    } catch(error) {
      errorToast(error.response.data.message)
    }
  }

  async function GetApplicationData() {
    try {
      const response = await axios.get(`http://localhost:4000/get-application?AppAcronym=${appacronym}`, { withCredentials: true });
      if (response.data.error) {
        console.log(response.data.error);
      } else if (!response.data.error) {
        appData = response.data
      }
    } catch (error) {
      console.log(error)
    }
  }
  
  $: GetApplicationData()
</script>
 
<style>
</style>

<Form>
  <Row>
    <Col>
      <FormGroup>
        <Label>App Acronym:</Label>
        <Input type="text" value={appData.app_acronym} readonly/>
      </FormGroup>
    </Col>
    <Col>
      <FormGroup>
        <Label>App Running Number:</Label>
        <Input type="text" value={appData.app_Rnum} readonly/>
      </FormGroup>
    </Col>
    <Row>
      <FormGroup>
        <Label>App Description:</Label>
        <Input type="textarea" placeholder="App Description" rows={2} value={appData.app_description} readonly/>
      </FormGroup>
    </Row>
    <Row>
      <Col>
        <FormGroup>
          <Label>App Start Date:</Label>
          <Input type="date" bind:value={app_startDate} />
        </FormGroup>
      </Col>
      <Col>
        <FormGroup>
          <Label>App End Date:</Label>
          <Input type="date" bind:value={app_endDate} />
        </FormGroup>
      </Col>
    </Row>
    <Row>
      <Col>
        <FormGroup>
          <Dropdown>
            <DropdownToggle style="width:100%" caret>App Permit Create</DropdownToggle>
            <DropdownMenu>
              {#each appData as app}
                <DropdownItem on:click={() => (permitCreate = app.app_permitCreate)} placeholder={app}>
                  {app.app_permitCreate} {app.app}
                </DropdownItem>
              {/each}
            </DropdownMenu>
          </Dropdown>
          <FormGroup>
            <Input value={permitCreate} type="text" readonly />
          </FormGroup>
        </FormGroup>
      </Col>
    </Row>
  </Row>
</Form>


