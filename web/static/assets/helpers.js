// helper functions

// fetch entries and add them to the page
async function updateEntries() {
    await fetchEntries()
    processEntries()
}

// add all entries to the page
function processEntries() {
    // reset entry list
    document.getElementById("entryList").replaceChildren()

    entries.forEach(entry => {
        console.log(entry)
        addEntry(entry)
    });
}

// push new rectangle onto array, render new map
function pushRectangle(xCoord, yCoord) {
    clickedRectangles.push({ x: xCoord, y: yCoord })
    renderMap()
}

// reset all rectangles, render new map
function resetRectangles() {
    clickedRectangles = []
    renderMap()
}

// resets the canvas by redrawing the background
function resetCanvas() {
    ctx.clearRect(0, 0, canvas.width, canvas.height)
    drawBackground()
}

// reacts to new mouse clicks on the canvas
function canvasClick(event) {
    pos = findClickedRectangle(event)
    pushRectangle(pos.x, pos.y)
}

// gets mouse position on the canvas after a click
function getMousePos(canvas, evt) {
    var rect = canvas.getBoundingClientRect(), // abs. size of element
        scaleX = canvas.width / rect.width,    // relationship bitmap vs. element for x
        scaleY = canvas.height / rect.height;  // relationship bitmap vs. element for y

    return {
        x: (evt.clientX - rect.left) * scaleX,   // scale mouse coordinates after they have
        y: (evt.clientY - rect.top) * scaleY     // been adjusted to be relative to element
    }
}

// gets mouse position on canvas and calculates in which rectangle the user clicked
function findClickedRectangle(event) {
    pos = getMousePos(canvas, event)

    xCoord = Math.floor(pos.x / gridSize)
    yCoord = Math.floor(pos.y / gridSize)

    return { x: xCoord, y: yCoord }
}

// adds a div for each entry in the array of entries
function addEntry(entry) {
    let entryDiv = document.createElement("div")
    entryDiv.classList.add("entry")
    entryDiv.id = entry.id
    
    // entry heading
    let entryHeading = document.createElement("h3")
    entryHeading.textContent = entry.name

    // entry body
    let entryDescription = document.createElement("p")
    entryDescription.textContent = entry.description

    // entry view button
    let entryActions = document.createElement("div")
    let view = document.createElement("img")
    view.src = "./assets/icons8-map-96.png"
    view.classList.add("entryButton")
    view.onclick = function(){showRectanglesForEntry(entry)}
    entryActions.appendChild(view)

    // entry edit button
    let edit = document.createElement("img")
    edit.src = "./assets/icons8-edit-384.png"
    edit.classList.add("entryButton")
    edit.onclick = function(){addEditEntryButton(entry.id)}
    entryActions.appendChild(edit)

    // build entry div
    entryDiv.appendChild(entryHeading)
    entryDiv.appendChild(entryDescription)
    entryDiv.appendChild(entryActions)
    document.getElementById("entryList").appendChild(entryDiv)
}