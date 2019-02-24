package yaml

import (
	"bufio"
	"io"
	"regexp"
)

type objType int

const (
	typeObject objType = iota
	typeList
)

type stack struct {
	objType    objType
	spaces     int
	text       string
	listNumber int
}

// Parser ...
type Parser struct {
	stacks  []stack
	scanner *bufio.Scanner
}

var leftSpaceRegex *regexp.Regexp
var startObjectRegex *regexp.Regexp
var propertyRegex *regexp.Regexp
var listRegex *regexp.Regexp

func init() {
	var err error
	leftSpaceRegex, _ = regexp.Compile("^([\\s]+)[\\S]")
	startObjectRegex, _ = regexp.Compile("^([^\\s:]+)[\\s]*[:][\\s]*$")
	propertyRegex, err = regexp.Compile("^([^\\s:]+)[\\s]*[:][\\s]*[\\S]")
	listRegex, err = regexp.Compile("^-[\\s]*[\\s]*([\\S]+)")
	if err != nil {
		panic(err.Error())
	}
}

// NewParser ...
func NewParser(origin io.Reader) Parser {
	p := Parser{}
	p.stacks = []stack{stack{}}
	p.scanner = bufio.NewScanner(origin)
	return p
}

func (p *Parser) Read() string {
	scaned := p.scanner.Scan()
	if !scaned {
		return ""
	}
	line := p.scanner.Text()
	bottom := p.stacks[len(p.stacks)-1]
	matches := leftSpaceRegex.FindStringSubmatch(line)
	text := line
	spaces := ""
	if len(matches) > 1 {
		spaces = matches[1]
		text = line[len(spaces):]
	}
	if len(p.stacks) > 1 && len(spaces) < bottom.spaces {
		for bottom.spaces >= len(spaces) {
			p.stacks = p.stacks[:len(p.stacks)-1]
			bottom = p.stacks[len(p.stacks)-1]
		}
	}
	if (len(p.stacks) == 1 && len(spaces) == bottom.spaces) ||
		(len(p.stacks) > 1 && len(spaces) > bottom.spaces) {
		print(text)
		matches := propertyRegex.FindStringSubmatch(text)
		if len(matches) > 0 {
			path := ""
			if len(p.stacks) > 1 {
				path = bottom.text + "." + matches[1]
			} else {
				path = matches[1]
			}
			return path
		}
		matches = startObjectRegex.FindStringSubmatch(text)
		if len(matches) > 0 {
			path := ""
			if len(p.stacks) > 1 {
				path = bottom.text + "." + matches[1]
			} else {
				path = matches[1]
			}
			n := stack{
				objType: typeObject,
				spaces:  len(spaces),
				text:    path,
			}
			p.stacks = append(p.stacks, n)
			return path
		}

	}
	if len(spaces) > bottom.spaces {

	}
	return ""
}
