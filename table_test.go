package ui_test

import (
	"testing"

	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
	"maragu.dev/is"

	ui "maragu.dev/gomponents-shadcn-ui"
)

func TestTable(t *testing.T) {
	t.Run("renders with default styling", func(t *testing.T) {
		got := render(t, ui.Table(ui.TableProps{}))
		want := `<div data-slot="table-container" class="relative w-full overflow-x-auto"><table data-slot="table" class="w-full caption-bottom text-sm"></table></div>`
		is.Equal(t, want, got)
	})

	t.Run("merges custom class into single attribute", func(t *testing.T) {
		got := render(t, ui.Table(ui.TableProps{}, Class("my-table")))
		want := `<div data-slot="table-container" class="relative w-full overflow-x-auto"><table data-slot="table" class="w-full caption-bottom text-sm my-table"></table></div>`
		is.Equal(t, want, got)
	})
}

func TestTableHeader(t *testing.T) {
	t.Run("renders with default styling", func(t *testing.T) {
		got := render(t, ui.TableHeader(ui.TableHeaderProps{}))
		want := `<thead data-slot="table-header" class="[&amp;_tr]:border-b"></thead>`
		is.Equal(t, want, got)
	})
}

func TestTableBody(t *testing.T) {
	t.Run("renders with default styling", func(t *testing.T) {
		got := render(t, ui.TableBody(ui.TableBodyProps{}))
		want := `<tbody data-slot="table-body" class="[&amp;_tr:last-child]:border-0"></tbody>`
		is.Equal(t, want, got)
	})
}

func TestTableFooter(t *testing.T) {
	t.Run("renders with default styling", func(t *testing.T) {
		got := render(t, ui.TableFooter(ui.TableFooterProps{}))
		want := `<tfoot data-slot="table-footer" class="bg-muted/50 border-t font-medium [&amp;&gt;tr]:last:border-b-0"></tfoot>`
		is.Equal(t, want, got)
	})
}

func TestTableRow(t *testing.T) {
	t.Run("renders with default styling", func(t *testing.T) {
		got := render(t, ui.TableRow(ui.TableRowProps{}))
		want := `<tr data-slot="table-row" class="hover:bg-muted/50 data-[state=selected]:bg-muted border-b transition-colors"></tr>`
		is.Equal(t, want, got)
	})
}

func TestTableHead(t *testing.T) {
	t.Run("renders with default styling", func(t *testing.T) {
		got := render(t, ui.TableHead(ui.TableHeadProps{}, Text("Header")))
		want := `<th data-slot="table-head" class="text-foreground h-10 px-2 text-left align-middle font-medium whitespace-nowrap [&amp;:has([role=checkbox])]:pr-0 [&amp;&gt;[role=checkbox]]:translate-y-[2px]">Header</th>`
		is.Equal(t, want, got)
	})
}

func TestTableCell(t *testing.T) {
	t.Run("renders with default styling", func(t *testing.T) {
		got := render(t, ui.TableCell(ui.TableCellProps{}, Text("Cell")))
		want := `<td data-slot="table-cell" class="p-2 align-middle whitespace-nowrap [&amp;:has([role=checkbox])]:pr-0 [&amp;&gt;[role=checkbox]]:translate-y-[2px]">Cell</td>`
		is.Equal(t, want, got)
	})
}

func TestTableCaption(t *testing.T) {
	t.Run("renders with default styling", func(t *testing.T) {
		got := render(t, ui.TableCaption(ui.TableCaptionProps{}, Text("Caption")))
		want := `<caption data-slot="table-caption" class="text-muted-foreground mt-4 text-sm">Caption</caption>`
		is.Equal(t, want, got)
	})
}
