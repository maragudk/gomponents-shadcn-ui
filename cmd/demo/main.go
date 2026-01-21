// Command demo generates a static demo site showcasing all components.
package main

import (
	"os"

	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/components"
	data "maragu.dev/gomponents-datastar"
	. "maragu.dev/gomponents/html"

	ui "maragu.dev/gomponents-shadcn-ui"
)

func main() {
	if err := run(); err != nil {
		_, _ = os.Stderr.WriteString(err.Error() + "\n")
		os.Exit(1)
	}
}

func run() error {
	if err := os.MkdirAll("demo", 0755); err != nil {
		return err
	}

	f, err := os.Create("demo/index.html")
	if err != nil {
		return err
	}

	if err := page().Render(f); err != nil {
		_ = f.Close()
		return err
	}
	return f.Close()
}

func page() Node {
	return HTML5(HTML5Props{
		Title:       "gomponents-shadcn-ui Demo",
		Description: "A demo of all gomponents-shadcn-ui components",
		Language:    "en",
		Head: []Node{
			Link(Rel("stylesheet"), Href("styles.css")),
			Script(Type("module"), Src("https://cdn.jsdelivr.net/gh/starfederation/datastar@v1.0.0-RC.7/bundles/datastar.js")),
		},
		Body: []Node{
			Class("bg-background text-foreground min-h-screen"),
			Div(
				Class("container mx-auto py-12 px-4 max-w-4xl"),
				header(),
				accordionSection(),
				alertSection(),
				alertDialogSection(),
				aspectRatioSection(),
				avatarSection(),
				badgeSection(),
				breadcrumbSection(),
				buttonSection(),
				cardSection(),
				checkboxSection(),
				collapsibleSection(),
				dialogSection(),
				dropdownMenuSection(),
				inputSection(),
				labelSection(),
				paginationSection(),
				popoverSection(),
				progressSection(),
				radioGroupSection(),
				selectSection(),
				separatorSection(),
				skeletonSection(),
				sliderSection(),
				switchSection(),
				tableSection(),
				tabsSection(),
				textareaSection(),
				toggleSection(),
				toggleGroupSection(),
				tooltipSection(),
			),
		},
	})
}

func header() Node {
	return Header(
		Class("mb-12"),
		H1(
			Class("text-4xl font-bold mb-4"),
			Text("gomponents-shadcn-ui"),
		),
		P(
			Class("text-muted-foreground text-lg mb-4"),
			A(
				Href("https://ui.shadcn.com"),
				Class("underline hover:text-foreground"),
				Text("shadcn/ui"),
			),
			Text(" components for Go, built with "),
			A(
				Href("https://www.gomponents.com"),
				Class("underline hover:text-foreground"),
				Text("gomponents"),
			),
			Text("."),
		),
		P(
			Class("text-muted-foreground"),
			A(
				Href("https://github.com/maragudk/gomponents-shadcn-ui"),
				Class("underline hover:text-foreground"),
				Text("View on GitHub"),
			),
		),
	)
}

func accordionSection() Node {
	return sectionWithSource("Accordion", "accordion.go",
		subsection("Default",
			ui.Accordion(ui.AccordionProps{},
				data.Signals(map[string]any{"item1Open": false, "item2Open": false, "item3Open": false}),
				ui.AccordionItem(ui.AccordionItemProps{},
					ui.AccordionTrigger(ui.AccordionTriggerProps{},
						data.On("click", "item1Open = !item1Open"),
						data.Attr("data-state", "item1Open ? 'open' : 'closed'"),
						Text("Is it accessible?"),
						Raw(`<svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="text-muted-foreground shrink-0 transition-transform duration-200"><path d="m6 9 6 6 6-6"/></svg>`),
					),
					ui.AccordionContent(ui.AccordionContentProps{},
						data.Show("item1Open"),
						Text("Yes. It adheres to the WAI-ARIA design pattern."),
					),
				),
				ui.AccordionItem(ui.AccordionItemProps{},
					ui.AccordionTrigger(ui.AccordionTriggerProps{},
						data.On("click", "item2Open = !item2Open"),
						data.Attr("data-state", "item2Open ? 'open' : 'closed'"),
						Text("Is it styled?"),
						Raw(`<svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="text-muted-foreground shrink-0 transition-transform duration-200"><path d="m6 9 6 6 6-6"/></svg>`),
					),
					ui.AccordionContent(ui.AccordionContentProps{},
						data.Show("item2Open"),
						Text("Yes. It comes with default styles that match the shadcn/ui aesthetic."),
					),
				),
				ui.AccordionItem(ui.AccordionItemProps{},
					ui.AccordionTrigger(ui.AccordionTriggerProps{},
						data.On("click", "item3Open = !item3Open"),
						data.Attr("data-state", "item3Open ? 'open' : 'closed'"),
						Text("Is it animated?"),
						Raw(`<svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="text-muted-foreground shrink-0 transition-transform duration-200"><path d="m6 9 6 6 6-6"/></svg>`),
					),
					ui.AccordionContent(ui.AccordionContentProps{},
						data.Show("item3Open"),
						Text("Yes. It's animated by default with CSS transitions."),
					),
				),
			),
		),
	)
}

func alertSection() Node {
	return sectionWithSource("Alert", "alert.go",
		P(
			Class("text-muted-foreground mb-6"),
			Text("Displays a callout for user attention."),
		),

		subsection("Default",
			ui.Alert(ui.AlertProps{},
				Raw(`<svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><circle cx="12" cy="12" r="10"/><path d="M12 16v-4"/><path d="M12 8h.01"/></svg>`),
				ui.AlertTitle(ui.AlertTitleProps{}, Text("Heads up!")),
				ui.AlertDescription(ui.AlertDescriptionProps{}, Text("You can add components to your app using the cli.")),
			),
		),

		subsection("Destructive",
			ui.Alert(ui.AlertProps{Variant: ui.AlertVariantDestructive},
				Raw(`<svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><circle cx="12" cy="12" r="10"/><line x1="12" x2="12" y1="8" y2="12"/><line x1="12" x2="12.01" y1="16" y2="16"/></svg>`),
				ui.AlertTitle(ui.AlertTitleProps{}, Text("Error")),
				ui.AlertDescription(ui.AlertDescriptionProps{}, Text("Your session has expired. Please log in again.")),
			),
		),

		subsection("Auto-Dismiss Notification",
			Div(
				data.Signals(map[string]any{"notification": false}),
				Class("flex flex-col gap-4"),
				ui.Button(ui.ButtonProps{},
					data.On("click", "$notification = true; setTimeout(() => $notification = false, 3000)"),
					Text("Show Notification"),
				),
				Div(
					data.Show("$notification"),
					ui.Alert(ui.AlertProps{},
						Raw(`<svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M22 11.08V12a10 10 0 1 1-5.93-9.14"/><polyline points="22 4 12 14.01 9 11.01"/></svg>`),
						ui.AlertTitle(ui.AlertTitleProps{}, Text("Success!")),
						ui.AlertDescription(ui.AlertDescriptionProps{}, Text("Your action was completed. This alert will dismiss in 3 seconds.")),
					),
				),
			),
		),
	)
}

