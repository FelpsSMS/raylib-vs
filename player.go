package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Sprite struct {
	FrameCounter int
	CurrentFrame int
	FrameSpeed   int
	FrameRec     rl.Rectangle
	Position     rl.Vector2
	Texture      rl.Texture2D
}

type State int

const (
	IDLE State = iota
	MOVING
	PAUSED
	DEAD
)

type Player struct {
	X             float32
	Y             float32
	Width         float32
	Height        float32
	Hitbox        rl.Rectangle
	State         State
	Sprite        Sprite
	Position      rl.Vector2
	Velocity      rl.Vector2
	isInvunerable bool
	originalY     float32
	//Weapon             Weapon
	HPBar           rl.Rectangle
	originalHPWidth float32
	HP              float32
	MaxHP           float32
	//Inventory       []*Item
	//projectileSlot     Projectile
	projectileQuantity int32
}

type OffsetParams struct {
	X      float32
	Y      float32
	Width  float32
	Height float32
}

func NewPlayer() *Player {
	//playerTexture := rl.LoadTexture("assets/player/player-spritemap-v9.png")

	return &Player{
		State:         IDLE,
		Position:      rl.NewVector2(float32(SCREEN_WIDTH)/2, float32(SCREEN_HEIGHT)/2),
		Velocity:      rl.NewVector2(0, 0),
		Width:         16,
		Height:        38,
		isInvunerable: false,
		HP:            100,
		MaxHP:         200,
		// Sprite: Sprite{
		// 	Texture:  playerTexture,
		// 	FrameRec: rl.NewRectangle(0, 0, float32(playerTexture.Width/8), float32(playerTexture.Height/4)),
		// 	Position: rl.Vector2{X: 100, Y: 100},
		// },
	}
}

func (p *Player) Draw() {
	rect := rl.Rectangle{X: p.Position.X, Y: p.Position.Y, Width: p.Width, Height: p.Height}
	p.Hitbox = rect

	if p.HP > 0 {
		p.HPBar.Width = p.originalHPWidth * (p.HP / p.MaxHP)

		p.HPBar.X = p.Position.X - p.Width
		p.HPBar.Y = p.Position.Y - p.Height/2

		rl.DrawRectangleRec(p.HPBar, rl.Red)
		rl.DrawRectangleRec(rect, rl.DarkBlue)
		//rl.DrawTextureRec(p.Sprite.Texture, p.Sprite.FrameRec, rl.Vector2{X: p.X - p.Hitbox.Width, Y: p.Y - p.Hitbox.Height/4}, rl.White)
	} else {
		p.State = DEAD
	}

}
