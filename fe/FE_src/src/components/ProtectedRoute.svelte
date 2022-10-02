<script>
  import { navigate, Route } from 'svelte-routing';
  import Login from './Login.svelte';

  export let path;
  export let component;

  const user = localStorage.getItem('isAdmin');
  $: if (user === null) navigate('/');
</script>

{#if user}
  <Route {path} {component} />
{:else}
  <Route {path} component={Login} />
{/if}
