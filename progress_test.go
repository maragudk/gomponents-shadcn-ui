package ui_test

import (
	"testing"

	. "maragu.dev/gomponents/html"
	"maragu.dev/is"

	ui "maragu.dev/gomponents-shadcn-ui"
)

func TestProgress(t *testing.T) {
	t.Run("renders with zero value", func(t *testing.T) {
		got := render(t, ui.Progress(ui.ProgressProps{Value: 0}))
		want := `<div role="progressbar" data-slot="progress" aria-valuenow="0" aria-valuemin="0" aria-valuemax="100" class="bg-primary/20 relative h-2 w-full overflow-hidden rounded-full"><div data-slot="progress-indicator" class="bg-primary h-full w-full flex-1 transition-all" style="transform: translateX(-100%);"></div></div>`
		is.Equal(t, want, got)
	})

	t.Run("renders with value of 50", func(t *testing.T) {
		got := render(t, ui.Progress(ui.ProgressProps{Value: 50}))
		want := `<div role="progressbar" data-slot="progress" aria-valuenow="50" aria-valuemin="0" aria-valuemax="100" class="bg-primary/20 relative h-2 w-full overflow-hidden rounded-full"><div data-slot="progress-indicator" class="bg-primary h-full w-full flex-1 transition-all" style="transform: translateX(-50%);"></div></div>`
		is.Equal(t, want, got)
	})

	t.Run("renders with value of 100", func(t *testing.T) {
		got := render(t, ui.Progress(ui.ProgressProps{Value: 100}))
		want := `<div role="progressbar" data-slot="progress" aria-valuenow="100" aria-valuemin="0" aria-valuemax="100" class="bg-primary/20 relative h-2 w-full overflow-hidden rounded-full"><div data-slot="progress-indicator" class="bg-primary h-full w-full flex-1 transition-all" style="transform: translateX(-0%);"></div></div>`
		is.Equal(t, want, got)
	})

	t.Run("merges custom class into single attribute", func(t *testing.T) {
		got := render(t, ui.Progress(ui.ProgressProps{Value: 33}, Class("w-60")))
		want := `<div role="progressbar" data-slot="progress" aria-valuenow="33" aria-valuemin="0" aria-valuemax="100" class="bg-primary/20 relative h-2 w-full overflow-hidden rounded-full w-60"><div data-slot="progress-indicator" class="bg-primary h-full w-full flex-1 transition-all" style="transform: translateX(-67%);"></div></div>`
		is.Equal(t, want, got)
	})
}
