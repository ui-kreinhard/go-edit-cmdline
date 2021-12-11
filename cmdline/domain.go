package cmdline

import (
	"strings"
)

type CmdLineParameters struct {
	params []Content
}

func (c *CmdLineParameters) Add(newContent Content) {
	c.params = append(c.params, newContent)
}

func (c *CmdLineParameters) AddOrEdit(newContent Content) {
	newParams := []Content{}
	found := false
	for _, currentElement := range c.params {
		if currentElement.GetName() == newContent.GetName() {
			newParams = append(newParams, newContent)
			found = true
		} else {
			newParams = append(newParams, currentElement)
		}
	}
	if !found {
		newParams = append(newParams, newContent)
	}
	c.params = newParams
}

func (c *CmdLineParameters) Delete(elementToDelete Content) {
	newParams := []Content{}
	for _, currentElement := range c.params {
		if currentElement.GetName() != elementToDelete.GetName() {
			newParams = append(newParams, currentElement)
		}
	}
	c.params = newParams
}

func (c *CmdLineParameters) String() string {
	stringElements := []string{}
	for _, element := range c.params {
		stringElements = append(stringElements, element.String())
	}
	return strings.Join(stringElements, " ") + "\n"
}

func NewCmdLineParameters() *CmdLineParameters {
	return &CmdLineParameters{
		[]Content{},
	}
}

type Flag struct {
	Name string
}

func (f *Flag) String() string {
	return f.Name
}

func (f *Flag) GetName() string {
	return f.Name
}

type Value struct {
	Name  string
	Value string
}

func (v *Value) GetName() string {
	return v.Name
}

func (v *Value) String() string {
	return v.Name + "=" + v.Value
}

type Content interface {
	String() string
	GetName() string
}
