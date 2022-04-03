document.addEventListener('astilectron-ready', function() {
    // This will listen to messages sent by GO
    astilectron.onMessage(function(message) {
        // Process message
        if (message === "hello") {
            console.log(message)
            return message
        }
    });
})


function addRow(moveList){
    moveUl = document.querySelector("body > ul")

    let moveHTML = "<li class=\"list-group-item-bitch\"><div class=\"row\">";

    for (const button of moveList){
        moveHTML += `<div class=\"column\"><img src=\"images/${button}\" ></div>`
    }

    moveHTML += "</div></li>"

    moveUl.innerHTML += moveHTML
}