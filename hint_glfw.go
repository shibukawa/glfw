// +build !js

package glfw

import "github.com/shibukawa/glfw-2/v3.2/glfw"

type Hint int

const (
	AlphaBits   = Hint(glfw.AlphaBits)
	DepthBits   = Hint(glfw.DepthBits)
	StencilBits = Hint(glfw.StencilBits)
	Samples     = Hint(glfw.Samples)

	// These hints used for WebGL contexts, ignored on desktop.
	PremultipliedAlpha = noopHint
	PreserveDrawingBuffer
	PreferLowPowerToHighPerformance
	FailIfMajorPerformanceCaveat
)

// noopHint is ignored.
const noopHint Hint = -1

func WindowHint(target Hint, hint int) {
	if target == noopHint {
		return
	}

	glfw.WindowHint(glfw.Hint(target), hint)
}
