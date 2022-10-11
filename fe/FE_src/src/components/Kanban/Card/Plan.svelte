<script>
  import { Card, CardBody, CardHeader, CardSubtitle, CardText, CardTitle, Col, Row, Button } from "sveltestrap"

  export let color = null
</script>

<Col>
  <Card class="mb-3">
    <CardHeader style="background-color: #A9A9A9;">
      <CardSubtitle style="font-size: 18px; color: white;">
        <slot name="plan-name" />
      </CardSubtitle>
    </CardHeader>
    <CardBody style={color ? `border-left:5px solid ${color}` : ""}>
      <CardSubtitle>
        <Row>
          <Col>
            <slot name="plan-startdate" />
          </Col>
          <Col>
            <slot name="plan-enddate" />
          </Col>
        </Row>
      </CardSubtitle>
    </CardBody>
  </Card>
</Col>
