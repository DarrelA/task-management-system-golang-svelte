<script>
  import axios from "axios";
  import { errorToast, successToast } from "../../toast";
  import {  Button, Card, CardBody, CardHeader, CardSubtitle, CardText, CardTitle, Row, Col } from "sveltestrap";
  import Icon from '@iconify/svelte';
</script>   

<!-- TO BE DONE BY BEATRICE -->

<div class="text-center">
  <Card class="mb-3">
    <CardHeader>
      <CardTitle>Plan</CardTitle>
    </CardHeader>
    <CardBody>
      <CardSubtitle>
        <Button>
          <Icon icon="bi:plus-lg" width="25" height="25"/>
        </Button>
      </CardSubtitle>
      <CardText>
          <!-- All plans will be displayed here -->
          <br/>
          
          <Card>
            <CardHeader>
              <CardTitle>Sprint 1</CardTitle>
            </CardHeader>
            <CardBody>
              Description ...
            </CardBody>
          </Card>
      </CardText>
    </CardBody>
  </Card>
</div>

<style>
</style>
