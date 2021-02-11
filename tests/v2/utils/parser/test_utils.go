package api

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/devfile/library/tests/v2/utils/common"

	schema "github.com/devfile/api/v2/pkg/apis/workspaces/v1alpha2"
	devfilepkg "github.com/devfile/library/pkg/devfile"
	"github.com/devfile/library/pkg/devfile/parser"
	devfileCtx "github.com/devfile/library/pkg/devfile/parser/context"
	devfileData "github.com/devfile/library/pkg/devfile/parser/data"
	"github.com/devfile/library/pkg/devfile/parser/data/v2/common"
	"sigs.k8s.io/yaml"
)


// WriteDevfile create a devifle on disk for use in a test.
// If useParser is true the parser library is used to generate the file, otherwise "sigs.k8s.io/yaml" is used.
func (devfile *TestDevfile) WriteDevfile(useParser bool) error {
	var err error

	fileName := devfile.FileName
	if !strings.HasSuffix(fileName, ".yaml") {
		fileName += ".yaml"
	}

	if useParser {
		LogInfoMessage(fmt.Sprintf("Use Parser to write devfile %s", fileName))

		ctx := devfileCtx.NewDevfileCtx(fileName)

		err = ctx.SetAbsPath()
		if err != nil {
			LogErrorMessage(fmt.Sprintf("Setting devfile path : %v", err))
		} else {
			devObj := parser.DevfileObj{
				Ctx:  ctx,
				Data: devfile.ParserData,
			}
			err = devObj.WriteYamlDevfile()
			if err != nil {
				LogErrorMessage(fmt.Sprintf("Writing devfile : %v", err))
			} else {
				devfile.SchemaParsed = false
			}
		}

	} else {
		LogInfoMessage(fmt.Sprintf("Marshall and write devfile %s", devfile.FileName))
		c, marshallErr := yaml.Marshal(&(devfile.SchemaDevFile))

		if marshallErr != nil {
			err = errors.New(LogErrorMessage(fmt.Sprintf("Marshall devfile %s : %v", devfile.FileName, marshallErr)))
		} else {
			err = ioutil.WriteFile(fileName, c, 0644)
			if err != nil {
				LogErrorMessage(fmt.Sprintf("Write devfile %s : %v", devfile.FileName, err))
			} else {
				devfile.SchemaParsed = false
			}
		}
	}
	return err
}

// parseSchema uses the parser to parse a devfile on disk
func (devfile *TestDevfile) parseSchema() error {

	var err error
	if !devfile.SchemaParsed {
		err = devfile.WriteDevfile(true)
		if err != nil {
			LogErrorMessage(fmt.Sprintf("From WriteDevfile %v : ", err))
		} else {
			LogInfoMessage(fmt.Sprintf("Parse and Validate %s : ", devfile.FileName))
			parsedSchemaObj, parse_err := devfilepkg.ParseAndValidate(devfile.FileName)
			if parse_err != nil {
				err = parse_err
				LogErrorMessage(fmt.Sprintf("From ParseAndValidate %v : ", err))
			}
			devfile.SchemaParsed = true
			devfile.ParserData = parsedSchemaObj.Data
		}
	}
	return err
}

// Verify verifies the contents of the specified devfile with the expected content
func (devfile *TestDevfile) Verify() error {

	LogInfoMessage(fmt.Sprintf("Verify %s : ", devfile.FileName))

	var errorString []string

	err := devfile.parseSchema()

	if err != nil {
		errorString = append(errorString, LogErrorMessage(fmt.Sprintf("parsing schema %s : %v", devfile.FileName, err)))
	} else {
		LogInfoMessage(fmt.Sprintf("Get commands %s : ", devfile.FileName))
		commands, _ := devfile.ParserData.GetCommands(common.DevfileOptions{})
		if commands != nil && len(commands) > 0 {
			err = devfile.VerifyCommands(commands)
			if err != nil {
				errorString = append(errorString, LogErrorMessage(fmt.Sprintf("Verfify Commands %s : %v", devfile.FileName, err)))
			}
		} else {
			LogInfoMessage(fmt.Sprintf("No command found in %s : ", devfile.FileName))
		}

		LogInfoMessage(fmt.Sprintf("Get components %s : ", devfile.FileName))
		components, _ := devfile.ParserData.GetComponents(common.DevfileOptions{})
		if components != nil && len(components) > 0 {
			err = devfile.VerifyComponents(components)
			if err != nil {
				errorString = append(errorString, LogErrorMessage(fmt.Sprintf("Verfify Commands %s : %v", devfile.FileName, err)))
			}
		} else {
			LogInfoMessage(fmt.Sprintf("No components found in %s : ", devfile.FileName))
		}
	}
	var returnError error
	if len(errorString) > 0 {
		returnError = errors.New(fmt.Sprint(errorString))
	}
	return returnError

}

