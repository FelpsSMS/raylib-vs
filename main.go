package main

import (
	"log"
	"os"
	"reflect"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var SCREEN_WIDTH = 1024
var SCREEN_HEIGHT = 800
var GAME_TITLE = "RAYLIB VS"
var logger = log.New(os.Stdout, "LOG: ", log.Ldate|log.Ltime|log.Lshortfile)
var camera rl.Camera2D
var playerInstance *Player
var mobs []*Mob

func FindElementIndex[T any](slice []T, element T) int {
	for index, elementInSlice := range slice {
		if reflect.DeepEqual(elementInSlice, element) {
			return index
		}
	}

	return -1
}

func RemoveFromSlice[T any](slice []T, index int) []T {
	return append(slice[:index], slice[index+1:]...)
}

func main() {
	rl.InitWindow(int32(SCREEN_WIDTH), int32(SCREEN_HEIGHT), GAME_TITLE)
	defer rl.CloseWindow()

	//Disable esc key for closing the game
	rl.SetExitKey(0)

	playerInstance = startDebugPlayer()
	startDebugItemsAndMobs()

	camera = rl.NewCamera2D(rl.NewVector2(float32(SCREEN_WIDTH)/2, float32(SCREEN_HEIGHT)/2), rl.NewVector2(playerInstance.X, playerInstance.Y), 0, 1)

	movementSpeed := float32(4)

	rl.SetTargetFPS(60)
	playerInstance.Velocity = rl.NewVector2(0, 0)

	for !rl.WindowShouldClose() {
		// Reset the velocity
		playerInstance.Velocity = rl.NewVector2(0, 0)

		// Handle input
		if rl.IsKeyDown(rl.KeyW) || rl.IsKeyDown(rl.KeyUp) {
			playerInstance.Velocity.Y -= 1
		}
		if rl.IsKeyDown(rl.KeyS) || rl.IsKeyDown(rl.KeyDown) {
			playerInstance.Velocity.Y += 1
		}
		if rl.IsKeyDown(rl.KeyA) || rl.IsKeyDown(rl.KeyLeft) {
			playerInstance.Velocity.X -= 1
		}
		if rl.IsKeyDown(rl.KeyD) || rl.IsKeyDown(rl.KeyRight) {
			playerInstance.Velocity.X += 1
		}

		// Normalize the velocity to ensure the playerInstance moves at the same speed in all directions
		playerInstance.Velocity = rl.Vector2Normalize(playerInstance.Velocity)

		// Scale the velocity by the speed factor
		playerInstance.Velocity = rl.Vector2Scale(playerInstance.Velocity, movementSpeed)

		// Apply the velocity to the position
		playerInstance.Position = rl.Vector2Add(playerInstance.Position, playerInstance.Velocity)

		camera.Target = playerInstance.Position

		rl.BeginDrawing()

		rl.ClearBackground(rl.RayWhite)

		rl.BeginMode2D(camera)

		playerInstance.Draw()

		for _, mob := range mobs {
			mob.Draw()
		}

		rl.EndMode2D()

		rl.EndDrawing()
	}
}

func startDebugPlayer() *Player {
	p := NewPlayer()

	p.HPBar = rl.NewRectangle(p.X, p.Y+p.Height, 20, 4)
	p.originalHPWidth = 100

	return p
}

func startDebugItemsAndMobs() {
	//Basic mob
	basicMob := Spawn(Mob{Name: "Test", Position: rl.Vector2{X: 400, Y: 350}, Width: 30, Height: 40, HP: 100, MoveSpeed: 2, MovePattern: FIXED_HORIZONTAL, Damage: 5})

	mobs = append(mobs, basicMob)
}
