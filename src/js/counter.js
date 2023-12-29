function counterIncrement(dataID) {
    var input = document.getElementById(dataID)
    input.setAttribute("value", Number(input.value) + 1)
    console.log(input.value)  
}

function counterDecrement(dataID) {
    var input = document.getElementById(dataID)
    input.setAttribute("value", Number(input.value) - 1)
    console.log(input.value)
}