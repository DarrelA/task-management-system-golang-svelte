<script>
  import axios from 'axios'
  import { toast } from '@zerodevx/svelte-toast'
</script>

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

<div class="container-fluid">
  <div class="row no-gutter">
      <!-- The image half -->
      <div class="col-md-6 d-none d-md-flex bg-image"></div>


      <!-- The content half -->
      <div class="col-md-6 bg-light">
          <div class="login d-flex align-items-center py-5">

              <!-- Demo content-->
              <div class="container">
                  <div class="row">
                      <div class="col-lg-10 col-xl-7 mx-auto">
                          <h3 class="display-4">Login into TMS</h3>
                          <p class="text-muted mb-4">Start Managing With Us</p>
                          <form>
                              <div class="form-group mb-3">
                                  <input id="username" type="text" placeholder="Username" required="" autofocus="" class="form-control rounded-pill border-0 shadow-sm px-4">
                              </div>
                              <div class="form-group mb-3">
                                  <input id="password" type="password" placeholder="Password" required="" class="form-control rounded-pill border-0 shadow-sm px-4 text-primary">
                              </div>
                              <div class="custom-control custom-checkbox mb-3">
                                  <input id="customCheck1" type="checkbox" checked class="custom-control-input">
                                  <label for="customCheck1" class="custom-control-label">Remember password</label>
                              </div>
                              <button type="submit" class="btn btn-primary btn-block text-uppercase mb-2 rounded-pill shadow-sm">Sign in</button>
                          </form>
                      </div>
                  </div>
              </div><!-- End -->

          </div>
      </div><!-- End -->

  </div>
</div>


