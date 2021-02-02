package utils

import (
	"errors"
	"fmt"
	"io/ioutil"

	"github.com/google/go-cmp/cmp"
	"sigs.k8s.io/yaml"

	schema "github.com/devfile/api/v2/pkg/apis/workspaces/v1alpha2"
	)

func getCommandsArray() []string {

	numCommands := GetRandonNumber(8)
	var commands []string

	for i:=0; i<numCommands; i++ {
		commands = append(commands,GetRandomUniqueString(63,true))
	}

	return commands
}


func (devfile *TestDevfile) AddEvent() {
	LogInfoMessage("Add Event")
	devfile.SchemaDevFile.Events = schema.Events{}
	setEventValues(&devfile.SchemaDevFile.Events)
}

func setEventValues(event *schema.Events) {
	if GetBinaryDecision() {
		event.PreStart = getCommandsArray()
		LogInfoMessage(fmt.Sprintf("Add %d PreStart commands",len(event.PreStart)))
	}
	if GetBinaryDecision() {
		event.PostStart = getCommandsArray()
		LogInfoMessage(fmt.Sprintf("Add %d PostStart commands",len(event.PostStart)))
	}
	if GetBinaryDecision() {
		event.PreStop = getCommandsArray()
		LogInfoMessage(fmt.Sprintf("Add %d PreStop commands",len(event.PreStop)))
	}
	if GetBinaryDecision() {
		event.PostStop = getCommandsArray()
		LogInfoMessage(fmt.Sprintf("Add %d PostStop commands",len(event.PostStop)))
	}
}

func (devfile *TestDevfile) UpdateEvent(parserEvent schema.Events) {

	LogInfoMessage("Update Event")
	setEventValues(&parserEvent)
	devfile.SchemaDevFile.Events = parserEvent

}

func (devfile TestDevfile) VerifyEvents(parserEvents []schema.Events) error {

	LogInfoMessage("Enter VerifyEvents")
	var errorString []string

	// Compare entire array of commands
	if !cmp.Equal(parserEvents, devfile.SchemaDevFile.Events) {
		errorString = append(errorString,LogErrorMessage(fmt.Sprintf("Events compare failed : %s",devfile.FileName)))
		parserFilename := AddSuffixToFileName(devfile.FileName, "_Parser")
		LogInfoMessage(fmt.Sprintf(".......write parser copy for events :  %s", parserFilename))
		c, err := yaml.Marshal(parserEvents)
		if err != nil {
			errorString = append(errorString, LogErrorMessage(fmt.Sprintf("......failed to marshall parser events %s : %v", parserFilename,err)))
		} else {
			err = ioutil.WriteFile(parserFilename, c, 0644)
			if err != nil {
				errorString = append(errorString, LogErrorMessage(fmt.Sprintf(".......write to write parser events %s : %v", parserFilename,err)))
			}
		}

	}

	var err error
	if len(errorString) > 0 {
		err = errors.New(fmt.Sprint(errorString))
	}
	return err
}
