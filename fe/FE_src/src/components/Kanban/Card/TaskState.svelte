<script>
  import { Col, Card, CardHeader, CardTitle, CardBody, CardSubtitle, Button, CardText } from "sveltestrap"
  import Icon from "@iconify/svelte"

  export let title
  export let backgroundcolor = null
</script>

<Card class="mb-3">
  <CardHeader style={`background-color: ${backgroundcolor}`}>
    <CardTitle style="text-align: center; color: white;">{title}</CardTitle>
  </CardHeader>
  <CardBody>
    <CardSubtitle>
      <slot name="button" />
    </CardSubtitle>
    <CardText>
      <slot />
    </CardText>
  </CardBody>
</Card>
