package shader

var Camera_Shader = `//kage:unit pixels
	package main

	func Fragment(targetCoords vec4, srcPos vec2, _ vec4) vec4 {
		col := imageSrc0At(srcPos.xy)
		retun_col := col
		check_col_close := imageSrc0At(vec2(srcPos.x-2, srcPos.y-2))


		if col.w == 0 {
			if check_col_close.w != 0 {
				retun_col = vec4(0, 0, 0, check_col_close.w/2)
			} else {
				retun_col = col
			}
		}

		return retun_col
	}
`

var Water_shader = `//kage:unit pixels
	package main

	var Water_Level float
	var Camera_Y float
	var Camera_X float
	var Game_Time float

	var R float 
	var G float
	var B float

	var RR float
	var GG float
	var BB float

	func Fragment(targetCoords vec4, srcPos vec2, _ vec4) vec4 {
		col := imageSrc0At(srcPos.xy)
		retun_col := vec4(0)

		if (col.w != 0) {
				retun_col.r = ((R+RR)/1.5) 
				retun_col.g = ((G+GG)/1.5) 
				retun_col.b = ((B+BB)/1.5) 

				retun_col.r /= 2
				retun_col.g /= 2
				retun_col.b /= 2
				retun_col.a = 0.5
			return retun_col
		} else {
			if (srcPos.y+sin(Game_Time/20)-Camera_Y+cos(srcPos.x-Camera_X-Game_Time) > Water_Level) {
				retun_col.r = ((R+RR)/1.5) 
				retun_col.g = ((G+GG)/1.5) 
				retun_col.b = ((B+BB)/1.5) 

				retun_col.r /= 2
				retun_col.g /= 2
				retun_col.b /= 2
				retun_col.a = 0.5
			} else {
				retun_col = vec4(0)
			}
		}

		return retun_col
	}

`
