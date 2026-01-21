package ui_test

import (
	"testing"

	. "maragu.dev/gomponents/html"
	"maragu.dev/is"

	ui "maragu.dev/gomponents-shadcn-ui"
)

func TestSwitch(t *testing.T) {
	t.Run("renders with default styling", func(t *testing.T) {
		got := render(t, ui.Switch(ui.SwitchProps{}))
		want := `<input type="checkbox" role="switch" data-slot="switch" class="peer appearance-none inline-flex h-5 w-9 shrink-0 items-center rounded-full border border-transparent shadow-xs transition-all outline-none cursor-pointer bg-input dark:bg-input/80 checked:bg-primary focus-visible:border-ring focus-visible:ring-ring/50 focus-visible:ring-[3px] disabled:cursor-not-allowed disabled:opacity-50 before:content-[&#39;&#39;] before:block before:size-4 before:rounded-full before:bg-background before:shadow-xs before:ring-0 before:transition-transform before:translate-x-0 checked:before:translate-x-4 checked:before:bg-primary-foreground dark:before:bg-foreground dark:checked:before:bg-primary-foreground">`
		is.Equal(t, want, got)
	})

	t.Run("renders with ID attribute", func(t *testing.T) {
		got := render(t, ui.Switch(ui.SwitchProps{}, ID("airplane-mode")))
		want := `<input type="checkbox" role="switch" data-slot="switch" class="peer appearance-none inline-flex h-5 w-9 shrink-0 items-center rounded-full border border-transparent shadow-xs transition-all outline-none cursor-pointer bg-input dark:bg-input/80 checked:bg-primary focus-visible:border-ring focus-visible:ring-ring/50 focus-visible:ring-[3px] disabled:cursor-not-allowed disabled:opacity-50 before:content-[&#39;&#39;] before:block before:size-4 before:rounded-full before:bg-background before:shadow-xs before:ring-0 before:transition-transform before:translate-x-0 checked:before:translate-x-4 checked:before:bg-primary-foreground dark:before:bg-foreground dark:checked:before:bg-primary-foreground" id="airplane-mode">`
		is.Equal(t, want, got)
	})

	t.Run("renders checked state", func(t *testing.T) {
		got := render(t, ui.Switch(ui.SwitchProps{}, Checked()))
		want := `<input type="checkbox" role="switch" data-slot="switch" class="peer appearance-none inline-flex h-5 w-9 shrink-0 items-center rounded-full border border-transparent shadow-xs transition-all outline-none cursor-pointer bg-input dark:bg-input/80 checked:bg-primary focus-visible:border-ring focus-visible:ring-ring/50 focus-visible:ring-[3px] disabled:cursor-not-allowed disabled:opacity-50 before:content-[&#39;&#39;] before:block before:size-4 before:rounded-full before:bg-background before:shadow-xs before:ring-0 before:transition-transform before:translate-x-0 checked:before:translate-x-4 checked:before:bg-primary-foreground dark:before:bg-foreground dark:checked:before:bg-primary-foreground" checked>`
		is.Equal(t, want, got)
	})

	t.Run("renders disabled state", func(t *testing.T) {
		got := render(t, ui.Switch(ui.SwitchProps{}, Disabled()))
		want := `<input type="checkbox" role="switch" data-slot="switch" class="peer appearance-none inline-flex h-5 w-9 shrink-0 items-center rounded-full border border-transparent shadow-xs transition-all outline-none cursor-pointer bg-input dark:bg-input/80 checked:bg-primary focus-visible:border-ring focus-visible:ring-ring/50 focus-visible:ring-[3px] disabled:cursor-not-allowed disabled:opacity-50 before:content-[&#39;&#39;] before:block before:size-4 before:rounded-full before:bg-background before:shadow-xs before:ring-0 before:transition-transform before:translate-x-0 checked:before:translate-x-4 checked:before:bg-primary-foreground dark:before:bg-foreground dark:checked:before:bg-primary-foreground" disabled>`
		is.Equal(t, want, got)
	})
}
