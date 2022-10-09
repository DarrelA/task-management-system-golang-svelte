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
    <CardTitle>{title}</CardTitle>
  </CardHeader>
  <CardBody>
    <CardSubtitle>
      <slot name="button" />
    </CardSubtitle>
    <CardText>
      <slot />
      <br />
    </CardText>
  </CardBody>
</Card>
