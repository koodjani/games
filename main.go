package main

import (
	"image"
	"image/color"
	"log"
	"github.com/koodjani/games/scripts"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Sprite struct{
	Name string
	Img *ebiten.Image
	X, Y float64
	Alive bool
}

type Game struct{
	player *Sprite
	sprites map[string]*Sprite
	objects map[string]*Sprite
}

func (g *Game) Update() error {

	// react to key presses
	if ebiten.IsKeyPressed(ebiten.KeyD) {
		g.player.X += 2
	}
	if ebiten.IsKeyPressed(ebiten.KeyA) {
		g.player.X -= 2
	}
	if ebiten.IsKeyPressed(ebiten.KeyW) {
		g.player.Y -= 2
	}
	if ebiten.IsKeyPressed(ebiten.KeyS) {
		g.player.Y += 2
	}

	if scripts.IsClose(g.player.X, g.player.Y, g.sprites["shaman"].X, g.sprites["shaman"].Y, 16.0) {
		g.objects["dialogbox"].Alive = true
	} else {
		g.objects["dialogbox"].Alive = false
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {

	screen.Fill(color.RGBA{43, 63, 27, 255})
	ebitenutil.DebugPrint(screen, "Hello, World!")
	

	opts := ebiten.DrawImageOptions{}
	opts.GeoM.Translate(g.player.X, g.player.Y)
	// draw our player
	screen.DrawImage(
		g.player.Img.SubImage(
			image.Rect(0, 0, 16, 16),
		).(*ebiten.Image),
		&opts,
	)
	opts.GeoM.Reset()

	opts.GeoM.Translate(g.sprites["shaman"].X, g.sprites["shaman"].Y)

		screen.DrawImage(
			g.sprites["shaman"].Img.SubImage(
				image.Rect(0,0,16,16),
			).(*ebiten.Image),
			&opts,
		)
		
		opts.GeoM.Reset()
	
	if g.objects["dialogbox"].Alive {
		opts.GeoM.Scale(0.3,0.3)
		opts.GeoM.Translate(g.sprites["shaman"].X-30, g.sprites["shaman"].Y -20)
		screen.DrawImage(
			g.objects["dialogbox"].Img,
			&opts,)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Hello, World!")
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)

	playerImg, _, err := ebitenutil.NewImageFromFile("assets/images/knight.png")
	if err != nil {
		log.Fatal(err)
	}
	shamanImg, _, err := ebitenutil.NewImageFromFile("assets/images/shaman.png")
	if err != nil {
		log.Fatal(err)
	}
	dialogboxImg, _, err := ebitenutil.NewImageFromFile("assets/images/dialogbox.png")
	if err != nil {
		log.Fatal(err)
	}

	game := Game{
		player: &Sprite{
			Img: playerImg,
			X: 100,
			Y: 100,
		},
		sprites: map[string]*Sprite {
			"shaman": {
				Img: shamanImg,
				X: 50.0,
				Y: 50.0,		
			},
		},
		objects: map[string]*Sprite {
			"dialogbox": {
				Img: dialogboxImg,
				X: 55,
				Y: 66,
				Alive: false,
			},
		},
	}

	if err := ebiten.RunGame(&game); err != nil {
		log.Fatal(err)
	}
}