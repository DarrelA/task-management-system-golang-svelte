<script>
    import { navigate, Route } from 'svelte-routing';
    import AccessDenied from './Pages/AccessDenied.svelte';
    import Login from './Pages/Login.svelte';

    export let path;
    export let component;

    const user = Boolean(localStorage.getItem("isAdmin"))
    $: console.log(user)

</script>

{#if user}
    <Route path={path} component={component} />
{:else}
    <Route path={path} component={AccessDenied} />
{/if}