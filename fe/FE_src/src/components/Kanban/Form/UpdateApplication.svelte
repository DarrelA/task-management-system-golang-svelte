<script>
  import axios from "axios";
  import { onMount } from "svelte";
  import { Form, FormGroup, Input, Label, Col, Row, Styles, Icon, Accordion, AccordionItem, Button, Modal, ModalHeader, ModalBody } from "sveltestrap";
  import { errorToast, successToast } from "../../toast";

  let app_startDate = ""
  let app_endDate = ""
  let app_permitCreate = ""
  let app_permitOpen = ""
  let app_permitToDo = ""
  let app_permitDoing = ""
  let app_permitDone = ""
  let app_acronym = ""

  let size = "xl";
  let open = false; 

  export async function handleSubmit(event) {
    event.preventDefault()
    const json = {app_acronym, app_startDate, app_endDate, app_permitCreate, app_permitOpen, app_permitToDo, app_permitDoing, app_permitDone}
    try {
      const response = await axios.post("http://localhost:4000/update-application", json, {withCredentials: true});
      setTimeout(() => {
        if (response) {
         successToast(response.data.message);
         app_startDate = ""
         app_endDate = ""
         app_permitCreate = ""
         app_permitOpen = ""
         app_permitToDo = ""
         app_permitDoing = "" 
         app_permitDone = ""
        }
      }, 500) 
    } catch(error) {
      errorToast(error.response.data.message)
    }
  }

  async function GetApplicationData() {
    try {
      const response = await axios.get("http://localhost:4000/get-application", { withCredentials: true });
      console.log(response);
     
    } catch (error) {}
  }
  
  function handleClick() {
    open = !open
  }

  function toggle(e) {
    e.preventDefault()
    open = !open;
  }
</script>

<style>
</style>

<Button style="font-weight: bold; color: black"; color="warning" on:click={handleClick} >Update Application</Button>

<Modal isOpen={open} {toggle} {size}>
  <ModalHeader {toggle}>Update Application</ModalHeader>
  <ModalBody>
    <Form>
      <Row>
        <Col>
          <FormGroup>
            <Label>Application Name</Label>
            <Input type="text" bind:value={app_acronym} placeholder="Application Name" />
          </FormGroup>
        </Col>
        <Col>
          <FormGroup>
            <Label>Application Description</Label>
            <Input type="text" bind:value={}
          </FormGroup>
        </Col>
      </Row>
    </Form>
  </ModalBody>
</Modal>




