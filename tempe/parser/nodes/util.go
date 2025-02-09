package nodes

import (
	"strings"
)

const TextAlignWidth int = 60
const FormatStringMaxLength int = 60

func formatText(text string) string {
	str := replaceNewline(text)
	if len(str) > FormatStringMaxLength {
		str = str[:FormatStringMaxLength-3] + "..."
	}
	return str
}

func replaceNewline(text string) string {
	text = strings.Replace(text, "\r", `\r`, -1)
	return strings.Replace(text, "\n", `\n`, -1)
}

func replaceSpace(text string) string {
	text = strings.Replace(text, "\t", `\t`, -1)
	return strings.Replace(text, " ", "␣", -1)
}

func replaceWhitespace(text string) string {
	text = replaceNewline(text)
	return replaceSpace(text)
}

type nodeSpec struct {
	Type string `json:"type"`
	Spec any    `json:"spec"`
}

func jsonNodeSpec(typeName string, spec any) nodeSpec {
	return nodeSpec{
		Type: typeName,
		Spec: spec,
	}
}