func alertDialogSection() Node {
	return sectionWithSource("Alert Dialog", "alertdialog.go",
		subsection("Default",
			ui.AlertDialog(ui.AlertDialogProps{},
				data.Signals(map[string]any{"alertOpen": false}),
				ui.AlertDialogTrigger(ui.AlertDialogTriggerProps{},
					ui.Button(ui.ButtonProps{Variant: ui.ButtonVariantOutline},
						data.On("click", "alertOpen = true"),
						Text("Show Dialog"),
					),
				),
				Group([]Node{
					ui.AlertDialogOverlay(ui.AlertDialogOverlayProps{},
						data.Show("alertOpen"),
					),
					ui.AlertDialogContent(ui.AlertDialogContentProps{},
						data.Show("alertOpen"),
						ui.AlertDialogHeader(ui.AlertDialogHeaderProps{},
							ui.AlertDialogTitle(ui.AlertDialogTitleProps{}, Text("Are you absolutely sure?")),
							ui.AlertDialogDescription(ui.AlertDialogDescriptionProps{}, Text("This action cannot be undone. This will permanently delete your account and remove your data from our servers.")),
						),
						ui.AlertDialogFooter(ui.AlertDialogFooterProps{},
							ui.AlertDialogCancel(ui.AlertDialogCancelProps{},
								data.On("click", "alertOpen = false"),
								Text("Cancel"),
							),
							ui.AlertDialogAction(ui.AlertDialogActionProps{},
								data.On("click", "alertOpen = false"),
								Text("Continue"),
							),
						),
					),
				}),
			),
		),
	)
}

func aspectRatioSection() Node {
	return sectionWithSource("Aspect Ratio", "aspectratio.go",
		P(
			Class("text-muted-foreground mb-6"),
			Text("Displays content within a desired ratio."),
		),

		subsection("16:9",
			Div(
				Class("w-[450px]"),
				ui.AspectRatio(ui.AspectRatioProps{Ratio: 16.0 / 9.0},
					Img(
						Src("https://images.unsplash.com/photo-1588345921523-c2dcdb7f1dcd?w=800&dpr=2&q=80"),
						Alt("Photo by Drew Beamer"),
						Class("rounded-md object-cover"),
					),
				),
			),
		),

		subsection("Square",
			Div(
				Class("w-[200px]"),
				ui.AspectRatio(ui.AspectRatioProps{Ratio: 1},
					Div(Class("bg-muted flex items-center justify-center rounded-md"), Text("1:1")),
				),
			),
		),
	)
}

func avatarSection() Node {
	return sectionWithSource("Avatar", "avatar.go",
		P(
			Class("text-muted-foreground mb-6"),
			Text("An image element with a fallback for representing the user."),
		),

		subsection("With Image",
			Div(
				Class("flex flex-wrap gap-4"),
				ui.Avatar(ui.AvatarProps{},
					ui.AvatarImage(ui.AvatarImageProps{}, Src("https://github.com/shadcn.png"), Alt("@shadcn")),
				),
			),
		),

		subsection("With Fallback",
			Div(
				Class("flex flex-wrap gap-4"),
				ui.Avatar(ui.AvatarProps{},
					ui.AvatarFallback(ui.AvatarFallbackProps{}, Text("CN")),
				),
			),
		),

		subsection("Sizes",
			Div(
				Class("flex flex-wrap items-center gap-4"),
				ui.Avatar(ui.AvatarProps{}, Class("size-6"),
					ui.AvatarFallback(ui.AvatarFallbackProps{}, Class("text-xs"), Text("XS")),
				),
				ui.Avatar(ui.AvatarProps{},
					ui.AvatarFallback(ui.AvatarFallbackProps{}, Text("SM")),
				),
				ui.Avatar(ui.AvatarProps{}, Class("size-12"),
					ui.AvatarFallback(ui.AvatarFallbackProps{}, Text("MD")),
				),
				ui.Avatar(ui.AvatarProps{}, Class("size-16"),
					ui.AvatarFallback(ui.AvatarFallbackProps{}, Class("text-lg"), Text("LG")),
				),
			),
		),
	)
}

func badgeSection() Node {
	return sectionWithSource("Badge", "badge.go",
		P(
			Class("text-muted-foreground mb-6"),
			Text("Displays a badge or a component that looks like a badge."),
		),

		subsection("Variants",
			Div(
				Class("flex flex-wrap gap-4"),
				ui.Badge(ui.BadgeProps{Variant: ui.BadgeVariantDefault}, Text("Default")),
				ui.Badge(ui.BadgeProps{Variant: ui.BadgeVariantSecondary}, Text("Secondary")),
				ui.Badge(ui.BadgeProps{Variant: ui.BadgeVariantDestructive}, Text("Destructive")),
				ui.Badge(ui.BadgeProps{Variant: ui.BadgeVariantOutline}, Text("Outline")),
			),
		),
	)
}

func breadcrumbSection() Node {
	return sectionWithSource("Breadcrumb", "breadcrumb.go",
		P(
			Class("text-muted-foreground mb-6"),
			Text("Displays the path to the current resource using a hierarchy of links."),
		),

		subsection("Default",
			ui.Breadcrumb(ui.BreadcrumbProps{},
				ui.BreadcrumbList(ui.BreadcrumbListProps{},
					ui.BreadcrumbItem(ui.BreadcrumbItemProps{},
						ui.BreadcrumbLink(ui.BreadcrumbLinkProps{}, Href("#"), Text("Home")),
					),
					ui.BreadcrumbSeparator(ui.BreadcrumbSeparatorProps{}),
					ui.BreadcrumbItem(ui.BreadcrumbItemProps{},
						ui.BreadcrumbLink(ui.BreadcrumbLinkProps{}, Href("#"), Text("Components")),
					),
					ui.BreadcrumbSeparator(ui.BreadcrumbSeparatorProps{}),
					ui.BreadcrumbItem(ui.BreadcrumbItemProps{},
						ui.BreadcrumbPage(ui.BreadcrumbPageProps{}, Text("Breadcrumb")),
					),
				),
			),
		),

		subsection("With Ellipsis",
			ui.Breadcrumb(ui.BreadcrumbProps{},
				ui.BreadcrumbList(ui.BreadcrumbListProps{},
					ui.BreadcrumbItem(ui.BreadcrumbItemProps{},
						ui.BreadcrumbLink(ui.BreadcrumbLinkProps{}, Href("#"), Text("Home")),
					),
					ui.BreadcrumbSeparator(ui.BreadcrumbSeparatorProps{}),
					ui.BreadcrumbItem(ui.BreadcrumbItemProps{},
						ui.BreadcrumbEllipsis(ui.BreadcrumbEllipsisProps{}),
					),
					ui.BreadcrumbSeparator(ui.BreadcrumbSeparatorProps{}),
					ui.BreadcrumbItem(ui.BreadcrumbItemProps{},
						ui.BreadcrumbLink(ui.BreadcrumbLinkProps{}, Href("#"), Text("Components")),
					),
					ui.BreadcrumbSeparator(ui.BreadcrumbSeparatorProps{}),
					ui.BreadcrumbItem(ui.BreadcrumbItemProps{},
						ui.BreadcrumbPage(ui.BreadcrumbPageProps{}, Text("Breadcrumb")),
					),
				),
			),
		),
	)
}

