package shaders

var Enemy_Shader = `//kage:unit pixels
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
