<script>
  import { navigate } from 'svelte-routing';
  import {
    Collapse,
    Navbar,
    NavbarToggler,
    NavbarBrand,
    Nav,
    NavItem,
    NavLink,
    Dropdown,
    DropdownToggle,
    DropdownMenu,
    DropdownItem
  } from 'sveltestrap'

  import axios from "axios";

  let isOpen = false;

  function handleUpdate(event) {
    isOpen = event.detail.isOpen;
  }

  // Need to do handleLogout
  const handleLogOut = async (e) =>{
    e.preventDefault()
    localStorage.removeItem("username")
    localStorage.removeItem("isAdmin")
    await axios.get("http://localhost:4000/logout", {
        withCredentials: true,
      });
    navigate("/")
  }

  // handled/disabled go back functionality in browser
</script>

<Navbar color="light" light expand="md">
  <NavbarBrand href="/home">TMS</NavbarBrand>
  <NavbarToggler on:click={() => (isOpen = !isOpen)} />
  <Collapse {isOpen} navbar expand="md" on:update={handleUpdate}>
    <Nav class="ms-auto" navbar>
      <Dropdown nav inNavbar>
        <DropdownToggle nav caret>User Management</DropdownToggle>
        <DropdownMenu end>
          <DropdownItem href="/user-management">Accounts Table</DropdownItem>
          <DropdownItem href="/add-user-to-group">Add Users To Group</DropdownItem>
        </DropdownMenu>
      </Dropdown>
      <NavItem>
        <NavLink href="/group-management">Group Management</NavLink>
      </NavItem>
      <NavItem>
        <NavLink href="/" on:click={handleLogOut}>Log out</NavLink>
      </NavItem>
    </Nav>
  </Collapse>
</Navbar>