func buttonSection() Node {
	return sectionWithSource("Button", "button.go",
		P(
			Class("text-muted-foreground mb-6"),
			Text("Displays a button or a component that looks like a button."),
		),

		subsection("Variants",
			Div(
				Class("flex flex-wrap gap-4"),
				ui.Button(ui.ButtonProps{Variant: ui.ButtonVariantDefault}, Text("Default")),
				ui.Button(ui.ButtonProps{Variant: ui.ButtonVariantSecondary}, Text("Secondary")),
				ui.Button(ui.ButtonProps{Variant: ui.ButtonVariantDestructive}, Text("Destructive")),
				ui.Button(ui.ButtonProps{Variant: ui.ButtonVariantOutline}, Text("Outline")),
				ui.Button(ui.ButtonProps{Variant: ui.ButtonVariantGhost}, Text("Ghost")),
				ui.Button(ui.ButtonProps{Variant: ui.ButtonVariantLink}, Text("Link")),
			),
		),

		subsection("Sizes",
			Div(
				Class("flex flex-wrap items-center gap-4"),
				ui.Button(ui.ButtonProps{Size: ui.ButtonSizeSm}, Text("Small")),
				ui.Button(ui.ButtonProps{Size: ui.ButtonSizeDefault}, Text("Default")),
				ui.Button(ui.ButtonProps{Size: ui.ButtonSizeLg}, Text("Large")),
				ui.Button(ui.ButtonProps{Size: ui.ButtonSizeIcon},
					Raw(`<svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M5 12h14"/><path d="M12 5v14"/></svg>`),
				),
			),
		),

		subsection("As Link",
			Div(
				Class("flex flex-wrap gap-4"),
				ui.ButtonA(ui.ButtonProps{}, Href("#"), Text("Link Button")),
				ui.ButtonA(ui.ButtonProps{Variant: ui.ButtonVariantOutline}, Href("#"), Text("Outline Link")),
			),
		),

		subsection("Disabled",
			Div(
				Class("flex flex-wrap gap-4"),
				ui.Button(ui.ButtonProps{}, Disabled(), Text("Disabled")),
				ui.Button(ui.ButtonProps{Variant: ui.ButtonVariantSecondary}, Disabled(), Text("Disabled")),
				ui.Button(ui.ButtonProps{Variant: ui.ButtonVariantOutline}, Disabled(), Text("Disabled")),
			),
		),

		subsection("Counter",
			Div(
				data.Signals(map[string]any{"count": 0}),
				Class("flex items-center gap-4"),
				ui.Button(ui.ButtonProps{Variant: ui.ButtonVariantOutline, Size: ui.ButtonSizeIcon},
					data.On("click", "$count--"),
					Raw(`<svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M5 12h14"/></svg>`),
				),
				Span(Class("text-2xl font-bold w-12 text-center"), data.Text("$count")),
				ui.Button(ui.ButtonProps{Variant: ui.ButtonVariantOutline, Size: ui.ButtonSizeIcon},
					data.On("click", "$count++"),
					Raw(`<svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M5 12h14"/><path d="M12 5v14"/></svg>`),
				),
			),
		),
	)
}

func cardSection() Node {
	return sectionWithSource("Card", "card.go",
		P(
			Class("text-muted-foreground mb-6"),
			Text("Displays a card with header, content, and footer."),
		),

		subsection("Default",
			ui.Card(ui.CardProps{}, Class("w-96"),
				ui.CardHeader(ui.CardHeaderProps{},
					ui.CardTitle(ui.CardTitleProps{}, Text("Card Title")),
					ui.CardDescription(ui.CardDescriptionProps{}, Text("Card description goes here.")),
				),
				ui.CardContent(ui.CardContentProps{},
					P(Text("Card content goes here. You can put any content inside.")),
				),
				ui.CardFooter(ui.CardFooterProps{},
					ui.Button(ui.ButtonProps{}, Text("Action")),
				),
			),
		),

		subsection("With Action",
			ui.Card(ui.CardProps{}, Class("w-96"),
				ui.CardHeader(ui.CardHeaderProps{},
					ui.CardTitle(ui.CardTitleProps{}, Text("Notifications")),
					ui.CardDescription(ui.CardDescriptionProps{}, Text("You have 3 unread messages.")),
					ui.CardAction(ui.CardActionProps{},
						ui.Button(ui.ButtonProps{Variant: ui.ButtonVariantOutline, Size: ui.ButtonSizeSm}, Text("Mark all read")),
					),
				),
				ui.CardContent(ui.CardContentProps{},
					P(Text("Your notifications will appear here.")),
				),
			),
		),
	)
}

func checkboxSection() Node {
	return sectionWithSource("Checkbox", "checkbox.go",
		P(
			Class("text-muted-foreground mb-6"),
			Text("A control that allows the user to toggle between checked and not checked."),
		),

		subsection("Default",
			Div(
				Class("flex items-center space-x-2"),
				ui.Checkbox(ui.CheckboxProps{}, ID("terms")),
				ui.Label(ui.LabelProps{}, For("terms"), Text("Accept terms and conditions")),
			),
		),

		subsection("Checked",
			Div(
				Class("flex items-center space-x-2"),
				ui.Checkbox(ui.CheckboxProps{}, ID("newsletter"), Checked()),
				ui.Label(ui.LabelProps{}, For("newsletter"), Text("Subscribe to newsletter")),
			),
		),

		subsection("Disabled",
			Div(
				Class("flex items-center space-x-2"),
				ui.Checkbox(ui.CheckboxProps{}, ID("disabled-checkbox"), Disabled()),
				ui.Label(ui.LabelProps{}, For("disabled-checkbox"), Text("Disabled")),
			),
		),

		subsection("Interactive",
			Div(
				data.Signals(map[string]any{"checked": false}),
				Class("flex flex-col gap-4"),
				Div(
					Class("flex items-center space-x-2"),
					ui.Checkbox(ui.CheckboxProps{},
						ID("interactive-checkbox"),
						data.Bind("checked"),
					),
					ui.Label(ui.LabelProps{}, For("interactive-checkbox"), Text("I agree to the terms")),
				),
				Span(
					Class("text-sm text-muted-foreground"),
					data.Text("$checked ? 'You have agreed!' : 'Please check the box'"),
				),
			),
		),
	)
}

func collapsibleSection() Node {
	return sectionWithSource("Collapsible", "collapsible.go",
		P(
			Class("text-muted-foreground mb-6"),
			Text("An interactive component which expands/collapses content."),
		),

		subsection("Default",
			Div(
				data.Signals(map[string]any{"open": false}),
				Class("w-[350px] space-y-2"),
				Div(
					Class("flex items-center justify-between space-x-4 px-4"),
					H4(Class("text-sm font-semibold"), Text("@peduarte starred 3 repositories")),
					ui.Button(ui.ButtonProps{Variant: ui.ButtonVariantGhost, Size: ui.ButtonSizeIcon},
						data.On("click", "$open = !$open"),
						Raw(`<svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><rect width="7" height="7" x="14" y="3" rx="1"/><path d="M10 21V8a1 1 0 0 0-1-1H4a1 1 0 0 0-1 1v12a1 1 0 0 0 1 1h12a1 1 0 0 0 1-1v-5a1 1 0 0 0-1-1H3"/></svg>`),
					),
				),
				Div(
					Class("rounded-md border px-4 py-2 font-mono text-sm shadow-sm"),
					Text("@radix-ui/primitives"),
				),
				Div(
					data.Show("$open"),
					Class("space-y-2"),
					Div(
						Class("rounded-md border px-4 py-2 font-mono text-sm shadow-sm"),
						Text("@radix-ui/colors"),
					),
					Div(
						Class("rounded-md border px-4 py-2 font-mono text-sm shadow-sm"),
						Text("@stitches/react"),
					),
				),
			),
		),
	)
}

