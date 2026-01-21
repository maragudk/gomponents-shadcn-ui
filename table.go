package ui

import (
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/components"
	h "maragu.dev/gomponents/html"
)

// TableProps for [Table].
type TableProps struct{}

// Table renders a table element with shadcn/ui styling wrapped in a scrollable container.
// Pass additional attributes and children as needed.
func Table(props TableProps, children ...Node) Node {
	return h.Div(
		h.Data("slot", "table-container"),
		h.Class("relative w-full overflow-x-auto"),
		h.Table(
			h.Data("slot", "table"),
			JoinAttrs("class",
				h.Class(tableBaseClass),
				Group(children),
			),
		),
	)
}

const tableBaseClass = "w-full caption-bottom text-sm"

// TableHeaderProps for [TableHeader].
type TableHeaderProps struct{}

// TableHeader renders a thead element with shadcn/ui styling.
// Pass additional attributes and children as needed.
func TableHeader(props TableHeaderProps, children ...Node) Node {
	return h.THead(
		h.Data("slot", "table-header"),
		JoinAttrs("class",
			h.Class(tableHeaderBaseClass),
			Group(children),
		),
	)
}

const tableHeaderBaseClass = "[&_tr]:border-b"

// TableBodyProps for [TableBody].
type TableBodyProps struct{}

// TableBody renders a tbody element with shadcn/ui styling.
// Pass additional attributes and children as needed.
func TableBody(props TableBodyProps, children ...Node) Node {
	return h.TBody(
		h.Data("slot", "table-body"),
		JoinAttrs("class",
			h.Class(tableBodyBaseClass),
			Group(children),
		),
	)
}

const tableBodyBaseClass = "[&_tr:last-child]:border-0"

// TableFooterProps for [TableFooter].
type TableFooterProps struct{}

// TableFooter renders a tfoot element with shadcn/ui styling.
// Pass additional attributes and children as needed.
func TableFooter(props TableFooterProps, children ...Node) Node {
	return h.TFoot(
		h.Data("slot", "table-footer"),
		JoinAttrs("class",
			h.Class(tableFooterBaseClass),
			Group(children),
		),
	)
}

const tableFooterBaseClass = "bg-muted/50 border-t font-medium [&>tr]:last:border-b-0"

// TableRowProps for [TableRow].
type TableRowProps struct{}

// TableRow renders a tr element with shadcn/ui styling.
// Pass additional attributes and children as needed.
func TableRow(props TableRowProps, children ...Node) Node {
	return h.Tr(
		h.Data("slot", "table-row"),
		JoinAttrs("class",
			h.Class(tableRowBaseClass),
			Group(children),
		),
	)
}

const tableRowBaseClass = "hover:bg-muted/50 data-[state=selected]:bg-muted border-b transition-colors"

// TableHeadProps for [TableHead].
type TableHeadProps struct{}

// TableHead renders a th element with shadcn/ui styling.
// Pass additional attributes and children as needed.
func TableHead(props TableHeadProps, children ...Node) Node {
	return h.Th(
		h.Data("slot", "table-head"),
		JoinAttrs("class",
			h.Class(tableHeadBaseClass),
			Group(children),
		),
	)
}

const tableHeadBaseClass = "text-foreground h-10 px-2 text-left align-middle font-medium whitespace-nowrap [&:has([role=checkbox])]:pr-0 [&>[role=checkbox]]:translate-y-[2px]"

// TableCellProps for [TableCell].
type TableCellProps struct{}

// TableCell renders a td element with shadcn/ui styling.
// Pass additional attributes and children as needed.
func TableCell(props TableCellProps, children ...Node) Node {
	return h.Td(
		h.Data("slot", "table-cell"),
		JoinAttrs("class",
			h.Class(tableCellBaseClass),
			Group(children),
		),
	)
}

const tableCellBaseClass = "p-2 align-middle whitespace-nowrap [&:has([role=checkbox])]:pr-0 [&>[role=checkbox]]:translate-y-[2px]"

// TableCaptionProps for [TableCaption].
type TableCaptionProps struct{}

// TableCaption renders a caption element with shadcn/ui styling.
// Pass additional attributes and children as needed.
func TableCaption(props TableCaptionProps, children ...Node) Node {
	return h.Caption(
		h.Data("slot", "table-caption"),
		JoinAttrs("class",
			h.Class(tableCaptionBaseClass),
			Group(children),
		),
	)
}

const tableCaptionBaseClass = "text-muted-foreground mt-4 text-sm"
