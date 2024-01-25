package main

import (
	"math"
	"math/rand"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type MovePattern int

const (
	STILL MovePattern = iota
	FIXED_HORIZONTAL
)

type AttackPattern int

const (
	PACIFIST AttackPattern = iota
	RANGED_BOTH_SIDES
	RANGED_BOTH_SIDES_RANDOM
	RANGED_FRONT
	MELEE_FRONT
	MELEE_BOTH_SIDES
)

type Mob struct {
	Position        rl.Vector2
	Velocity        rl.Vector2
	Width           float32
	Height          float32
	Hitbox          rl.Rectangle
	State           State
	Sprite          Sprite
	RightSide       bool
	HP              float32
	MaxHP           float32
	MovePattern     MovePattern
	SpawnX          float32
	SpawnY          float32
	Name            string
	MoveSpeed       float32
	HPBar           rl.Rectangle
	originalHPWidth float32
	Damage          float32
	Color           rl.Color
	// dropTable       []ItemDrop
	// projectile      Projectile
	//attackPattern AttackPattern
}

func Spawn(mob Mob) *Mob {
	mob.Velocity = rl.NewVector2(0, 0)
	mob.State = MOVING
	mob.MaxHP = mob.HP
	mob.HPBar = rl.NewRectangle(mob.Position.X-mob.Width, mob.Position.Y-mob.Height/2, 20, 4)
	mob.originalHPWidth = 100
	mob.Color = RandomColor()

	return &mob
}

func (mob *Mob) Move() {
	const speed = 2
	stuck := true
	leader := &Mob{}                            // The leader of the mob
	closestDistance := float32(math.MaxFloat32) // The shortest distance to a leader

	for _, spawnedMob := range mobs {
		mobGroupingHitbox := rl.NewRectangle(mob.Hitbox.X, mob.Hitbox.Y, mob.Hitbox.Width, mob.Hitbox.Height)
		spawnedMobGroupingHitbox := rl.NewRectangle(spawnedMob.Hitbox.X, spawnedMob.Hitbox.Y, spawnedMob.Hitbox.Width, spawnedMob.Hitbox.Height)

		if !rl.CheckCollisionRecs(spawnedMobGroupingHitbox, mobGroupingHitbox) {
			stuck = false
			direction := rl.Vector2Subtract(playerInstance.Position, mob.Position)

			direction.X += float32(rand.Intn(5)) / 10
			direction.Y += float32(rand.Intn(5)) / 10

			direction = rl.Vector2Normalize(direction)

			mob.Position = rl.Vector2Add(mob.Position, rl.Vector2Scale(direction, speed))
			break

		} else {
			distance := rl.Vector2Length(rl.Vector2Subtract(mob.Position, spawnedMob.Position))
			if distance < closestDistance {
				closestDistance = distance
				leader = spawnedMob
			}
		}
	}

	if stuck {
		direction := rl.Vector2Subtract(leader.Position, mob.Position)

		direction.X += float32(rand.Intn(5)) / 10
		direction.Y += float32(rand.Intn(5)) / 10

		direction = rl.Vector2Normalize(direction)

		mob.Position = rl.Vector2Add(mob.Position, rl.Vector2Scale(direction, speed))
	}
}

func (mob *Mob) CheckForCollision() {
	if rl.CheckCollisionRecs(playerInstance.Hitbox, mob.Hitbox) {
		playerInstance.HP -= 0.1
	}
}

// func (mob *Mob) Move() {
// 	maxDistanceX := float32(50)

// 	switch mob.MovePattern {
// 	// case STILL:

// 	case FIXED_HORIZONTAL:

// 		if mob.X >= mob.SpawnX+maxDistanceX {
// 			mob.RightSide = true
// 		}

// 		if mob.X <= mob.SpawnX-maxDistanceX {
// 			mob.RightSide = false
// 		}

// 		if mob.RightSide {
// 			mob.X -= mob.MoveSpeed

// 		} else {
// 			mob.X += mob.MoveSpeed
// 		}
// 	}
// }

func (mob *Mob) OffsetHitbox(offset OffsetParams) rl.Rectangle {
	return rl.Rectangle{X: mob.Hitbox.X + offset.X, Y: mob.Hitbox.Y + offset.Y, Width: mob.Hitbox.Width + offset.Width, Height: mob.Hitbox.Height + offset.Height}
}

// func (mob *Mob) dropItems() {
// 	roll := rand.Intn(100)

// 	for _, itemDrop := range mob.dropTable {

// 		if roll <= itemDrop.chance {
// 			rect := rl.Rectangle{X: mob.X + mob.Width/2, Y: mob.Y + mob.Height/2, Width: itemDrop.item.hitbox.Width, Height: itemDrop.item.hitbox.Height}

// 			itemDrop.item.hitbox = rect
// 			itemsInMap = append(itemsInMap, &itemDrop.item)
// 		}
// 	}
// }

func (mob *Mob) Draw() {
	rect := rl.Rectangle{X: mob.Position.X, Y: mob.Position.Y, Width: mob.Width, Height: mob.Height}

	mob.Hitbox = rect

	if mob.HP > 0 {
		// mob.HPBar.Width = mob.originalHPWidth * (mob.HP / mob.MaxHP)

		// mob.HPBar.X = mob.Position.X - mob.Width
		// mob.HPBar.Y = mob.Position.Y - mob.Height/2

		// rl.DrawRectangleRec(mob.HPBar, rl.Red)
		rl.DrawRectangleRec(rect, mob.Color)
	} else {
		// index := FindElementIndex(mobs, mob)

		// if index != -1 {
		// 	mobs = RemoveFromSlice(mobs, index)
		// }

		// mob.dropItems()
	}
}
