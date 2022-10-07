<script>
  import axios from "axios";
  import { errorToast, successToast } from "../../toast";
  import {  Button, Card, CardBody, CardHeader, CardSubtitle, CardText, CardTitle, Row, Col } from "sveltestrap";
  import Icon from '@iconify/svelte';
  import Plan from "../Card/Plan.svelte";

  export let appacronym = null;
  let plansData = [];

  async function GetAllPlans() {
    try {
      const response = await axios.get(
        `http://localhost:4000/get-all-plans?AppAcronym=${appacronym}`,
        { withCredentials: true }
      );

      if (response.data) {
        plansData = response.data;
      }
    } catch (error) {
      console.log("error");
    }
  }

  $: GetAllPlans();
</script>   

<div class="text-center">
  <Card class="mb-3">
    <CardHeader>
      <CardTitle>Plan</CardTitle>
    </CardHeader>
    <CardBody>
      <CardSubtitle>
      </CardSubtitle>
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

<style>
</style>
