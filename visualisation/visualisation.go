package visualisation

import (
	"fmt"
	"image/color"

	"github.com/221bytes/tsp/plan"

	"engo.io/ecs"
	"engo.io/engo"
	"engo.io/engo/common"
)

type DefaultScene struct {
	m *plan.Map
}

var (
	zoomSpeed   float32 = -0.125
	scrollSpeed float32 = 700

	worldWidth  int = 800
	worldHeight int = 800
)

type MyShape struct {
	ecs.BasicEntity
	common.RenderComponent
	common.SpaceComponent
}

func (scene *DefaultScene) DisplayMap() []*MyShape {
	m := scene.m
	rectangles := make([]*MyShape, 0, len(m.Grid)*len(m.Grid[0]))
	for x := 0; x < len(m.Grid); x++ {
		for y := 0; y < len(m.Grid[0]); y++ {
			rect := &MyShape{BasicEntity: ecs.NewBasic()}

			var c color.Color

			switch m.Grid[x][y] {
			case 119:
				fmt.Printf("%s", "w  ")
				c = color.RGBA{0, 0, 255, 255}
				break
			case 32:
				fmt.Printf("%s", "   ")
				c = color.RGBA{0, 255, 0, 255}
				break
			default:
				fmt.Printf("%2d ", m.Grid[x][y])
				c = color.RGBA{255, 0, 0, 255}
			}

			rect.SpaceComponent = common.SpaceComponent{Position: engo.Point{X: float32(y) * 10, Y: float32(x) * 10}, Width: 10, Height: 10}
			rect.RenderComponent = common.RenderComponent{Drawable: common.Rectangle{}, Color: c}
			rectangles = append(rectangles, rect)
		}
		fmt.Printf("\n")
	}
	return rectangles
}

func (*DefaultScene) Preload() {}

// Setup is called before the main loop is started
func (scene *DefaultScene) Setup(w *ecs.World) {
	common.SetBackground(color.RGBA{55, 55, 55, 255})
	w.AddSystem(&common.RenderSystem{})

	// Adding camera controllers so we can verify it doesn't break when we move
	w.AddSystem(common.NewKeyboardScroller(scrollSpeed, engo.DefaultHorizontalAxis, engo.DefaultVerticalAxis))
	w.AddSystem(&common.MouseZoomer{zoomSpeed})
	w.AddSystem(&common.MouseRotator{RotationSpeed: 0.125})

	rectangles := scene.DisplayMap()
	for _, system := range w.Systems() {
		switch sys := system.(type) {
		case *common.RenderSystem:
			for _, rect := range rectangles {
				sys.Add(&rect.BasicEntity, &rect.RenderComponent, &rect.SpaceComponent)
			}
		}
	}

}

func (*DefaultScene) Type() string { return "Game" }

func Run(m *plan.Map) {
	scene := &DefaultScene{m: m}
	opts := engo.RunOptions{
		Title:          "Shapes Demo",
		Width:          worldWidth,
		Height:         worldHeight,
		StandardInputs: true,
		MSAA:           4, // This one is not mandatory, but makes the shapes look so much better when rotating the camera
	}
	engo.Run(opts, scene)
}
