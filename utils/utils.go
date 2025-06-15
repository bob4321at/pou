package utils

import "math"

var Mouse_X float64
var Mouse_Y float64
var GameTime float64

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

func CalculateVolume(buf []byte) float64 {
	var sum int64
	for i := 0; i < len(buf); i += 2 {
		if i+1 >= len(buf) {
			break
		}
		// Convert 2 bytes (little-endian) into a signed 16-bit value
		sample := int16(buf[i]) | int16(buf[i+1])<<8
		sum += int64(sample) * int64(sample)
	}

	count := len(buf) / 2
	if count == 0 {
		return 0
	}
	mean := float64(sum) / float64(count)
	return math.Sqrt(mean) / 32768.0 // Normalize to range [0.0, 1.0]
}
