package styles

import (
	"github.com/softpunks/chroma"
)

// Manni style.
var Manni = Register(chroma.MustNewStyle("manni", chroma.StyleEntries{
	chroma.TextWhitespace:        "#bbbbbb",
	chroma.Comment:               "italic #0099FF",
	chroma.CommentPreproc:        "noitalic #009999",
	chroma.CommentSpecial:        "bold",
	chroma.Keyword:               "bold #006699",
	chroma.KeywordPseudo:         "nobold",
	chroma.KeywordType:           "#007788",
	chroma.Operator:              "#555555",
	chroma.OperatorWord:          "bold #000000",
	chroma.NameBuiltin:           "#336666",
	chroma.NameFunction:          "#CC00FF",
	chroma.NameClass:             "bold #00AA88",
	chroma.NameNamespace:         "bold #00CCFF",
	chroma.NameException:         "bold #CC0000",
	chroma.NameVariable:          "#003333",
	chroma.NameConstant:          "#336600",
	chroma.NameLabel:             "#9999FF",
	chroma.NameEntity:            "bold #999999",
	chroma.NameAttribute:         "#330099",
	chroma.NameTag:               "bold #330099",
	chroma.NameDecorator:         "#9999FF",
	chroma.LiteralString:         "#CC3300",
	chroma.LiteralStringDoc:      "italic",
	chroma.LiteralStringInterpol: "#AA0000",
	chroma.LiteralStringEscape:   "bold #CC3300",
	chroma.LiteralStringRegex:    "#33AAAA",
	chroma.LiteralStringSymbol:   "#FFCC33",
	chroma.LiteralStringOther:    "#CC3300",
	chroma.LiteralNumber:         "#FF6600",
	chroma.GenericHeading:        "bold #003300",
	chroma.GenericSubheading:     "bold #003300",
	chroma.GenericDeleted:        "border:#CC0000 bg:#FFCCCC",
	chroma.GenericInserted:       "border:#00CC00 bg:#CCFFCC",
	chroma.GenericError:          "#FF0000",
	chroma.GenericEmph:           "italic",
	chroma.GenericStrong:         "bold",
	chroma.GenericPrompt:         "bold #000099",
	chroma.GenericOutput:         "#AAAAAA",
	chroma.GenericTraceback:      "#99CC66",
	chroma.GenericUnderline:      "underline",
	chroma.Error:                 "bg:#FFAAAA #AA0000",
}))
