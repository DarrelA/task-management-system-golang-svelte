<script>
    import axios from "axios";
    import { errorToast, successToast } from "../toast";
    import { Form, FormGroup, Input, Button, Modal, ModalBody, ModalHeader, ModalFooter } from "sveltestrap";
    
    export let groupname;

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
      } catch (e) {
        errorToast(e.response.data.message);
      }
    } 
</script>
  
<style>
    /* input {
      color: purple;
    } */
</style>

<Form on:submit={handleAddGroup}>
   <FormGroup>
    <label for="group">Groupname</label>
    <Input placeholder="groupname" type="text" bind:value={groupname} autofocus />
  </FormGroup>
</Form>
