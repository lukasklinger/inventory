// resets the map and shows rectangles for an entry
function showRectanglesForEntry(entry) {
    // reset selected rectangles
    resetRectangles()

    // draw all rectangles
    entry.rectangles.forEach(rect => {
        pushRectangle(rect.x, rect.y)
    })

    // scroll to map
    window.scrollTo(xCoord, yCoord);
}

// redirect to edit page
function addEditEntryButton(id) {
    window.location.href = "./entry?id=" + id
}