func dialogSection() Node {
	return sectionWithSource("Dialog", "dialog.go",
		subsection("Default",
			ui.Dialog(ui.DialogProps{},
				data.Signals(map[string]any{"dialogOpen": false}),
				ui.DialogTrigger(ui.DialogTriggerProps{},
					ui.Button(ui.ButtonProps{Variant: ui.ButtonVariantOutline},
						data.On("click", "dialogOpen = true"),
						Text("Edit Profile"),
					),
				),
				Group([]Node{
					ui.DialogOverlay(ui.DialogOverlayProps{},
						data.Show("dialogOpen"),
						data.On("click", "dialogOpen = false"),
					),
					ui.DialogContent(ui.DialogContentProps{},
						data.Show("dialogOpen"),
						ui.DialogClose(ui.DialogCloseProps{},
							data.On("click", "dialogOpen = false"),
							Raw(`<svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M18 6 6 18"/><path d="m6 6 12 12"/></svg>`),
							Span(Class("sr-only"), Text("Close")),
						),
						ui.DialogHeader(ui.DialogHeaderProps{},
							ui.DialogTitle(ui.DialogTitleProps{}, Text("Edit profile")),
							ui.DialogDescription(ui.DialogDescriptionProps{}, Text("Make changes to your profile here. Click save when you're done.")),
						),
						Div(
							Class("grid gap-4 py-4"),
							Div(
								Class("grid grid-cols-4 items-center gap-4"),
								ui.Label(ui.LabelProps{}, For("name"), Class("text-right"), Text("Name")),
								ui.Input(ui.InputProps{}, ID("name"), Value("Pedro Duarte"), Class("col-span-3")),
							),
							Div(
								Class("grid grid-cols-4 items-center gap-4"),
								ui.Label(ui.LabelProps{}, For("username"), Class("text-right"), Text("Username")),
								ui.Input(ui.InputProps{}, ID("username"), Value("@peduarte"), Class("col-span-3")),
							),
						),
						ui.DialogFooter(ui.DialogFooterProps{},
							ui.Button(ui.ButtonProps{},
								data.On("click", "dialogOpen = false"),
								Text("Save changes"),
							),
						),
					),
				}),
			),
		),
	)
}

func dropdownMenuSection() Node {
	return sectionWithSource("Dropdown Menu", "dropdownmenu.go",
		subsection("Default",
			ui.DropdownMenu(ui.DropdownMenuProps{},
				data.Signals(map[string]any{"menuOpen": false}),
				ui.DropdownMenuTrigger(ui.DropdownMenuTriggerProps{},
					ui.Button(ui.ButtonProps{Variant: ui.ButtonVariantOutline},
						data.On("click", "menuOpen = !menuOpen"),
						Text("Open Menu"),
					),
				),
				ui.DropdownMenuContent(ui.DropdownMenuContentProps{},
					data.Show("menuOpen"),
					ui.DropdownMenuLabel(ui.DropdownMenuLabelProps{}, Text("My Account")),
					ui.DropdownMenuSeparator(ui.DropdownMenuSeparatorProps{}),
					ui.DropdownMenuGroup(ui.DropdownMenuGroupProps{},
						ui.DropdownMenuItem(ui.DropdownMenuItemProps{},
							Text("Profile"),
							ui.DropdownMenuShortcut(ui.DropdownMenuShortcutProps{}, Text("⇧⌘P")),
						),
						ui.DropdownMenuItem(ui.DropdownMenuItemProps{},
							Text("Billing"),
							ui.DropdownMenuShortcut(ui.DropdownMenuShortcutProps{}, Text("⌘B")),
						),
						ui.DropdownMenuItem(ui.DropdownMenuItemProps{},
							Text("Settings"),
							ui.DropdownMenuShortcut(ui.DropdownMenuShortcutProps{}, Text("⌘S")),
						),
					),
					ui.DropdownMenuSeparator(ui.DropdownMenuSeparatorProps{}),
					ui.DropdownMenuItem(ui.DropdownMenuItemProps{Variant: ui.DropdownMenuItemVariantDestructive},
						Text("Log out"),
						ui.DropdownMenuShortcut(ui.DropdownMenuShortcutProps{}, Text("⇧⌘Q")),
					),
				),
			),
		),
	)
}

func inputSection() Node {
	return sectionWithSource("Input", "input.go",
		P(
			Class("text-muted-foreground mb-6"),
			Text("Displays a form input field or a component that looks like an input field."),
		),

		subsection("Default",
			ui.Input(ui.InputProps{}, Class("max-w-sm"), Placeholder("Email")),
		),

		subsection("With Label",
			Div(
				Class("flex flex-col gap-2 max-w-sm"),
				ui.Label(ui.LabelProps{}, For("email-input"), Text("Email")),
				ui.Input(ui.InputProps{}, Type("email"), ID("email-input"), Placeholder("Enter your email")),
			),
		),

		subsection("Disabled",
			ui.Input(ui.InputProps{}, Class("max-w-sm"), Disabled(), Placeholder("Disabled")),
		),

		subsection("File Input",
			Div(
				Class("flex flex-col gap-2 max-w-sm"),
				ui.Label(ui.LabelProps{}, For("file-input"), Text("Picture")),
				ui.Input(ui.InputProps{}, Type("file"), ID("file-input")),
			),
		),

		subsection("Live Binding",
			Div(
				data.Signals(map[string]any{"name": ""}),
				Class("flex flex-col gap-4 max-w-sm"),
				ui.Input(ui.InputProps{},
					Placeholder("Enter your name"),
					data.Bind("name"),
				),
				P(
					Class("text-sm text-muted-foreground"),
					Text("Hello, "),
					Span(data.Text("$name || 'stranger'")),
					Text("!"),
				),
			),
		),
	)
}

func labelSection() Node {
	return sectionWithSource("Label", "label.go",
		P(
			Class("text-muted-foreground mb-6"),
			Text("Renders an accessible label associated with controls."),
		),

		subsection("Default",
			Div(
				Class("flex flex-col gap-2"),
				ui.Label(ui.LabelProps{}, For("email"), Text("Email")),
				ui.Input(ui.InputProps{}, Type("email"), ID("email"), Placeholder("Enter your email"), Class("max-w-sm")),
			),
		),
	)
}

