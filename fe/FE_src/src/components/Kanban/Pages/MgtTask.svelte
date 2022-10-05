<script>
  import axios from "axios";
  import { errorToast, successToast } from "../../toast";
  import {  Button, Card, CardBody, CardFooter, CardHeader, CardSubtitle, CardText, CardTitle, Row, Col } from "sveltestrap";
</script>   

<style>
</style>


<div class="text-center">
  <Row>
    <!-- Task: Open -->
    <Col>
      <Card class="mb-3">
        <CardHeader>
          <CardTitle>Open</CardTitle>
        </CardHeader>
        <CardBody>
          <CardSubtitle>Card subtitle</CardSubtitle>
          <CardText>
            Some quick example text to build on the card title and make up the bulk of
            the card's content.
          </CardText>
          <Button>Button</Button>
        </CardBody>
      </Card>
    </Col>
    <!-- Task: To Do -->
    <Col>
      <Card class="mb-3">
        <CardHeader>
          <CardTitle>To Do</CardTitle>
        </CardHeader>
        <CardBody>
          <CardSubtitle>Card subtitle</CardSubtitle>
          <CardText>
            Some quick example text to build on the card title and make up the bulk of
            the card's content.
          </CardText>
          <Button>Button</Button>
        </CardBody>
      </Card>
    </Col>
    <!-- Task: Doing -->
    <Col>
      <Card class="mb-3">
        <CardHeader>
          <CardTitle>Doing</CardTitle>
        </CardHeader>
        <CardBody>
          <CardSubtitle>Card subtitle</CardSubtitle>
          <CardText>
            Some quick example text to build on the card title and make up the bulk of
            the card's content.
          </CardText>
          <Button>Button</Button>
        </CardBody>
      </Card>
    </Col>
    <!-- Task: Done -->
    <Col>
      <Card class="mb-3">
        <CardHeader>
          <CardTitle>Done</CardTitle>
        </CardHeader>
        <CardBody>
          <CardSubtitle>Card subtitle</CardSubtitle>
          <CardText>
            Some quick example text to build on the card title and make up the bulk of
            the card's content.
          </CardText>
          <Button>Button</Button>
        </CardBody>
      </Card>
    </Col>
    <!-- Task: Closed -->
    <Col>
      <Card class="mb-3">
        <CardHeader>
          <CardTitle>Closed</CardTitle>
        </CardHeader>
        <CardBody>
          <CardSubtitle>Card subtitle</CardSubtitle>
          <CardText>
            Some quick example text to build on the card title and make up the bulk of
            the card's content.
          </CardText>
          <Button>Button</Button>
        </CardBody>
      </Card>
    </Col>
  </Row>
</div>