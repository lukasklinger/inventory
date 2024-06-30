// working variables
var mapAssetPath
var gridSize
var canvas
var ctx
var clickedRectangles = []
var entries = []
var editEntry = {}

// init website state
function init() {
    canvas = document.getElementById("map");
    ctx = canvas.getContext('2d')

    gridSize = parseInt(document.getElementById("map").getAttribute("data-grid-size"))
    mapAssetPath = document.getElementById("map").getAttribute("data-asset-path")

    drawBackground()
    canvas.addEventListener('click', function (e) { canvasClick(e) }, false)
    load()
}

// load information when editing an existing entry
async function load() {
    let urlParams = new URLSearchParams(window.location.search)
    let id = urlParams.get('id')

    // if and ID is part of the query string, get it from the entries
    if (id) {
        await fetchEntries()

        // find entry
        entries.forEach(entry => {
            if (entry.id == id) {
                editEntry = entry
            }
        });

        // if an entry has been found, load in the details
        if (editEntry.id) {
            document.getElementById("name").textContent = editEntry.name
            document.getElementById("description").textContent = editEntry.description
            document.getElementById("notes").textContent = editEntry.note

            // and draw the rectangles onto the map
            editEntry.rectangles.forEach(rect => {
                pushRectangle(rect.x, rect.y)
            });
        } else {
            alert("Cloud not load entry details!")
        }
    }
}

// save changes back to the database
async function save() {
    let entry = {}
    document.getElementById("save").style.display = "none"
    document.getElementById("loader").style.display = "block"

    // fill object
    entry.name = document.getElementById("name").value
    entry.description = document.getElementById("description").value
    entry.note = document.getElementById("notes").value
    entry.rectangles = clickedRectangles

    // set ID
    if (editEntry.id) {
        entry.id = editEntry.id
        console.log("bloop")
    } else {
        entry.id = "dummy"
    }

    // POST/PUT
    let response = await fetch('./api/entry', {
        method: editEntry.id ? "PUT" : "POST",
        headers: { 'Content-Type': 'application/json; charset=utf-8' },
        body: JSON.stringify(entry)
    })

    if (response.ok) {
        // navigate back to overview page, if everything worked as expected
        window.location.href = '../'
    } else {
        alert(`Failed updating/adding entry, status: ${response.status}, text ${response.text}`)
    }
}

// delete an entry from the database
async function deleteEntry() {
    // only if we are editing an existing entry
    if (editEntry.id) {
        let response = await fetch(`./api/entry?id=${editEntry.id}`, {
            method: "DELETE",
        })

        if (response.ok) {
            window.location.href = '../'
        } else {
            alert(`Failed deleting entry, status: ${response.status}, text ${response.text}`)
        }
    } else {
        // if we are not editing an existing entry, simply go back to the overview
        window.location.href = '../'
    }
}