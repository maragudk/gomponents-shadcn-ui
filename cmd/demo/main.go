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
			h.Script(h.Src("https://cdn.tailwindcss.com")),
			h.Script(g.Raw(tailwindConfig)),
			h.StyleEl(g.Raw(styles)),
		},
		Body: []g.Node{
			h.Class("bg-background text-foreground min-h-screen"),
			h.Div(
				h.Class("container mx-auto py-12 px-4 max-w-4xl"),
				header(),
				buttonSection(),
				labelSection(),
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

const tailwindConfig = `
tailwind.config = {
  darkMode: 'class',
  theme: {
    extend: {
      colors: {
        border: 'hsl(var(--border))',
        input: 'hsl(var(--input))',
        ring: 'hsl(var(--ring))',
        background: 'hsl(var(--background))',
        foreground: 'hsl(var(--foreground))',
        primary: {
          DEFAULT: 'hsl(var(--primary))',
          foreground: 'hsl(var(--primary-foreground))',
        },
        secondary: {
          DEFAULT: 'hsl(var(--secondary))',
          foreground: 'hsl(var(--secondary-foreground))',
        },
        destructive: {
          DEFAULT: 'hsl(var(--destructive))',
          foreground: 'hsl(var(--destructive-foreground))',
        },
        muted: {
          DEFAULT: 'hsl(var(--muted))',
          foreground: 'hsl(var(--muted-foreground))',
        },
        accent: {
          DEFAULT: 'hsl(var(--accent))',
          foreground: 'hsl(var(--accent-foreground))',
        },
      },
      borderRadius: {
        lg: 'var(--radius)',
        md: 'calc(var(--radius) - 2px)',
        sm: 'calc(var(--radius) - 4px)',
      },
    },
  },
}
`

const styles = `
:root {
  --background: 0 0% 100%;
  --foreground: 240 10% 3.9%;
  --muted: 240 4.8% 95.9%;
  --muted-foreground: 240 3.8% 46.1%;
  --border: 240 5.9% 90%;
  --input: 240 5.9% 90%;
  --ring: 240 5.9% 10%;
  --primary: 240 5.9% 10%;
  --primary-foreground: 0 0% 98%;
  --secondary: 240 4.8% 95.9%;
  --secondary-foreground: 240 5.9% 10%;
  --destructive: 0 84.2% 60.2%;
  --destructive-foreground: 0 0% 98%;
  --accent: 240 4.8% 95.9%;
  --accent-foreground: 240 5.9% 10%;
  --radius: 0.5rem;
}
`
