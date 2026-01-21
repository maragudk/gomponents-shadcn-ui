package ui_test

import (
	"testing"

	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
	"maragu.dev/is"

	ui "maragu.dev/gomponents-shadcn-ui"
)

func TestBreadcrumb(t *testing.T) {
	t.Run("renders with default styling", func(t *testing.T) {
		got := render(t, ui.Breadcrumb(ui.BreadcrumbProps{}))
		want := `<nav aria-label="breadcrumb" data-slot="breadcrumb"></nav>`
		is.Equal(t, want, got)
	})
}

func TestBreadcrumbList(t *testing.T) {
	t.Run("renders with default styling", func(t *testing.T) {
		got := render(t, ui.BreadcrumbList(ui.BreadcrumbListProps{}))
		want := `<ol data-slot="breadcrumb-list" class="text-muted-foreground flex flex-wrap items-center gap-1.5 text-sm break-words sm:gap-2.5"></ol>`
		is.Equal(t, want, got)
	})

	t.Run("merges custom class into single attribute", func(t *testing.T) {
		got := render(t, ui.BreadcrumbList(ui.BreadcrumbListProps{}, Class("custom")))
		want := `<ol data-slot="breadcrumb-list" class="text-muted-foreground flex flex-wrap items-center gap-1.5 text-sm break-words sm:gap-2.5 custom"></ol>`
		is.Equal(t, want, got)
	})
}

func TestBreadcrumbItem(t *testing.T) {
	t.Run("renders with default styling", func(t *testing.T) {
		got := render(t, ui.BreadcrumbItem(ui.BreadcrumbItemProps{}))
		want := `<li data-slot="breadcrumb-item" class="inline-flex items-center gap-1.5"></li>`
		is.Equal(t, want, got)
	})
}

func TestBreadcrumbLink(t *testing.T) {
	t.Run("renders with default styling", func(t *testing.T) {
		got := render(t, ui.BreadcrumbLink(ui.BreadcrumbLinkProps{}, Href("/"), Text("Home")))
		want := `<a data-slot="breadcrumb-link" class="hover:text-foreground transition-colors" href="/">Home</a>`
		is.Equal(t, want, got)
	})
}

func TestBreadcrumbPage(t *testing.T) {
	t.Run("renders with default styling", func(t *testing.T) {
		got := render(t, ui.BreadcrumbPage(ui.BreadcrumbPageProps{}, Text("Current")))
		want := `<span data-slot="breadcrumb-page" role="link" aria-disabled="true" aria-current="page" class="text-foreground font-normal">Current</span>`
		is.Equal(t, want, got)
	})
}

func TestBreadcrumbSeparator(t *testing.T) {
	t.Run("renders with default chevron icon", func(t *testing.T) {
		got := render(t, ui.BreadcrumbSeparator(ui.BreadcrumbSeparatorProps{}))
		want := `<li data-slot="breadcrumb-separator" role="presentation" aria-hidden="true" class="[&amp;&gt;svg]:size-3.5"><svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="size-3.5"><path d="m9 18 6-6-6-6"/></svg></li>`
		is.Equal(t, want, got)
	})
}

func TestBreadcrumbEllipsis(t *testing.T) {
	t.Run("renders with default styling", func(t *testing.T) {
		got := render(t, ui.BreadcrumbEllipsis(ui.BreadcrumbEllipsisProps{}))
		want := `<span data-slot="breadcrumb-ellipsis" role="presentation" aria-hidden="true" class="flex size-9 items-center justify-center"><svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="size-4"><circle cx="12" cy="12" r="1"/><circle cx="19" cy="12" r="1"/><circle cx="5" cy="12" r="1"/></svg><span class="sr-only">More</span></span>`
		is.Equal(t, want, got)
	})
}
