package ui_test

import (
	"testing"

	. "maragu.dev/gomponents/html"
	"maragu.dev/is"

	ui "maragu.dev/gomponents-shadcn-ui"
)

func TestSkeleton(t *testing.T) {
	t.Run("renders with default styling", func(t *testing.T) {
		got := render(t, ui.Skeleton(ui.SkeletonProps{}))
		want := `<div class="bg-accent animate-pulse rounded-md"></div>`
		is.Equal(t, want, got)
	})

	t.Run("merges custom class into single attribute", func(t *testing.T) {
		got := render(t, ui.Skeleton(ui.SkeletonProps{}, Class("h-4 w-[250px]")))
		want := `<div class="bg-accent animate-pulse rounded-md h-4 w-[250px]"></div>`
		is.Equal(t, want, got)
	})
}
