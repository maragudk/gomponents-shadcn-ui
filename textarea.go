package ui

import (
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/components"
	h "maragu.dev/gomponents/html"
)

// TextareaProps for [Textarea].
type TextareaProps struct{}

// Textarea renders a textarea element with shadcn/ui styling.
// Pass additional attributes (like [h.Placeholder], [h.Rows]) and children as needed.
func Textarea(props TextareaProps, children ...Node) Node {
	return h.Textarea(
		h.Data("slot", "textarea"),
		JoinAttrs("class",
			h.Class(textareaBaseClass),
			Group(children),
		),
	)
}

const textareaBaseClass = "border-input placeholder:text-muted-foreground focus-visible:border-ring focus-visible:ring-ring/50 aria-invalid:ring-destructive/20 dark:aria-invalid:ring-destructive/40 aria-invalid:border-destructive dark:bg-input/30 flex field-sizing-content min-h-16 w-full rounded-md border bg-transparent px-3 py-2 text-base shadow-xs transition-[color,box-shadow] outline-none focus-visible:ring-[3px] disabled:cursor-not-allowed disabled:opacity-50 md:text-sm"
