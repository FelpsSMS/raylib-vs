package main

import (
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
	// dropTable       []ItemDrop
	// projectile      Projectile
	attackPattern AttackPattern
}

func Spawn(mob Mob) *Mob {
	mob.Velocity = rl.NewVector2(0, 0)
	mob.State = MOVING
	mob.MaxHP = mob.HP
	mob.HPBar = rl.NewRectangle(mob.Position.X-mob.Width, mob.Position.Y-mob.Height/2, 20, 4)
	mob.originalHPWidth = 100

	return &mob
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
		mob.HPBar.Width = mob.originalHPWidth * (mob.HP / mob.MaxHP)

		mob.HPBar.X = mob.Position.X - mob.Width
		mob.HPBar.Y = mob.Position.Y - mob.Height/2

		rl.DrawRectangleRec(mob.HPBar, rl.Red)
		rl.DrawRectangleRec(rect, rl.DarkGreen)
	} else {
		// index := FindElementIndex(mobs, mob)

		// if index != -1 {
		// 	mobs = RemoveFromSlice(mobs, index)
		// }

		// mob.dropItems()
	}
}
