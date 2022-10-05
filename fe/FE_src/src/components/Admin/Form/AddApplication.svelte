<script>
  import axios from "axios";
  import { onMount } from "svelte";
  import { Form, FormGroup, Input, Label, Col, Row, Spinner, Styles } from "sveltestrap";
  import { errorToast, successToast } from "../../toast";

  //   import FetchGroups from "../../Home.svelte";

  let app_acronym = "";
  let app_description = "";
  let app_Rnum = "";
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
        }
      }, 500);
    } catch (error) {
      errorToast(error.response.data.message);
    }
  }

  let groups = [];
  let data;

  export async function FetchGroups() {
    try {
      const response = await axios.get("http://localhost:4000/get-user-groups", { withCredentials: true });
      console.log(response);
      //   data.forEach((group) => {
      //     groups.push(group);
      //   });
      //   groups = groups;
    } catch (error) {}
  }

  //   $: FetchGroups();

  onMount(() => {});
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
      <Col />

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
          <Label>Permit Create:</Label>
          <Input type="select" bind:value={app_permitCreate}>
            <option>1</option>
          </Input>
        </FormGroup>
      </Col>

      <Col>
        <FormGroup>
          <Label>Permit Open:</Label>
          <Input type="select" bind:value={app_permitOpen} />
        </FormGroup>
      </Col>

      <Col>
        <FormGroup>
          <Label>Permit To Do:</Label>
          <Input type="select" bind:value={app_permitTodo} />
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
