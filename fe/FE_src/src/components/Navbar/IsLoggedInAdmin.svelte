<script>
  import {
    Collapse,
    Navbar,
    NavbarToggler,
    NavbarBrand,
    Nav,
    NavItem,
    NavLink
  } from 'sveltestrap'

  let isOpen = false;

  function handleUpdate(event) {
    isOpen = event.detail.isOpen;
  }

  // Need to do handleLogout

  // handled/disabled go back functionality in browser
</script>

<Navbar color="light" light expand="md">
  <NavbarBrand href="/home">TMS</NavbarBrand>
  <NavbarToggler on:click={() => (isOpen = !isOpen)} />
  <Collapse {isOpen} navbar expand="md" on:update={handleUpdate}>
    <Nav class="ms-auto" navbar>
      <NavItem>
        <NavLink href="/user-management">User Management</NavLink>
      </NavItem>
      <NavItem>
        <NavLink href="/group-management">Group Management</NavLink>
      </NavItem>
      <NavItem>
        <NavLink href="/">Log out</NavLink>
      </NavItem>
    </Nav>
  </Collapse>
</Navbar>