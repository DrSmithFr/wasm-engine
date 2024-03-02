package canvas

import (
    "syscall/js"
)

type Canvas interface {
    Create(width, height int)
    Bind(canvas *js.Value, width int, height int)
    Js() *js.Value

    SetSize(width, height int)
    Size() (int, int)

    SetCssProperty(property, value string)
}