func paginationSection() Node {
	return sectionWithSource("Pagination", "pagination.go",
		P(
			Class("text-muted-foreground mb-6"),
			Text("Pagination with page navigation, next and previous links."),
		),

		subsection("Default",
			ui.Pagination(ui.PaginationProps{},
				ui.PaginationContent(ui.PaginationContentProps{},
					ui.PaginationItem(ui.PaginationItemProps{},
						ui.PaginationPrevious(ui.PaginationPreviousProps{}, Href("#")),
					),
					ui.PaginationItem(ui.PaginationItemProps{},
						ui.PaginationLink(ui.PaginationLinkProps{}, Href("#"), Text("1")),
					),
					ui.PaginationItem(ui.PaginationItemProps{},
						ui.PaginationLink(ui.PaginationLinkProps{IsActive: true}, Href("#"), Text("2")),
					),
					ui.PaginationItem(ui.PaginationItemProps{},
						ui.PaginationLink(ui.PaginationLinkProps{}, Href("#"), Text("3")),
					),
					ui.PaginationItem(ui.PaginationItemProps{},
						ui.PaginationEllipsis(ui.PaginationEllipsisProps{}),
					),
					ui.PaginationItem(ui.PaginationItemProps{},
						ui.PaginationNext(ui.PaginationNextProps{}, Href("#")),
					),
				),
			),
		),
	)
}

func popoverSection() Node {
	return sectionWithSource("Popover", "popover.go",
		subsection("Default",
			ui.Popover(ui.PopoverProps{},
				data.Signals(map[string]any{"popoverOpen": false}),
				ui.PopoverTrigger(ui.PopoverTriggerProps{},
					ui.Button(ui.ButtonProps{Variant: ui.ButtonVariantOutline},
						data.On("click", "popoverOpen = !popoverOpen"),
						Text("Open popover"),
					),
				),
				ui.PopoverContent(ui.PopoverContentProps{},
					data.Show("popoverOpen"),
					Div(
						Class("grid gap-4"),
						Div(
							Class("space-y-2"),
							H4(Class("font-medium leading-none"), Text("Dimensions")),
							P(Class("text-sm text-muted-foreground"), Text("Set the dimensions for the layer.")),
						),
						Div(
							Class("grid gap-2"),
							Div(
								Class("grid grid-cols-3 items-center gap-4"),
								ui.Label(ui.LabelProps{}, For("width"), Text("Width")),
								ui.Input(ui.InputProps{}, ID("width"), Value("100%"), Class("col-span-2 h-8")),
							),
							Div(
								Class("grid grid-cols-3 items-center gap-4"),
								ui.Label(ui.LabelProps{}, For("height"), Text("Height")),
								ui.Input(ui.InputProps{}, ID("height"), Value("25px"), Class("col-span-2 h-8")),
							),
						),
					),
				),
			),
		),
	)
}

func progressSection() Node {
	return sectionWithSource("Progress", "progress.go",
		P(
			Class("text-muted-foreground mb-6"),
			Text("Displays an indicator showing the completion progress of a task."),
		),

		subsection("Default",
			Div(
				Class("w-60"),
				ui.Progress(ui.ProgressProps{Value: 33}),
			),
		),

		subsection("Different Values",
			Div(
				Class("flex flex-col gap-4 w-60"),
				ui.Progress(ui.ProgressProps{Value: 0}),
				ui.Progress(ui.ProgressProps{Value: 25}),
				ui.Progress(ui.ProgressProps{Value: 50}),
				ui.Progress(ui.ProgressProps{Value: 75}),
				ui.Progress(ui.ProgressProps{Value: 100}),
			),
		),

		subsection("Interactive",
			Div(
				data.Signals(map[string]any{"progress": 33}),
				Class("flex flex-col gap-4 w-60"),
				Div(
					Class("bg-primary/20 relative h-2 w-full overflow-hidden rounded-full"),
					Div(
						Class("bg-primary h-full transition-all"),
						data.Attr("style", "'width:' + $progress + '%'"),
					),
				),
				Div(
					Class("flex gap-2"),
					ui.Button(ui.ButtonProps{Size: ui.ButtonSizeSm},
						data.On("click", "$progress = Math.max(0, $progress - 10)"),
						Text("-10"),
					),
					ui.Button(ui.ButtonProps{Size: ui.ButtonSizeSm},
						data.On("click", "$progress = Math.min(100, $progress + 10)"),
						Text("+10"),
					),
				),
			),
		),
	)
}

func radioGroupSection() Node {
	return sectionWithSource("Radio Group", "radiogroup.go",
		P(
			Class("text-muted-foreground mb-6"),
			Text("A set of checkable buttons where only one can be checked at a time."),
		),

		subsection("Default",
			ui.RadioGroup(ui.RadioGroupProps{},
				Div(
					Class("flex items-center space-x-2"),
					ui.RadioGroupItem(ui.RadioGroupItemProps{}, ID("option-one"), Name("options"), Value("1"), Checked()),
					ui.Label(ui.LabelProps{}, For("option-one"), Text("Option One")),
				),
				Div(
					Class("flex items-center space-x-2"),
					ui.RadioGroupItem(ui.RadioGroupItemProps{}, ID("option-two"), Name("options"), Value("2")),
					ui.Label(ui.LabelProps{}, For("option-two"), Text("Option Two")),
				),
				Div(
					Class("flex items-center space-x-2"),
					ui.RadioGroupItem(ui.RadioGroupItemProps{}, ID("option-three"), Name("options"), Value("3")),
					ui.Label(ui.LabelProps{}, For("option-three"), Text("Option Three")),
				),
			),
		),

		subsection("Disabled",
			ui.RadioGroup(ui.RadioGroupProps{},
				Div(
					Class("flex items-center space-x-2"),
					ui.RadioGroupItem(ui.RadioGroupItemProps{}, ID("disabled-one"), Name("disabled"), Value("1"), Disabled()),
					ui.Label(ui.LabelProps{}, For("disabled-one"), Text("Disabled option")),
				),
			),
		),
	)
}

func selectSection() Node {
	return sectionWithSource("Select", "selectcomponent.go",
		subsection("Default",
			Div(
				Class("w-[180px]"),
				ui.Select(ui.SelectProps{},
					data.Signals(map[string]any{"selectOpen": false, "selectedFruit": ""}),
					ui.SelectTrigger(ui.SelectTriggerProps{},
						data.On("click", "selectOpen = !selectOpen"),
						ui.SelectValue(ui.SelectValueProps{Placeholder: "Select a fruit"},
							data.Text("$selectedFruit || 'Select a fruit'"),
						),
						Raw(`<svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="opacity-50"><path d="m6 9 6 6 6-6"/></svg>`),
					),
					ui.SelectContent(ui.SelectContentProps{},
						data.Show("selectOpen"),
						ui.SelectGroup(ui.SelectGroupProps{},
							ui.SelectLabel(ui.SelectLabelProps{}, Text("Fruits")),
							ui.SelectItem(ui.SelectItemProps{},
								data.On("click", "selectedFruit = 'Apple'; selectOpen = false"),
								Text("Apple"),
							),
							ui.SelectItem(ui.SelectItemProps{},
								data.On("click", "selectedFruit = 'Banana'; selectOpen = false"),
								Text("Banana"),
							),
							ui.SelectItem(ui.SelectItemProps{},
								data.On("click", "selectedFruit = 'Orange'; selectOpen = false"),
								Text("Orange"),
							),
							ui.SelectItem(ui.SelectItemProps{},
								data.On("click", "selectedFruit = 'Mango'; selectOpen = false"),
								Text("Mango"),
							),
						),
					),
				),
			),
		),
	)
}

