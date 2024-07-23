// files defines functions needed to draw to the canvas

// render the full map onto the canvas, including clicked rectangles
function renderMap() {
    resetCanvas()
    drawRectangles()
}

// draws all rectables based on coordinates in the array
function drawRectangles() {
    clickedRectangles.forEach(rectangle => {
        drawRectangle(rectangle.x, rectangle.y)
    })
}

// fills in one of the rectangles on the map based on x and y coordinates
function drawRectangle(xCoord, yCoord) {
    x = xCoord * gridSize
    y = yCoord * gridSize

    ctx.fillStyle = 'rgba(225,225,0,0.5)';
    ctx.beginPath();
    ctx.fillRect(x, y, gridSize, gridSize);
    ctx.stroke();
    ctx.fillStyle = 'rgba(225,225,255,1.0)';
}

// draws a grey grid on top of the map
function drawBackground() {
    // draw map svg into canvas
    var background = new Image();
    background.src = mapAssetPath;
    background.onload = function () {
        ctx.drawImage(background, 0, 0, canvas.width, canvas.height);
    }

    // draw grid lines
    ctx.strokeStyle = "rgba(0, 0, 0, 0.04)";
    ctx.lineWidth = 1;

    for (let x = 0; x <= canvas.width; x += gridSize) {
        ctx.moveTo(x, 0);
        ctx.lineTo(x, canvas.height);
        ctx.stroke();
    }

    for (let y = 0; y <= canvas.height; y += gridSize) {
        ctx.moveTo(0, y);
        ctx.lineTo(canvas.width, y);
        ctx.stroke();
    }
}
