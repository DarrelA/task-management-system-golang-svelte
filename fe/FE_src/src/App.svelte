<script>
  //export let name;
  import { Router, Route } from "svelte-routing";
  import { SvelteToast } from "@zerodevx/svelte-toast";
  import ProtectedRoute from "./components/ProtectedRoute.svelte";
  import Login from "./components/Pages/Login.svelte";
  import Home from "./components/Pages/Home.svelte";
  import MgtUser from "./components/Pages/MgtUser.svelte";
  import MgtGroup from "./components/Pages/MgtGroup.svelte";
  import AddUserToGroup from "./components/Pages/AddUserToGroup.svelte";
</script>

<SvelteToast />

<main>
  <Router>
    <Route path="/" component={Login} />

		<!-- Need to change on protected route -->
		<ProtectedRoute path="/add-user-to-group" component={AddUserToGroup} />
		<ProtectedRoute path="/home" component={Home} />
		<ProtectedRoute path="/user-management" component={MgtUser} />
		<ProtectedRoute path="/group-management" component={MgtGroup} />
	</Router>
</main>

<style>
  /* main {
		text-align: center;
		padding: 1em;
		max-width: 240px;
		margin: 0 auto;
	}

	@media (min-width: 640px) {
		main {
			max-width: none;
		}
	} */
</style>
