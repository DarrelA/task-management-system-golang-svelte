<script>
    import { errorToast, successToast } from "../../toast";
    import { Form, FormGroup, Input, Label, Button, Modal, ModalHeader, ModalFooter, Col, Row, Spinner, ModalBody, Styles } from "sveltestrap";
    import MultiSelect from "svelte-multiselect";
    import axios from "axios";

    let task_name = ""
    let task_description = ""
    let task_notes = ""
    let task_state = "Open"
    let task_plan = ""
    let task_app_acronym = "durian"
    
    let username = localStorage.getItem("username")
    let message = ""

    let size = "xl";
    let open = false;

    async function handleSubmit(event) {
        event.preventDefault()
        const json = {task_app_acronym, task_name, task_description, task_notes, task_plan}
        try {
            const response = await axios.post("http://localhost:4000/create-task", json, {withCredentials: true});
            if (response) {
                message = response.data.message;
                successToast(message);
                task_name = ""
                task_description = ""
                task_notes = ""
                task_plan = ""
            }
        } catch(error) {
            errorToast(error.response.data.message);
        }
    }

    function handleClick() {
        open = !open
    }

    function toggle(e) {
        e.preventDefault()
        open = !open;
    }
</script>

<Button style="font-weight: bold; color: black;" color="warning" on:click={handleClick} >Create Task</Button>

<Modal isOpen={open} {toggle} {size}>
    <ModalHeader {toggle}>Create Task</ModalHeader>
    <ModalBody>
        <Form>
            <Row>
                <Col>
                    <FormGroup>
                        <Label>Task Name:</Label>
                        <Input type="text" bind:value={task_name} placeholder="Enter a Task Name" autofocus />
                    </FormGroup>
                </Col>
                <Col>
                    <FormGroup>
                        <Label>Task Name:</Label>
                        <Input type="select" bind:value={task_plan} placeholder="Select a Plan">
                            <option value="Sprint 1">Sprint 1</option>
                            <option value="Sprint 2">Sprint 2</option>
                        </Input>
                    </FormGroup>
                </Col>
            </Row>
            <Row>
                <FormGroup>
                    <Label>Task Description:</Label>
                    <Input type="textarea" bind:value={task_description} placeholder="Task Description" rows={2} />
                </FormGroup>
            </Row>
            <Row>
                <FormGroup>
                    <Label>Task Notes:</Label>
                    <Input type="textarea" bind:value={task_notes} placeholder="Task Notes" rows={2} />
                </FormGroup>
            </Row>
            <Row>
                <Col>
                    <FormGroup>
                        <Label>Task State:</Label>
                        <Input bind:value={task_state} readonly />
                    </FormGroup>
                </Col>
                <Col>
                    <FormGroup>
                        <Label>Task Creator:</Label>
                        <Input bind:value={username} readonly />
                    </FormGroup>
                </Col>
                <Col>
                    <FormGroup>
                        <Label>Task Owner:</Label>
                        <Input bind:value={username} readonly />
                    </FormGroup>
                </Col>
            </Row>
        </Form>
    </ModalBody>

    <ModalFooter>
      <Button style="color: #fffbf0;" color="warning" on:click={(e) => handleSubmit(e)}>Create Task</Button>
      <Button class="back-button" color="danger" on:click={toggle}>Back</Button>
    </ModalFooter>
  </Modal>



