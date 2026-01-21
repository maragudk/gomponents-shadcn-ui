package ui_test

import (
	"testing"

	. "maragu.dev/gomponents"
	"maragu.dev/is"

	ui "maragu.dev/gomponents-shadcn-ui"
)

func TestToggle(t *testing.T) {
	t.Run("renders with default styling", func(t *testing.T) {
		got := render(t, ui.Toggle(ui.ToggleProps{}, Text("Toggle")))
		want := `<button type="button" data-slot="toggle" aria-pressed="false" data-state="off" class="inline-flex items-center justify-center gap-2 rounded-md text-sm font-medium hover:bg-muted hover:text-muted-foreground disabled:pointer-events-none disabled:opacity-50 data-[state=on]:bg-accent data-[state=on]:text-accent-foreground [&amp;_svg]:pointer-events-none [&amp;_svg:not([class*=&#39;size-&#39;])]:size-4 [&amp;_svg]:shrink-0 focus-visible:border-ring focus-visible:ring-ring/50 focus-visible:ring-[3px] outline-none transition-[color,box-shadow] aria-invalid:ring-destructive/20 dark:aria-invalid:ring-destructive/40 aria-invalid:border-destructive whitespace-nowrap cursor-pointer bg-transparent h-9 px-2 min-w-9">Toggle</button>`
		is.Equal(t, want, got)
	})

	t.Run("renders pressed state", func(t *testing.T) {
		got := render(t, ui.Toggle(ui.ToggleProps{Pressed: true}, Text("Toggle")))
		want := `<button type="button" data-slot="toggle" aria-pressed="true" data-state="on" class="inline-flex items-center justify-center gap-2 rounded-md text-sm font-medium hover:bg-muted hover:text-muted-foreground disabled:pointer-events-none disabled:opacity-50 data-[state=on]:bg-accent data-[state=on]:text-accent-foreground [&amp;_svg]:pointer-events-none [&amp;_svg:not([class*=&#39;size-&#39;])]:size-4 [&amp;_svg]:shrink-0 focus-visible:border-ring focus-visible:ring-ring/50 focus-visible:ring-[3px] outline-none transition-[color,box-shadow] aria-invalid:ring-destructive/20 dark:aria-invalid:ring-destructive/40 aria-invalid:border-destructive whitespace-nowrap cursor-pointer bg-transparent h-9 px-2 min-w-9">Toggle</button>`
		is.Equal(t, want, got)
	})

	t.Run("renders outline variant", func(t *testing.T) {
		got := render(t, ui.Toggle(ui.ToggleProps{Variant: ui.ToggleVariantOutline}, Text("Toggle")))
		want := `<button type="button" data-slot="toggle" aria-pressed="false" data-state="off" class="inline-flex items-center justify-center gap-2 rounded-md text-sm font-medium hover:bg-muted hover:text-muted-foreground disabled:pointer-events-none disabled:opacity-50 data-[state=on]:bg-accent data-[state=on]:text-accent-foreground [&amp;_svg]:pointer-events-none [&amp;_svg:not([class*=&#39;size-&#39;])]:size-4 [&amp;_svg]:shrink-0 focus-visible:border-ring focus-visible:ring-ring/50 focus-visible:ring-[3px] outline-none transition-[color,box-shadow] aria-invalid:ring-destructive/20 dark:aria-invalid:ring-destructive/40 aria-invalid:border-destructive whitespace-nowrap cursor-pointer border border-input bg-transparent shadow-xs hover:bg-accent hover:text-accent-foreground h-9 px-2 min-w-9">Toggle</button>`
		is.Equal(t, want, got)
	})

	t.Run("renders small size", func(t *testing.T) {
		got := render(t, ui.Toggle(ui.ToggleProps{Size: ui.ToggleSizeSm}, Text("Toggle")))
		want := `<button type="button" data-slot="toggle" aria-pressed="false" data-state="off" class="inline-flex items-center justify-center gap-2 rounded-md text-sm font-medium hover:bg-muted hover:text-muted-foreground disabled:pointer-events-none disabled:opacity-50 data-[state=on]:bg-accent data-[state=on]:text-accent-foreground [&amp;_svg]:pointer-events-none [&amp;_svg:not([class*=&#39;size-&#39;])]:size-4 [&amp;_svg]:shrink-0 focus-visible:border-ring focus-visible:ring-ring/50 focus-visible:ring-[3px] outline-none transition-[color,box-shadow] aria-invalid:ring-destructive/20 dark:aria-invalid:ring-destructive/40 aria-invalid:border-destructive whitespace-nowrap cursor-pointer bg-transparent h-8 px-1.5 min-w-8">Toggle</button>`
		is.Equal(t, want, got)
	})

	t.Run("renders large size", func(t *testing.T) {
		got := render(t, ui.Toggle(ui.ToggleProps{Size: ui.ToggleSizeLg}, Text("Toggle")))
		want := `<button type="button" data-slot="toggle" aria-pressed="false" data-state="off" class="inline-flex items-center justify-center gap-2 rounded-md text-sm font-medium hover:bg-muted hover:text-muted-foreground disabled:pointer-events-none disabled:opacity-50 data-[state=on]:bg-accent data-[state=on]:text-accent-foreground [&amp;_svg]:pointer-events-none [&amp;_svg:not([class*=&#39;size-&#39;])]:size-4 [&amp;_svg]:shrink-0 focus-visible:border-ring focus-visible:ring-ring/50 focus-visible:ring-[3px] outline-none transition-[color,box-shadow] aria-invalid:ring-destructive/20 dark:aria-invalid:ring-destructive/40 aria-invalid:border-destructive whitespace-nowrap cursor-pointer bg-transparent h-10 px-2.5 min-w-10">Toggle</button>`
		is.Equal(t, want, got)
	})
}
