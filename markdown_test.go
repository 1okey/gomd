package gomd

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHeading(t *testing.T) {
	h := Heading{Text: Paragraph{[]Text{{Content: "AN UNEXPECTED PARTY", Style: Normal}}}, Size: 3}

	assert.Equal(t, "### AN UNEXPECTED PARTY"+LINE_RETURN, h.ToString())
}

func TestMarkdown(t *testing.T) {
	md := Markdown{}

	h := Heading{Text: Paragraph{TextElements: []Text{{Content: "AN UNEXPECTED PARTY", Style: Normal}}}, Size: 3}
	text := Text{Content: "In a hole in the ground there lived a hobbit. Not a nasty, dirty, wet hole, filled with the ends of worms and an oozy smell, nor yet a dry, bare, sandy hole with nothing in it to sit down on or to eat: it was a hobbit-hole, and that means comfort.", Style: Normal}
	md.Add(h, text)

	assert.Len(t, strings.Split(md.Print(), LINE_RETURN), 2)

	one := Text{Content: "In a hole in the ground there lived a hobbit. ", Style: Normal}
	other := Text{Content: "Not a nasty, dirty, wet hole,", Style: Normal}
	other_bold := Text{Content: "filled with the ends of worms", Style: Bold}
	other_second := Text{Content: "and an oozy smell, nor yet a dry, bare, sandy hole with nothing in it to sit down on or to eat: it was a hobbit-hole,", Style: Normal}
	other_second_italic := Text{Content: "and that means comfort.", Style: Italic}

	p := Paragraph{TextElements: []Text{one, other, other_bold, other_second, other_second_italic}}
	assert.Equal(t, "In a hole in the ground there lived a hobbit. Not a nasty, dirty, wet hole, __filled with the ends of worms__ and an oozy smell, nor yet a dry, bare, sandy hole with nothing in it to sit down on or to eat: it was a hobbit-hole, *and that means comfort.* "+LINE_RETURN, p.ToString())

	md.Add(p)
	assert.Len(t, strings.Split(md.Print(), LINE_RETURN), 3)
}
