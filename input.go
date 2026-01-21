package ui

import (
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/components"
	h "maragu.dev/gomponents/html"
)

// InputProps for [Input].
type InputProps struct{}

// Input renders an input element with shadcn/ui styling.
// Pass additional attributes (like [h.Type], [h.Placeholder]) and children as needed.
func Input(props InputProps, children ...Node) Node {
	return h.Input(
		h.Data("slot", "input"),
		JoinAttrs("class",
			h.Class(inputBaseClass),
			Group(children),
		),
	)
}

const inputBaseClass = "file:text-foreground placeholder:text-muted-foreground selection:bg-primary selection:text-primary-foreground dark:bg-input/30 border-input h-9 w-full min-w-0 rounded-md border bg-transparent px-3 py-1 text-base shadow-xs transition-[color,box-shadow] outline-none file:inline-flex file:h-7 file:border-0 file:bg-transparent file:text-sm file:font-medium disabled:pointer-events-none disabled:cursor-not-allowed disabled:opacity-50 md:text-sm focus-visible:border-ring focus-visible:ring-ring/50 focus-visible:ring-[3px] aria-invalid:ring-destructive/20 dark:aria-invalid:ring-destructive/40 aria-invalid:border-destructive"
