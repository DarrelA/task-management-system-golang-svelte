<script>
  //export let name;
  import { SvelteToast } from '@zerodevx/svelte-toast';
  import { Route, Router } from 'svelte-routing';
  import Home from './components/Home.svelte';
  import Login from './components/Login.svelte';
  import MgtGroup from './components/Admin/Pages/MgtGroup.svelte';
  import MgtUser from './components/Admin/Pages/MgtUser.svelte';
  import User from './components/User/Pages/MgtUser.svelte';
  import ProtectedRoute from './components/ProtectedRoute.svelte';
</script>

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


<SvelteToast />

<main>
  <Router>
    <Route path="/" component={Login} />

    <!-- Need to change on protected route -->
    <ProtectedRoute path="/home" component={Home} />
    <ProtectedRoute path="/user" component={User} />
    <ProtectedRoute path="/user-management" component={MgtUser} />
    <ProtectedRoute path="/group-management" component={MgtGroup} />
    <ProtectedRoute path="*" component={Login} />
  </Router>
</main>