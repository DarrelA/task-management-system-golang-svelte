<script>
    import { Collapse, Navbar, NavbarToggler, NavbarBrand, Nav, NavItem, NavLink } from 'sveltestrap'
    import axios from "axios";
    import { navigate } from 'svelte-routing';

    let isOpen = false;
  
    function handleUpdate(event) {
      isOpen = event.detail.isOpen;
    }

    const handleLogOut = async (e) =>{
    e.preventDefault()
    localStorage.removeItem("username")
    localStorage.removeItem("isAdmin")
    await axios.get("http://localhost:4000/logout", {
        withCredentials: true,
      });
    navigate("/")
  }

</script>
  
<Navbar color="dark" dark expand="md">
  <NavbarBrand href="/home">TMS</NavbarBrand>
  <NavbarToggler on:click={() => (isOpen = !isOpen)} />
  <Collapse {isOpen} navbar expand="md" on:update={handleUpdate}>
    <Nav class="ms-auto" navbar>
      <NavLink href="/user">User Management</NavLink>
      <NavItem>
        <NavLink href="/" on:click={handleLogOut}>Log out</NavLink>
      </NavItem>
    </Nav>
  </Collapse>
</Navbar>