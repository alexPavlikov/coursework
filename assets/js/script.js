
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
    enter.id = "ent";
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

function addPurchase() {
    var select_id = document.getElementById("selectuser");
    select_id.options[select_id.selectedIndex].text;
    var product = document.getElementById("selectproduct");
    select_id.options[select_id.selectedIndex].text;
    let count = document.getElementById('count');

    if(!user || user.value.length < 10) {
        // user.classList.add("incorrect");
        user.oninput = ClearIncorrect;
        isCorrect = false;
    }

    let xhr = new XMLHttpRequest();
    xhr.open("POST", "/purchaseList/addPruchase/buy");
    let purchase = JSON.stringify({
        User: user.value,
        Product: product.value,
        Count: count.value,
    });
    console.log(purchase);
    xhr.send(purchase)
}
