package gomd

import (
	"bytes"
	"log"
	"slices"
	"strings"
)

// ThematicBreak

type ThematicBreak struct {
	character byte
}

var VALID_TB_CHARS []byte = []byte{'+', '*', '-'}

func makeThematicBreak(ch byte) ThematicBreak {
	if slices.Index(VALID_TB_CHARS, ch) == -1 {
		log.Fatal("Thematic break can only be made of " + string(bytes.Join([][]byte{VALID_TB_CHARS}, []byte(", "))))
	}

	return ThematicBreak{ch}
}

func (tb ThematicBreak) ToString() string {
	var sb strings.Builder
	for range 3 {
		sb.WriteByte(byte(tb.character))
	}
	sb.WriteByte(byte(LINE_FEED))
	return sb.String()
}

func (tb ThematicBreak) ToUnicode() []rune {
	log.Fatal("Not implemented")
	return []rune{}
}

// ATX heading

type Heading struct {
	Text Paragraph
	Size uint8
}

func (h Heading) ToString() string {
	return strings.Repeat("#", int(h.Size)) + " " + h.Text.ToString()
}

func (h Heading) ToUnicode() []rune {
	log.Fatal("Not implemented")
	return []rune{}
}

// Setext heading (multiline)
type MultilineHeading struct {
	text Paragraph
	size uint8
}

func makeMultilineHeading(text Paragraph, size uint8) MultilineHeading {
	if size > 2 || size == 0 {
		log.Fatal("Multiline heading can only be of size 1 or 2")
	}

	return MultilineHeading{text, size}
}

var sizeToUndeline = map[uint8]string{
	1: "=",
	2: "-",
}

func (mh MultilineHeading) ToString() string {
	var sb strings.Builder
	sb.Write([]byte(mh.text.ToString()))
	sb.Write([]byte(sizeToUndeline[mh.size]))
	sb.WriteByte(LINE_FEED)
	return sb.String()
}

func (mh MultilineHeading) ToUnicode() []rune {
	log.Fatal("Not implemented")
	return []rune{}
}
