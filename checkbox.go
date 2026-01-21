package ui

import (
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/components"
	h "maragu.dev/gomponents/html"
)

// CheckboxProps for [Checkbox].
type CheckboxProps struct{}

// Checkbox renders an input[type="checkbox"] element with shadcn/ui styling.
// Pass additional attributes (like [h.ID], [h.Name], [h.Checked]) and children as needed.
func Checkbox(props CheckboxProps, children ...Node) Node {
	return h.Input(
		h.Type("checkbox"),
		h.Data("slot", "checkbox"),
		JoinAttrs("class",
			h.Class(checkboxBaseClass),
			Group(children),
		),
	)
}

const checkboxBaseClass = "peer border-input dark:bg-input/30 checked:bg-primary checked:text-primary-foreground dark:checked:bg-primary checked:border-primary focus-visible:border-ring focus-visible:ring-ring/50 aria-invalid:ring-destructive/20 dark:aria-invalid:ring-destructive/40 aria-invalid:border-destructive size-4 shrink-0 rounded-[4px] border shadow-xs transition-shadow outline-none focus-visible:ring-[3px] disabled:cursor-not-allowed disabled:opacity-50 accent-primary"
