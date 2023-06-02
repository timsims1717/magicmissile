#version 330 core

in vec2  vTexCoords;
out vec4 fragColor;

uniform vec4 uTexBounds;
uniform sampler2D uTexture;

// custom uniforms
uniform float uRed;
uniform float uGreen;
uniform float uBlue;

// explosion texture
uniform sampler2D uExpTexture;

void main() {
    // Get our current screen coordinate
    vec2 t = (vTexCoords - uTexBounds.xy) / uTexBounds.zw;

    vec4 col = texture(uTexture, t);

    // any color
    bool b = col.a > 0.9;
    // highlight color
    bool h = col.r >= 0.99 && col.g <= 0.01 && col.b >= 174.0/255.0 && col.b <= 176.0/255.0;

    // set it to the layer's color if b && !h
    vec4 color = vec4(0, 0, 0, 0);
    if (b) {
        if (h) {
            color = vec4(0, 0, 0, 1.);
        } else {
            color = vec4(uRed, uGreen, uBlue, 1.);
        }
    }
    fragColor = color;
}
