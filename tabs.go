package ui

import (
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/components"
	h "maragu.dev/gomponents/html"
)

// TabsProps for [Tabs].
type TabsProps struct{}

// Tabs renders a container for tab components with shadcn/ui styling.
// Use with [TabsList], [TabsTrigger], and [TabsContent].
func Tabs(props TabsProps, children ...Node) Node {
	return h.Div(
		h.Data("slot", "tabs"),
		JoinAttrs("class",
			h.Class(tabsBaseClass),
			Group(children),
		),
	)
}

const tabsBaseClass = "flex flex-col gap-2"

// TabsListProps for [TabsList].
type TabsListProps struct{}

// TabsList renders a container for tab triggers with shadcn/ui styling.
func TabsList(props TabsListProps, children ...Node) Node {
	return h.Div(
		h.Role("tablist"),
		h.Data("slot", "tabs-list"),
		JoinAttrs("class",
			h.Class(tabsListBaseClass),
			Group(children),
		),
	)
}

const tabsListBaseClass = "bg-muted text-muted-foreground inline-flex h-9 w-fit items-center justify-center rounded-lg p-[3px]"

// TabsTriggerProps for [TabsTrigger].
type TabsTriggerProps struct {
	// Active indicates if this tab is currently selected.
	Active bool
}

// TabsTrigger renders a button that activates a tab panel with shadcn/ui styling.
// Use Datastar to manage active state.
func TabsTrigger(props TabsTriggerProps, children ...Node) Node {
	attrs := []Node{
		h.Type("button"),
		h.Role("tab"),
		h.Data("slot", "tabs-trigger"),
	}

	if props.Active {
		attrs = append(attrs, h.Aria("selected", "true"), h.Data("state", "active"))
	} else {
		attrs = append(attrs, h.Aria("selected", "false"), h.Data("state", "inactive"))
	}

	attrs = append(attrs,
		JoinAttrs("class",
			h.Class(tabsTriggerBaseClass),
			Group(children),
		),
	)

	return h.Button(attrs...)
}

const tabsTriggerBaseClass = "data-[state=active]:bg-background dark:data-[state=active]:text-foreground focus-visible:border-ring focus-visible:ring-ring/50 focus-visible:outline-ring dark:data-[state=active]:border-input dark:data-[state=active]:bg-input/30 text-foreground dark:text-muted-foreground inline-flex h-[calc(100%-1px)] flex-1 items-center justify-center gap-1.5 rounded-md border border-transparent px-2 py-1 text-sm font-medium whitespace-nowrap transition-[color,box-shadow] focus-visible:ring-[3px] focus-visible:outline-1 disabled:pointer-events-none disabled:opacity-50 data-[state=active]:shadow-sm [&_svg]:pointer-events-none [&_svg]:shrink-0 [&_svg:not([class*='size-'])]:size-4 cursor-pointer"

// TabsContentProps for [TabsContent].
type TabsContentProps struct{}

// TabsContent renders a panel for tab content with shadcn/ui styling.
// Use Datastar's data-show directive to toggle visibility.
func TabsContent(props TabsContentProps, children ...Node) Node {
	return h.Div(
		h.Role("tabpanel"),
		h.Data("slot", "tabs-content"),
		JoinAttrs("class",
			h.Class(tabsContentBaseClass),
			Group(children),
		),
	)
}

const tabsContentBaseClass = "flex-1 outline-none"
