<script>
  import { Col, Card, CardHeader, CardTitle, CardBody, CardSubtitle, Button, CardText } from "sveltestrap"
  import Icon from "@iconify/svelte"

  export let title
</script>

<Card class="mb-3">
  <CardHeader style="background-color: #A9A9A9;">
    <CardTitle style="text-align: center; color: white;">{title}</CardTitle>
  </CardHeader>
  <CardBody>
    <CardSubtitle />
    <CardText>
      <slot />
    </CardText>
  </CardBody>
</Card>
