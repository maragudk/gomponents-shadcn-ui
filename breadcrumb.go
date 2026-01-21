package ui

import (
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/components"
	h "maragu.dev/gomponents/html"
)

// BreadcrumbProps for [Breadcrumb].
type BreadcrumbProps struct{}

// Breadcrumb renders a nav element as a breadcrumb navigation with shadcn/ui styling.
// Pass additional attributes and children as needed.
func Breadcrumb(props BreadcrumbProps, children ...Node) Node {
	return h.Nav(
		h.Aria("label", "breadcrumb"),
		h.Data("slot", "breadcrumb"),
		Group(children),
	)
}

// BreadcrumbListProps for [BreadcrumbList].
type BreadcrumbListProps struct{}

// BreadcrumbList renders an ol element as a breadcrumb list with shadcn/ui styling.
// Pass additional attributes and children as needed.
func BreadcrumbList(props BreadcrumbListProps, children ...Node) Node {
	return h.Ol(
		h.Data("slot", "breadcrumb-list"),
		JoinAttrs("class",
			h.Class(breadcrumbListBaseClass),
			Group(children),
		),
	)
}

const breadcrumbListBaseClass = "text-muted-foreground flex flex-wrap items-center gap-1.5 text-sm break-words sm:gap-2.5"

// BreadcrumbItemProps for [BreadcrumbItem].
type BreadcrumbItemProps struct{}

// BreadcrumbItem renders a li element as a breadcrumb item with shadcn/ui styling.
// Pass additional attributes and children as needed.
func BreadcrumbItem(props BreadcrumbItemProps, children ...Node) Node {
	return h.Li(
		h.Data("slot", "breadcrumb-item"),
		JoinAttrs("class",
			h.Class(breadcrumbItemBaseClass),
			Group(children),
		),
	)
}

const breadcrumbItemBaseClass = "inline-flex items-center gap-1.5"

// BreadcrumbLinkProps for [BreadcrumbLink].
type BreadcrumbLinkProps struct{}

// BreadcrumbLink renders an a element as a breadcrumb link with shadcn/ui styling.
// Pass additional attributes (like [h.Href]) and children as needed.
func BreadcrumbLink(props BreadcrumbLinkProps, children ...Node) Node {
	return h.A(
		h.Data("slot", "breadcrumb-link"),
		JoinAttrs("class",
			h.Class(breadcrumbLinkBaseClass),
			Group(children),
		),
	)
}

const breadcrumbLinkBaseClass = "hover:text-foreground transition-colors"

// BreadcrumbPageProps for [BreadcrumbPage].
type BreadcrumbPageProps struct{}

// BreadcrumbPage renders a span element for the current breadcrumb page with shadcn/ui styling.
// Pass additional attributes and children as needed.
func BreadcrumbPage(props BreadcrumbPageProps, children ...Node) Node {
	return h.Span(
		h.Data("slot", "breadcrumb-page"),
		h.Role("link"),
		h.Aria("disabled", "true"),
		h.Aria("current", "page"),
		JoinAttrs("class",
			h.Class(breadcrumbPageBaseClass),
			Group(children),
		),
	)
}

const breadcrumbPageBaseClass = "text-foreground font-normal"

// BreadcrumbSeparatorProps for [BreadcrumbSeparator].
type BreadcrumbSeparatorProps struct{}

// BreadcrumbSeparator renders a li element as a separator between breadcrumb items.
// By default renders a chevron right icon. Pass custom children to override.
func BreadcrumbSeparator(props BreadcrumbSeparatorProps, children ...Node) Node {
	// Default to chevron right SVG if no children provided
	if len(children) == 0 {
		children = []Node{Raw(`<svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="size-3.5"><path d="m9 18 6-6-6-6"/></svg>`)}
	}
	return h.Li(
		h.Data("slot", "breadcrumb-separator"),
		h.Role("presentation"),
		h.Aria("hidden", "true"),
		JoinAttrs("class",
			h.Class(breadcrumbSeparatorBaseClass),
			Group(children),
		),
	)
}

const breadcrumbSeparatorBaseClass = "[&>svg]:size-3.5"

// BreadcrumbEllipsisProps for [BreadcrumbEllipsis].
type BreadcrumbEllipsisProps struct{}

// BreadcrumbEllipsis renders a span element as an ellipsis indicator with shadcn/ui styling.
// Pass additional attributes and children as needed.
func BreadcrumbEllipsis(props BreadcrumbEllipsisProps, children ...Node) Node {
	return h.Span(
		h.Data("slot", "breadcrumb-ellipsis"),
		h.Role("presentation"),
		h.Aria("hidden", "true"),
		JoinAttrs("class",
			h.Class(breadcrumbEllipsisBaseClass),
			Group(children),
		),
		Raw(`<svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="size-4"><circle cx="12" cy="12" r="1"/><circle cx="19" cy="12" r="1"/><circle cx="5" cy="12" r="1"/></svg>`),
		h.Span(h.Class("sr-only"), Text("More")),
	)
}

const breadcrumbEllipsisBaseClass = "flex size-9 items-center justify-center"
