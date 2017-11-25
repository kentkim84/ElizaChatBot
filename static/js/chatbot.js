$("#myForm").submit(function(event){
    event.preventDefault();
    // Get some values from elements on the page:				
    var $form = $( this ),
    message = $form.find( "input[name='message']" ).val(),
    url = $form.attr("action");

    // user perspective						
    var classAttr = "list-group-item text-right";
    var styleAttr = "background-color: rgba(175, 243, 255, 0.18); border-color: rgba(120, 120, 120, 0.15);"
    botMachine(message, classAttr, styleAttr);

    $.ajax({
        type: "POST",
        url: $(this).attr("action"),
        data: $('#myForm').serialize() // serializes the form's elements.
    }).done(function(data){
        // chatbot perspective        
        var regex = /<.*>/gi;
        var body = data;
        var result = body.replace(regex, "");
        console.log("before trim: " + result);
        console.log("after trim: " + $.trim(result));
        classAttr = "list-group-item text-left";
        styleAttr = "background-color: rgba(243, 175, 180, 0.18); border-color: rgba(120, 120, 120, 0.15);";
        botMachine($.trim(result), classAttr, styleAttr);
    });
    // clean up the input text box
    var form = document.getElementById("myForm");
    form.reset();
});

function botMachine(output, classAttr, styleAttr) {
    // message perspective
    console.log("data in"+output)
    var liNode = document.createElement("li");
    var txtNode = document.createTextNode(output);
    liNode.setAttribute("class", classAttr);
    liNode.setAttribute("style", styleAttr);
    
    // append a list
    var outVal = document.getElementById("displayMessage");
    liNode.appendChild(txtNode);
    outVal.appendChild(liNode);
}