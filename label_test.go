package ui_test

import (
	"strings"
	"testing"

	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
	"maragu.dev/is"

	ui "maragu.dev/gomponents-shadcn-ui"
)

func TestLabel(t *testing.T) {
	t.Run("renders a label with default styling", func(t *testing.T) {
		l := ui.Label(ui.LabelProps{}, g.Text("Email"))
		html := render(t, l)

		is.True(t, strings.Contains(html, "<label"))
		is.True(t, strings.Contains(html, "Email"))
		is.True(t, strings.Contains(html, "text-sm"))
		is.True(t, strings.Contains(html, "font-medium"))
		is.True(t, strings.Contains(html, "leading-none"))
		is.True(t, strings.Contains(html, "peer-disabled:cursor-not-allowed"))
		is.True(t, strings.Contains(html, "peer-disabled:opacity-70"))
	})

	t.Run("merges custom class into single attribute", func(t *testing.T) {
		l := ui.Label(ui.LabelProps{}, h.Class("my-custom-class"), g.Text("Password"))
		html := render(t, l)

		// Both classes present
		is.True(t, strings.Contains(html, "my-custom-class"))
		is.True(t, strings.Contains(html, "text-sm"))
		// Merged into single class attribute
		is.Equal(t, 1, strings.Count(html, "class="))
	})

	t.Run("accepts for attribute", func(t *testing.T) {
		l := ui.Label(ui.LabelProps{}, h.For("email-input"), g.Text("Email"))
		html := render(t, l)

		is.True(t, strings.Contains(html, `for="email-input"`))
	})

	t.Run("accepts additional attributes", func(t *testing.T) {
		l := ui.Label(ui.LabelProps{}, h.ID("my-label"), g.Text("Username"))
		html := render(t, l)

		is.True(t, strings.Contains(html, `id="my-label"`))
	})
}
