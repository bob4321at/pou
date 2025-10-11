package shaders

var Flash_Shader = `//kage:unit pixels
			package main

			var I float

			func Fragment(targetCoords vec4, srcPos vec2, _ vec4) vec4 {
				col := imageSrc0At(srcPos.xy)
				if col.w != 0 {
					return vec4((col.x + I), (col.y + I), (col.z + I), col.w)
				} else {
					return col
				}
			}
`

var Air_Bubble_Shader = `//kage:unit pixels
	package main

	var Percent float

	func Fragment(targetCoords vec4, srcPos vec2, _ vec4) vec4 {
		col := imageSrc0At(srcPos.xy)

		Pos := imageSrc0Origin()
		Size := imageSrc0Size()

		var rot float = atan2((Size.x/2+(float(Pos.x)-float(srcPos.x))),(Size.y/2+(float(Pos.y)-float(srcPos.y))))
		rot_deg := mod(rot*(180.0/3.14159), 360)  

		if abs(rot_deg) > 360.0 * Percent {
			col.a = 0
			return vec4(0)
		}

		return col
	}
`

var Test_Refraction_Shader = `//kage:unit pixels
	package main

	func Fragment(targetCoords vec4, srcPos vec2, _ vec4) vec4 {
		col := imageSrc0At(srcPos.xy)
		check_col_close := imageSrc0At(vec2(srcPos.x-2, srcPos.y-2))


		if col.w == 0 {
			if check_col_close.w != 0 {
				return vec4(0, 0, 0, check_col_close.w/2)
			} else {
				return col
			}
		}

		return col
	}
`

var Fill_Shader = `//kage:unit pixels
	package main

	var Percent float

	func Fragment(targetCoords vec4, srcPos vec2, _ vec4) vec4 {
		// size := imageSrc0Size()
		color := imageSrc0At(srcPos.xy)

		aPercent := Percent

		return vec4(color.xyz-(aPercent/2), color.w)

		// if srcPos.y < size.y-(size.y*aPercent) {
		// 	return vec4(color.x/2, color.y/2, color.z/2, color.w)
		// } else {
		// 	return color
		// }
	}
`

var Chunk_Shader = `//kage:unit pixels
			package main

			var R float 
			var G float
			var B float

			var RR float
			var GG float
			var BB float

			func Fragment(targetCoords vec4, srcPos vec2, _ vec4) vec4 {
				col := imageSrc0At(srcPos.xy)
				if col.x >= 163.0/255 &&col.x <= 165.0/255{
					return vec4(R, G, B, 255)
				}
				if col.x >= 132.0/255 && col.x <= 134.0/255{
					return vec4(RR, GG, BB, 255)
				}

				if col.x >= 126.0/255 && col.x <= 128.0/255 {
					return vec4((R + RR) / 1.5, (G + GG) / 1.5, (B + BB) / 1.5, 1)
				}

				return col
			}
`