func separatorSection() Node {
	return sectionWithSource("Separator", "separator.go",
		P(
			Class("text-muted-foreground mb-6"),
			Text("Visually or semantically separates content."),
		),

		subsection("Horizontal",
			Div(
				Class("space-y-4"),
				Div(
					Class("space-y-1"),
					H4(Class("text-sm font-medium leading-none"), Text("Radix Primitives")),
					P(Class("text-sm text-muted-foreground"), Text("An open-source UI component library.")),
				),
				ui.Separator(ui.SeparatorProps{}),
				Div(
					Class("flex h-5 items-center space-x-4 text-sm"),
					Div(Text("Blog")),
					ui.Separator(ui.SeparatorProps{Orientation: ui.SeparatorOrientationVertical}),
					Div(Text("Docs")),
					ui.Separator(ui.SeparatorProps{Orientation: ui.SeparatorOrientationVertical}),
					Div(Text("Source")),
				),
			),
		),
	)
}

func skeletonSection() Node {
	return sectionWithSource("Skeleton", "skeleton.go",
		P(
			Class("text-muted-foreground mb-6"),
			Text("Use to show a placeholder while content is loading."),
		),

		subsection("Default",
			Div(
				Class("flex items-center space-x-4"),
				ui.Skeleton(ui.SkeletonProps{}, Class("h-12 w-12 rounded-full")),
				Div(
					Class("space-y-2"),
					ui.Skeleton(ui.SkeletonProps{}, Class("h-4 w-[250px]")),
					ui.Skeleton(ui.SkeletonProps{}, Class("h-4 w-[200px]")),
				),
			),
		),

		subsection("Card",
			Div(
				Class("flex flex-col space-y-3"),
				ui.Skeleton(ui.SkeletonProps{}, Class("h-[125px] w-[250px] rounded-xl")),
				Div(
					Class("space-y-2"),
					ui.Skeleton(ui.SkeletonProps{}, Class("h-4 w-[250px]")),
					ui.Skeleton(ui.SkeletonProps{}, Class("h-4 w-[200px]")),
				),
			),
		),

		subsection("Loading State",
			Div(
				data.Signals(map[string]any{"loading": true}),
				Class("flex flex-col gap-4"),
				ui.Button(ui.ButtonProps{Size: ui.ButtonSizeSm},
					data.On("click", "$loading = !$loading"),
					data.Text("$loading ? 'Show Content' : 'Show Skeleton'"),
				),
				Div(
					data.Show("$loading"),
					Div(
						Class("flex items-center space-x-4"),
						ui.Skeleton(ui.SkeletonProps{}, Class("h-12 w-12 rounded-full")),
						Div(
							Class("space-y-2"),
							ui.Skeleton(ui.SkeletonProps{}, Class("h-4 w-[250px]")),
							ui.Skeleton(ui.SkeletonProps{}, Class("h-4 w-[200px]")),
						),
					),
				),
				Div(
					data.Show("!$loading"),
					Div(
						Class("flex items-center space-x-4"),
						ui.Avatar(ui.AvatarProps{},
							ui.AvatarFallback(ui.AvatarFallbackProps{}, Text("CN")),
						),
						Div(
							Class("space-y-1"),
							P(Class("text-sm font-medium leading-none"), Text("Sofia Davis")),
							P(Class("text-sm text-muted-foreground"), Text("sofia@example.com")),
						),
					),
				),
			),
		),
	)
}

func sliderSection() Node {
	return sectionWithSource("Slider", "slider.go",
		P(
			Class("text-muted-foreground mb-6"),
			Text("An input where the user selects a value from within a given range."),
		),

		subsection("Default",
			Div(
				Class("w-60"),
				ui.Slider(ui.SliderProps{}, Min("0"), Max("100"), Value("50")),
			),
		),

		subsection("Different Values",
			Div(
				Class("flex flex-col gap-6 w-60"),
				ui.Slider(ui.SliderProps{}, Min("0"), Max("100"), Value("0")),
				ui.Slider(ui.SliderProps{}, Min("0"), Max("100"), Value("25")),
				ui.Slider(ui.SliderProps{}, Min("0"), Max("100"), Value("50")),
				ui.Slider(ui.SliderProps{}, Min("0"), Max("100"), Value("75")),
				ui.Slider(ui.SliderProps{}, Min("0"), Max("100"), Value("100")),
			),
		),

		subsection("Disabled",
			Div(
				Class("w-60"),
				ui.Slider(ui.SliderProps{}, Min("0"), Max("100"), Value("50"), Disabled()),
			),
		),

		subsection("Interactive",
			Div(
				data.Signals(map[string]any{"sliderValue": 50}),
				Class("flex flex-col gap-4 w-60"),
				ui.Slider(ui.SliderProps{},
					Min("0"), Max("100"),
					data.Bind("sliderValue"),
				),
				Span(
					Class("text-sm text-muted-foreground"),
					Text("Value: "),
					Span(data.Text("$sliderValue")),
				),
			),
		),
	)
}

func switchSection() Node {
	return sectionWithSource("Switch", "switch.go",
		P(
			Class("text-muted-foreground mb-6"),
			Text("A control that allows the user to toggle between on and off."),
		),

		subsection("Default",
			Div(
				Class("flex items-center space-x-2"),
				ui.Switch(ui.SwitchProps{}, ID("airplane-mode")),
				ui.Label(ui.LabelProps{}, For("airplane-mode"), Text("Airplane Mode")),
			),
		),

		subsection("Checked",
			Div(
				Class("flex items-center space-x-2"),
				ui.Switch(ui.SwitchProps{}, ID("wifi"), Checked()),
				ui.Label(ui.LabelProps{}, For("wifi"), Text("Wi-Fi")),
			),
		),

		subsection("Disabled",
			Div(
				Class("flex flex-col gap-4"),
				Div(
					Class("flex items-center space-x-2"),
					ui.Switch(ui.SwitchProps{}, ID("disabled-off"), Disabled()),
					ui.Label(ui.LabelProps{}, For("disabled-off"), Text("Disabled (off)")),
				),
				Div(
					Class("flex items-center space-x-2"),
					ui.Switch(ui.SwitchProps{}, ID("disabled-on"), Disabled(), Checked()),
					ui.Label(ui.LabelProps{}, For("disabled-on"), Text("Disabled (on)")),
				),
			),
		),

		subsection("Interactive",
			Div(
				data.Signals(map[string]any{"enabled": false}),
				Class("flex flex-col gap-4"),
				Div(
					Class("flex items-center space-x-2"),
					ui.Switch(ui.SwitchProps{},
						ID("interactive-switch"),
						data.Bind("enabled"),
					),
					ui.Label(ui.LabelProps{}, For("interactive-switch"), Text("Enable notifications")),
				),
				Span(
					Class("text-sm text-muted-foreground"),
					data.Text("$enabled ? 'Notifications enabled' : 'Notifications disabled'"),
				),
			),
		),
	)
}

