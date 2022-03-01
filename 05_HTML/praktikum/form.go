<!DOCTYPE html>
<html lang="en">
  <head>
    <meta charset="UTF-8" />
    <meta http-equiv="X-UA-Compatible" content="IE=edge" />
    <meta name="viewport" content="width=device-width, initial-scale=1.0" />
    <title>Gisela Ramadinawati</title>
  </head>
  <body>
    <form action="">
      <h1>Buat Akun Baru!</h1>
      <h3>Sign Up Form</h3>
      <div>
        <label for="first-name"> First Name: </label><br />
        <input id="first-name" name="firstName" type="text" />
      </div>
      <br />
      <div>
        <label for="last-name"> Last Name: </label>
        <br />
        <input id="last-name" name="lastName" type="text" />
      </div>
      <br />
      <div>
        <label for="gender"> Gender </label>
        <br />
        <input id="male" name="male" value="male" type="radio" />
        <label for="male">male</label>
        <br />
        <input id="female" name="female" value="female" type="radio" />
        <label for="female">female</label>
        <br />
        <input id="other" name="other" value="other" type="radio" />
        <label for="other">Other</label>
      </div>
      <br />
      <div>
        <label for="gender"> Nasionality </label>
        <br />
        <select class="form-select" aria-label="Default select example">
          <option value="1">Indonesia</option>
          <option value="2">Inggris</option>
          <option value="3">Prancis</option>
        </select>
      </div>
      <br />
      <div class="form-check">
        Language Spoken: <br />
        <input class="form-check-input" type="checkbox" value="" id="flexCheckDefault" />
        <label class="form-check-label" for="flexCheckDefault"> Bahasa Indonesia </label>
      </div>
      <div class="form-check">
        <input class="form-check-input" type="checkbox" value="" id="flexCheckDefault" />
        <label class="form-check-label" for="flexCheckChecked"> English </label>
      </div>
      <div class="form-check">
        <input class="form-check-input" type="checkbox" value="" id="flexCheckDefault" />
        <label class="form-check-label" for="flexCheckChecked"> other </label>
      </div>
      <br />
      <div class="form-floating">
        Bio: <br />
        <textarea class="form-control" id="floatingTextarea2" style="height: 100px"></textarea>
      </div>
      <div>
        <a href="welcome.html" target="_blank"><button>Sign Up</button></a>
      </div>
    </form>
  </body>
</html>
