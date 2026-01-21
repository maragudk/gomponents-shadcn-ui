package ui

import (
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/components"
	h "maragu.dev/gomponents/html"
)

// AccordionProps for [Accordion].
type AccordionProps struct{}

// Accordion renders a container for accordion items.
// Use with Datastar for open/close state management.
func Accordion(props AccordionProps, children ...Node) Node {
	return h.Div(
		h.Data("slot", "accordion"),
		JoinAttrs("class",
			h.Class(accordionBaseClass),
			Group(children),
		),
	)
}

const accordionBaseClass = ""

// AccordionItemProps for [AccordionItem].
type AccordionItemProps struct{}

// AccordionItem renders a single accordion item.
func AccordionItem(props AccordionItemProps, children ...Node) Node {
	return h.Div(
		h.Data("slot", "accordion-item"),
		JoinAttrs("class",
			h.Class(accordionItemBaseClass),
			Group(children),
		),
	)
}

const accordionItemBaseClass = "border-b last:border-b-0"

// AccordionTriggerProps for [AccordionTrigger].
type AccordionTriggerProps struct{}

// AccordionTrigger renders the button that toggles an accordion item.
// Include a chevron icon as a child for the rotation animation.
func AccordionTrigger(props AccordionTriggerProps, children ...Node) Node {
	return h.Div(
		h.Class("flex"),
		h.Button(
			h.Type("button"),
			h.Data("slot", "accordion-trigger"),
			JoinAttrs("class",
				h.Class(accordionTriggerBaseClass),
				Group(children),
			),
		),
	)
}

const accordionTriggerBaseClass = "focus-visible:border-ring focus-visible:ring-ring/50 flex flex-1 items-start justify-between gap-4 rounded-md py-4 text-left text-sm font-medium transition-all outline-none hover:underline focus-visible:ring-[3px] disabled:pointer-events-none disabled:opacity-50 [&[data-state=open]>svg]:rotate-180 cursor-pointer"

// AccordionContentProps for [AccordionContent].
type AccordionContentProps struct{}

// AccordionContent renders the expandable content of an accordion item.
// Use with Datastar's data-show directive to control visibility.
func AccordionContent(props AccordionContentProps, children ...Node) Node {
	return h.Div(
		h.Data("slot", "accordion-content"),
		JoinAttrs("class",
			h.Class(accordionContentBaseClass),
			Group(children),
		),
	)
}

const accordionContentBaseClass = "overflow-hidden text-sm pb-4"
