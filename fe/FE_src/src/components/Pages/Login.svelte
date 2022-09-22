<script>
  import axios from 'axios';
  import { navigate } from 'svelte-routing';

  let message = '';
  let username;
  let password;

  async function handleSubmit() {
    const json = { username, password };

    try {
      const response = await axios.post('http://localhost:4000/login', json, {
        withCredentials: true,
      });
      if (response) {
        navigate('http://localhost:3000/home');
      }
    } catch (error) {
      error.response && error.response.data.message
        ? console.log(error.response.data.message)
        : console.log(error.message);
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
              <form on:submit|preventDefault={handleSubmit}>
                <div class="form-group mb-3">
                  <input
                    id="username"
                    type="text"
                    placeholder="Username"
                    required=""
                    autofocus=""
                    class="form-control rounded-pill border-0 shadow-sm px-4"
                    bind:value={username}
                  />
                </div>
                <div class="form-group mb-3">
                  <input
                    id="password"
                    type="password"
                    placeholder="Password"
                    required=""
                    class="form-control rounded-pill border-0 shadow-sm px-4 text-primary"
                    bind:value={password}
                  />
                </div>
                <button
                  type="submit"
                  class="btn btn-primary btn-block text-uppercase mb-2 rounded-pill shadow-sm"
                  >Sign in</button
                >
              </form>
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
    background-image: url('https://res.cloudinary.com/mhmd/image/upload/v1555917661/art-colorful-contemporary-2047905_dxtao7.jpg');
    background-size: cover;
    background-position: center center;
  }
</style>
