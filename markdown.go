package gomd

import (
	"log"
	"strings"
)

const LINE_RETURN = "\n"

const (
	TAB           = '\u0009'
	SPACE         = '\u0020'
	LINE_FEED     = '\u000A'
	CARRIAGE_FEED = '\u000A'
)

type Serializer interface {
	ToString() string
	ToUnicode() []rune
}

type TextStyle string

const (
	Normal TextStyle = "normal"
	Code   TextStyle = "code"
	Bold   TextStyle = "bold"
	Italic TextStyle = "italic"
)

type Text struct {
	Content string
	Style   TextStyle
}

func (t Text) ToString() string {
	switch t.Style {
	default:
	case Normal:
		return t.Content

	case Code:
		return " `" + t.Content + "` "

	case Bold:
		return " __" + t.Content + "__ "

	case Italic:
		return " *" + t.Content + "* "
	}

	return ""
}

func (t Text) ToUnicode() []rune {
	log.Fatal("Not implemented")
	return []rune{}
}

type Paragraph struct {
	TextElements []Text
}

func (p Paragraph) ToString() string {
	elements := make([]string, 0)

	for _, el := range p.TextElements {
		elements = append(elements, el.ToString())
	}

	return strings.Join(elements, "") + string(LINE_FEED)
}

func (p Paragraph) ToUnicode() []rune {
	log.Fatal("Not implemented")
	return []rune{}
}

type Markdown struct {
	elements []Serializer
}

func (d *Markdown) Add(elements ...Serializer) *Markdown {
	d.elements = append(d.elements, elements...)
	return d
}

func (d *Markdown) Flush() *Markdown {
	clear(d.elements)
	return d
}

func (d *Markdown) GetElements() []Serializer {
	return d.elements
}

func (d *Markdown) Print() string {
	output := make([]string, 0)

	for _, e := range d.elements {
		output = append(output, e.ToString())
	}

	return strings.Join(output, "")
}

func (d *Markdown) Bytes() []byte {
	return []byte(d.Print())
}
