package ui

import (
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/components"
	h "maragu.dev/gomponents/html"
)

// SwitchProps for [Switch].
type SwitchProps struct{}

// Switch renders an input[type="checkbox"] element styled as a toggle switch with shadcn/ui styling.
// Pass additional attributes (like [h.ID], [h.Name], [h.Checked]) and children as needed.
func Switch(props SwitchProps, children ...Node) Node {
	return h.Input(
		h.Type("checkbox"),
		h.Role("switch"),
		h.Data("slot", "switch"),
		JoinAttrs("class",
			h.Class(switchBaseClass),
			Group(children),
		),
	)
}

const switchBaseClass = "peer appearance-none inline-flex h-5 w-9 shrink-0 items-center rounded-full border border-transparent shadow-xs transition-all outline-none cursor-pointer bg-input dark:bg-input/80 checked:bg-primary focus-visible:border-ring focus-visible:ring-ring/50 focus-visible:ring-[3px] disabled:cursor-not-allowed disabled:opacity-50 before:content-[''] before:block before:size-4 before:rounded-full before:bg-background before:shadow-xs before:ring-0 before:transition-transform before:translate-x-0 checked:before:translate-x-4 checked:before:bg-primary-foreground dark:before:bg-foreground dark:checked:before:bg-primary-foreground"
