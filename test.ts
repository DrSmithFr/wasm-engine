const canvas: HTMLCanvasElement = document.getElementById("myCanvas");
const ctx = canvas.getContext("2d");

// Start a new Path
ctx.beginPath();
ctx.moveTo(0, 0);
ctx.lineTo(300, 150);

// Draw the Path
ctx.closePath()
