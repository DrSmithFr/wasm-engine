package browser

type CssProperties interface {
    SetProperty(property, value string)
    GetProperty(property string) string
    Style() string
}

type Css struct {
    properties map[string]string
}

func (c *Css) SetProperty(property, value string) {
    old := c.properties[property]

    if old != value {
        c.properties[property] = value
    }
}

func (c *Css) GetProperty(property string) string {
    return c.properties[property]
}

func (c *Css) Style() string {
    style := ""

    for property, value := range c.properties {
        style += property + ":" + value + ";"
    }

    return style
}
