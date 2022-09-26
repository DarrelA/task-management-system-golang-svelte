<script>
  import axios from "axios";
  import { Form, FormGroup, Input, Button } from "sveltestrap";
  import { navigate } from "svelte-routing";
  import { errorToast } from "../toast";

  let username;
  let password;

  async function handleSubmit(e) {
    e.preventDefault();
    const json = { username, password };

    try {
      const response = await axios.post("http://localhost:4000/login", json, {
        withCredentials: true,
      });
      if (response) {
        navigate("http://localhost:3000/home");
        localStorage.setItem("username", json.username);
      }
    } catch (e) {
      errorToast("Invalid login");
      e.response && e.response.data.message ? console.log(e.response.data.message) : console.log(e.message);
    }
  }
</script>

<div class="container-fluid">
  <div class="row no-gutter">
    <!-- The image half -->
    <div class="col-md-6 d-none d-md-flex bg-image" />

    <!-- The content half -->
    <div class="col-md-6 bg-light">
      <div class="login d-flex align-items-center py-5">
        <!-- Demo content-->
        <div class="container">
          <div class="row">
            <div class="col-lg-10 col-xl-7 mx-auto">
              <h3 class="display-4">Login into TMS</h3>
              <p class="text-muted mb-4">Start Managing With Us</p>
              <Form on:submit={handleSubmit}>
                <FormGroup>
                  <div class="form-group mb-3">
                    <Input
                      autofocus
                      id="username"
                      type="text"
                      placeholder="Username"
                      required=""
                      class="form-control rounded-pill border-0 shadow-sm px-4"
                      bind:value={username}
                    />
                  </div>
                </FormGroup>

                <FormGroup>
                  <div class="form-group mb-3">
                    <Input
                      id="password"
                      type="password"
                      placeholder="Password"
                      required=""
                      class="form-control rounded-pill border-0 shadow-sm px-4 text-primary"
                      bind:value={password}
                    />
                  </div>
                </FormGroup>
                <Button type="submit" class="btn btn-primary btn-block text-uppercase mb-2 rounded-pill shadow-sm" color="success">Sign in</Button>
              </Form>
            </div>
          </div>
        </div>
        <!-- End -->
      </div>
    </div>
    <!-- End -->
  </div>
</div>

<style>
  .login {
    min-height: 100vh;
  }

  .bg-image {
    min-height: 100vh;
    background-image: url("https://res.cloudinary.com/mhmd/image/upload/v1555917661/art-colorful-contemporary-2047905_dxtao7.jpg");
    background-size: cover;
    background-position: center center;
  }
</style>
