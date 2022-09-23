<script>
  import axios from "axios";
  import { onMount } from "svelte";
  import { Form, FormGroup, Input, Label, Button, Modal, ModalHeader, ModalFooter, Col, Row, Spinner, ModalBody, Styles } from "sveltestrap";
  import MultiSelect from "svelte-multiselect";

  import { errorToast, successToast } from "../toast";

  let username = "";
  let email = "";
  let password = "";
  let status = "Active";

  let groupsArray = [];
  let selected = [];

  let loading = false;

  async function CreateUser(e) {
    e.preventDefault();

    const loggedInUser = localStorage.getItem("username");
    const json = { loggedInUser, username, email, password, user_group: selected, status };
    try {
      const response = await axios.post("http://localhost:4000/admin-create-user", json, { withCredentials: true });
      loading = true;

      setTimeout(() => {
        if (response.data.error) {
          errorToast(response.data.error);
          loading = false;

          username = "";
          password = "";
          email = "";
          selected = [];
          status = "Active";
        } else if (!response.data.error) {
          successToast("New user created");
          loading = false;

          username = "";
          password = "";
          email = "";
          selected = [];
          status = "Active";
        }
      }, 500);
    } catch (e) {
      console.error(e);
    }
  }

  onMount(() => {
    async function GetUserGroups() {
      try {
        const response = await axios.get("http://localhost:4000/get-user-groups");

        if (response.data.error) {
          console.error(response.data.error);
        } else if (!response.data.error) {
          const usergroups = response.data;

          usergroups.forEach((group) => {
            groupsArray.push(group);
          });
          groupsArray = groupsArray;
        }
      } catch (e) {
        console.error(e);
      }
    }
    GetUserGroups();
  });

  let openModal = false;
  let size;
  const toggle = (e) => {
    e.preventDefault();
    openModal = !openModal;
    size = "xl";

    username = "";
    password = "";
    email = "";
    selected = [];
    status = "Active";
  };
</script>

<Styles />
<div>
  <Button color="primary" on:click={toggle}>Add user</Button>
  <Modal isOpen={openModal} {toggle} {size}>
    <ModalHeader {toggle}>Add user</ModalHeader>

    {#if loading}
      <Spinner size="sm" />
    {/if}

    <ModalBody>
      <Form>
        <Row>
          <Col>
            <FormGroup>
              <Label>Username:</Label>
              <Input placeholder="username" bind:value={username} autofocus />
            </FormGroup>
          </Col>

          <Col>
            <FormGroup>
              <Label>Password:</Label>
              <Input placeholder="password" bind:value={password} />
            </FormGroup>
          </Col>
        </Row>

        <Row>
          <Col>
            <FormGroup>
              <Label>Email:</Label>
              <Input placeholder="example@email.com" bind:value={email} />
            </FormGroup>
          </Col>

          <Col>
            <FormGroup>
              <Label>User group(s)</Label>
              <MultiSelect bind:selected options={groupsArray} allowUserOptions={true} />
            </FormGroup>
          </Col>
        </Row>

        <Row>
          <Col>
            <FormGroup>
              <Input type="select" bind:value={status} placeholder="Status">
                <option>Inactive</option>
                <option>Active</option>
              </Input>
              <!-- <select bind:value={status}>
                  {#if statusPlaceholder}
                    <option value="" disabled selected>{statusPlaceholder}</option>
                  {/if}
  
                  <option value="Active">Active</option>
                  <option value="Inactive">Inactive</option>
                </select> -->
            </FormGroup>
          </Col>
        </Row>

        <ModalFooter>
          <Col>
            <Button on:click={CreateUser} style="background-color: #FCA311; border: none;">Create</Button>
            <Button color="danger" on:click={toggle}>Cancel</Button>
          </Col>
        </ModalFooter>
      </Form>
    </ModalBody>
  </Modal>
</div>

<style>
</style>
