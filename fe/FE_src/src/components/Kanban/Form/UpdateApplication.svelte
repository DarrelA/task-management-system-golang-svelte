<script>
  import axios from "axios";
  import { onMount, createEventDispatcher } from "svelte";
  import { Form, FormGroup, Input, Label, Col, Row, Styles, Icon, Accordion, AccordionItem, Button, Modal, ModalHeader, ModalBody, Dropdown, DropdownToggle, DropdownMenu, DropdownItem } from "sveltestrap";
  import { errorToast, successToast } from "../../toast";

  const dispatch = createEventDispatcher();

  export let app_startDate = "" ;
  export let app_endDate = "";
  export let app_permitCreate = "";
  export let app_permitOpen = "";
  export let app_permitTodo = "";
  export let app_permitDoing = "";
  export let app_permitDone = "";
  export let app_acronym = "";
  export let appacronym;

  let appData = "";

  $: console.log(appacronym);

  export async function handleSubmit(e) {
    e.preventDefault()
    const json = {app_acronym, app_startDate, app_endDate, app_permitCreate, app_permitOpen, app_permitTodo, app_permitDoing, app_permitDone};
    try {
      const response = await axios.post("http://localhost:4000/update-application", json, { withCredentials: true });  
        if (response) {
         successToast(response.data.message);
         app_startDate = "";
         app_endDate = "";
         app_permitCreate = "";
         app_permitOpen = "";
         app_permitTodo = "";
         app_permitDoing = "";
         app_permitDone = "";
         GetApplicationData();
        }
    } catch(error) {
      errorToast(error.response.data.message)
    }
  }

  let groups = [];
  async function FetchGroups() {
    try {
      const response = await axios.get("http://localhost:4000/get-user-groups", { withCredentials: true });
      console.log(response);
      response.data.forEach((group) => {
        groups.push(group);
      });
      groups = groups;
      dispatch("fetch", {
        response,
      });
    } catch (error) {}
  }

  onMount(() => {
    FetchGroups();
  });

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
        <Label>Running Number:</Label>
        <Input type="number" value={appData.app_Rnum} readonly/>
      </FormGroup>
    </Col>
  </Row>

    <Row>
      <Col>
        <FormGroup>
          <Label>Description:</Label>
          <Input type="textarea" placeholder="App Description" rows={5} value={appData.app_description} readonly/>
        </FormGroup>
      </Col>
    </Row>

    <Row>
      <Col>
        <FormGroup>
          <Label>Create:</Label>
          <Input type="select" bind:value={app_permitCreate}>
            {#each groups as group}
              <option>{group}</option>
            {/each}
          </Input>
        </FormGroup>
      </Col>

      <Col>
        <FormGroup>
          <Label>Open:</Label>
          <Input type="select" bind:value={app_permitOpen}>
            {#each groups as group}
              <option>{group}</option>
            {/each}
          </Input>
        </FormGroup>
      </Col>

      <Col>
        <FormGroup>
          <Label>To Do:</Label>
          <Input type="select" bind:value={app_permitTodo}>
            {#each groups as group}
              <option>{group}</option>
            {/each}
          </Input>
        </FormGroup>
      </Col>

      <Col>
        <FormGroup>
          <Label>Doing:</Label>
          <Input type="select" bind:value={app_permitDoing}>
            {#each groups as group}
              <option>{group}</option>
            {/each}
          </Input>
        </FormGroup>
      </Col>
      
      <Col>
        <FormGroup>
          <Label>Done:</Label>
          <Input type="select" bind:value={app_permitDone}>
            {#each groups as group}
              <option>{group}</option>
            {/each}
          </Input>
        </FormGroup>
      </Col>
    </Row>

    <Row class="justify-content-md-center">
      <Col xs lg="2">
        <FormGroup>
          <Label>Start:</Label>
          <Input type="date" bind:value={app_startDate} />
        </FormGroup>
      </Col>

      <Col xs lg="2">
        <FormGroup>
          <Label>End:</Label>
          <Input type="date" bind:value={app_endDate} />
        </FormGroup>
      </Col>
    </Row>
</Form>

