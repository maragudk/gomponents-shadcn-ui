package ui

import (
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/components"
	h "maragu.dev/gomponents/html"
)

// SliderProps for [Slider].
type SliderProps struct{}

// Slider renders an input[type="range"] element with shadcn/ui styling.
// Pass attributes like [h.Min], [h.Max], [h.Value], [h.Step] as needed.
func Slider(props SliderProps, children ...Node) Node {
	return h.Input(
		h.Type("range"),
		h.Data("slot", "slider"),
		JoinAttrs("class",
			h.Class(sliderBaseClass),
			Group(children),
		),
	)
}

const sliderBaseClass = "w-full h-1.5 cursor-pointer appearance-none rounded-full bg-muted outline-none disabled:opacity-50 disabled:cursor-not-allowed " +
	"[&::-webkit-slider-thumb]:appearance-none [&::-webkit-slider-thumb]:size-4 [&::-webkit-slider-thumb]:rounded-full [&::-webkit-slider-thumb]:border [&::-webkit-slider-thumb]:border-primary [&::-webkit-slider-thumb]:bg-white [&::-webkit-slider-thumb]:shadow-sm [&::-webkit-slider-thumb]:transition-shadow [&::-webkit-slider-thumb]:hover:ring-4 [&::-webkit-slider-thumb]:hover:ring-ring/50 [&::-webkit-slider-thumb]:focus:ring-4 [&::-webkit-slider-thumb]:focus:ring-ring/50 " +
	"[&::-moz-range-thumb]:appearance-none [&::-moz-range-thumb]:size-4 [&::-moz-range-thumb]:rounded-full [&::-moz-range-thumb]:border [&::-moz-range-thumb]:border-primary [&::-moz-range-thumb]:bg-white [&::-moz-range-thumb]:shadow-sm [&::-moz-range-thumb]:transition-shadow [&::-moz-range-thumb]:hover:ring-4 [&::-moz-range-thumb]:hover:ring-ring/50 [&::-moz-range-thumb]:focus:ring-4 [&::-moz-range-thumb]:focus:ring-ring/50"
