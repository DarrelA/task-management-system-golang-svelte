<script>
    import {Card, CardBody, CardHeader, CardSubtitle, CardText, CardTitle, Col, Button } from "sveltestrap";

  export let color = null;
</script>

<Col>
  <Card class="mb-3">
    <CardHeader>
      <CardTitle>
        <slot name="task-name" />
      </CardTitle>
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
