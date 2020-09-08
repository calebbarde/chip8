package display

type Displayer interface {
	Draw(x, y byte, data []byte)
	Update()
}
