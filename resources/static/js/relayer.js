
let app = { //Dict with all app functions
    init: function(){
        //Wait for astilectron to be ready
        document.addEventListener('astilectron-ready', function(){
            document.getElementsByClassName("button iconbutton")[0].addEventListener("click", app.onClickAddRow)
        })
    },
    onClickAddRow: function(){
        //Create new list item, and append it
        if (document.getElementById("movelist").lastElementChild == null || document.getElementById("movelist").lastElementChild.id != 'inputCell'){
            const ul = document.getElementById("movelist")
            const li = document.createElement("li")
            const input = document.createElement("input")

            input.type = "text"
            input.placeholder = "Type combo here"

            li.className = "list-group-item"
            li.id = "inputCell"

            li.appendChild(input)
            ul.appendChild(li)
            input.focus()

            input.addEventListener('keypress', function(e){
                if (e.key == 'Enter')
                    astilectron.sendMessage(input.value, function(movelist){
                        console.log(movelist)
                        ul.removeChild(li)
                        app.addRow(movelist)
                    })
            })
        }
    },
    addRow: function(moveList){
        moveUl = document.querySelector("body > ul")
    
        let moveHTML = "<li class=\"list-group-item\"><div class=\"row\">";
    
        for (const button of moveList){
            moveHTML += `<div class=\"column\"><img src=\"images/${button}\" ></div>`
        }
    
        moveHTML += "</div></li>"
    
        moveUl.innerHTML += moveHTML
    }
}