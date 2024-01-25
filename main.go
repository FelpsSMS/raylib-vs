package main

import (
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

func main() {
	rl.InitWindow(int32(SCREEN_WIDTH), int32(SCREEN_HEIGHT), GAME_TITLE)
	defer rl.CloseWindow()

	//Disable esc key for closing the game
	rl.SetExitKey(0)

	playerInstance = startDebugPlayer()
	//startDebugItemsAndMobs()

	camera = rl.NewCamera2D(rl.NewVector2(float32(SCREEN_WIDTH)/2, float32(SCREEN_HEIGHT)/2), rl.NewVector2(playerInstance.X, playerInstance.Y), 0, 1)

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

		playerInstance.Draw()

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

func spawnMobsDebug() {
	secondsToSpawn := time.Duration(5)
	numberOfMobs := 100

	go func() {
		for _ = range time.Tick(secondsToSpawn * time.Second) {
			logger.Print("spawn")

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
