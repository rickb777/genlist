package option

import "github.com/clipperhouse/typewriter"

var option = &typewriter.Template{
	Name: "option",
	Text: `// {{.OptionName}} is an optional of type {{.Type}}. Use it where you want to be explicit about the presence
	or absence of data.
interface {{.OptionName}} {
}

type Some{{.Type}} {{.Type}}

type No{{.Type}} struct{}
`,
}
