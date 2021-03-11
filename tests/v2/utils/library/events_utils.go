package utils

import (
	"errors"
	"fmt"
	"io/ioutil"

	"github.com/google/go-cmp/cmp"
	"sigs.k8s.io/yaml"

	schema "github.com/devfile/api/v2/pkg/apis/workspaces/v1alpha2"
	//commonUtils "github.com/devfile/api/v2/test/v200/utils/common"
	commonUtils "github.com/devfile/library/tests/v2/utils/common"
)

// UpdateEvents randomly updates attribute values of the Events in the devfile schema
func UpdateEvents(devfile *commonUtils.TestDevfile,) {
	devfile.SetEventsValues(devfile.SchemaDevFile.Events)
}

// VerifyCEvents verifies eventss returned by the parser are the same as those saved in the devfile schema
func VerifyEvents(devfile *commonUtils.TestDevfile, parserEvents schema.Events) error {

	commonUtils.LogInfoMessage("Enter VerifyEvents")
	var errorString []string

	// Compare entire array of commands
	if !cmp.Equal(parserEvents, devfile.SchemaDevFile.Events) {
		errorString = append(errorString, commonUtils.LogErrorMessage(fmt.Sprintf("Events compare failed.")))
		// Array compare failed. Narrow down by comparing indivdual commands
		if (!cmp.Equal(parserEvents.PreStart,devfile.SchemaDevFile.Events.PreStart)) {
			errorString = append(errorString, commonUtils.LogErrorMessage(fmt.Sprintf("preStart Events compare failed.")))
		}
		if (!cmp.Equal(parserEvents.PostStop,devfile.SchemaDevFile.Events.PostStop)) {
			errorString = append(errorString, commonUtils.LogErrorMessage(fmt.Sprintf("postStart Events compare failed.")))
		}
		if (!cmp.Equal(parserEvents.PreStop,devfile.SchemaDevFile.Events.PreStop)) {
			errorString = append(errorString, commonUtils.LogErrorMessage(fmt.Sprintf("preStop Events compare failed.")))
		}
		if (!cmp.Equal(parserEvents.PostStop,devfile.SchemaDevFile.Events.PostStop)) {
			errorString = append(errorString, commonUtils.LogErrorMessage(fmt.Sprintf("postStop Events compare failed.")))
		}
		parserFilename := commonUtils.AddSuffixToFileName(devfile.FileName, "_Events_Parser")
		testFilename := commonUtils.AddSuffixToFileName(devfile.FileName, "_Events_Test")

		commonUtils.LogInfoMessage(fmt.Sprintf(".......marshall and write devfile error snippet %s", parserFilename))
		c, err := yaml.Marshal(parserEvents)
		if err != nil {
			errorString = append(errorString, commonUtils.LogErrorMessage(fmt.Sprintf(".......marshall devfile %s : %v ", parserFilename,err)))
		} else {
			err = ioutil.WriteFile(parserFilename, c, 0644)
			if err != nil {
				errorString = append(errorString, commonUtils.LogErrorMessage(fmt.Sprintf(".......write devfile %s : %v ", parserFilename,err)))
			}
		}
		commonUtils.LogInfoMessage(fmt.Sprintf(".......marshall and write devfile erro snippet %s", testFilename))
		c, err = yaml.Marshal(devfile.SchemaDevFile.Events)
		if err != nil {
			errorString = append(errorString, commonUtils.LogErrorMessage(fmt.Sprintf(".......marshall devfile %s : %v", testFilename,err)))
		} else {
			err = ioutil.WriteFile(testFilename, c, 0644)
			if err != nil {
				errorString = append(errorString, commonUtils.LogErrorMessage(fmt.Sprintf(".......write devfile %s : %v", testFilename,err)))
			}
		}
	} else {
		commonUtils.LogInfoMessage(fmt.Sprintf(" --> Events matched"))
	}

	var err error
	if len(errorString) > 0 {
		err = errors.New(fmt.Sprint(errorString))
	}
	return err
}

