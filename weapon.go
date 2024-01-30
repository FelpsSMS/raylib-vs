package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Weapon struct {
	id       string
	hitbox   rl.Rectangle
	position rl.Vector2
}

func (w *Weapon) Draw() {
	// Calculate the angle between the rectangle's origin and the mouse position
	//angle := math.Atan2(float64(float32(rl.GetMouseY())-w.hitbox.Y), float64(float32(rl.GetMouseX())-w.hitbox.X))
	// Update the rectangle's rotation
	w.position = playerInstance.Position

	w.hitbox.X = w.position.X
	w.hitbox.Y = w.position.Y

	rl.DrawRectangleRec(w.hitbox, rl.Gray)
}
