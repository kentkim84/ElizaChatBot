// get the user input value and create a node		
var outVal = document.getElementById("displayUserMessage");
var inVal = document.getElementById("inputUserMessage");

$("#myform").submit(function(event){
    event.preventDefault();
    bot();

    // clean up the input text box
    inVal.value = "";
});

function bot() {
    var liNode = document.createElement("li");
    liNode.setAttribute("class", "list-group-item text-right");
    liNode.setAttribute("style", "background-color: rgba(175, 243, 255, 0.18); border-color: rgba(120, 120, 120, 0.15);");
    var txtNode = document.createTextNode(inVal.value);

    // append a list
    liNode.appendChild(txtNode);
    outVal.appendChild(liNode);
}