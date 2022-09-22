<script>
    import { Route } from 'svelte-routing';

    export let path;
    export let component;
</script>

<Route path={path} component={component} />