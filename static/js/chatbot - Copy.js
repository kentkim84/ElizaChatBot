$("#myForm").submit(function(event){
    event.preventDefault();
    // Get some values from elements on the page:				
    var $form = $( this ),
    message = $form.find( "input[name='message']" ).val(),
    url = $form.attr("action");

    // user perspective						
    /* var classAttr = "list-group-item text-right";
    var styleAttr = "background-color: rgba(175, 243, 255, 0.18); border-color: rgba(120, 120, 120, 0.15);"
    botMachine(message, classAttr, styleAttr); */

    $.ajax({
        type: "POST",
        url: $(this).attr("action"),
        data: $('#myForm').serialize(), // serializes the form's elements.
        success: function(data)
        {
            console.log('Data sent');
        },
        error: function (data) {
            console.log('An error occurred.');
            console.log(data);
        }
    }).done(function(data){
        // chatbot perspective
        console.log('Done: '+data);
        /* classAttr = "list-group-item text-left";
        styleAttr = "background-color: rgba(243, 175, 180, 0.18); border-color: rgba(120, 120, 120, 0.15);";
        botMachine(response, classAttr, styleAttr); */
    });
    // clean up the input text box
    var form = document.getElementById("myForm");
    form.reset();
});

function botMachine(output, classAttr, styleAttr) {
    console.log("text in: "+output);
    // message perspective
    var liNode = document.createElement("li");
    var txtNode = document.createTextNode(output);
    liNode.setAttribute("class", classAttr);
    liNode.setAttribute("style", styleAttr);
    
    // append a list
    var outVal = document.getElementById("displayMessage");
    liNode.appendChild(txtNode);
    outVal.appendChild(liNode);
}