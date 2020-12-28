function getLength(jsonData) {
    let count = 0;
    for (var i in jsonData) count++;
    return count;
}
function handleBlogResult(resultData) {
    let blogsElement = jQuery("#blogs_info");
    console.log(resultData)
    let record = ""
    for (let i = 0; i < getLength(resultData["data"]); ++i) {
        record += '<a href="/blogs/' + resultData["data"][i]["ID"] + '">'
        + resultData["data"][i]["title"] + "</a>";
    }
    blogsElement.append(record)
}
jQuery.ajax({
    dataType: "json", // Setting return data type
    method: "GET", // Setting request method
    url: "/blogs",
    success: (resultData) => handleBlogResult(resultData) // Setting callback function to handle data returned successfully by the StarsServlet
});