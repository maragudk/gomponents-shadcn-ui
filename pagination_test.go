package ui_test

import (
	"testing"

	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
	"maragu.dev/is"

	ui "maragu.dev/gomponents-shadcn-ui"
)

func TestPagination(t *testing.T) {
	t.Run("renders with default styling", func(t *testing.T) {
		got := render(t, ui.Pagination(ui.PaginationProps{}))
		want := `<nav role="navigation" aria-label="pagination" data-slot="pagination" class="mx-auto flex w-full justify-center"></nav>`
		is.Equal(t, want, got)
	})

	t.Run("merges custom class into single attribute", func(t *testing.T) {
		got := render(t, ui.Pagination(ui.PaginationProps{}, Class("my-4")))
		want := `<nav role="navigation" aria-label="pagination" data-slot="pagination" class="mx-auto flex w-full justify-center my-4"></nav>`
		is.Equal(t, want, got)
	})
}

func TestPaginationContent(t *testing.T) {
	t.Run("renders with default styling", func(t *testing.T) {
		got := render(t, ui.PaginationContent(ui.PaginationContentProps{}))
		want := `<ul data-slot="pagination-content" class="flex flex-row items-center gap-1"></ul>`
		is.Equal(t, want, got)
	})
}

func TestPaginationItem(t *testing.T) {
	t.Run("renders with default styling", func(t *testing.T) {
		got := render(t, ui.PaginationItem(ui.PaginationItemProps{}))
		want := `<li data-slot="pagination-item"></li>`
		is.Equal(t, want, got)
	})
}

func TestPaginationLink(t *testing.T) {
	t.Run("renders inactive link", func(t *testing.T) {
		got := render(t, ui.PaginationLink(ui.PaginationLinkProps{}, Href("#"), Text("1")))
		want := `<a data-slot="pagination-link" class="inline-flex items-center justify-center gap-2 whitespace-nowrap rounded-md text-sm font-medium transition-all disabled:pointer-events-none disabled:opacity-50 [&amp;_svg]:pointer-events-none [&amp;_svg:not([class*=&#39;size-&#39;])]:size-4 shrink-0 [&amp;_svg]:shrink-0 outline-none focus-visible:border-ring focus-visible:ring-ring/50 focus-visible:ring-[3px] size-9 hover:bg-accent hover:text-accent-foreground" href="#">1</a>`
		is.Equal(t, want, got)
	})

	t.Run("renders active link", func(t *testing.T) {
		got := render(t, ui.PaginationLink(ui.PaginationLinkProps{IsActive: true}, Href("#"), Text("1")))
		want := `<a data-slot="pagination-link" aria-current="page" data-active="true" class="inline-flex items-center justify-center gap-2 whitespace-nowrap rounded-md text-sm font-medium transition-all disabled:pointer-events-none disabled:opacity-50 [&amp;_svg]:pointer-events-none [&amp;_svg:not([class*=&#39;size-&#39;])]:size-4 shrink-0 [&amp;_svg]:shrink-0 outline-none focus-visible:border-ring focus-visible:ring-ring/50 focus-visible:ring-[3px] size-9 border bg-background shadow-xs hover:bg-accent hover:text-accent-foreground" href="#">1</a>`
		is.Equal(t, want, got)
	})
}

func TestPaginationPrevious(t *testing.T) {
	t.Run("renders with default styling", func(t *testing.T) {
		got := render(t, ui.PaginationPrevious(ui.PaginationPreviousProps{}, Href("#")))
		want := `<a aria-label="Go to previous page" data-slot="pagination-link" class="inline-flex items-center justify-center gap-2 whitespace-nowrap rounded-md text-sm font-medium transition-all disabled:pointer-events-none disabled:opacity-50 [&amp;_svg]:pointer-events-none [&amp;_svg:not([class*=&#39;size-&#39;])]:size-4 shrink-0 [&amp;_svg]:shrink-0 outline-none focus-visible:border-ring focus-visible:ring-ring/50 focus-visible:ring-[3px] hover:bg-accent hover:text-accent-foreground h-9 gap-1 px-2.5 sm:pl-2.5" href="#"><svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="size-4"><path d="m15 18-6-6 6-6"/></svg><span class="hidden sm:block">Previous</span></a>`
		is.Equal(t, want, got)
	})
}

func TestPaginationNext(t *testing.T) {
	t.Run("renders with default styling", func(t *testing.T) {
		got := render(t, ui.PaginationNext(ui.PaginationNextProps{}, Href("#")))
		want := `<a aria-label="Go to next page" data-slot="pagination-link" class="inline-flex items-center justify-center gap-2 whitespace-nowrap rounded-md text-sm font-medium transition-all disabled:pointer-events-none disabled:opacity-50 [&amp;_svg]:pointer-events-none [&amp;_svg:not([class*=&#39;size-&#39;])]:size-4 shrink-0 [&amp;_svg]:shrink-0 outline-none focus-visible:border-ring focus-visible:ring-ring/50 focus-visible:ring-[3px] hover:bg-accent hover:text-accent-foreground h-9 gap-1 px-2.5 sm:pr-2.5" href="#"><span class="hidden sm:block">Next</span><svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="size-4"><path d="m9 18 6-6-6-6"/></svg></a>`
		is.Equal(t, want, got)
	})
}

func TestPaginationEllipsis(t *testing.T) {
	t.Run("renders with default styling", func(t *testing.T) {
		got := render(t, ui.PaginationEllipsis(ui.PaginationEllipsisProps{}))
		want := `<span aria-hidden="true" data-slot="pagination-ellipsis" class="flex size-9 items-center justify-center"><svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="size-4"><circle cx="12" cy="12" r="1"/><circle cx="19" cy="12" r="1"/><circle cx="5" cy="12" r="1"/></svg><span class="sr-only">More pages</span></span>`
		is.Equal(t, want, got)
	})
}
