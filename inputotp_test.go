package ui_test

import (
	"testing"

	. "maragu.dev/gomponents"
	"maragu.dev/is"

	ui "maragu.dev/gomponents-shadcn-ui"
)

func TestInputOTP(t *testing.T) {
	t.Run("renders with default styling", func(t *testing.T) {
		got := render(t, ui.InputOTP(ui.InputOTPProps{}))
		want := `<div data-slot="input-otp" class="flex items-center gap-2 has-[:disabled]:opacity-50"></div>`
		is.Equal(t, want, got)
	})

	t.Run("renders with max length", func(t *testing.T) {
		got := render(t, ui.InputOTP(ui.InputOTPProps{MaxLength: 6}))
		want := `<div data-slot="input-otp" data-max-length="6" class="flex items-center gap-2 has-[:disabled]:opacity-50"></div>`
		is.Equal(t, want, got)
	})
}

func TestInputOTPGroup(t *testing.T) {
	t.Run("renders with default styling", func(t *testing.T) {
		got := render(t, ui.InputOTPGroup(ui.InputOTPGroupProps{}))
		want := `<div data-slot="input-otp-group" class="flex items-center"></div>`
		is.Equal(t, want, got)
	})
}

func TestInputOTPSlot(t *testing.T) {
	t.Run("renders with default styling", func(t *testing.T) {
		got := render(t, ui.InputOTPSlot(ui.InputOTPSlotProps{}, Text("1")))
		want := `<div data-slot="input-otp-slot" class="data-[active=true]:border-ring data-[active=true]:ring-ring/50 border-input relative flex h-9 w-9 items-center justify-center border-y border-r text-sm shadow-xs transition-all outline-none first:rounded-l-md first:border-l last:rounded-r-md data-[active=true]:z-10 data-[active=true]:ring-[3px]">1</div>`
		is.Equal(t, want, got)
	})

	t.Run("renders with active state", func(t *testing.T) {
		got := render(t, ui.InputOTPSlot(ui.InputOTPSlotProps{Active: true}))
		want := `<div data-slot="input-otp-slot" data-active="true" class="data-[active=true]:border-ring data-[active=true]:ring-ring/50 border-input relative flex h-9 w-9 items-center justify-center border-y border-r text-sm shadow-xs transition-all outline-none first:rounded-l-md first:border-l last:rounded-r-md data-[active=true]:z-10 data-[active=true]:ring-[3px]"></div>`
		is.Equal(t, want, got)
	})
}

func TestInputOTPSeparator(t *testing.T) {
	t.Run("renders with default styling", func(t *testing.T) {
		got := render(t, ui.InputOTPSeparator(ui.InputOTPSeparatorProps{}, Text("-")))
		want := `<div role="separator" data-slot="input-otp-separator" class="">-</div>`
		is.Equal(t, want, got)
	})
}