func tableSection() Node {
	return sectionWithSource("Table", "table.go",
		P(
			Class("text-muted-foreground mb-6"),
			Text("A responsive table component."),
		),

		subsection("Default",
			ui.Table(ui.TableProps{},
				ui.TableCaption(ui.TableCaptionProps{}, Text("A list of your recent invoices.")),
				ui.TableHeader(ui.TableHeaderProps{},
					ui.TableRow(ui.TableRowProps{},
						ui.TableHead(ui.TableHeadProps{}, Class("w-[100px]"), Text("Invoice")),
						ui.TableHead(ui.TableHeadProps{}, Text("Status")),
						ui.TableHead(ui.TableHeadProps{}, Text("Method")),
						ui.TableHead(ui.TableHeadProps{}, Class("text-right"), Text("Amount")),
					),
				),
				ui.TableBody(ui.TableBodyProps{},
					ui.TableRow(ui.TableRowProps{},
						ui.TableCell(ui.TableCellProps{}, Class("font-medium"), Text("INV001")),
						ui.TableCell(ui.TableCellProps{}, Text("Paid")),
						ui.TableCell(ui.TableCellProps{}, Text("Credit Card")),
						ui.TableCell(ui.TableCellProps{}, Class("text-right"), Text("$250.00")),
					),
					ui.TableRow(ui.TableRowProps{},
						ui.TableCell(ui.TableCellProps{}, Class("font-medium"), Text("INV002")),
						ui.TableCell(ui.TableCellProps{}, Text("Pending")),
						ui.TableCell(ui.TableCellProps{}, Text("PayPal")),
						ui.TableCell(ui.TableCellProps{}, Class("text-right"), Text("$150.00")),
					),
					ui.TableRow(ui.TableRowProps{},
						ui.TableCell(ui.TableCellProps{}, Class("font-medium"), Text("INV003")),
						ui.TableCell(ui.TableCellProps{}, Text("Unpaid")),
						ui.TableCell(ui.TableCellProps{}, Text("Bank Transfer")),
						ui.TableCell(ui.TableCellProps{}, Class("text-right"), Text("$350.00")),
					),
				),
				ui.TableFooter(ui.TableFooterProps{},
					ui.TableRow(ui.TableRowProps{},
						ui.TableCell(ui.TableCellProps{}, ColSpan("3"), Text("Total")),
						ui.TableCell(ui.TableCellProps{}, Class("text-right"), Text("$750.00")),
					),
				),
			),
		),
	)
}

func tabsSection() Node {
	return sectionWithSource("Tabs", "tabs.go",
		P(
			Class("text-muted-foreground mb-6"),
			Text("A set of layered sections of content—known as tab panels—that are displayed one at a time."),
		),

		subsection("Default",
			Div(
				data.Signals(map[string]any{"activeTab": "account"}),
				ui.Tabs(ui.TabsProps{}, Class("w-[400px]"),
					ui.TabsList(ui.TabsListProps{},
						Div(
							data.Attr("data-state", "$activeTab === 'account' ? 'active' : 'inactive'"),
							data.On("click", "$activeTab = 'account'"),
							Class("data-[state=active]:bg-background dark:data-[state=active]:text-foreground text-foreground dark:text-muted-foreground inline-flex h-[calc(100%-1px)] flex-1 items-center justify-center gap-1.5 rounded-md border border-transparent px-2 py-1 text-sm font-medium whitespace-nowrap data-[state=active]:shadow-sm cursor-pointer"),
							Text("Account"),
						),
						Div(
							data.Attr("data-state", "$activeTab === 'password' ? 'active' : 'inactive'"),
							data.On("click", "$activeTab = 'password'"),
							Class("data-[state=active]:bg-background dark:data-[state=active]:text-foreground text-foreground dark:text-muted-foreground inline-flex h-[calc(100%-1px)] flex-1 items-center justify-center gap-1.5 rounded-md border border-transparent px-2 py-1 text-sm font-medium whitespace-nowrap data-[state=active]:shadow-sm cursor-pointer"),
							Text("Password"),
						),
					),
					Div(
						data.Show("$activeTab === 'account'"),
						ui.TabsContent(ui.TabsContentProps{},
							ui.Card(ui.CardProps{},
								ui.CardHeader(ui.CardHeaderProps{},
									ui.CardTitle(ui.CardTitleProps{}, Text("Account")),
									ui.CardDescription(ui.CardDescriptionProps{}, Text("Make changes to your account here.")),
								),
								ui.CardContent(ui.CardContentProps{}, Class("space-y-2"),
									Div(
										Class("space-y-1"),
										ui.Label(ui.LabelProps{}, For("name"), Text("Name")),
										ui.Input(ui.InputProps{}, ID("name"), Value("Pedro Duarte")),
									),
									Div(
										Class("space-y-1"),
										ui.Label(ui.LabelProps{}, For("username"), Text("Username")),
										ui.Input(ui.InputProps{}, ID("username"), Value("@peduarte")),
									),
								),
								ui.CardFooter(ui.CardFooterProps{},
									ui.Button(ui.ButtonProps{}, Text("Save changes")),
								),
							),
						),
					),
					Div(
						data.Show("$activeTab === 'password'"),
						ui.TabsContent(ui.TabsContentProps{},
							ui.Card(ui.CardProps{},
								ui.CardHeader(ui.CardHeaderProps{},
									ui.CardTitle(ui.CardTitleProps{}, Text("Password")),
									ui.CardDescription(ui.CardDescriptionProps{}, Text("Change your password here.")),
								),
								ui.CardContent(ui.CardContentProps{}, Class("space-y-2"),
									Div(
										Class("space-y-1"),
										ui.Label(ui.LabelProps{}, For("current"), Text("Current password")),
										ui.Input(ui.InputProps{}, ID("current"), Type("password")),
									),
									Div(
										Class("space-y-1"),
										ui.Label(ui.LabelProps{}, For("new"), Text("New password")),
										ui.Input(ui.InputProps{}, ID("new"), Type("password")),
									),
								),
								ui.CardFooter(ui.CardFooterProps{},
									ui.Button(ui.ButtonProps{}, Text("Save password")),
								),
							),
						),
					),
				),
			),
		),
	)
}

func textareaSection() Node {
	return sectionWithSource("Textarea", "textarea.go",
		P(
			Class("text-muted-foreground mb-6"),
			Text("Displays a form textarea or a component that looks like a textarea."),
		),

		subsection("Default",
			ui.Textarea(ui.TextareaProps{}, Class("max-w-sm"), Placeholder("Type your message here.")),
		),

		subsection("With Label",
			Div(
				Class("flex flex-col gap-2 max-w-sm"),
				ui.Label(ui.LabelProps{}, For("message"), Text("Your message")),
				ui.Textarea(ui.TextareaProps{}, ID("message"), Placeholder("Type your message here.")),
			),
		),

		subsection("Disabled",
			ui.Textarea(ui.TextareaProps{}, Class("max-w-sm"), Disabled(), Placeholder("Disabled")),
		),
	)
}

