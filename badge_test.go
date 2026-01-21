package ui_test

import (
	"testing"

	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
	"maragu.dev/is"

	ui "maragu.dev/gomponents-shadcn-ui"
)

func TestBadge(t *testing.T) {
	t.Run("renders with default variant", func(t *testing.T) {
		got := render(t, ui.Badge(ui.BadgeProps{}, Text("Badge")))
		want := `<span class="inline-flex items-center justify-center rounded-full border px-2 py-0.5 text-xs font-medium w-fit whitespace-nowrap shrink-0 [&amp;&gt;svg]:size-3 gap-1 [&amp;&gt;svg]:pointer-events-none focus-visible:border-ring focus-visible:ring-ring/50 focus-visible:ring-[3px] aria-invalid:ring-destructive/20 dark:aria-invalid:ring-destructive/40 aria-invalid:border-destructive transition-[color,box-shadow] overflow-hidden border-transparent bg-primary text-primary-foreground [a&amp;]:hover:bg-primary/90">Badge</span>`
		is.Equal(t, want, got)
	})

	t.Run("renders secondary variant", func(t *testing.T) {
		got := render(t, ui.Badge(ui.BadgeProps{Variant: ui.BadgeVariantSecondary}, Text("Secondary")))
		want := `<span class="inline-flex items-center justify-center rounded-full border px-2 py-0.5 text-xs font-medium w-fit whitespace-nowrap shrink-0 [&amp;&gt;svg]:size-3 gap-1 [&amp;&gt;svg]:pointer-events-none focus-visible:border-ring focus-visible:ring-ring/50 focus-visible:ring-[3px] aria-invalid:ring-destructive/20 dark:aria-invalid:ring-destructive/40 aria-invalid:border-destructive transition-[color,box-shadow] overflow-hidden border-transparent bg-secondary text-secondary-foreground [a&amp;]:hover:bg-secondary/90">Secondary</span>`
		is.Equal(t, want, got)
	})

	t.Run("renders destructive variant", func(t *testing.T) {
		got := render(t, ui.Badge(ui.BadgeProps{Variant: ui.BadgeVariantDestructive}, Text("Destructive")))
		want := `<span class="inline-flex items-center justify-center rounded-full border px-2 py-0.5 text-xs font-medium w-fit whitespace-nowrap shrink-0 [&amp;&gt;svg]:size-3 gap-1 [&amp;&gt;svg]:pointer-events-none focus-visible:border-ring focus-visible:ring-ring/50 focus-visible:ring-[3px] aria-invalid:ring-destructive/20 dark:aria-invalid:ring-destructive/40 aria-invalid:border-destructive transition-[color,box-shadow] overflow-hidden border-transparent bg-destructive text-white [a&amp;]:hover:bg-destructive/90 focus-visible:ring-destructive/20 dark:focus-visible:ring-destructive/40 dark:bg-destructive/60">Destructive</span>`
		is.Equal(t, want, got)
	})

	t.Run("renders outline variant", func(t *testing.T) {
		got := render(t, ui.Badge(ui.BadgeProps{Variant: ui.BadgeVariantOutline}, Text("Outline")))
		want := `<span class="inline-flex items-center justify-center rounded-full border px-2 py-0.5 text-xs font-medium w-fit whitespace-nowrap shrink-0 [&amp;&gt;svg]:size-3 gap-1 [&amp;&gt;svg]:pointer-events-none focus-visible:border-ring focus-visible:ring-ring/50 focus-visible:ring-[3px] aria-invalid:ring-destructive/20 dark:aria-invalid:ring-destructive/40 aria-invalid:border-destructive transition-[color,box-shadow] overflow-hidden text-foreground [a&amp;]:hover:bg-accent [a&amp;]:hover:text-accent-foreground">Outline</span>`
		is.Equal(t, want, got)
	})

	t.Run("merges custom class into single attribute", func(t *testing.T) {
		got := render(t, ui.Badge(ui.BadgeProps{}, Class("my-custom-class"), Text("Custom")))
		want := `<span class="inline-flex items-center justify-center rounded-full border px-2 py-0.5 text-xs font-medium w-fit whitespace-nowrap shrink-0 [&amp;&gt;svg]:size-3 gap-1 [&amp;&gt;svg]:pointer-events-none focus-visible:border-ring focus-visible:ring-ring/50 focus-visible:ring-[3px] aria-invalid:ring-destructive/20 dark:aria-invalid:ring-destructive/40 aria-invalid:border-destructive transition-[color,box-shadow] overflow-hidden border-transparent bg-primary text-primary-foreground [a&amp;]:hover:bg-primary/90 my-custom-class">Custom</span>`
		is.Equal(t, want, got)
	})
}
