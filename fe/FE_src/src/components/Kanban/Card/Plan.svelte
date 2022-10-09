<script>
    import {Card, CardBody, CardHeader, CardSubtitle, CardText, CardTitle, Col, Button } from "sveltestrap";
  
    export let color = null;
  </script>
  
  <Col>
    <Card class="mb-3">
      <CardHeader>
        <CardTitle>
          <slot name="plan-name" />
        </CardTitle>
      </CardHeader>
      <CardBody style={color ? `border-left:5px solid ${color}` : ""}>
        <CardSubtitle>
          <slot name="plan-startdate" />
          <slot name="plan-enddate" />
        </CardSubtitle>
      </CardBody>
    </Card>
  </Col>
  