func toggleSection() Node {
	return sectionWithSource("Toggle", "toggle.go",
		P(
			Class("text-muted-foreground mb-6"),
			Text("A two-state button that can be either on or off."),
		),

		subsection("Default",
			ui.Toggle(ui.ToggleProps{},
				Raw(`<svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M6 4v6a6 6 0 0 0 12 0V4"/><line x1="4" x2="20" y1="20" y2="20"/></svg>`),
			),
		),

		subsection("Outline",
			ui.Toggle(ui.ToggleProps{Variant: ui.ToggleVariantOutline},
				Raw(`<svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M6 4v6a6 6 0 0 0 12 0V4"/><line x1="4" x2="20" y1="20" y2="20"/></svg>`),
			),
		),

		subsection("With Text",
			ui.Toggle(ui.ToggleProps{},
				Raw(`<svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M3 6h18"/><path d="M7 12h10"/><path d="M10 18h4"/></svg>`),
				Text("Toggle Italic"),
			),
		),

		subsection("Sizes",
			Div(
				Class("flex items-center gap-4"),
				ui.Toggle(ui.ToggleProps{Size: ui.ToggleSizeSm},
					Raw(`<svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M6 4v6a6 6 0 0 0 12 0V4"/><line x1="4" x2="20" y1="20" y2="20"/></svg>`),
				),
				ui.Toggle(ui.ToggleProps{Size: ui.ToggleSizeDefault},
					Raw(`<svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M6 4v6a6 6 0 0 0 12 0V4"/><line x1="4" x2="20" y1="20" y2="20"/></svg>`),
				),
				ui.Toggle(ui.ToggleProps{Size: ui.ToggleSizeLg},
					Raw(`<svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M6 4v6a6 6 0 0 0 12 0V4"/><line x1="4" x2="20" y1="20" y2="20"/></svg>`),
				),
			),
		),

		subsection("Interactive",
			Div(
				data.Signals(map[string]any{"bold": false}),
				Class("flex flex-col gap-4"),
				Div(
					data.Attr("data-state", "$bold ? 'on' : 'off'"),
					data.On("click", "$bold = !$bold"),
					Class("inline-flex items-center justify-center gap-2 rounded-md text-sm font-medium hover:bg-muted hover:text-muted-foreground data-[state=on]:bg-accent data-[state=on]:text-accent-foreground outline-none cursor-pointer bg-transparent h-9 px-2 min-w-9"),
					Raw(`<svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M6 4h8a4 4 0 0 1 4 4 4 4 0 0 1-4 4H6z"/><path d="M6 12h9a4 4 0 0 1 4 4 4 4 0 0 1-4 4H6z"/></svg>`),
				),
				Span(
					Class("text-sm text-muted-foreground"),
					data.Text("$bold ? 'Bold is on' : 'Bold is off'"),
				),
			),
		),
	)
}

func toggleGroupSection() Node {
	return sectionWithSource("Toggle Group", "togglegroup.go",
		P(
			Class("text-muted-foreground mb-6"),
			Text("A set of two-state buttons that can be toggled on or off."),
		),

		subsection("Default",
			ui.ToggleGroup(ui.ToggleGroupProps{},
				ui.ToggleGroupItem(ui.ToggleGroupItemProps{Variant: ui.ToggleVariantOutline},
					Raw(`<svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M6 4h8a4 4 0 0 1 4 4 4 4 0 0 1-4 4H6z"/><path d="M6 12h9a4 4 0 0 1 4 4 4 4 0 0 1-4 4H6z"/></svg>`),
				),
				ui.ToggleGroupItem(ui.ToggleGroupItemProps{Variant: ui.ToggleVariantOutline},
					Raw(`<svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M3 6h18"/><path d="M7 12h10"/><path d="M10 18h4"/></svg>`),
				),
				ui.ToggleGroupItem(ui.ToggleGroupItemProps{Variant: ui.ToggleVariantOutline},
					Raw(`<svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M6 4v6a6 6 0 0 0 12 0V4"/><line x1="4" x2="20" y1="20" y2="20"/></svg>`),
				),
			),
		),

		subsection("With Selection",
			ui.ToggleGroup(ui.ToggleGroupProps{},
				ui.ToggleGroupItem(ui.ToggleGroupItemProps{Variant: ui.ToggleVariantOutline, Pressed: true},
					Raw(`<svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M4 5h16"/><path d="M4 12h16"/><path d="M4 19h16"/></svg>`),
				),
				ui.ToggleGroupItem(ui.ToggleGroupItemProps{Variant: ui.ToggleVariantOutline},
					Raw(`<svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M4 5h16"/><path d="M4 12h10"/><path d="M4 19h6"/></svg>`),
				),
				ui.ToggleGroupItem(ui.ToggleGroupItemProps{Variant: ui.ToggleVariantOutline},
					Raw(`<svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M4 5h16"/><path d="M10 12h10"/><path d="M14 19h6"/></svg>`),
				),
			),
		),

		subsection("Sizes",
			Div(
				Class("flex flex-col gap-4"),
				ui.ToggleGroup(ui.ToggleGroupProps{},
					ui.ToggleGroupItem(ui.ToggleGroupItemProps{Variant: ui.ToggleVariantOutline, Size: ui.ToggleSizeSm}, Text("S")),
					ui.ToggleGroupItem(ui.ToggleGroupItemProps{Variant: ui.ToggleVariantOutline, Size: ui.ToggleSizeSm}, Text("M")),
					ui.ToggleGroupItem(ui.ToggleGroupItemProps{Variant: ui.ToggleVariantOutline, Size: ui.ToggleSizeSm}, Text("L")),
				),
				ui.ToggleGroup(ui.ToggleGroupProps{},
					ui.ToggleGroupItem(ui.ToggleGroupItemProps{Variant: ui.ToggleVariantOutline, Size: ui.ToggleSizeDefault}, Text("S")),
					ui.ToggleGroupItem(ui.ToggleGroupItemProps{Variant: ui.ToggleVariantOutline, Size: ui.ToggleSizeDefault}, Text("M")),
					ui.ToggleGroupItem(ui.ToggleGroupItemProps{Variant: ui.ToggleVariantOutline, Size: ui.ToggleSizeDefault}, Text("L")),
				),
				ui.ToggleGroup(ui.ToggleGroupProps{},
					ui.ToggleGroupItem(ui.ToggleGroupItemProps{Variant: ui.ToggleVariantOutline, Size: ui.ToggleSizeLg}, Text("S")),
					ui.ToggleGroupItem(ui.ToggleGroupItemProps{Variant: ui.ToggleVariantOutline, Size: ui.ToggleSizeLg}, Text("M")),
					ui.ToggleGroupItem(ui.ToggleGroupItemProps{Variant: ui.ToggleVariantOutline, Size: ui.ToggleSizeLg}, Text("L")),
				),
			),
		),
	)
}

func tooltipSection() Node {
	return sectionWithSource("Tooltip", "tooltip.go",
		subsection("Default",
			ui.Tooltip(ui.TooltipProps{},
				ui.TooltipTrigger(ui.TooltipTriggerProps{},
					ui.Button(ui.ButtonProps{Variant: ui.ButtonVariantOutline}, Text("Hover me")),
				),
				ui.TooltipContent(ui.TooltipContentProps{},
					Text("This is a tooltip"),
				),
			),
		),
	)
}

func sectionWithSource(title, sourceFile string, children ...Node) Node {
	headerContent := []Node{
		Span(Text(title)),
	}
	if sourceFile != "" {
		headerContent = append(headerContent,
			A(
				Href("https://github.com/maragudk/gomponents-shadcn-ui/blob/main/"+sourceFile),
				Class("text-sm font-normal text-muted-foreground hover:text-foreground ml-2"),
				Text("View Source"),
			),
		)
	}
	return Section(
		Class("mb-24"),
		H2(
			Class("text-2xl font-semibold mb-4 pb-2 border-b flex items-baseline"),
			Group(headerContent),
		),
		Group(children),
	)
}

func subsection(title string, children ...Node) Node {
	return Div(
		Class("mb-6"),
		H3(
			Class("text-lg font-medium mb-3"),
			Text(title),
		),
		Group(children),
	)
}

