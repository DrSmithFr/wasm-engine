package browser

import "syscall/js"

type Element struct {
    Js   js.Value
    Css  Css
    Size Size
}
