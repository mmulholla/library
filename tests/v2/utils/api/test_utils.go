package api

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"time"

	schema "github.com/devfile/api/v2/pkg/apis/workspaces/v1alpha2"
	devfilepkg "github.com/devfile/library/pkg/devfile"
	"github.com/devfile/library/pkg/devfile/parser"
	devfileCtx "github.com/devfile/library/pkg/devfile/parser/context"
	devfileData "github.com/devfile/library/pkg/devfile/parser/data"
	"github.com/devfile/library/pkg/devfile/parser/data/v2/common"
	"sigs.k8s.io/yaml"
)

const (
	defaultTempDir = "./tmp/"
	logFileName    = "test.log"
	// logToFileOnly - If set to false the log output will also be output to the console
	logToFileOnly = true // If set to false the log output will also be output to the console
)

var schemas = make(map[string]SchemaFile)

type SchemaFile struct {
	FileName string
	URL      string
	Schema   *jsonschema.Schema
}

// CheckWithSchema checks the validity of aa devfile againts the schema.
func (schemaFile *SchemaFile) CheckWithSchema(yamlFile string, expectedMessage string) error {

	// Read the created yaml file, ready for converison to json
	data, err := ioutil.ReadFile(yamlFile)
	if err != nil {
		LogErrorMessage(fmt.Sprintf("  FAIL: schema : unable to read %s: %v", yamlFile, err))
		return err
	}

	// Convert the yaml file to json
	yamldoc, err := yaml.YAMLToJSON(data)
	if err != nil {
		LogErrorMessage(fmt.Sprintf("  FAIL : %s : schema : failed to convert to json : %v", yamlFile, err))
		return err
	}

	validationErr := schemaFile.Schema.Validate(bytes.NewReader(yamldoc))
	if validationErr != nil {
		if len(expectedMessage) > 0 {
			if !strings.Contains(validationErr.Error(), expectedMessage) {
				err = errors.New(LogErrorMessage(fmt.Sprintf("  FAIL : schema : %s : Did not fail as expected : %s  got : %v", yamlFile, expectedMessage, validationErr)))
			} else {
				LogInfoMessage(fmt.Sprintf("PASS: schema :  Expected Error received : %s", expectedMessage))
			}
		} else {
			err = errors.New(LogErrorMessage(fmt.Sprintf("  FAIL : schema : %s : Did not pass as expected, got : %v", yamlFile, validationErr)))
		}
	} else {
		if len(expectedMessage) > 0 {
			err = errors.New(LogErrorMessage(fmt.Sprintf("  FAIL : schema : %s :  was valid - Expected Error not found : %v", yamlFile, validationErr)))
		} else {
			LogInfoMessage(fmt.Sprintf("  PASS : schema : %s : devfile was valid.", yamlFile))
		}
	}
	return err
}

// GetSchema downloads and saves a schema from the provided url
func GetSchema(URL string) (*SchemaFile, error) {

	schemaFile, found := schemas[URL]
	if !found {

		schemaFile = SchemaFile{}
		schemaFile.URL = URL
		fileName := URL[strings.LastIndex(URL, "/"):]
		schemaFile.FileName = filepath.Join("./tmp/", fileName)

		// Get the json schema from the URL
		resp, err := http.Get(URL)
		if err != nil {
			LogErrorMessage(fmt.Sprintf("FAIL : Failed to get from url %s, %v", URL, err))
			return nil, err
		}
		defer resp.Body.Close()

		// Create the file to contain the json schema
		out, err := os.Create(schemaFile.FileName)
		if err != nil {
			LogErrorMessage(fmt.Sprintf("FAIL : Failed to open file %s :  %v", schemaFile.FileName, err))
			return nil, err
		}
		defer out.Close()

		// Write the url content to the file
		_, err = io.Copy(out, resp.Body)
		if err != nil {
			LogErrorMessage(fmt.Sprintf("FAIL : Failed to copy the schema into a file  %v", err))
			return nil, err
		}

		// Prepare the schema file
		compiler := jsonschema.NewCompiler()
		compiler.Draft = jsonschema.Draft7
		schemaFile.Schema, err = compiler.Compile(schemaFile.FileName)
		if err != nil {
			//t.Fatalf("  FAIL : Schema compile failed : %s: %v", testJsonContent.SchemaFile, err)
			LogErrorMessage(fmt.Sprintf("FAIL : Failed to compile schema  %v", err))
			return nil, err
		} else {
			LogInfoMessage(fmt.Sprintf("Schema compiled from file: %s, url: %s)", schemaFile.FileName, URL))
			schemas[URL] = schemaFile
		}
	}
	return &schemaFile, nil
}

// WriteDevfile create a devifle on disk for use in a test.
// If useParser is true the parser library is used to generate the file, otherwise "sigs.k8s.io/yaml" is used.
func (devfile *TestDevfile) WriteDevfile(useParser bool) error {
	var err error

	fileName := devfile.FileName
	if !strings.HasSuffix(fileName, ".yaml") {
		fileName += ".yaml"
	}

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
	return err
}

// parseSchema uses the parser to parse a devfile on disk
func (devfile *TestDevfile) parseSchema() error {

	var err error
	devfile.SchemaParsed = true

	var schemaFile *SchemaFile
	schemaFile, err = GetSchema("https://raw.githubusercontent.com/devfile/api/master/schemas/latest/ide-targeted/devfile.json")
	if err != nil {
		LogErrorMessage(fmt.Sprintf("Failed to get devfile schema : %v", err))
	} else {
		err = schemaFile.CheckWithSchema(devfile.FileName, "")
		if err != nil {
			LogErrorMessage(fmt.Sprintf("Verification with devfile schema failed : %v", err))
		} else {
			LogInfoMessage(fmt.Sprintf("Devfile validated using JSONSchema schema : %s", devfile.FileName))
		}
	}

	devfile.SchemaParsed = true
	return err
}
