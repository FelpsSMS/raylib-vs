package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"reflect"
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var SCREEN_WIDTH = 1024
var SCREEN_HEIGHT = 800
var GAME_TITLE = "RAYLIB VS"
var logger = log.New(os.Stdout, "LOG: ", log.Ldate|log.Ltime|log.Lshortfile)
var camera rl.Camera2D
var playerInstance *Player
var mobs []*Mob
var timer *Timer = &Timer{timerControl: 0, minutes: 0, seconds: 0}

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

func RandomColor() rl.Color {
	r := uint8(rand.Intn(256))
	g := uint8(rand.Intn(256))
	b := uint8(rand.Intn(256))
	return rl.NewColor(r, g, b, 255)
}

func DrawOutlinedText(text string, posX int32, posY int32, fontSize int32, color rl.Color, outlineSize int32, outlineColor rl.Color) {
	rl.DrawText(text, posX+outlineSize, posY+outlineSize, fontSize, outlineColor)
	rl.DrawText(text, posX, posY, fontSize, color)
}

func main() {
	rl.InitWindow(int32(SCREEN_WIDTH), int32(SCREEN_HEIGHT), GAME_TITLE)
	defer rl.CloseWindow()

	//Disable esc key for closing the game
	rl.SetExitKey(0)

	playerInstance = startDebugPlayer()
	//startDebugItemsAndMobs()

	startDebugWeapons()

	camera = rl.NewCamera2D(rl.NewVector2(float32(SCREEN_WIDTH)/2, float32(SCREEN_HEIGHT)/2), rl.NewVector2(playerInstance.Position.X, playerInstance.Position.Y), 0, 1)

	movementSpeed := float32(4)

	rl.SetTargetFPS(60)
	playerInstance.Velocity = rl.NewVector2(0, 0)

	spawnMobsDebug()

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

		playerInstance.Velocity = rl.Vector2Normalize(playerInstance.Velocity)

		playerInstance.Velocity = rl.Vector2Scale(playerInstance.Velocity, movementSpeed)

		playerInstance.Position = rl.Vector2Add(playerInstance.Position, playerInstance.Velocity)

		camera.Target = playerInstance.Position

		rl.BeginDrawing()

		rl.ClearBackground(rl.RayWhite)

		rl.BeginMode2D(camera)

		for _, mob := range mobs {
			mob.Move()
			mob.CheckForCollision()
			mob.Draw()
		}

		for _, weapon := range playerInstance.weapons {
			weapon.Draw()
		}

		playerInstance.Draw()

		rl.EndMode2D()

		updateTimer()

		rl.EndDrawing()
	}
}

func updateTimer() {
	timer.timerControl++

	timer.seconds = timer.timerControl / 60

	if timer.seconds == 60 {
		timer.minutes++
		timer.timerControl = 0
	}

	timerStr := fmt.Sprintf("%02d:%02d", timer.minutes, timer.seconds)

	DrawOutlinedText(timerStr, int32(rl.GetScreenWidth())/2, int32(float32(rl.GetScreenHeight())/20), 40, rl.LightGray, 4, rl.Black)

}

func startDebugPlayer() *Player {
	p := NewPlayer()

	p.HPBar = rl.NewRectangle(p.Position.X, p.Position.Y+p.Height, 20, 4)
	p.originalHPWidth = 100

	return p
}

func startDebugWeapons() {
	basicWeaponPos := playerInstance.Position
	basicWeapon := Weapon{hitbox: rl.NewRectangle(basicWeaponPos.X, basicWeaponPos.Y, 200, 20), id: "basicWeapon", position: basicWeaponPos}

	playerInstance.weapons = append(playerInstance.weapons, &basicWeapon)
}

func spawnMobsDebug() {
	milisecondsToSpawn := time.Duration(50000)
	numberOfMobs := 10

	go func() {
		for range time.Tick(milisecondsToSpawn * time.Millisecond) {

			mobSpawnPointAux1 := rand.Intn(2)
			if mobSpawnPointAux1 == 0 {
				mobSpawnPointAux1 = -1
			}

			mobSpawnPointAux2 := rand.Intn(2)
			if mobSpawnPointAux2 == 0 {
				mobSpawnPointAux2 = -1
			}

			for i := 0; i < numberOfMobs; i++ {
				basicMob := Spawn(Mob{Name: "Test", Position: rl.Vector2{X: 400, Y: 350}, Width: 30, Height: 40, HP: 100, MoveSpeed: 2, MovePattern: FIXED_HORIZONTAL, Damage: 5})

				basicMob.Position.X = camera.Target.X + float32(mobSpawnPointAux1)*float32(rand.Intn(int(SCREEN_WIDTH)))
				basicMob.Position.Y = camera.Target.Y + float32(mobSpawnPointAux2)*float32(rand.Intn(int(SCREEN_HEIGHT)))
				basicMob.MoveSpeed = rand.Float32() * 2

				spawnedMob := Spawn(*basicMob)

				mobs = append(mobs, spawnedMob)
			}
		}
	}()
}
