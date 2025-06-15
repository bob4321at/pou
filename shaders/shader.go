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
