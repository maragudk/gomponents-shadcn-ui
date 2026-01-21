package ui_test

import (
	"testing"

	. "maragu.dev/gomponents/html"
	"maragu.dev/is"

	ui "maragu.dev/gomponents-shadcn-ui"
)

func TestAspectRatio(t *testing.T) {
	t.Run("renders with default square ratio", func(t *testing.T) {
		got := render(t, ui.AspectRatio(ui.AspectRatioProps{}))
		want := `<div data-slot="aspect-ratio" style="position: relative; width: 100%; aspect-ratio: 1;" class="[&amp;&gt;*]:absolute [&amp;&gt;*]:inset-0 [&amp;&gt;*]:size-full"></div>`
		is.Equal(t, want, got)
	})

	t.Run("renders with 16/9 ratio", func(t *testing.T) {
		got := render(t, ui.AspectRatio(ui.AspectRatioProps{Ratio: 16.0 / 9.0}))
		want := `<div data-slot="aspect-ratio" style="position: relative; width: 100%; aspect-ratio: 1.7777777777777777;" class="[&amp;&gt;*]:absolute [&amp;&gt;*]:inset-0 [&amp;&gt;*]:size-full"></div>`
		is.Equal(t, want, got)
	})

	t.Run("merges custom class into single attribute", func(t *testing.T) {
		got := render(t, ui.AspectRatio(ui.AspectRatioProps{Ratio: 1}, Class("w-[450px]")))
		want := `<div data-slot="aspect-ratio" style="position: relative; width: 100%; aspect-ratio: 1;" class="[&amp;&gt;*]:absolute [&amp;&gt;*]:inset-0 [&amp;&gt;*]:size-full w-[450px]"></div>`
		is.Equal(t, want, got)
	})
}
