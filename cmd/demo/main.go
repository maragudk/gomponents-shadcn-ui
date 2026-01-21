// Command demo generates a static demo site showcasing all components.
package main

import (
	"os"

	g "maragu.dev/gomponents"
	c "maragu.dev/gomponents/components"
	h "maragu.dev/gomponents/html"

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

func page() g.Node {
	return c.HTML5(c.HTML5Props{
		Title:       "gomponents-shadcn-ui Demo",
		Description: "A demo of all gomponents-shadcn-ui components",
		Language:    "en",
		Head: []g.Node{
			h.Link(h.Rel("stylesheet"), h.Href("styles.css")),
		},
		Body: []g.Node{
			h.Class("bg-background text-foreground min-h-screen"),
			h.Div(
				h.Class("container mx-auto py-12 px-4 max-w-4xl"),
				header(),
				alertSection(),
				avatarSection(),
				badgeSection(),
				buttonSection(),
				cardSection(),
				labelSection(),
				separatorSection(),
				skeletonSection(),
			),
		},
	})
}

func header() g.Node {
	return h.Header(
		h.Class("mb-12"),
		h.H1(
			h.Class("text-4xl font-bold mb-4"),
			g.Text("gomponents-shadcn-ui"),
		),
		h.P(
			h.Class("text-muted-foreground text-lg"),
			g.Text("shadcn/ui components for Go, built with "),
			h.A(
				h.Href("https://www.gomponents.com"),
				h.Class("underline hover:text-foreground"),
				g.Text("gomponents"),
			),
			g.Text("."),
		),
	)
}

func alertSection() g.Node {
	return section("Alert",
		h.P(
			h.Class("text-muted-foreground mb-6"),
			g.Text("Displays a callout for user attention."),
		),

		subsection("Default",
			ui.Alert(ui.AlertProps{},
				g.Raw(`<svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><circle cx="12" cy="12" r="10"/><path d="M12 16v-4"/><path d="M12 8h.01"/></svg>`),
				ui.AlertTitle(ui.AlertTitleProps{}, g.Text("Heads up!")),
				ui.AlertDescription(ui.AlertDescriptionProps{}, g.Text("You can add components to your app using the cli.")),
			),
		),

		subsection("Destructive",
			ui.Alert(ui.AlertProps{Variant: ui.AlertVariantDestructive},
				g.Raw(`<svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><circle cx="12" cy="12" r="10"/><line x1="12" x2="12" y1="8" y2="12"/><line x1="12" x2="12.01" y1="16" y2="16"/></svg>`),
				ui.AlertTitle(ui.AlertTitleProps{}, g.Text("Error")),
				ui.AlertDescription(ui.AlertDescriptionProps{}, g.Text("Your session has expired. Please log in again.")),
			),
		),
	)
}

func avatarSection() g.Node {
	return section("Avatar",
		h.P(
			h.Class("text-muted-foreground mb-6"),
			g.Text("An image element with a fallback for representing the user."),
		),

		subsection("With Image",
			h.Div(
				h.Class("flex flex-wrap gap-4"),
				ui.Avatar(ui.AvatarProps{},
					ui.AvatarImage(ui.AvatarImageProps{}, h.Src("https://github.com/shadcn.png"), h.Alt("@shadcn")),
				),
			),
		),

		subsection("With Fallback",
			h.Div(
				h.Class("flex flex-wrap gap-4"),
				ui.Avatar(ui.AvatarProps{},
					ui.AvatarFallback(ui.AvatarFallbackProps{}, g.Text("CN")),
				),
			),
		),

		subsection("Sizes",
			h.Div(
				h.Class("flex flex-wrap items-center gap-4"),
				ui.Avatar(ui.AvatarProps{}, h.Class("size-6"),
					ui.AvatarFallback(ui.AvatarFallbackProps{}, h.Class("text-xs"), g.Text("XS")),
				),
				ui.Avatar(ui.AvatarProps{},
					ui.AvatarFallback(ui.AvatarFallbackProps{}, g.Text("SM")),
				),
				ui.Avatar(ui.AvatarProps{}, h.Class("size-12"),
					ui.AvatarFallback(ui.AvatarFallbackProps{}, g.Text("MD")),
				),
				ui.Avatar(ui.AvatarProps{}, h.Class("size-16"),
					ui.AvatarFallback(ui.AvatarFallbackProps{}, h.Class("text-lg"), g.Text("LG")),
				),
			),
		),
	)
}

func badgeSection() g.Node {
	return section("Badge",
		h.P(
			h.Class("text-muted-foreground mb-6"),
			g.Text("Displays a badge or a component that looks like a badge."),
		),

		subsection("Variants",
			h.Div(
				h.Class("flex flex-wrap gap-4"),
				ui.Badge(ui.BadgeProps{Variant: ui.BadgeVariantDefault}, g.Text("Default")),
				ui.Badge(ui.BadgeProps{Variant: ui.BadgeVariantSecondary}, g.Text("Secondary")),
				ui.Badge(ui.BadgeProps{Variant: ui.BadgeVariantDestructive}, g.Text("Destructive")),
				ui.Badge(ui.BadgeProps{Variant: ui.BadgeVariantOutline}, g.Text("Outline")),
			),
		),
	)
}

func buttonSection() g.Node {
	return section("Button",
		h.P(
			h.Class("text-muted-foreground mb-6"),
			g.Text("Displays a button or a component that looks like a button."),
		),

		subsection("Variants",
			h.Div(
				h.Class("flex flex-wrap gap-4"),
				ui.Button(ui.ButtonProps{Variant: ui.ButtonVariantDefault}, g.Text("Default")),
				ui.Button(ui.ButtonProps{Variant: ui.ButtonVariantSecondary}, g.Text("Secondary")),
				ui.Button(ui.ButtonProps{Variant: ui.ButtonVariantDestructive}, g.Text("Destructive")),
				ui.Button(ui.ButtonProps{Variant: ui.ButtonVariantOutline}, g.Text("Outline")),
				ui.Button(ui.ButtonProps{Variant: ui.ButtonVariantGhost}, g.Text("Ghost")),
				ui.Button(ui.ButtonProps{Variant: ui.ButtonVariantLink}, g.Text("Link")),
			),
		),

		subsection("Sizes",
			h.Div(
				h.Class("flex flex-wrap items-center gap-4"),
				ui.Button(ui.ButtonProps{Size: ui.ButtonSizeSm}, g.Text("Small")),
				ui.Button(ui.ButtonProps{Size: ui.ButtonSizeDefault}, g.Text("Default")),
				ui.Button(ui.ButtonProps{Size: ui.ButtonSizeLg}, g.Text("Large")),
				ui.Button(ui.ButtonProps{Size: ui.ButtonSizeIcon},
					g.Raw(`<svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M5 12h14"/><path d="M12 5v14"/></svg>`),
				),
			),
		),

		subsection("As Link",
			h.Div(
				h.Class("flex flex-wrap gap-4"),
				ui.ButtonA(ui.ButtonProps{}, h.Href("#"), g.Text("Link Button")),
				ui.ButtonA(ui.ButtonProps{Variant: ui.ButtonVariantOutline}, h.Href("#"), g.Text("Outline Link")),
			),
		),

		subsection("Disabled",
			h.Div(
				h.Class("flex flex-wrap gap-4"),
				ui.Button(ui.ButtonProps{}, h.Disabled(), g.Text("Disabled")),
				ui.Button(ui.ButtonProps{Variant: ui.ButtonVariantSecondary}, h.Disabled(), g.Text("Disabled")),
				ui.Button(ui.ButtonProps{Variant: ui.ButtonVariantOutline}, h.Disabled(), g.Text("Disabled")),
			),
		),
	)
}

func cardSection() g.Node {
	return section("Card",
		h.P(
			h.Class("text-muted-foreground mb-6"),
			g.Text("Displays a card with header, content, and footer."),
		),

		subsection("Default",
			ui.Card(ui.CardProps{}, h.Class("w-96"),
				ui.CardHeader(ui.CardHeaderProps{},
					ui.CardTitle(ui.CardTitleProps{}, g.Text("Card Title")),
					ui.CardDescription(ui.CardDescriptionProps{}, g.Text("Card description goes here.")),
				),
				ui.CardContent(ui.CardContentProps{},
					h.P(g.Text("Card content goes here. You can put any content inside.")),
				),
				ui.CardFooter(ui.CardFooterProps{},
					ui.Button(ui.ButtonProps{}, g.Text("Action")),
				),
			),
		),

		subsection("With Action",
			ui.Card(ui.CardProps{}, h.Class("w-96"),
				ui.CardHeader(ui.CardHeaderProps{},
					ui.CardTitle(ui.CardTitleProps{}, g.Text("Notifications")),
					ui.CardDescription(ui.CardDescriptionProps{}, g.Text("You have 3 unread messages.")),
					ui.CardAction(ui.CardActionProps{},
						ui.Button(ui.ButtonProps{Variant: ui.ButtonVariantOutline, Size: ui.ButtonSizeSm}, g.Text("Mark all read")),
					),
				),
				ui.CardContent(ui.CardContentProps{},
					h.P(g.Text("Your notifications will appear here.")),
				),
			),
		),
	)
}

func labelSection() g.Node {
	return section("Label",
		h.P(
			h.Class("text-muted-foreground mb-6"),
			g.Text("Renders an accessible label associated with controls."),
		),

		subsection("Default",
			h.Div(
				h.Class("flex flex-col gap-2"),
				ui.Label(ui.LabelProps{}, h.For("email"), g.Text("Email")),
				h.Input(
					h.Type("email"),
					h.ID("email"),
					h.Placeholder("Enter your email"),
					h.Class("flex h-9 w-full max-w-sm rounded-md border border-input bg-transparent px-3 py-1 text-base shadow-sm transition-colors placeholder:text-muted-foreground focus-visible:outline-none focus-visible:ring-1 focus-visible:ring-ring"),
				),
			),
		),
	)
}

func separatorSection() g.Node {
	return section("Separator",
		h.P(
			h.Class("text-muted-foreground mb-6"),
			g.Text("Visually or semantically separates content."),
		),

		subsection("Horizontal",
			h.Div(
				h.Class("space-y-4"),
				h.Div(
					h.Class("space-y-1"),
					h.H4(h.Class("text-sm font-medium leading-none"), g.Text("Radix Primitives")),
					h.P(h.Class("text-sm text-muted-foreground"), g.Text("An open-source UI component library.")),
				),
				ui.Separator(ui.SeparatorProps{}),
				h.Div(
					h.Class("flex h-5 items-center space-x-4 text-sm"),
					h.Div(g.Text("Blog")),
					ui.Separator(ui.SeparatorProps{Orientation: ui.SeparatorOrientationVertical}),
					h.Div(g.Text("Docs")),
					ui.Separator(ui.SeparatorProps{Orientation: ui.SeparatorOrientationVertical}),
					h.Div(g.Text("Source")),
				),
			),
		),
	)
}

func skeletonSection() g.Node {
	return section("Skeleton",
		h.P(
			h.Class("text-muted-foreground mb-6"),
			g.Text("Use to show a placeholder while content is loading."),
		),

		subsection("Default",
			h.Div(
				h.Class("flex items-center space-x-4"),
				ui.Skeleton(ui.SkeletonProps{}, h.Class("h-12 w-12 rounded-full")),
				h.Div(
					h.Class("space-y-2"),
					ui.Skeleton(ui.SkeletonProps{}, h.Class("h-4 w-[250px]")),
					ui.Skeleton(ui.SkeletonProps{}, h.Class("h-4 w-[200px]")),
				),
			),
		),

		subsection("Card",
			h.Div(
				h.Class("flex flex-col space-y-3"),
				ui.Skeleton(ui.SkeletonProps{}, h.Class("h-[125px] w-[250px] rounded-xl")),
				h.Div(
					h.Class("space-y-2"),
					ui.Skeleton(ui.SkeletonProps{}, h.Class("h-4 w-[250px]")),
					ui.Skeleton(ui.SkeletonProps{}, h.Class("h-4 w-[200px]")),
				),
			),
		),
	)
}

func section(title string, children ...g.Node) g.Node {
	return h.Section(
		h.Class("mb-12"),
		h.H2(
			h.Class("text-2xl font-semibold mb-4 pb-2 border-b"),
			g.Text(title),
		),
		g.Group(children),
	)
}

func subsection(title string, children ...g.Node) g.Node {
	return h.Div(
		h.Class("mb-6"),
		h.H3(
			h.Class("text-lg font-medium mb-3"),
			g.Text(title),
		),
		g.Group(children),
	)
}

