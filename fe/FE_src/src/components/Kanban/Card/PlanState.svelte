<script>
  import {
    Col,
    Card,
    CardHeader,
    CardTitle,
    CardBody,
    CardSubtitle,
    Button,
    CardText,
  } from "sveltestrap";
  import Icon from "@iconify/svelte";

  export let title;
</script>

<Card class="mb-3">
  <CardHeader>
    <CardTitle style="text-align: center;">{title}</CardTitle>
  </CardHeader>
  <CardBody>
    <CardSubtitle>
    </CardSubtitle>
    <CardText>
      <slot />
      <br />
    </CardText>
  </CardBody>
</Card>
