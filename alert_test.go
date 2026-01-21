package ui_test

import (
	"testing"

	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
	"maragu.dev/is"

	ui "maragu.dev/gomponents-shadcn-ui"
)

func TestAlert(t *testing.T) {
	t.Run("renders with default variant", func(t *testing.T) {
		got := render(t, ui.Alert(ui.AlertProps{}, Text("Alert")))
		want := `<div role="alert" class="relative w-full rounded-lg border px-4 py-3 text-sm grid has-[&gt;svg]:grid-cols-[calc(var(--spacing)*4)_1fr] grid-cols-[0_1fr] has-[&gt;svg]:gap-x-3 gap-y-0.5 items-start [&amp;&gt;svg]:size-4 [&amp;&gt;svg]:translate-y-0.5 [&amp;&gt;svg]:text-current bg-card text-card-foreground">Alert</div>`
		is.Equal(t, want, got)
	})

	t.Run("renders destructive variant", func(t *testing.T) {
		got := render(t, ui.Alert(ui.AlertProps{Variant: ui.AlertVariantDestructive}, Text("Error")))
		want := `<div role="alert" class="relative w-full rounded-lg border px-4 py-3 text-sm grid has-[&gt;svg]:grid-cols-[calc(var(--spacing)*4)_1fr] grid-cols-[0_1fr] has-[&gt;svg]:gap-x-3 gap-y-0.5 items-start [&amp;&gt;svg]:size-4 [&amp;&gt;svg]:translate-y-0.5 [&amp;&gt;svg]:text-current text-destructive bg-card [&amp;&gt;svg]:text-current *:data-[slot=alert-description]:text-destructive/90">Error</div>`
		is.Equal(t, want, got)
	})

	t.Run("merges custom class into single attribute", func(t *testing.T) {
		got := render(t, ui.Alert(ui.AlertProps{}, Class("my-4")))
		want := `<div role="alert" class="relative w-full rounded-lg border px-4 py-3 text-sm grid has-[&gt;svg]:grid-cols-[calc(var(--spacing)*4)_1fr] grid-cols-[0_1fr] has-[&gt;svg]:gap-x-3 gap-y-0.5 items-start [&amp;&gt;svg]:size-4 [&amp;&gt;svg]:translate-y-0.5 [&amp;&gt;svg]:text-current bg-card text-card-foreground my-4"></div>`
		is.Equal(t, want, got)
	})
}

func TestAlertTitle(t *testing.T) {
	t.Run("renders with default styling", func(t *testing.T) {
		got := render(t, ui.AlertTitle(ui.AlertTitleProps{}, Text("Title")))
		want := `<div class="col-start-2 line-clamp-1 min-h-4 font-medium tracking-tight">Title</div>`
		is.Equal(t, want, got)
	})

	t.Run("merges custom class into single attribute", func(t *testing.T) {
		got := render(t, ui.AlertTitle(ui.AlertTitleProps{}, Class("text-lg")))
		want := `<div class="col-start-2 line-clamp-1 min-h-4 font-medium tracking-tight text-lg"></div>`
		is.Equal(t, want, got)
	})
}

func TestAlertDescription(t *testing.T) {
	t.Run("renders with default styling", func(t *testing.T) {
		got := render(t, ui.AlertDescription(ui.AlertDescriptionProps{}, Text("Description")))
		want := `<div class="text-muted-foreground col-start-2 grid justify-items-start gap-1 text-sm [&amp;_p]:leading-relaxed">Description</div>`
		is.Equal(t, want, got)
	})

	t.Run("merges custom class into single attribute", func(t *testing.T) {
		got := render(t, ui.AlertDescription(ui.AlertDescriptionProps{}, Class("mt-2")))
		want := `<div class="text-muted-foreground col-start-2 grid justify-items-start gap-1 text-sm [&amp;_p]:leading-relaxed mt-2"></div>`
		is.Equal(t, want, got)
	})
}
