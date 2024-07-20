// file provides functionality for fulltext search on the page
let debounceTimer

// initialize search functionality, attach event listeners
function initSearch() {
    document.getElementById("search").addEventListener(
        "input",
        (event) => {
            const query = event.target.value;
            debounce(() => handleSearchPosts(query), 300);
        },
        false
    );
}

// debouncing helper
function debounce(callback, time) {
    window.clearTimeout(debounceTimer);
    debounceTimer = window.setTimeout(callback, time);
}

// filter list of entries when a query is typed
function handleSearchPosts(query) {
    let searchQuery = query.trim().toLowerCase();
    let entries = document.getElementsByClassName("entry")

    // if the search query is empty, show all entries
    if (!searchQuery) {
        for (let entry of entries) {
            entry.style.display = "table-row"
        }
    }

    // hide entries that do not contain the entered text
    for (let entry of entries) {
        if (entry.innerText.includes(searchQuery)) {
            entry.style.display = "table-row"
        } else {
            entry.style.display = "none"
        }
    }
}