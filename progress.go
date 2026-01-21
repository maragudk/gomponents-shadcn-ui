package ui

import (
	"strconv"

	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/components"
	h "maragu.dev/gomponents/html"
)

// ProgressProps for [Progress].
type ProgressProps struct {
	// Value is the progress value from 0 to 100.
	Value int
}

// Progress renders a progress bar with shadcn/ui styling.
// Pass additional attributes (like [h.Class]) and children as needed.
func Progress(props ProgressProps, children ...Node) Node {
	return h.Div(
		h.Role("progressbar"),
		h.Data("slot", "progress"),
		h.Aria("valuenow", strconv.Itoa(props.Value)),
		h.Aria("valuemin", "0"),
		h.Aria("valuemax", "100"),
		JoinAttrs("class",
			h.Class(progressBaseClass),
			Group(children),
		),
		h.Div(
			h.Data("slot", "progress-indicator"),
			h.Class(progressIndicatorClass),
			h.Style("transform: translateX(-"+strconv.Itoa(100-props.Value)+"%);"),
		),
	)
}

const progressBaseClass = "bg-primary/20 relative h-2 w-full overflow-hidden rounded-full"
const progressIndicatorClass = "bg-primary h-full w-full flex-1 transition-all"
