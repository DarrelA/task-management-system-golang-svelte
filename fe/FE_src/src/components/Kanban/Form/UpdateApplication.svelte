<script>
  import axios from "axios";
  import { onMount, createEventDispatcher } from "svelte";
  import { Form, FormGroup, Input, Label, Col, Row, Styles, Icon, Accordion, AccordionItem, Button, Modal, ModalHeader, ModalBody, Dropdown, DropdownToggle, DropdownMenu, DropdownItem } from "sveltestrap";
  import { errorToast, successToast } from "../../toast";

  const dispatch = createEventDispatcher();

  export let rnumber = 0;
  export let app_description = "";
  export let start_date = "" ;
  export let end_date = "";
  export let app_permitCreate = "";
  export let app_permitOpen = "";
  export let app_permitTodo = "";
  export let app_permitDoing = "";
  export let app_permitDone = "";
  export let app_acronym = "";
  export let appacronym;

  app_acronym = appacronym;

  let appData = "";

  export async function handleSubmit(e) {
    e.preventDefault()
    const json = {app_acronym, app_Rnum:rnumber, app_description, start_date, end_date, app_permitCreate, app_permitOpen, app_permitTodo, app_permitDoing, app_permitDone};
    try {
      console.log(json)
      const response = await axios.post(`http://localhost:4000/update-application?AppAcronym=${appacronym}`, json, { withCredentials: true });  
        if (response) {
         successToast(response.data.message);
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
      //console.log(response);
      response.data.forEach((group) => {
        groups.push(group);
      });
      groups = groups;
      dispatch("fetch", {
        response,
      });
    } catch (error) {}
  }

  // onMount(() => {
  //   FetchGroups();
  // });

  async function GetApplicationData() {
    try {
      const response = await axios.get(`http://localhost:4000/get-application?AppAcronym=${appacronym}`, { withCredentials: true });
      if (response.data.error) {
        console.log(response.data.error);
      } else if (response) {
         appData = response.data
         app_description = appData.app_description;
         start_date = appData.start_date;
         end_date = appData.end_date;
         rnumber = appData.app_Rnum;
         app_permitCreate = appData.app_permitCreate;
         app_permitOpen = appData.app_permitOpen;
         app_permitTodo = appData.app_permitTodo;
         app_permitDoing = appData.app_permitDoing;
         app_permitDone = appData.app_permitDone;
        //console.log(appData)
      }
    } catch (error) {
      console.log(error)
    }
  }
  
  $: FetchGroups()
  $: GetApplicationData()
</script>

<Form>
  <Row>
    <Col>
      <FormGroup>
        <Label>App Acronym:</Label>
        <Input type="text" value={app_acronym} readonly/>
      </FormGroup>
    </Col>

    <Col>
      <FormGroup>
        <Label>Running Number:</Label>
        <Input type="number" value={rnumber} readonly/>
      </FormGroup>
    </Col>
  </Row>

    <Row>
      <Col>
        <FormGroup>
          <Label>Description:</Label>
          <Input type="textarea" placeholder="App Description" rows={5} bind:value={app_description}/>
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
          <Input type="date" bind:value={start_date} />
        </FormGroup>
      </Col>

      <Col xs lg="2">
        <FormGroup>
          <Label>End:</Label>
          <Input type="date" bind:value={end_date} />
        </FormGroup>
      </Col>
    </Row>
</Form>

