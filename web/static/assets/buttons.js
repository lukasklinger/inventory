// resets the map and shows rectangles for an entry
function showRectanglesForEntry(entry) {
    // reset map
    resetRectangles()
    resetCanvas()

    // draw all rectangles
    entry.rectangles.forEach(rect => {
        pushRectangle(rect.x, rect.y)
    })
}

// redirect to edit page
function addEditEntryButton(id) {
    window.location.href = "./entry?id=" + id
}