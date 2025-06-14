package utils

import "math"

var Mouse_X float64
var Mouse_Y float64

type Vec2 struct {
	X, Y float64
}

func Collide(pos1, size1, pos2, size2 Vec2) bool {
	if pos1.X < pos2.X+size2.X && pos1.X+size1.X > pos2.X {
		if pos1.Y < pos2.Y+size2.Y && pos1.Y+size1.Y > pos2.Y {
			return true
		}
	}
	return false
}

func Deg2Rad(num float64) float64 {
	return num * (3.14159 / 180)
}

func Rad2Deg(num float64) float64 {
	return num * (180 / 3.14159)
}

func GetAngle(point_1, point_2 Vec2) float64 {
	offset_x := point_1.X - point_2.X
	offset_y := point_1.Y - point_2.Y

	return math.Atan2(offset_x, offset_y)
}

func RemoveArrayElement[T any](index_to_remove int, slice *[]T) {
	*slice = append((*slice)[:index_to_remove], (*slice)[index_to_remove+1:]...)
}
