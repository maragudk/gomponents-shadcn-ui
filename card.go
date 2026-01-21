package ui

import (
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/components"
	h "maragu.dev/gomponents/html"
)

// CardProps for [Card].
type CardProps struct{}

// Card renders a div element as a card container with shadcn/ui styling.
// Pass additional attributes (like [h.Class]) and children as needed.
func Card(props CardProps, children ...Node) Node {
	return h.Div(
		JoinAttrs("class",
			h.Class(cardBaseClass),
			Group(children),
		),
	)
}

const cardBaseClass = "bg-card text-card-foreground flex flex-col gap-6 rounded-xl border py-6 shadow-sm"

// CardHeaderProps for [CardHeader].
type CardHeaderProps struct{}

// CardHeader renders a div element as the card header with shadcn/ui styling.
// Pass additional attributes (like [h.Class]) and children as needed.
func CardHeader(props CardHeaderProps, children ...Node) Node {
	return h.Div(
		JoinAttrs("class",
			h.Class(cardHeaderBaseClass),
			Group(children),
		),
	)
}

const cardHeaderBaseClass = "@container/card-header grid auto-rows-min grid-rows-[auto_auto] items-start gap-2 px-6 has-data-[slot=card-action]:grid-cols-[1fr_auto] [.border-b]:pb-6"

// CardTitleProps for [CardTitle].
type CardTitleProps struct{}

// CardTitle renders a div element as the card title with shadcn/ui styling.
// Pass additional attributes (like [h.Class]) and children as needed.
func CardTitle(props CardTitleProps, children ...Node) Node {
	return h.Div(
		JoinAttrs("class",
			h.Class(cardTitleBaseClass),
			Group(children),
		),
	)
}

const cardTitleBaseClass = "leading-none font-semibold"

// CardDescriptionProps for [CardDescription].
type CardDescriptionProps struct{}

// CardDescription renders a div element as the card description with shadcn/ui styling.
// Pass additional attributes (like [h.Class]) and children as needed.
func CardDescription(props CardDescriptionProps, children ...Node) Node {
	return h.Div(
		JoinAttrs("class",
			h.Class(cardDescriptionBaseClass),
			Group(children),
		),
	)
}

const cardDescriptionBaseClass = "text-muted-foreground text-sm"

// CardActionProps for [CardAction].
type CardActionProps struct{}

// CardAction renders a div element for card actions with shadcn/ui styling.
// Pass additional attributes (like [h.Class]) and children as needed.
func CardAction(props CardActionProps, children ...Node) Node {
	return h.Div(
		JoinAttrs("class",
			h.Class(cardActionBaseClass),
			Group(children),
		),
	)
}

const cardActionBaseClass = "col-start-2 row-span-2 row-start-1 self-start justify-self-end"

// CardContentProps for [CardContent].
type CardContentProps struct{}

// CardContent renders a div element as the card content with shadcn/ui styling.
// Pass additional attributes (like [h.Class]) and children as needed.
func CardContent(props CardContentProps, children ...Node) Node {
	return h.Div(
		JoinAttrs("class",
			h.Class(cardContentBaseClass),
			Group(children),
		),
	)
}

const cardContentBaseClass = "px-6"

// CardFooterProps for [CardFooter].
type CardFooterProps struct{}

// CardFooter renders a div element as the card footer with shadcn/ui styling.
// Pass additional attributes (like [h.Class]) and children as needed.
func CardFooter(props CardFooterProps, children ...Node) Node {
	return h.Div(
		JoinAttrs("class",
			h.Class(cardFooterBaseClass),
			Group(children),
		),
	)
}

const cardFooterBaseClass = "flex items-center px-6 [.border-t]:pt-6"
