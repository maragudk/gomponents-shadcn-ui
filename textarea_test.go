package ui_test

import (
	"testing"

	. "maragu.dev/gomponents/html"
	"maragu.dev/is"

	ui "maragu.dev/gomponents-shadcn-ui"
)

func TestTextarea(t *testing.T) {
	t.Run("renders with default styling", func(t *testing.T) {
		got := render(t, ui.Textarea(ui.TextareaProps{}))
		want := `<textarea data-slot="textarea" class="border-input placeholder:text-muted-foreground focus-visible:border-ring focus-visible:ring-ring/50 aria-invalid:ring-destructive/20 dark:aria-invalid:ring-destructive/40 aria-invalid:border-destructive dark:bg-input/30 flex field-sizing-content min-h-16 w-full rounded-md border bg-transparent px-3 py-2 text-base shadow-xs transition-[color,box-shadow] outline-none focus-visible:ring-[3px] disabled:cursor-not-allowed disabled:opacity-50 md:text-sm"></textarea>`
		is.Equal(t, want, got)
	})

	t.Run("merges custom class into single attribute", func(t *testing.T) {
		got := render(t, ui.Textarea(ui.TextareaProps{}, Class("max-w-sm")))
		want := `<textarea data-slot="textarea" class="border-input placeholder:text-muted-foreground focus-visible:border-ring focus-visible:ring-ring/50 aria-invalid:ring-destructive/20 dark:aria-invalid:ring-destructive/40 aria-invalid:border-destructive dark:bg-input/30 flex field-sizing-content min-h-16 w-full rounded-md border bg-transparent px-3 py-2 text-base shadow-xs transition-[color,box-shadow] outline-none focus-visible:ring-[3px] disabled:cursor-not-allowed disabled:opacity-50 md:text-sm max-w-sm"></textarea>`
		is.Equal(t, want, got)
	})

	t.Run("renders with placeholder", func(t *testing.T) {
		got := render(t, ui.Textarea(ui.TextareaProps{}, Placeholder("Enter message")))
		want := `<textarea data-slot="textarea" class="border-input placeholder:text-muted-foreground focus-visible:border-ring focus-visible:ring-ring/50 aria-invalid:ring-destructive/20 dark:aria-invalid:ring-destructive/40 aria-invalid:border-destructive dark:bg-input/30 flex field-sizing-content min-h-16 w-full rounded-md border bg-transparent px-3 py-2 text-base shadow-xs transition-[color,box-shadow] outline-none focus-visible:ring-[3px] disabled:cursor-not-allowed disabled:opacity-50 md:text-sm" placeholder="Enter message"></textarea>`
		is.Equal(t, want, got)
	})
}
