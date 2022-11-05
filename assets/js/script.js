/* <script type='text/javascript'>
document.addEventListener("DOMContentLoaded", () => {
 const auth = document.getElementById("auth");
 const main = document.getElementById("main");
 const user = document.getElementById("user");

 user.addEventListener("click", () => {
 main.classList.toggle("is-visible");
 user.classList.toggle("active");
 });
});
</script> */

// <main id="main">
// <label for="">Добро пожаловать!</label>
// <input name="login" id="userLog" type="text" placeholder="Введите ваш логин">
// <input name="password" id="userPass" type="password" placeholder="Введите ваш пароль">
// <div class="b_list">
//     <button name="enter" formaction="/enter" class="pre-order__button">Войти</button>
//     <button name="registration" formaction="/registration" class="pre-order__button">Регистрация</button>
// </div>
//<div class="bgPage"></div>
// </main>

function authWindow() {
    let main =  document.createElement("main");
    main.className = "main";
    main.id = "main";

    let closeMain = document.createElement("i");
    closeMain.classList.add("fas", "fa-times");
    closeMain.onclick = function () {
        this.parentElement.parentElement.remove();
    };

    let welcome = document.createElement("lable");
    welcome.textContent = "Добро пожаловать!";

    let login = document.createElement("input");
    login.type = "text";
    login.placeholder = "Введите ваш логин";
    login.name = "lgn";

    let password = document.createElement("input");
    password.type = "password";
    password.placeholder = "Введите ваш пароль";
    password.name = "pass";

    let b_list = document.createElement("div");
    b_list.className = "b_list";

    let enter = document.createElement("button");
    enter.textContent = "Войти";
    enter.className = "pre-order__button";
    enter.formAction = "/login";
    enter.onclick = Login.bind(enter, login, password);


    if(enter==true) {
        dialogWindow();
    }


    let registration = document.createElement("a");
    registration.textContent = "Регистрация";
    registration.className = "pre-order__button";
    registration.href = "/registration";

    let bgPage = document.createElement("div");
    bgPage.className = "bgPage";

    b_list.append(enter, registration);

    main.append(closeMain, welcome, login, password, b_list);
    bgPage.append(main);
    document.body.append(bgPage);

}
/** 
* Авторизация log.value.length < 3 && log.value.length > 50
*@param {HTMLInputElement} log
*@param {HTMLInputElement} pass
*/
function Login(log, pass) {

    let isCorrect = true;

    if(!log || log.value.length < 10) {
        log.classList.add("incorrect");
        log.oninput = ClearIncorrect;
        isCorrect = false;
    }

    if(!pass || pass.value.length < 8) {
        pass.classList.add("incorrect");
        pass.oninput = ClearIncorrect;
        isCorrect = false;
    }

    if (!isCorrect) {
        return;
    }

    let xhr = new XMLHttpRequest();
    xhr.open("POST", "/login");
    let data = JSON.stringify({
        Login: log.value,
        Password: pass.value,
    });
    console.log(data);
    xhr.send(data)
     
}

function ClearIncorrect() {
    this.classList.remove("incorrect");
    this,oninput = undefined;
}


function dialogWindow() {
    let dialog =  document.createElement("div");
    dialog.className = "dialWin";
    dialog.id = "dialWin";

    let err = document.createElement("lable");
    err.textContent = `Такого пользователя не существует.
                            Проверьте логин/пароль`;

    let close = document.createElement("button");
    close.textContent = "Закрыть";
    close.className = "pre-order__button";
    close.formAction = "/logerr";


    dialog.append(err, close);
    document.body.append(dialog);

}
