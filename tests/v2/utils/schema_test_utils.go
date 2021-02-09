package utils

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/santhosh-tekuri/jsonschema"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"sigs.k8s.io/yaml"
	"strings"
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