// EditCommands modifies random attributes for each of the commands in the devfile.
func (devfile *TestDevfile) EditCommands() error {

	LogInfoMessage(fmt.Sprintf("Edit %s : ", devfile.FileName))

	err := devfile.parseSchema()
	if err != nil {
		LogErrorMessage(fmt.Sprintf("From parser : %v", err))
	} else {
		LogInfoMessage(fmt.Sprintf(" -> Get commands %s : ", devfile.FileName))
		commands, _ := devfile.ParserData.GetCommands(common.DevfileOptions{})
		for _, command := range commands {
			err = devfile.UpdateCommand(command.Id)
			if err != nil {
				LogErrorMessage(fmt.Sprintf("Updating command : %v", err))
			}
		}
	}
	return err
}

// EditComponents modifies random attributes for each of the components in the devfile.
func (devfile *TestDevfile) EditComponents() error {

	LogInfoMessage(fmt.Sprintf("Edit %s : ", devfile.FileName))

	err := devfile.parseSchema()
	if err != nil {
		LogErrorMessage(fmt.Sprintf("From parser : %v", err))
	} else {
		LogInfoMessage(fmt.Sprintf(" -> Get commands %s : ", devfile.FileName))
		components, _ := devfile.ParserData.GetComponents(common.DevfileOptions{})
		for _, component := range components {
			err = devfile.UpdateComponent(component.Name)
			if err != nil {
				LogErrorMessage(fmt.Sprintf("Updating component : %v", err))
			}
		}
	}
	return err
}

// runMultiThreadTest : Runs the same test on multiple threads, the test is based on the content of the specified TestContent
func runMultiThreadTest(testContent TestContent, t *testing.T) {

	utils.LogMessage(fmt.Sprintf("Start Threaded test for %s", testContent.FileName))

	devfileName := testContent.FileName
	var i int
	for i = 1; i < numThreads; i++ {
		testContent.FileName = utils.AddSuffixToFileName(devfileName, strconv.Itoa(i))
		go runTest(testContent, t)
	}
	testContent.FileName = utils.AddSuffixToFileName(devfileName, strconv.Itoa(i))
	runTest(testContent, t)

	utils.LogMessage(fmt.Sprintf("Sleep 3 seconds to allow all threads to complete : %s", devfileName))
	time.Sleep(3 * time.Second)
	utils.LogMessage(fmt.Sprintf("Sleep complete : %s", devfileName))

}

// runTest : Runs a test beased on the content of the specified TestContent
func runTest(testContent common.TestContent, t *testing.T) {

	common.runTest(testContent,t)

	if testContent.EditContent {
		if len(testContent.CommandTypes) > 0 {
			err = testDevfile.EditCommands()
			if err != nil {
				t.Fatalf(utils.LogErrorMessage(fmt.Sprintf("ERROR editing commands :  %s : %v", testContent.FileName, err)))
			}
		}
		if len(testContent.ComponentTypes) > 0 {
			err = testDevfile.EditComponents()
			if err != nil {
				t.Fatalf(utils.LogErrorMessage(fmt.Sprintf("ERROR editing components :  %s : %v", testContent.FileName, err)))
			}
		}
		err = testDevfile.Verify()
		if err != nil {
			t.Fatalf(utils.LogErrorMessage(fmt.Sprintf("ERROR verifying devfile content : %s : %v", testContent.FileName, err)))
		}

	}

}

