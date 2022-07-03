package main

import (
	"image/color"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"github.com/notnil/chess"
)

var (
	grid  *fyne.Container
	over  *canvas.Image
	start *canvas.Rectangle
	win   fyne.Window
)

func main() {
	a := app.New()
	win = a.NewWindow("Chess")
	win.Resize(fyne.NewSize(500, 500))

	game := chess.NewGame()
	grid = createGrid(game)

	over = canvas.NewImageFromResource(nil)
	over.Hide()

	start = canvas.NewRectangle(color.Transparent)
	start.StrokeWidth = 4

	win.SetContent(container.NewMax(grid, container.NewWithoutLayout(start, over)))
	win.ShowAndRun()
}

func move(m *chess.Move, game *chess.Game, grid *fyne.Container, over *canvas.Image) {
	off := squareToOffset(m.S1())
	cell := grid.Objects[off].(*fyne.Container)
	img := cell.Objects[2].(*piece)
	pos1 := cell.Position()

	over.Resource = img.Resource
	over.Move(pos1)
	over.Resize(img.Size())

	over.Show()
	img.Resource = nil
	img.Refresh()

	off = squareToOffset(m.S2())
	cell = grid.Objects[off].(*fyne.Container)
	pos2 := cell.Position()

	animation := canvas.NewPositionAnimation(pos1, pos2, time.Millisecond*400, func(p fyne.Position) {
		over.Move(p)
		over.Refresh()
	})
	animation.Start()
	time.Sleep(time.Millisecond * 550)

	game.Move(m)
	refreshGrid(grid, game.Position().Board())
	over.Hide()
	over.Resource = nil
	over.Refresh()
	if game.Outcome() != chess.NoOutcome {
		result := "draw"
		switch game.Outcome().String() {
		case "1-0":
			result = "white Won"
		case "0-1":
			result = "Black Won"
		}
		dialog.ShowInformation("Game Ended", result, win)
	}
}

func squareToOffset(sq chess.Square) int {
	x := sq % 8
	y := 7 - ((sq - x) / 8)
	return int(x + y*8)
}
