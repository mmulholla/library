package errorsTest

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"testing"

	devfilepkg "github.com/devfile/library/pkg/devfile"
	"github.com/devfile/library/tests/v2/utils"
	"sigs.k8s.io/yaml"
)

const (
	jsonDir      = "../json/errors/"
)

type TestToRun struct {
	SchemaVersion string
	FileDirectory string
	Label         string   `json:"Label"`
	ParserMessage string   `json:"Parser_Message"`
	SchemaMessage string   `json:"JSONSchema_Message"`
	YamlContents  []string `json:"Yaml"`
}

type TestJsonFile struct {
	FileInfo      os.FileInfo
	TempDirectory string
	SchemaVersion string      `json:"SchemaVersion"`
	SchemaURL     string      `json:"SchemaURL"`
	Tests         []TestToRun `json:"Tests"`
}

// GetJsonFile returns an array of TestJsonFile objects, one for each json file containing tests
func GetJsonFiles(directory string) ([]TestJsonFile, error) {

	var jsonFiles []TestJsonFile

	// Read the content of the json directory to find test files
	files, err := ioutil.ReadDir(directory)
	if err != nil {
		utils.LogErrorMessage(fmt.Sprintf("Error finding test json files in : %s :  %v", jsonDir, err))
	} else {
		for _, testJsonFile := range files {
			// if the file ends with -test.json it can be processed
			if strings.HasSuffix(testJsonFile.Name(), "-tests.json") {
				testJson := TestJsonFile{}
				testJson.FileInfo = testJsonFile
				dirName := testJsonFile.Name()[0:strings.LastIndex(testJsonFile.Name(), ".json")]
				testJson.TempDirectory = utils.CreateTempDir(dirName)
				jsonFiles = append(jsonFiles, testJson)
			}
		}
	}
	return jsonFiles, err
}

// GetTests returns an array of TestToRun objects, one for each test contained in a json file
func (testJsonFile *TestJsonFile) GetTests() ([]TestToRun, error) {

	var err error
	if len(testJsonFile.Tests) < 1 {
		// Open the json file which defines the tests to run
		testJson, err := os.Open(filepath.Join(jsonDir, testJsonFile.FileInfo.Name()))
		if err != nil {
			utils.LogErrorMessage(fmt.Sprintf("Failed to open %s : %s", testJsonFile.FileInfo.Name(), err))
		} else {
			// Read contents of the json file which defines the tests to run
			byteValue, err := ioutil.ReadAll(testJson)
			if err != nil {
				utils.LogErrorMessage(fmt.Sprintf("Failed to read : %s : %v", testJsonFile.FileInfo.Name(), err))
			} else {
				// Unmarshall the contents of the json file which defines the tests to run for each test
				err = json.Unmarshal(byteValue, &testJsonFile)
				if err != nil {
					utils.LogErrorMessage(fmt.Sprintf("Failed to unmarshal : %s : %v", testJsonFile.FileInfo.Name(), err))
				}
			}
		}
		testJson.Close()

		for testNum, _ := range testJsonFile.Tests {
			(&testJsonFile.Tests[testNum]).FileDirectory = testJsonFile.TempDirectory
			(&testJsonFile.Tests[testNum]).SchemaVersion = testJsonFile.SchemaVersion
		}
	}
	return testJsonFile.Tests, err
}

