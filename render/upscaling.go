package render

type Resolution struct {
    Width  int
    Height int
}

type Upscale struct {
    Source Resolution
    Target Resolution
}

func (u *Upscale) Scale(x, y float64) (float64, float64) {
    return x * float64(u.Target.Width) / float64(u.Source.Width), y * float64(u.Target.Height) / float64(u.Source.Height)
}

func (u *Upscale) ScaleWidth(x float64) float64 {
    return x * float64(u.Target.Width) / float64(u.Source.Width)
}

func (u *Upscale) ScaleHeight(y float64) float64 {
    return y * float64(u.Target.Height) / float64(u.Source.Height)
}
