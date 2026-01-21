package ui

import (
	"strconv"

	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/components"
	h "maragu.dev/gomponents/html"
)

// AspectRatioProps for [AspectRatio].
type AspectRatioProps struct {
	// Ratio is the aspect ratio as width/height (e.g., 16/9 = 1.777...).
	// Defaults to 1 (square) if not set.
	Ratio float64
}

// AspectRatio renders a div element that maintains a specific aspect ratio with shadcn/ui styling.
// Pass additional attributes (like [h.Class]) and children as needed.
func AspectRatio(props AspectRatioProps, children ...Node) Node {
	ratio := props.Ratio
	if ratio == 0 {
		ratio = 1
	}
	return h.Div(
		h.Data("slot", "aspect-ratio"),
		h.Style("position: relative; width: 100%; aspect-ratio: "+strconv.FormatFloat(ratio, 'f', -1, 64)+";"),
		JoinAttrs("class",
			h.Class(aspectRatioBaseClass),
			Group(children),
		),
	)
}

const aspectRatioBaseClass = "[&>*]:absolute [&>*]:inset-0 [&>*]:size-full"
