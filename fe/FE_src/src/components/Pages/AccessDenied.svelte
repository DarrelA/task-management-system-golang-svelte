<script>
  import { Link } from 'svelte-routing';
  import {
      Collapse,
      Navbar,
      NavbarToggler,
      NavbarBrand,
      Nav,
      NavItem,
      NavLink
    } from 'sveltestrap'
</script>

<Navbar color="light" light expand="md">
  <NavbarBrand href="/home">TMS</NavbarBrand>
</Navbar>
<h1 style="text-align:center">Unauthorized Access 🤕</h1>
<h4 style="text-align:center">
  Go back to <Link to="/">Login</Link>
</h4>
