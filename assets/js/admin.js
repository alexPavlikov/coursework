/**
 * Функция выбора изображения для добавления на страницу Блог
 * @param {HTMLInputElement} elem
 */
function changeImages(elem) {

    let preview = document.querySelector(".image-check");
    if(!preview){
        console.warn("Элемент [.image-check] не найден");
        return;
    }

    for (const file of elem.files) {
        let img = document.createElement("img");
        img.src = window.URL.createObjectURL(file);
        img.width = 100;
        img.height = 100;

        preview.append(img);
    }
}


function createModalAddPost() {
    
    
    let form = document.querySelector(".admin-add");
    if(!form){
        console.warn("Элемент [.admin-add] не найден");
        return;
    }
    let closeMain = document.createElement("i");
    closeMain.classList.add("fas", "fa-times");
    closeMain.onclick = function () {
        // this.parentElement.parentElement.remove();
        this.parentElement.parentElement.classList.remove("visual");
    };

    let blog = document.querySelector(".blog");
    blog.classList.add("visual");
   

    if (form.append(closeMain)) {
        window.parent.location = window.parent.location.href;
    }
    
    
    // document.body.append(form);
}

function createModalDatabase() {
    
    let db = document.querySelector(".admin-database");
    if(!db){
        console.warn("Элемент [.admin-database] не найден");
        return;
    }
    let closeMain = document.createElement("i");
    if(closeMain) {
        closeMain.classList.add("fas", "fa-times");
    }
   
    closeMain.onclick = function () {
        // this.parentElement.parentElement.remove();
        this.parentElement.parentElement.classList.remove("visual");
    };

    let blog = document.querySelector(".db");
    blog.classList.add("visual");
   

    if (db.append(closeMain)) {
        window.parent.location = window.parent.location.href;
    }
    
    
    // document.body.append(form);
}


function createModalDelete() {
    let form = document.querySelector(".admin-del");
    if(!form){
        console.warn("Элемент [.admin-del] не найден");
        return;
    }
    let closeMain = document.createElement("i");
    closeMain.classList.add("fas", "fa-times");
    closeMain.onclick = function () {
        // this.parentElement.parentElement.remove();
        this.parentElement.parentElement.classList.remove("visual");
    };

    let blog = document.querySelector(".bl");
    blog.classList.add("visual");
   

    if (form.append(closeMain)) {
        window.parent.location = window.parent.location.href;
    }
}

function createModalAdd() {
    let form = document.querySelector(".admin-add");
    if(!form){
        console.warn("Элемент [.admin-add] не найден");
        return;
    }
    let closeMain = document.createElement("i");
    closeMain.classList.add("fas", "fa-times");
    closeMain.onclick = function () {
        // this.parentElement.parentElement.remove();
        this.parentElement.parentElement.classList.remove("visual");
    };

    let blog = document.querySelector(".blog");
    blog.classList.add("visual");
   

    if (form.append(closeMain)) {
        window.parent.location = window.parent.location.href;
    }
}
