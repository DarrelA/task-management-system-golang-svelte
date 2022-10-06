<script>
      import {Card, CardBody, CardHeader, CardSubtitle, CardText, CardTitle, Col, Button} from "sveltestrap"
      import {createEventDispatcher} from 'svelte'

      const dispatch = createEventDispatcher()
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
        <slot name="move-left"><Button on:click={() => dispatch("left", "darrel")}></Button></slot>
        <slot name="update-task"><span on:click={() => dispatch("update-task")}></span></slot>
        <slot name="move-right"><Button on:click={() => dispatch("left", "darrel")}></Button></slot>
        </CardBody>
    </Card>
</Col>