<script>
  import { Card, CardBody, CardHeader, CardSubtitle, CardText, CardTitle, Col, Button } from "sveltestrap"

  export let color = null
  export let backgroundcolor = null
</script>

<Col>
  <Card class="mb-3">
    <CardHeader style={`background-color: ${backgroundcolor}`}>
      <CardSubtitle style="font-size: 20px; color: white;">
        <slot name="task-name" />
        <slot name="button" />
      </CardSubtitle>
    </CardHeader>
    <CardBody style={color ? `border-left:5px solid ${color}` : ""}>
      <CardSubtitle>
        <slot name="task-owner" />
      </CardSubtitle>
      <CardText>
        <slot name="task-description" />
      </CardText>
      <slot name="task-actions" />
    </CardBody>
  </Card>
</Col>