func Test_ErrorsInDevfiles(t *testing.T) {

	parser_passTests := 0
	schema_passTests := 0
	parser_totalTests := 0
	schema_totalTests := 0

	run_parser := true
	run_schema := true

	jsonFiles, error := GetJsonFiles(jsonDir)
	if error != nil {
		t.Fatalf(utils.LogErrorMessage(fmt.Sprintf("An error occurred reading tests from %s : %v", jsonDir, error)))
	}

	for _, jsonFile := range jsonFiles {
		tests, getError := jsonFile.GetTests()
		if getError != nil {
			t.Fatalf(utils.LogErrorMessage(fmt.Sprintf("An error occurred reading test from json content from : %v", getError)))
		}
		for _, test := range tests {
			utils.LogInfoMessage("--------------------------")
			utils.LogInfoMessage(fmt.Sprintf("    Label : %s", test.Label))
			utils.LogInfoMessage(fmt.Sprintf("    Parser Message : %s", test.ParserMessage))
			utils.LogInfoMessage(fmt.Sprintf("    Schema Message : %s", test.SchemaMessage))

			for testNum, yamlContent := range test.YamlContents {
				utils.LogInfoMessage(fmt.Sprintf("    Yaml : %s", yamlContent))
				yamlFile, err := test.CreateTestYaml(yamlContent, testNum)
				if err != nil {
					t.Fatalf(utils.LogErrorMessage(fmt.Sprintf("An error occurred reading test from json content from : %v", error)))
				}
				utils.LogInfoMessage(fmt.Sprintf("    Yaml : %s", yamlFile))

				if run_parser {
					parser_totalTests++
					// Parse and validate the devfile
					_, err = devfilepkg.ParseAndValidate(yamlFile)
					if err != nil {
						if !strings.Contains(err.Error(), test.ParserMessage) {
							t.Error(utils.LogErrorMessage(fmt.Sprintf("  FAIL : parser :  %s : Did not fail as expected : %s  got : %v", yamlFile, test.ParserMessage, err)))
						} else {
							utils.LogInfoMessage(fmt.Sprintf("PASS : parser : Expected Error received : %s", err.Error()))
							parser_passTests++
						}
					} else {
						t.Error(utils.LogErrorMessage(fmt.Sprintf("  FAIL : parser : %s : devfile was valid - Expected Error not found :  %s", yamlFile, test.ParserMessage)))
					}
				}
				if run_schema {

					schema_totalTests++
					schemaFile, err := utils.GetSchema(jsonFile.SchemaURL)
					if err != nil {
						t.Errorf(utils.LogErrorMessage(fmt.Sprintf("FAIL : schema : Failed to get devfile schema : %v", err)))
					} else {
						err = schemaFile.CheckWithSchema(yamlFile, test.SchemaMessage)
						if err != nil {
							t.Errorf(utils.LogErrorMessage(fmt.Sprintf("FAIL : schema : Verification failed : %v", err)))
						} else {
							schema_passTests++
						}
					}
				}

			}
		}
	}

	if parser_totalTests > 0 {
		parser_failedTests := parser_totalTests - parser_passTests
		if parser_failedTests > 0 {
			t.Errorf(utils.LogMessage(fmt.Sprintf("PARSER TESTS OVERALL FAIL :  %d tests passed. %d tests failed.", parser_passTests, parser_failedTests)))
		} else {
			t.Log(utils.LogMessage(fmt.Sprintf("PARSER TESTS OVERALL PASS : %d tests passed.", parser_totalTests)))
		}

	}
	if schema_totalTests > 0 {
		schema_failedTests := schema_totalTests - schema_passTests
		if schema_failedTests > 0 {
			t.Errorf(utils.LogMessage(fmt.Sprintf("SCHEMA TESTS OVERALL FAIL :  %d tests passed. %d tests failed.", schema_passTests, schema_failedTests)))
		} else {
			t.Log(utils.LogMessage(fmt.Sprintf("SCHEMA TESTS OVERALL PASS : %d tests passed.", schema_totalTests)))
		}
	}

}

// CreatTestYaml creates a devfile.yaml file on disk as required by a test
func (testToRun *TestToRun) CreateTestYaml(yamlContentAsJson string, testNum int) (string, error) {

	var yamlFileName string

	fileName := fmt.Sprintf("Error_%s%d.yaml", testToRun.Label, testNum)
	yamlFileName = filepath.Join(testToRun.FileDirectory, fileName)
	// Open the file to contain the generated test yaml'

	yamlFile, err := os.OpenFile(yamlFileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		utils.LogErrorMessage(fmt.Sprintf("FAIL : Failed to open %s : %v", yamlFileName, err))
	} else {

		yamlFile.WriteString("schemaVersion: \"" + testToRun.SchemaVersion + "\"\n")

		testContentAsJSON := []byte(yamlContentAsJson)
		testContentAsYaml, err := yaml.JSONToYAML(testContentAsJSON)
		if err != nil {
			fmt.Printf("err: %v\n", err)
			return yamlFileName, err
		}

		yamlFile.Write(testContentAsYaml)
		yamlFile.WriteString("\n")
		yamlFile.Close()
	}
	return yamlFileName, nil
}
