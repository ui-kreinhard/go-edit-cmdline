package cmdline

import (
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func Load(fileName string) (string, error) {
	bytes, err := ioutil.ReadFile(fileName)
	return string(bytes), err
}

func Parse(cmdlineContent string) CmdLineParameters {
	cmdLineParameters := NewCmdLineParameters()
	elementsRaw := strings.Split(cmdlineContent, " ")
	for _, element := range elementsRaw {
		element = strings.Trim(element, "\n ")
		cmdLineParameters.Add(CreateFlagOrValuePair(element))
	}

	return *cmdLineParameters
}

func Save(fileName string, newContent string) error {
	return ioutil.WriteFile(fileName, []byte(newContent), 0755)
}

func CreateFlagOrValuePair(rawString string) Content {
	elements := strings.Split(rawString, "=")
	if len(elements) <= 1 {
		return &Flag{
			elements[0],
		}
	} else {
		return &Value{
			elements[0],
			elements[1],
		}
	}
}

func ParseFlags(cmdLineParameters CmdLineParameters) CmdLineParameters {
	for _, arg := range os.Args[1:] {
		switch arg[0] {
		case '+':
			cmdLineParameters.AddOrEdit(CreateFlagOrValuePair(arg[1:]))
		case '-':
			cmdLineParameters.Delete(CreateFlagOrValuePair(arg[1:]))
		}
	}
	return cmdLineParameters
}

func defaultEmptyString(toTest, defaultStr string) string {
	if toTest == "" {
		return defaultStr
	}
	return toTest
}

func Run() {
	in := defaultEmptyString(os.Getenv("cmdline"), "/boot/cmdline.txt")
	out := defaultEmptyString(os.Getenv("target"), "/boot/cmdline.txt")

	raw, err := Load(in)
	if err != nil {
		log.Fatalln("Cannot read file", err)
	}

	parsedCmdlineParams := Parse(raw)
	newCmdLineParams := ParseFlags(parsedCmdlineParams)
	err = Save(out, newCmdLineParams.String())
	if err != nil {
		log.Fatalln("cannot write", err)
	}
}
