<script>
  import CreateTask from "../components/Kanban/Form/CreateTask.svelte";
  import axios from "axios";
  import { onMount } from "svelte";
  import { Button, Modal, ModalBody, ModalHeader, ModalFooter } from "sveltestrap";
  import Icon from "@iconify/svelte";
  import AdminNavbar from "./Admin/NavBar/IsLoggedInAdmin.svelte";
  import UserNavbar from "./User/Navbar/IsLoggedInUser.svelte";
  import AddApplication from "./Admin/Form/AddApplication.svelte";

  const isAdmin = localStorage.getItem("isAdmin");
  let username = localStorage.getItem("username");

  let openModal = false;
  let size = "lg";
  let addButton;
  const toggle = (e) => {
    e.preventDefault();
    openModal = !openModal;
    size = "xl";
  };

  let applications = [];

  async function fetchApplications() {
    try {
      const response = await axios.get("http://localhost:4000/get-all-applications", { withCredentials: true });
      const data = response.data;

      data.forEach((app) => {
        applications.push(app);
      });
      applications = applications;
    } catch (e) {}
  }

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

  onMount(() => {
    fetchApplications();
  });
</script>

{#if isAdmin === "true"}
  <AdminNavbar />
{:else if isAdmin === "false"}
  <UserNavbar />
  <!-- TO BE DONE BY ALFRED & AMOS -->
  <!-- This is where application(s) will be displayed -->
  <!-- 1. Add App -->
  <!-- 2. Update App -->
  <!-- 3. Display App -->
{/if}

<br />
<br/>

<div class="masthead">
  <h2>Welcome {username} &#x1F642;</h2>
  <p>Do you have any tasks to complete today?</p>
</div>

<h2 style="text-align: center;">Applications</h2>
<div class="applications">
  {#each applications as application}
    <div class="application">
      <h3>{application.app_acronym}</h3>
      <div class="description">
        <p>{application.app_description}</p>
      </div>

      <br />
      <p>{application.start_date}</p>
      <p>{application.end_date}</p>
    </div>
  {/each}

  <div class="add-button">
    <Button style="background-color: #e9c46a; border: none;" size="lg" on:click={toggle}>
      <Icon icon="bi:plus-lg" color="#000" />
    </Button>
  </div>

  <Modal isOpen={openModal} {toggle} {size}>
    <ModalHeader {toggle}>Add Application</ModalHeader>
    <ModalBody>
      <AddApplication bind:this={addButton} />
    </ModalBody>

    <ModalFooter>
      <Button style="color: #fffbf0; background-color: #2a9d8f;" on:click={(e) => addButton.CreateApp(e)}>Add Application</Button>
      <Button class="back-button" color="danger" on:click={toggle}>Back</Button>
    </ModalFooter>
  </Modal>
</div>

<style>
  .masthead {
    border-radius: 20px;
    background-color: grey;
    height: 15%;
    width: 75%;
    margin: 20px auto;
    display: flex;
    flex-direction: column;
    justify-content: center;
    align-items: center;
    text-align: center;
    color: #fffbf0;
  }

  .applications {
    margin: 2% auto;
    display: flex;
    justify-content: center;
    align-items: center;
    gap: 20px;
    flex-direction: row;
  }

  .application {
    background-color: #bde0fe;
    padding: 10px;
    width: 300px;
    height: 200px;
    border-radius: 10px;
    box-shadow: 6px 9px 10px 5px rgba(0, 0, 0, 0.08);
  }

  .description {
    height: 30px;
    width: 100%;
    white-space: nowrap;
    overflow: hidden;
    text-overflow: ellipsis;
  }

  .add-button {
    display: flex;
    justify-content: center;
    align-items: center;
    width: 60px;
    height: 60px;
  }
</style>
