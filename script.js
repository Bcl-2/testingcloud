var btnLogin = document.getElementById('do-login');
var idLogin = document.getElementById('login');
var username = document.getElementById('username');
var pass = document.getElementById('password')
btnLogin.onclick = async function(){
  await fetch('http://188.235.75.170:8080/user', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json;charset=utf-8'
    },
    body: JSON.stringify(pass.innerText)
  });

}