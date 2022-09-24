<script>
  import axios from "axios";
  import { toast } from "@zerodevx/svelte-toast";
  import { errorToast, successToast } from "../toast";
  import Navbar from "../Navbar/IsLoggedInAdmin.svelte";
  import AdminUpdateUser from "./AdminUpdateUser.svelte";
  import AddUserToGroup from "./AddUserToGroup.svelte";
  import AdminUpdateUserForm from "./AdminUpdateUserForm.svelte";

  let groupname = "";
  let message = "";

  async function CreateGroup() {
    const json = { user_group: groupname };
    console.log(json);
    try {
      const response = await axios.post("http://localhost:4000/admin-create-group", json);
      console.log(response.data);
      if (response.data) {
        message = response.data.message;
        successToast(message);
      }
    } catch (e) {
      console.log(e);
    }
  }
</script>

<Navbar />
<h3>Group Management</h3>
<AddUserToGroup />

<style>

</style>
