// working variables
var mapAssetPath
var gridSize
var canvas
var ctx
var clickedRectangles = []
var entries = []

// init website state
function init() {
    canvas = document.getElementById("map")
    ctx = canvas.getContext('2d')

    gridSize = parseInt(document.getElementById("map").getAttribute("data-grid-size"))
    mapAssetPath = document.getElementById("map").getAttribute("data-asset-path")

    drawBackground()
    updateEntries()
    initSearch()
}