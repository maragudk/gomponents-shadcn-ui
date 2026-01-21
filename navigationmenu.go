package ui

import (
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/components"
	h "maragu.dev/gomponents/html"
)

// NavigationMenuProps for [NavigationMenu].
type NavigationMenuProps struct{}

// NavigationMenu renders a navigation menu container.
// Use with Datastar for open/close state management of menu items.
func NavigationMenu(props NavigationMenuProps, children ...Node) Node {
	return h.Nav(
		h.Data("slot", "navigation-menu"),
		JoinAttrs("class",
			h.Class(navigationMenuBaseClass),
			Group(children),
		),
	)
}

const navigationMenuBaseClass = "relative flex max-w-max flex-1 items-center justify-center"

// NavigationMenuListProps for [NavigationMenuList].
type NavigationMenuListProps struct{}

// NavigationMenuList renders a list container for navigation items.
func NavigationMenuList(props NavigationMenuListProps, children ...Node) Node {
	return h.Ul(
		h.Data("slot", "navigation-menu-list"),
		JoinAttrs("class",
			h.Class(navigationMenuListBaseClass),
			Group(children),
		),
	)
}

const navigationMenuListBaseClass = "group flex flex-1 list-none items-center justify-center gap-1"

// NavigationMenuItemProps for [NavigationMenuItem].
type NavigationMenuItemProps struct{}

// NavigationMenuItem renders a single navigation menu item container.
func NavigationMenuItem(props NavigationMenuItemProps, children ...Node) Node {
	return h.Li(
		h.Data("slot", "navigation-menu-item"),
		JoinAttrs("class",
			h.Class(navigationMenuItemBaseClass),
			Group(children),
		),
	)
}

const navigationMenuItemBaseClass = "relative"

// NavigationMenuTriggerProps for [NavigationMenuTrigger].
type NavigationMenuTriggerProps struct{}

// NavigationMenuTrigger renders a button that triggers the display of content.
func NavigationMenuTrigger(props NavigationMenuTriggerProps, children ...Node) Node {
	return h.Button(
		h.Type("button"),
		h.Data("slot", "navigation-menu-trigger"),
		JoinAttrs("class",
			h.Class(navigationMenuTriggerBaseClass),
			Group(children),
		),
	)
}

const navigationMenuTriggerBaseClass = "group inline-flex h-9 w-max items-center justify-center rounded-md bg-background px-4 py-2 text-sm font-medium hover:bg-accent hover:text-accent-foreground focus:bg-accent focus:text-accent-foreground disabled:pointer-events-none disabled:opacity-50 data-[state=open]:bg-accent/50 data-[state=open]:text-accent-foreground outline-none transition-colors cursor-pointer"

// NavigationMenuContentProps for [NavigationMenuContent].
type NavigationMenuContentProps struct{}

// NavigationMenuContent renders the dropdown content panel.
// Use with Datastar's data-show directive to control visibility.
func NavigationMenuContent(props NavigationMenuContentProps, children ...Node) Node {
	return h.Div(
		h.Data("slot", "navigation-menu-content"),
		JoinAttrs("class",
			h.Class(navigationMenuContentBaseClass),
			Group(children),
		),
	)
}

const navigationMenuContentBaseClass = "bg-popover text-popover-foreground absolute top-full left-0 z-50 mt-1.5 w-auto overflow-hidden rounded-md border p-2 shadow"

// NavigationMenuLinkProps for [NavigationMenuLink].
type NavigationMenuLinkProps struct {
	// Active indicates if this link is currently active.
	Active bool
}

// NavigationMenuLink renders a link inside the navigation menu.
func NavigationMenuLink(props NavigationMenuLinkProps, children ...Node) Node {
	attrs := []Node{
		h.Data("slot", "navigation-menu-link"),
	}

	if props.Active {
		attrs = append(attrs, h.Data("active", "true"))
	}

	attrs = append(attrs,
		JoinAttrs("class",
			h.Class(navigationMenuLinkBaseClass),
			Group(children),
		),
	)

	return h.A(attrs...)
}

const navigationMenuLinkBaseClass = "data-[active=true]:bg-accent/50 data-[active=true]:text-accent-foreground hover:bg-accent hover:text-accent-foreground focus:bg-accent focus:text-accent-foreground flex flex-col gap-1 rounded-sm p-2 text-sm transition-colors outline-none"

// NavigationMenuIndicatorProps for [NavigationMenuIndicator].
type NavigationMenuIndicatorProps struct{}

// NavigationMenuIndicator renders a visual indicator for the active menu item.
func NavigationMenuIndicator(props NavigationMenuIndicatorProps, children ...Node) Node {
	return h.Div(
		h.Data("slot", "navigation-menu-indicator"),
		JoinAttrs("class",
			h.Class(navigationMenuIndicatorBaseClass),
			Group(children),
		),
		h.Div(h.Class("bg-border relative top-[60%] h-2 w-2 rotate-45 rounded-tl-sm shadow-md")),
	)
}

const navigationMenuIndicatorBaseClass = "top-full z-[1] flex h-1.5 items-end justify-center overflow-hidden"

// NavigationMenuViewportProps for [NavigationMenuViewport].
type NavigationMenuViewportProps struct{}

// NavigationMenuViewport renders a viewport container for content.
// This is optional and used for more complex navigation menu layouts.
func NavigationMenuViewport(props NavigationMenuViewportProps, children ...Node) Node {
	return h.Div(
		h.Data("slot", "navigation-menu-viewport"),
		JoinAttrs("class",
			h.Class(navigationMenuViewportBaseClass),
			Group(children),
		),
	)
}

const navigationMenuViewportBaseClass = "bg-popover text-popover-foreground absolute top-full left-0 z-50 mt-1.5 overflow-hidden rounded-md border shadow"
