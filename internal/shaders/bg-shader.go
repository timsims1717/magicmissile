package shaders

var BGShader = `
#version 330 core

in vec2  vTexCoords;
out vec4 fragColor;

uniform vec4 uTexBounds;
uniform sampler2D uTexture;

// custom uniforms
uniform float uRed;
uniform float uGreen;
uniform float uBlue;

void main() {
	// Get our current screen coordinate
	vec2 t = (vTexCoords - uTexBounds.xy) / uTexBounds.zw;

	vec4 col = texture(uTexture, t);

	bool b = col.a > 0.;

	// 
	vec4 color = vec4(0, 0, 0, 0);
	if (b) {
		color = vec4(uRed, uGreen, uBlue, 1.);
	}
	fragColor = color;
}
`
