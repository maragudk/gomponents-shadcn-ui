package ui_test

import (
	"strings"
	"testing"

	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
	"maragu.dev/is"

	ui "maragu.dev/gomponents-shadcn-ui"
)

func TestLabel(t *testing.T) {
	t.Run("renders a label with default styling", func(t *testing.T) {
		l := ui.Label(ui.LabelProps{}, Text("Email"))
		got := render(t, l)

		is.True(t, strings.Contains(got, "<label"))
		is.True(t, strings.Contains(got, "Email"))
		is.True(t, strings.Contains(got, "text-sm"))
		is.True(t, strings.Contains(got, "font-medium"))
		is.True(t, strings.Contains(got, "leading-none"))
		is.True(t, strings.Contains(got, "peer-disabled:cursor-not-allowed"))
		is.True(t, strings.Contains(got, "peer-disabled:opacity-70"))
	})

	t.Run("merges custom class into single attribute", func(t *testing.T) {
		l := ui.Label(ui.LabelProps{}, Class("my-custom-class"), Text("Password"))
		got := render(t, l)

		// Both classes present
		is.True(t, strings.Contains(got, "my-custom-class"))
		is.True(t, strings.Contains(got, "text-sm"))
		// Merged into single class attribute
		is.Equal(t, 1, strings.Count(got, "class="))
	})

	t.Run("accepts additional attributes", func(t *testing.T) {
		l := ui.Label(ui.LabelProps{}, For("email-input"), ID("my-label"), Text("Email"))
		got := render(t, l)

		is.True(t, strings.Contains(got, `for="email-input"`))
		is.True(t, strings.Contains(got, `id="my-label"`))
	})
}
