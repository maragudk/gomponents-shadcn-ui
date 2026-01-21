package ui

import (
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/components"
	h "maragu.dev/gomponents/html"
)

// PaginationProps for [Pagination].
type PaginationProps struct{}

// Pagination renders a nav element as a pagination container with shadcn/ui styling.
// Pass additional attributes and children as needed.
func Pagination(props PaginationProps, children ...Node) Node {
	return h.Nav(
		h.Role("navigation"),
		h.Aria("label", "pagination"),
		h.Data("slot", "pagination"),
		JoinAttrs("class",
			h.Class(paginationBaseClass),
			Group(children),
		),
	)
}

const paginationBaseClass = "mx-auto flex w-full justify-center"

// PaginationContentProps for [PaginationContent].
type PaginationContentProps struct{}

// PaginationContent renders a ul element for pagination items with shadcn/ui styling.
// Pass additional attributes and children as needed.
func PaginationContent(props PaginationContentProps, children ...Node) Node {
	return h.Ul(
		h.Data("slot", "pagination-content"),
		JoinAttrs("class",
			h.Class(paginationContentBaseClass),
			Group(children),
		),
	)
}

const paginationContentBaseClass = "flex flex-row items-center gap-1"

// PaginationItemProps for [PaginationItem].
type PaginationItemProps struct{}

// PaginationItem renders a li element for a pagination item.
// Pass additional attributes and children as needed.
func PaginationItem(props PaginationItemProps, children ...Node) Node {
	return h.Li(
		h.Data("slot", "pagination-item"),
		Group(children),
	)
}

// PaginationLinkProps for [PaginationLink].
type PaginationLinkProps struct {
	// IsActive marks this link as the current page.
	IsActive bool
}

// PaginationLink renders an anchor element as a pagination link with shadcn/ui styling.
// Pass additional attributes (like [h.Href]) and children as needed.
func PaginationLink(props PaginationLinkProps, children ...Node) Node {
	attrs := []Node{
		h.Data("slot", "pagination-link"),
	}
	if props.IsActive {
		attrs = append(attrs, h.Aria("current", "page"), h.Data("active", "true"))
	}
	variant := "ghost"
	if props.IsActive {
		variant = "outline"
	}
	attrs = append(attrs,
		JoinAttrs("class",
			h.Class(paginationLinkBaseClass+" "+paginationLinkVariantClasses[variant]),
			Group(children),
		),
	)
	return h.A(attrs...)
}

const paginationLinkBaseClass = "inline-flex items-center justify-center gap-2 whitespace-nowrap rounded-md text-sm font-medium transition-all disabled:pointer-events-none disabled:opacity-50 [&_svg]:pointer-events-none [&_svg:not([class*='size-'])]:size-4 shrink-0 [&_svg]:shrink-0 outline-none focus-visible:border-ring focus-visible:ring-ring/50 focus-visible:ring-[3px] size-9"

var paginationLinkVariantClasses = map[string]string{
	"ghost":   "hover:bg-accent hover:text-accent-foreground",
	"outline": "border bg-background shadow-xs hover:bg-accent hover:text-accent-foreground",
}

// PaginationPreviousProps for [PaginationPrevious].
type PaginationPreviousProps struct{}

// PaginationPrevious renders a pagination link for going to the previous page.
// Pass additional attributes (like [h.Href]) and children as needed.
func PaginationPrevious(props PaginationPreviousProps, children ...Node) Node {
	return h.A(
		h.Aria("label", "Go to previous page"),
		h.Data("slot", "pagination-link"),
		JoinAttrs("class",
			h.Class(paginationPreviousBaseClass),
			Group(children),
		),
		Raw(`<svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="size-4"><path d="m15 18-6-6 6-6"/></svg>`),
		h.Span(h.Class("hidden sm:block"), Text("Previous")),
	)
}

const paginationPreviousBaseClass = "inline-flex items-center justify-center gap-2 whitespace-nowrap rounded-md text-sm font-medium transition-all disabled:pointer-events-none disabled:opacity-50 [&_svg]:pointer-events-none [&_svg:not([class*='size-'])]:size-4 shrink-0 [&_svg]:shrink-0 outline-none focus-visible:border-ring focus-visible:ring-ring/50 focus-visible:ring-[3px] hover:bg-accent hover:text-accent-foreground h-9 gap-1 px-2.5 sm:pl-2.5"

// PaginationNextProps for [PaginationNext].
type PaginationNextProps struct{}

// PaginationNext renders a pagination link for going to the next page.
// Pass additional attributes (like [h.Href]) and children as needed.
func PaginationNext(props PaginationNextProps, children ...Node) Node {
	return h.A(
		h.Aria("label", "Go to next page"),
		h.Data("slot", "pagination-link"),
		JoinAttrs("class",
			h.Class(paginationNextBaseClass),
			Group(children),
		),
		h.Span(h.Class("hidden sm:block"), Text("Next")),
		Raw(`<svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="size-4"><path d="m9 18 6-6-6-6"/></svg>`),
	)
}

const paginationNextBaseClass = "inline-flex items-center justify-center gap-2 whitespace-nowrap rounded-md text-sm font-medium transition-all disabled:pointer-events-none disabled:opacity-50 [&_svg]:pointer-events-none [&_svg:not([class*='size-'])]:size-4 shrink-0 [&_svg]:shrink-0 outline-none focus-visible:border-ring focus-visible:ring-ring/50 focus-visible:ring-[3px] hover:bg-accent hover:text-accent-foreground h-9 gap-1 px-2.5 sm:pr-2.5"

// PaginationEllipsisProps for [PaginationEllipsis].
type PaginationEllipsisProps struct{}

// PaginationEllipsis renders an ellipsis indicator for pagination.
// Pass additional attributes and children as needed.
func PaginationEllipsis(props PaginationEllipsisProps, children ...Node) Node {
	return h.Span(
		h.Aria("hidden", "true"),
		h.Data("slot", "pagination-ellipsis"),
		JoinAttrs("class",
			h.Class(paginationEllipsisBaseClass),
			Group(children),
		),
		Raw(`<svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="size-4"><circle cx="12" cy="12" r="1"/><circle cx="19" cy="12" r="1"/><circle cx="5" cy="12" r="1"/></svg>`),
		h.Span(h.Class("sr-only"), Text("More pages")),
	)
}

const paginationEllipsisBaseClass = "flex size-9 items-center justify-center"
