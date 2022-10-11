<script>
  import axios from "axios";
  import {
    Button,
    Row,
    Col,
  } from "sveltestrap";
  import { Modal, ModalHeader, ModalBody, ModalFooter } from "sveltestrap";
  import Icon from "@iconify/svelte";
  import Plan from "../Card/Plan.svelte";
  import AddPlan from "../Form/AddPlan.svelte";
  import PlanState from "../Card/PlanState.svelte";

  export let appacronym = null;
  let plansData = [];

  let openPlanModal = false;
  let createPlanButton;
  let size = "lg";

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
        plansData = response.data.plans;
      }
    } catch (error) {
      console.log("error in get plans");
    }
  }

  export function toggleCreatePlan(e) {
    e.preventDefault();
    randomHexGenerator();
    openPlanModal = !openPlanModal;
    plan_name = "";
    plan_start = "";
    plan_end = "";
    GetAllPlans();
  }

  $: GetAllPlans();

  // Random Hex Color Generator (for tagging plan to task)
  function randomHexGenerator() {
    let useRandomMathGenerator = (Math.random() * 0xffff * 1000000).toString(
      16
    );
    let randomHex = "#" + useRandomMathGenerator.slice(0, 6);
    plan_color = randomHex;
  }
</script>

<PlanState title="Plan">
  {#each plansData as plan}
    <Plan color={plan.plan_color}>
      <span slot="plan-name">{plan.plan_name}</span>

      <span slot="plan-startdate" style="font-size: 15px;">
        <Icon icon="bi:calendar-event" /> {plan.plan_start}
      </span>
      <span slot="plan-enddate" style="font-size: 15px;">
        <Icon icon="bi:calendar-x" /> {plan.plan_end}
      </span>
    </Plan>
  {/each}
</PlanState>

<!-- Modal for Create Plan -->
<Modal isOpen={openPlanModal} {toggleCreatePlan} {size}>
  <ModalHeader {toggleCreatePlan}>Create Plan</ModalHeader>
  <ModalBody>
    <AddPlan bind:this={createPlanButton} {appacronym} {plan_color} />
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
