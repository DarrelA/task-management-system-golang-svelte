<script>
  //export let name;
  import { SvelteToast } from '@zerodevx/svelte-toast';
  import { Route, Router } from 'svelte-routing';
  import AddUserToGroup from './components/Pages/AddUserToGroup.svelte';
  import Home from './components/Pages/Home.svelte';
  import Login from './components/Pages/Login.svelte';
  import MgtGroup from './components/Pages/MgtGroup.svelte';
  import MgtUser from './components/Pages/MgtUser.svelte';
  import User from './components/Pages/UpdateUser.svelte';
  import ProtectedRoute from './components/ProtectedRoute.svelte';
</script>

<SvelteToast />

<main>
  <Router>
    <Route path="/" component={Login} />

    <!-- Need to change on protected route -->
    <ProtectedRoute path="/add-user-to-group" component={AddUserToGroup} />
    <ProtectedRoute path="/home" component={Home} />
    <ProtectedRoute path="/user" component={User} />
    <ProtectedRoute path="/user-management" component={MgtUser} />
    <ProtectedRoute path="/group-management" component={MgtGroup} />
    <ProtectedRoute path="*" component={Login} />
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
