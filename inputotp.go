package ui

import (
	"strconv"

	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/components"
	h "maragu.dev/gomponents/html"
)

// InputOTPProps for [InputOTP].
type InputOTPProps struct {
	// MaxLength sets the maximum number of characters.
	MaxLength int
}

// InputOTP renders a container for OTP (one-time password) input slots.
// Use with Datastar for managing input values across slots.
func InputOTP(props InputOTPProps, children ...Node) Node {
	attrs := []Node{
		h.Data("slot", "input-otp"),
	}

	if props.MaxLength > 0 {
		attrs = append(attrs, h.Data("max-length", intToString(props.MaxLength)))
	}

	attrs = append(attrs,
		JoinAttrs("class",
			h.Class(inputOTPBaseClass),
			Group(children),
		),
	)

	return h.Div(attrs...)
}

const inputOTPBaseClass = "flex items-center gap-2 has-[:disabled]:opacity-50"

// InputOTPGroupProps for [InputOTPGroup].
type InputOTPGroupProps struct{}

// InputOTPGroup renders a group container for OTP slots.
func InputOTPGroup(props InputOTPGroupProps, children ...Node) Node {
	return h.Div(
		h.Data("slot", "input-otp-group"),
		JoinAttrs("class",
			h.Class(inputOTPGroupBaseClass),
			Group(children),
		),
	)
}

const inputOTPGroupBaseClass = "flex items-center"

// InputOTPSlotProps for [InputOTPSlot].
type InputOTPSlotProps struct {
	// Active indicates if this slot is currently focused.
	Active bool
}

// InputOTPSlot renders an individual OTP digit slot.
// Use native input elements inside for actual input handling.
func InputOTPSlot(props InputOTPSlotProps, children ...Node) Node {
	attrs := []Node{
		h.Data("slot", "input-otp-slot"),
	}

	if props.Active {
		attrs = append(attrs, h.Data("active", "true"))
	}

	attrs = append(attrs,
		JoinAttrs("class",
			h.Class(inputOTPSlotBaseClass),
			Group(children),
		),
	)

	return h.Div(attrs...)
}

const inputOTPSlotBaseClass = "data-[active=true]:border-ring data-[active=true]:ring-ring/50 border-input relative flex h-9 w-9 items-center justify-center border-y border-r text-sm shadow-xs transition-all outline-none first:rounded-l-md first:border-l last:rounded-r-md data-[active=true]:z-10 data-[active=true]:ring-[3px]"

// InputOTPSeparatorProps for [InputOTPSeparator].
type InputOTPSeparatorProps struct{}

// InputOTPSeparator renders a visual separator between OTP slot groups.
func InputOTPSeparator(props InputOTPSeparatorProps, children ...Node) Node {
	return h.Div(
		h.Role("separator"),
		h.Data("slot", "input-otp-separator"),
		JoinAttrs("class",
			h.Class(inputOTPSeparatorBaseClass),
			Group(children),
		),
	)
}

const inputOTPSeparatorBaseClass = ""

// intToString converts an int to a string for attribute values.
func intToString(i int) string {
	return strconv.Itoa(i)
}
