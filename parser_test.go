package yaml

import (
	"bytes"
	"fmt"
	"regexp"
	"strings"
	"testing"
)

func TestParser(t *testing.T) {
	yamlText := `first:
	second:
		third:
			third1: third1-text
			third2: third2-text
	second2: second2-text
	second3:
	- aaa
	- bbb
`
	yamlText = strings.Replace(yamlText, "\t", "  ", -1)
	input := bytes.NewBufferString(yamlText)
	p := NewParser(input)

	r := p.Read()
	if r != "first" {
		t.Errorf("return [%s]", r)
	}

	r = p.Read()
	if r != "first.second" {
		t.Errorf("return [%s]", r)
	}

	r = p.Read()
	if r != "first.second.third" {
		t.Errorf("return [%s]", r)
	}

	r = p.Read()
	if r != "first.second.third.third1" {
		t.Errorf("return [%s]", r)
	}

	r = p.Read()
	if r != "first.second.third.third2" {
		t.Errorf("return [%s]", r)
	}

	r = p.Read()
	if r != "first.second2" {
		t.Errorf("return [%s]", r)
	}

	r = p.Read()
	if r != "first.second3" {
		t.Errorf("return [%s]", r)
	}

	r = p.Read()
	if r != "" {
		t.Errorf("return [%s]", r)
	}
}

func TestRegex(t *testing.T) {
	re, err := regexp.Compile("^([^\\s:]+)[\\s]?[:]")
	if err != nil {
		panic(err.Error())
	}

	tx := re.FindStringSubmatch("aaa:")
	if len(tx) > 0 {
		fmt.Printf("match %s", tx)
	} else {
		print("not match")
	}
}
