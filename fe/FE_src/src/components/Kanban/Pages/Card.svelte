<script>
  import {
    Card,
    CardBody,
    CardHeader,
    CardSubtitle,
    CardText,
    CardTitle,
    Col,
    Button,
  } from 'sveltestrap';
  import { createEventDispatcher } from 'svelte';
</script>

<Col>
  <Card class="mb-3">
    <CardHeader>
      <CardTitle>
        <slot name="task-name" />
      </CardTitle>
    </CardHeader>
    <CardBody>
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
