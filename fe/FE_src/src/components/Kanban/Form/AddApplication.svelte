<script>
  import axios from "axios";
  import { onMount, createEventDispatcher } from "svelte";
  import { Form, FormGroup, Input, Label, Col, Row, Spinner, Styles } from "sveltestrap";
  import { errorToast, successToast } from "../../toast";
  import Select from 'svelte-select'

  const dispatch = createEventDispatcher();

  let app_acronym = "";
  let app_description = "";
  let app_Rnum = 1;
  let start_date = "";
  let end_date = "";
  let app_permitCreate = "";
  let app_permitOpen = "";
  let app_permitTodo = "";
  let app_permitDoing = "";
  let app_permitDone = "";

  let loading = false;

  export async function CreateApp(e) {
    e.preventDefault();
    app_permitCreate ? app_permitCreate = app_permitCreate.value : app_permitCreate = ""
    app_permitOpen ? app_permitOpen = app_permitOpen.value : app_permitOpen = ""
    app_permitTodo ? app_permitTodo = app_permitTodo.value : app_permitTodo = ""
    app_permitDoing ? app_permitDoing = app_permitDoing.value : app_permitDoing = ""
    app_permitDone ? app_permitDone = app_permitDone.value : app_permitDone = ""
    const json = {
      app_acronym,
      app_description,
      app_Rnum,
      start_date,
      end_date,
      app_permitCreate,
      app_permitOpen,
      app_permitTodo,
      app_permitDoing,
      app_permitDone,
    };
    try {
      const response = await axios.post("http://localhost:4000/create-new-application", json, { withCredentials: true });
      loading = true;

      setTimeout(() => {
        if (!response.data.error) {
          successToast(response.data.message);
          loading = false;

          app_acronym = "";
          app_description = "";
          app_Rnum = 1;
          start_date = "";
          end_date = "";
          app_permitCreate = "";
          app_permitOpen = "";
          app_permitTodo = "";
          app_permitDoing = "";
          app_permitDone = "";
        }
      }, 500);
    } catch (error) {
      errorToast(error.response.data.message);
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

  $: appItems = groups.map(app => ({
    value: app,
    label: app
  }))

  onMount(() => {
    FetchGroups();
  });
</script>

<Styles />

<div>
  {#if loading}
    <div class="loading-spinner">
      <Spinner size="lg" />
    </div>
  {/if}
  <Form>
    <Row>
      <Col>
        <FormGroup>
          <Label>App Acronym</Label>
          <Input placeholder="apple" bind:value={app_acronym} autofocus />
        </FormGroup>
      </Col>

      <Col>
        <FormGroup>
          <Label>Running Number</Label>
          <Input type="number" placeholder="1" min="1" bind:value={app_Rnum} autofocus />
        </FormGroup>
      </Col>
    </Row>

    <Row>
      <Col>
        <FormGroup>
          <Label>Description:</Label>
          <Input rows="5" type="textarea" placeholder="Description for application" bind:value={app_description} />
        </FormGroup>
      </Col>
    </Row>

    <Row>
      <Col>
        <FormGroup>
          <Label>Create:</Label>
          <Select items={appItems} bind:value={app_permitCreate}></Select>
        </FormGroup>
      </Col>

      <Col>
        <FormGroup>
          <Label>Open:</Label>
          <Select items={appItems} bind:value={app_permitOpen}></Select>
        </FormGroup>
      </Col>

      <Col>
        <FormGroup>
          <Label>To Do:</Label>
          <Select items={appItems} bind:value={app_permitTodo}></Select>
        </FormGroup>
      </Col>

      <Col>
        <FormGroup>
          <Label>Doing:</Label>
          <Select items={appItems} bind:value={app_permitDoing}></Select>
        </FormGroup>
      </Col>

      <Col>
        <FormGroup>
          <Label>Done:</Label>
          <Select items={appItems} bind:value={app_permitDone}></Select>
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
</div>

<style>
  .loading-spinner {
    position: relative;
    left: 50%;
    top: 50%;
  }
</style>
