package gomd

import (
	"strings"
)

const LINE_RETURN = "\n"

type Serializer interface {
	ToString() string
}

type Heading struct {
	Text Paragraph
	Size uint8
}

func (h Heading) ToString() string {
	return strings.Repeat("#", int(h.Size)) + " " + h.Text.ToString()
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

func (p Text) ToString() string {
	switch p.Style {
	default:
	case Normal:
		return p.Content

	case Code:
		return " `" + p.Content + "` "

	case Bold:
		return " __" + p.Content + "__ "

	case Italic:
		return " *" + p.Content + "* "
	}

	return ""
}

type Paragraph struct {
	TextElements []Text
}

func (p Paragraph) ToString() string {
	elements := make([]string, 0)

	for _, el := range p.TextElements {
		elements = append(elements, el.ToString())
	}

	return strings.Join(elements, "") + LINE_RETURN
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
