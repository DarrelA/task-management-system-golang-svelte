<script>
    import axios from "axios";
    import { errorToast, successToast } from "../../toast";
    import { Form, FormGroup, Input, Label } from "sveltestrap";
    
    let task_name = ""
    let task_description = ""
    let task_notes = ""
    let task_state = ""
    let task_plan = ""
    let task_app_acronym = ""

</script>
  
<style>
</style>

<!-- TO BE DONE BY BEATRICE -->
<Form>
   <FormGroup>
    <Label for="group">Task Name</Label>
    <Input placeholder="groupname" type="text" bind:value={task_name} autofocus />
  </FormGroup>
</Form>
