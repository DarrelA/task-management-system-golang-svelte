<script>
  import axios from "axios";
  import { errorToast, successToast } from "../../toast";
  import {
    Button,
    Card,
    CardBody,
    CardHeader,
    CardSubtitle,
    CardText,
    CardTitle,
    Row,
    Col,
  } from "sveltestrap";
  import { Modal, ModalHeader, ModalBody, ModalFooter } from "sveltestrap";
  import Icon from "@iconify/svelte";
  import Plan from "../Card/Plan.svelte";
  import AddPlan from "../Form/AddPlan.svelte";

  export let appacronym = null;
  let plansData = [];

  let openPlanModal = false;
  let createPlanButton;
  let buttonVisible = true;
  let size = "xl";

  let plan_name = "";
  let plan_acronym = appacronym;
  let plan_color = "";
  let plan_start = "";
  let plan_end = "";

  async function GetAllPlans() {
    try {
      const response = await axios.get(
        `http://localhost:4000/get-all-plans?AppAcronym=${appacronym}`,
        { withCredentials: true }
      );

      if (response.data) {
        console.log(response.data);
        plansData = response.data;
      }
    } catch (error) {
      console.log("error");
    }
  }

  function toggleCreatePlan(e) {
    e.preventDefault();
    openPlanModal = !openPlanModal;
    plan_name = "";
    plan_color = "";
    plan_start = "";
    plan_end = "";
    GetAllPlans();
  }

  $: GetAllPlans();
</script>

<div class="text-center">
  <Card class="mb-3">
    <CardHeader>
      <CardTitle>Plan</CardTitle>
    </CardHeader>
    <CardBody>
      <!-- create plan test button -->
      <Button size="small" on:click={toggleCreatePlan}>Create plan</Button><br
      /><br />
      <CardSubtitle />
      <CardText>
        <!-- All plans will be displayed here -->
        {#each plansData as plan}
          <Plan color={plan.plan_color}>
            <span slot="plan-name">{plan.plan_name}</span>
            <span slot="plan-startdate">{plan.plan_start}</span>
            <span slot="plan-enddate">{plan.plan_end}</span>
          </Plan>
        {/each}
      </CardText>
    </CardBody>
  </Card>
</div>

<!-- Modal for Create Plan -->
<Modal isOpen={openPlanModal} {toggleCreatePlan} {size}>
  <ModalHeader {toggleCreatePlan}>Create Plan</ModalHeader>
  <ModalBody>
    <AddPlan bind:this={createPlanButton} {appacronym} />
  </ModalBody>
  <ModalFooter>
    <Button
      style="color: #fffbf0;"
      color="warning"
      on:click={(e) => createPlanButton.handleSubmit(e)}
    >
      Create Plan
    </Button>
    <Button class="back-button" color="danger" on:click={toggleCreatePlan}>
      Back
    </Button>
  </ModalFooter>
</Modal>

<style>
</style>
