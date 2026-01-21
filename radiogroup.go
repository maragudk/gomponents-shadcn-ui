package ui

import (
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/components"
	h "maragu.dev/gomponents/html"
)

// RadioGroupProps for [RadioGroup].
type RadioGroupProps struct{}

// RadioGroup renders a container for radio items with shadcn/ui styling.
// Pass [RadioGroupItem] children with the same name attribute to group them.
func RadioGroup(props RadioGroupProps, children ...Node) Node {
	return h.Div(
		h.Role("radiogroup"),
		h.Data("slot", "radio-group"),
		JoinAttrs("class",
			h.Class(radioGroupBaseClass),
			Group(children),
		),
	)
}

const radioGroupBaseClass = "grid gap-3"

// RadioGroupItemProps for [RadioGroupItem].
type RadioGroupItemProps struct{}

// RadioGroupItem renders an input[type="radio"] element with shadcn/ui styling.
// Pass attributes like [h.Name], [h.Value], [h.ID], and [h.Checked] as needed.
func RadioGroupItem(props RadioGroupItemProps, children ...Node) Node {
	return h.Input(
		h.Type("radio"),
		h.Data("slot", "radio-group-item"),
		JoinAttrs("class",
			h.Class(radioGroupItemBaseClass),
			Group(children),
		),
	)
}

const radioGroupItemBaseClass = "peer border-input text-primary focus-visible:border-ring focus-visible:ring-ring/50 aria-invalid:ring-destructive/20 dark:aria-invalid:ring-destructive/40 aria-invalid:border-destructive dark:bg-input/30 aspect-square size-4 shrink-0 rounded-full border shadow-xs transition-[color,box-shadow] outline-none focus-visible:ring-[3px] disabled:cursor-not-allowed disabled:opacity-50 appearance-none checked:bg-primary checked:border-primary cursor-pointer"
