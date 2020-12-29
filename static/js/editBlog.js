let path = window.location.pathname;
console.log(path)
let queryString = window.location.search;
console.log(queryString);
let urlParams = new URLSearchParams(queryString);
let saveButton = document.getElementById("btn_save");
let title = document.getElementById("title");
let body = document.getElementById("post");
console.log(saveButton)
if (path === "/edit-blog") {
    jQuery.ajax({
        dataType: "json", // Setting return data type
        method: "GET", // Setting request method
        url: "/blogs/" + urlParams.get('id') ,
        success: (resultData) => handleBlogResult(resultData) // Setting callback function to handle data returned successfully by the StarsServlet
    });
    saveButton.addEventListener("click", handleUpdatePost, false);
} else {
    saveButton.addEventListener("click", handleCreatePost, false);
}

function handleBlogResult(resultData) {
    console.log(resultData)
    title.value = resultData["data"]["title"];
    body.textContent = resultData["data"]["body"];
}
function handleUpdatePost() {
    jQuery.ajax({
        dataType: "json", // Setting return data type
        method: "PUT", // Setting request method
        url: "/blogs" ,
        data: {
            id: urlParams.get('id'),
            title: title.value,
            body: body.textContent
        },
        success: (resultData) => handleUpdatePostResult(resultData) // Setting callback function to handle data returned successfully by the StarsServlet
    });
}
function handleUpdatePostResult(resultData) {
    console.log(resultData)
}
function handleCreatePost() {
    let postData = {
        "title": title.value,
        "body": body.value,
    }
    console.log(postData)
    jQuery.ajax({
        dataType: "json", // Setting return data type
        method: "POST", // Setting request method
        url: "/blogs" ,
        data: JSON.stringify(postData),
        success: (resultData) => handleCreatePostResult(resultData) // Setting callback function to handle data returned successfully by the StarsServlet
    });
}
function handleCreatePostResult(resultData) {
    console.log(resultData)
}