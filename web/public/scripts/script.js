document.body.addEventListener("htmx:responseError", function(e) {
    error = e.detail.xhr.response;
    ERRORCONT.innerHTML = error;
    ERRORCONT.classList.remove("visually-hidden");
});
