package ctx

import (
	"go-webgl/dom/wasm"
	"syscall/js"
)

type Context2D struct {
	*wasm.Entity
}

func NewContext2D(ctx js.Value) *Context2D {
	if ctx.IsNull() || ctx.IsUndefined() {
		return nil
	}

	return &Context2D{
		Entity: wasm.New(ctx),
	}
}

//
// Context2D methods
//

// IsContextLost Returns a Boolean that is true if the context has been lost, or false otherwise.
func (c *Context2D) IsContextLost() bool {
	return c.Js().Call("isContextLost").Bool()
}

// ClearRect Clears all pixels on the canvas in the given rectangle to transparent black.
func (c *Context2D) ClearRect(x, y, width, height int) {
	c.Js().Call("clearRect", x, y, width, height)
}

// FillRect Draws a filled rectangle at (x, y) position whose size is determined by width and height.
func (c *Context2D) FillRect(x, y, width, height int) {
	c.Js().Call("fillRect", x, y, width, height)
}

// StrokeRect Draws a rectangle at (x, y) position whose size is determined by width and height.
func (c *Context2D) StrokeRect(x, y, width, height int) {
	c.Js().Call("strokeRect", x, y, width, height)
}

// FillText Draws (fills) a text string at the specified coordinates on the canvas.
func (c *Context2D) FillText(text string, x, y int) {
	c.Js().Call("fillText", text, x, y)
}

// StrokeText Draws (strokes) a text string at the specified coordinates on the canvas.
func (c *Context2D) StrokeText(text string, x, y int) {
	c.Js().Call("strokeText", text, x, y)
}

// MeasureText Returns a TextMetrics object that contains information about the measured text.
func (c *Context2D) MeasureText(text string) js.Value {
	return c.Js().Call("measureText", text)
}

// LineWidth Sets the thickness of lines in space units.
func (c *Context2D) LineWidth(width int) {
	c.Js().Set("lineWidth", width)
}

type CapStyle string

const (
	CapStyleButt   CapStyle = "butt"
	CapStyleRound           = "round"
	CapStyleSquare          = "square"
)

// LineCap Sets the appearance of the ends of lines.
func (c *Context2D) LineCap(cap CapStyle) {
	c.Js().Call("lineCap", string(cap))
}

type JoinStyle string

const (
	JoinStyleBevel JoinStyle = "bevel"
	JoinStyleRound           = "round"
	JoinStyleMiter           = "miter"
)

// LineJoin Sets the appearance of the "corners" where lines meet.
func (c *Context2D) LineJoin(join JoinStyle) {
	c.Js().Call("lineJoin", string(join))
}

// MiterLimit Sets the miter limit ratio in space units.
func (c *Context2D) MiterLimit(limit float64) {
	c.Js().Call("miterLimit", limit)
}

// GetLineDash Returns a sequence of numbers that specifies distances to alternately draw a line and a gap (in coordinate space units).
func (c *Context2D) GetLineDash() []float64 {
	sequence := c.Js().Call("getLineDash")
	length := sequence.Length()

	result := make([]float64, length)
	for i := 0; i < length; i++ {
		result[i] = sequence.Index(i).Float()
	}

	return result
}

// SetLineDash Sets the line dash pattern used when stroking lines.
func (c *Context2D) SetLineDash(sequence []float64) {
	c.Js().Call("setLineDash", sequence)
}

// LineDashOffset Sets the line dash pattern offset or "phase" (in coordinate space units).
func (c *Context2D) LineDashOffset(offset float64) {
	c.Js().Call("lineDashOffset", offset)
}

//
// Path methods
//

// BeginPath Starts a new path by emptying the list of sub-paths.
func (c *Context2D) BeginPath() {
	c.Js().Call("beginPath")
}

// ClosePath Attempts to add a straight line from the current point to the start of the current sub-path.
func (c *Context2D) ClosePath() {
	c.Js().Call("closePath")
}

// MoveTo Moves the starting point of a new sub-path to the (x, y) coordinates.
func (c *Context2D) MoveTo(x, y int) {
	c.Js().Call("moveTo", x, y)
}

// LineTo Connects the last point in the sub-path to the x, y coordinates with a straight line.
func (c *Context2D) LineTo(x, y int) {
	c.Js().Call("lineTo", x, y)
}

// BezierCurveTo Adds a cubic Bézier curve to the path.
func (c *Context2D) BezierCurveTo(cp1x, cp1y, cp2x, cp2y, x, y int) {
	c.Js().Call("bezierCurveTo", cp1x, cp1y, cp2x, cp2y, x, y)
}

// QuadraticCurveTo Adds a quadratic Bézier curve to the path.
func (c *Context2D) QuadraticCurveTo(cpx, cpy, x, y int) {
	c.Js().Call("quadraticCurveTo", cpx, cpy, x, y)
}

// Arc Adds an arc to the path with the given control points and radius, connected to the previous point by a straight line.
func (c *Context2D) Arc(x, y, radius, startAngle, endAngle float64, anticlockwise bool) {
	c.Js().Call("arc", x, y, radius, startAngle, endAngle, anticlockwise)
}

// ArcTo Adds an arc to the path with the given control points and radius, connected to the previous point by a straight line.
func (c *Context2D) ArcTo(x1, y1, x2, y2, radius float64) {
	c.Js().Call("arcTo", x1, y1, x2, y2, radius)
}

// Ellipse Adds an ellipse to the path.
func (c *Context2D) Ellipse(x, y, radiusX, radiusY, rotation, startAngle, endAngle float64, anticlockwise bool) {
	c.Js().Call("ellipse", x, y, radiusX, radiusY, rotation, startAngle, endAngle, anticlockwise)
}

// Rect Adds a rectangle to the path.
func (c *Context2D) Rect(x, y, width, height int) {
	c.Js().Call("rect", x, y, width, height)
}

// RoundRect Adds a rectangle with rounded corners to the path.
func (c *Context2D) RoundRect(x, y, width, height, radius int) {
	c.Js().Call("roundRect", x, y, width, height, radius)
}

// Fill Draws a filled rectangle at (x, y) position whose size is determined by width and height.
func (c *Context2D) Fill() {
	c.Js().Call("fill")
}

// Stroke Draws a rectangle at (x, y) position whose size is determined by width and height.
func (c *Context2D) Stroke() {
	c.Js().Call("stroke")
}

func (c *Context2D) SetFillStyle(hex string) {
	c.Js().Set("fillStyle", hex)
}

func (c *Context2D) SetStrokeStyle(hex string) {
	c.Js().Set("strokeStyle", hex)
}

func (c *Context2D) CreateImageData(width int, height int) js.Value {
	return c.Js().Call("createImageData", width, height)
}

func (c *Context2D) PutImageData(img js.Value, x, y int) {
	c.Js().Call("putImageData", img, x, y)
}
