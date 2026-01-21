package ui_test

import (
	"testing"

	. "maragu.dev/gomponents/html"
	"maragu.dev/is"

	ui "maragu.dev/gomponents-shadcn-ui"
)

func TestRadioGroup(t *testing.T) {
	t.Run("renders with default styling", func(t *testing.T) {
		got := render(t, ui.RadioGroup(ui.RadioGroupProps{}))
		want := `<div role="radiogroup" data-slot="radio-group" class="grid gap-3"></div>`
		is.Equal(t, want, got)
	})
}

func TestRadioGroupItem(t *testing.T) {
	t.Run("renders with default styling", func(t *testing.T) {
		got := render(t, ui.RadioGroupItem(ui.RadioGroupItemProps{}))
		want := `<input type="radio" data-slot="radio-group-item" class="peer border-input text-primary focus-visible:border-ring focus-visible:ring-ring/50 aria-invalid:ring-destructive/20 dark:aria-invalid:ring-destructive/40 aria-invalid:border-destructive dark:bg-input/30 aspect-square size-4 shrink-0 rounded-full border shadow-xs transition-[color,box-shadow] outline-none focus-visible:ring-[3px] disabled:cursor-not-allowed disabled:opacity-50 appearance-none checked:bg-primary checked:border-primary cursor-pointer">`
		is.Equal(t, want, got)
	})

	t.Run("renders with name and value", func(t *testing.T) {
		got := render(t, ui.RadioGroupItem(ui.RadioGroupItemProps{}, Name("option"), Value("1")))
		want := `<input type="radio" data-slot="radio-group-item" class="peer border-input text-primary focus-visible:border-ring focus-visible:ring-ring/50 aria-invalid:ring-destructive/20 dark:aria-invalid:ring-destructive/40 aria-invalid:border-destructive dark:bg-input/30 aspect-square size-4 shrink-0 rounded-full border shadow-xs transition-[color,box-shadow] outline-none focus-visible:ring-[3px] disabled:cursor-not-allowed disabled:opacity-50 appearance-none checked:bg-primary checked:border-primary cursor-pointer" name="option" value="1">`
		is.Equal(t, want, got)
	})

	t.Run("renders checked state", func(t *testing.T) {
		got := render(t, ui.RadioGroupItem(ui.RadioGroupItemProps{}, Name("option"), Checked()))
		want := `<input type="radio" data-slot="radio-group-item" class="peer border-input text-primary focus-visible:border-ring focus-visible:ring-ring/50 aria-invalid:ring-destructive/20 dark:aria-invalid:ring-destructive/40 aria-invalid:border-destructive dark:bg-input/30 aspect-square size-4 shrink-0 rounded-full border shadow-xs transition-[color,box-shadow] outline-none focus-visible:ring-[3px] disabled:cursor-not-allowed disabled:opacity-50 appearance-none checked:bg-primary checked:border-primary cursor-pointer" name="option" checked>`
		is.Equal(t, want, got)
	})
}
