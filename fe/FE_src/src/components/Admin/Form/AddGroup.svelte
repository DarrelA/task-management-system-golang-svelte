<script>
  import axios from "axios";
  import { errorToast, successToast } from "../../toast";
  import { Form, FormGroup, Input, Label } from "sveltestrap";
    
  let groupname;

  let loggedInUser = localStorage.getItem("username")

  export async function handleAddGroup(e) {
    e.preventDefault();
    const json = { loggedInUser, user_group: groupname };

    try {
      const response = await axios.post("http://localhost:4000/admin-create-group", json, { withCredentials: true });
            
      if (response) {
        successToast(response.data.message);
        groupname = "";
      }
    } catch (error) {
      errorToast(error.response.data.message);
    }
  } 
</script>
  
<style>
</style>

<Form on:submit={handleAddGroup}>
   <FormGroup>
    <Label for="group">Groupname</Label>
    <Input placeholder="groupname" type="text" bind:value={groupname} autofocus />
  </FormGroup>
</Form>
