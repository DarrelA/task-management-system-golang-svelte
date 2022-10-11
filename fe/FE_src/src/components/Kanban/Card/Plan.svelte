<script>
    import {Card, CardBody, CardHeader, CardSubtitle, CardText, CardTitle, Col, Row, Button } from "sveltestrap";
  
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
  