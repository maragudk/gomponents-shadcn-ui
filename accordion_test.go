package ui_test

import (
	"testing"

	. "maragu.dev/gomponents"
	"maragu.dev/is"

	ui "maragu.dev/gomponents-shadcn-ui"
)

func TestAccordion(t *testing.T) {
	t.Run("renders with default styling", func(t *testing.T) {
		got := render(t, ui.Accordion(ui.AccordionProps{}))
		want := `<div data-slot="accordion" class=""></div>`
		is.Equal(t, want, got)
	})
}

func TestAccordionItem(t *testing.T) {
	t.Run("renders with default styling", func(t *testing.T) {
		got := render(t, ui.AccordionItem(ui.AccordionItemProps{}))
		want := `<div data-slot="accordion-item" class="border-b last:border-b-0"></div>`
		is.Equal(t, want, got)
	})
}

func TestAccordionTrigger(t *testing.T) {
	t.Run("renders with default styling", func(t *testing.T) {
		got := render(t, ui.AccordionTrigger(ui.AccordionTriggerProps{}, Text("Toggle")))
		want := `<div class="flex"><button type="button" data-slot="accordion-trigger" class="focus-visible:border-ring focus-visible:ring-ring/50 flex flex-1 items-start justify-between gap-4 rounded-md py-4 text-left text-sm font-medium transition-all outline-none hover:underline focus-visible:ring-[3px] disabled:pointer-events-none disabled:opacity-50 [&amp;[data-state=open]&gt;svg]:rotate-180 cursor-pointer">Toggle</button></div>`
		is.Equal(t, want, got)
	})
}

func TestAccordionContent(t *testing.T) {
	t.Run("renders with default styling", func(t *testing.T) {
		got := render(t, ui.AccordionContent(ui.AccordionContentProps{}, Text("Content")))
		want := `<div data-slot="accordion-content" class="overflow-hidden text-sm pb-4">Content</div>`
		is.Equal(t, want, got)
	})
}
