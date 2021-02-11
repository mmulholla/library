package common

import (
	"errors"
	"fmt"
	"io/ioutil"

	"github.com/google/go-cmp/cmp"
	"sigs.k8s.io/yaml"

	"github.com/devfile/library/tests/v2/utils/common"
	schema "github.com/devfile/api/v2/pkg/apis/workspaces/v1alpha2"
)

// UpdateCommand randomly updates attribute values of a specified command in the devfile schema
func (devfile *TestDevfile) UpdateCommand(commandId string) error {

	var err error
	testCommand, found := getSchemaCommand(devfile.SchemaDevFile.Commands, commandId)
	if found {
		LogInfoMessage(fmt.Sprintf("Updating command id: %s", commandId))
		if testCommand.Exec != nil {
			devfile.setExecCommandValues(testCommand)
		} else if testCommand.Composite != nil {
			devfile.setCompositeCommandValues(testCommand)
		} else if testCommand.Apply != nil {
			devfile.setApplyCommandValues(testCommand)
		} else if testCommand.VscodeTask != nil {
			devfile.setVscodeTaskCommandValues(testCommand)
		} else if testCommand.VscodeLaunch != nil {
			devfile.setVscodeLaunchCommandValues(testCommand)
		}
	} else {
		err = errors.New(LogErrorMessage(fmt.Sprintf("Command not found in test : %s", commandId)))
	}
	return err
}


// VerifyCommands verifies commands returned by the parser are the same as those saved in the devfile schema
func (devfile *TestDevfile) VerifyCommands(parserCommands []schema.Command) error {

	LogInfoMessage("Enter VerifyCommands")
	var errorString []string

	// Compare entire array of commands
	if !cmp.Equal(parserCommands, devfile.SchemaDevFile.Commands) {
		errorString = append(errorString, LogErrorMessage(fmt.Sprintf("Command array compare failed.")))
		// Array compare failed. Narrow down by comparing indivdual commands
		for _, command := range parserCommands {
			if testCommand, found := getSchemaCommand(devfile.SchemaDevFile.Commands, command.Id); found {
				if !cmp.Equal(command, *testCommand) {
					parserFilename := AddSuffixToFileName(devfile.FileName, "_"+command.Id+"_Parser")
					testFilename := AddSuffixToFileName(devfile.FileName, "_"+command.Id+"_Test")
					LogInfoMessage(fmt.Sprintf(".......marshall and write devfile %s", devfile.FileName))
					c, err := yaml.Marshal(command)
					if err != nil {
						errorString = append(errorString, LogErrorMessage(fmt.Sprintf(".......marshall devfile %s", parserFilename)))
					} else {
						err = ioutil.WriteFile(parserFilename, c, 0644)
						if err != nil {
							errorString = append(errorString, LogErrorMessage(fmt.Sprintf(".......write devfile %s", parserFilename)))
						}
					}
					LogInfoMessage(fmt.Sprintf(".......marshall and write devfile %s", testFilename))
					c, err = yaml.Marshal(testCommand)
					if err != nil {
						errorString = append(errorString, LogErrorMessage(fmt.Sprintf(".......marshall devfile %s", testFilename)))
					} else {
						err = ioutil.WriteFile(testFilename, c, 0644)
						if err != nil {
							errorString = append(errorString, LogErrorMessage(fmt.Sprintf(".......write devfile %s", testFilename)))
						}
					}
					errorString = append(errorString, LogInfoMessage(fmt.Sprintf("Command %s did not match, see files : %s and %s", command.Id, parserFilename, testFilename)))
				} else {
					LogInfoMessage(fmt.Sprintf(" --> Command  matched : %s", command.Id))
				}
			} else {
				errorString = append(errorString, LogErrorMessage(fmt.Sprintf("Command from parser not known to test - id : %s ", command.Id)))
			}

		}
		for _, command := range devfile.SchemaDevFile.Commands {
			if _, found := getSchemaCommand(parserCommands, command.Id); !found {
				errorString = append(errorString, LogErrorMessage(fmt.Sprintf("Command from test not returned by parser : %s ", command.Id)))
			}
		}
	} else {
		LogInfoMessage(fmt.Sprintf(" --> Command structures matched"))
	}

	var err error
	if len(errorString) > 0 {
		err = errors.New(fmt.Sprint(errorString))
	}
	return err
}
