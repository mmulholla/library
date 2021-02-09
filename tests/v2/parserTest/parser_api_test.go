package parserTest

import (
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/devfile/library/tests/v2/utils"

	schema "github.com/devfile/api/v2/pkg/apis/workspaces/v1alpha2"
)

const (
	// numThreads :  Number of threads used by multi-thread tests
	numThreads = 5
	// maxCommands : The maximum number of commands to include in a generated devfile
	maxCommands = 10
	// maxComponents : The maximum number of components to include in a generated devfile
	maxComponents = 10
	// maxProjects : The maximum number of projects to include in a generated devfile
	maxProjects = 10
	// maxProjects : The maximum number of starter projects to include in a generated devfile
	maxStarterProjects = 10
)

// TestContent - structure used by a test to configure the tests to run
type TestContent struct {
	CommandTypes        []schema.CommandType
	ComponentTypes      []schema.ComponentType
	ProjectTypes        []schema.ProjectSourceType
	StarterProjectTypes []schema.ProjectSourceType
	FileName            string
	CreateWithParser    bool
	EditContent         bool
}

func Test_ExecCommand(t *testing.T) {
	testContent := TestContent{}
	testContent.CommandTypes = []schema.CommandType{schema.ExecCommandType}
	testContent.CreateWithParser = false
	testContent.EditContent = false
	testContent.FileName = utils.GetDevFileName()
	runTest(testContent, t)
	runMultiThreadTest(testContent, t)
}
func Test_ExecCommandEdit(t *testing.T) {
	testContent := TestContent{}
	testContent.CommandTypes = []schema.CommandType{schema.ExecCommandType}
	testContent.CreateWithParser = false
	testContent.EditContent = true
	testContent.FileName = utils.GetDevFileName()
	runTest(testContent, t)
	runMultiThreadTest(testContent, t)
}

func Test_ExecCommandParserCreate(t *testing.T) {
	testContent := TestContent{}
	testContent.CommandTypes = []schema.CommandType{schema.ExecCommandType}
	testContent.CreateWithParser = true
	testContent.EditContent = false
	testContent.FileName = utils.GetDevFileName()
	runTest(testContent, t)
	runMultiThreadTest(testContent, t)
}

func Test_ExecCommandEditParserCreate(t *testing.T) {
	testContent := TestContent{}
	testContent.CommandTypes = []schema.CommandType{schema.ExecCommandType}
	testContent.CreateWithParser = true
	testContent.EditContent = true
	testContent.FileName = utils.GetDevFileName()
	runTest(testContent, t)
	runMultiThreadTest(testContent, t)
}

func Test_CompositeCommand(t *testing.T) {
	testContent := TestContent{}
	testContent.CommandTypes = []schema.CommandType{schema.CompositeCommandType}
	testContent.CreateWithParser = false
	testContent.EditContent = false
	testContent.FileName = utils.GetDevFileName()
	runTest(testContent, t)
	runMultiThreadTest(testContent, t)
}
func Test_CompositeCommandEdit(t *testing.T) {
	testContent := TestContent{}
	testContent.CommandTypes = []schema.CommandType{schema.CompositeCommandType}
	testContent.CreateWithParser = false
	testContent.EditContent = true
	testContent.FileName = utils.GetDevFileName()
	runTest(testContent, t)
	runMultiThreadTest(testContent, t)
}

func Test_CompositeCommandParserCreate(t *testing.T) {
	testContent := TestContent{}
	testContent.CommandTypes = []schema.CommandType{schema.CompositeCommandType}
	testContent.CreateWithParser = true
	testContent.EditContent = false
	testContent.FileName = utils.GetDevFileName()
	runTest(testContent, t)
	runMultiThreadTest(testContent, t)
}

func Test_CompositeCommandEditParserCreate(t *testing.T) {
	testContent := TestContent{}
	testContent.CommandTypes = []schema.CommandType{schema.CompositeCommandType}
	testContent.CreateWithParser = true
	testContent.EditContent = true
	testContent.FileName = utils.GetDevFileName()
	runTest(testContent, t)
	runMultiThreadTest(testContent, t)
}

func Test_MultiCommand(t *testing.T) {
	testContent := TestContent{}
	testContent.CommandTypes = []schema.CommandType{schema.ExecCommandType, schema.CompositeCommandType}
	testContent.CreateWithParser = true
	testContent.EditContent = true
	testContent.FileName = utils.GetDevFileName()
	runTest(testContent, t)
	runMultiThreadTest(testContent, t)
}

func Test_ContainerComponent(t *testing.T) {
	testContent := TestContent{}
	testContent.ComponentTypes = []schema.ComponentType{schema.ContainerComponentType}
	testContent.CreateWithParser = false
	testContent.EditContent = false
	testContent.FileName = utils.GetDevFileName()
	runTest(testContent, t)
	runMultiThreadTest(testContent, t)
}

func Test_ContainerComponentEdit(t *testing.T) {
	testContent := TestContent{}
	testContent.ComponentTypes = []schema.ComponentType{schema.ContainerComponentType}
	testContent.CreateWithParser = false
	testContent.EditContent = true
	testContent.FileName = utils.GetDevFileName()
	runTest(testContent, t)
	runMultiThreadTest(testContent, t)
}

func Test_ContainerComponentCreateWithParser(t *testing.T) {
	testContent := TestContent{}
	testContent.ComponentTypes = []schema.ComponentType{schema.ContainerComponentType}
	testContent.CreateWithParser = true
	testContent.EditContent = false
	testContent.FileName = utils.GetDevFileName()
	runTest(testContent, t)
	runMultiThreadTest(testContent, t)
}

func Test_ContainerComponentEditCreateWithParser(t *testing.T) {
	testContent := TestContent{}
	testContent.ComponentTypes = []schema.ComponentType{schema.ContainerComponentType}
	testContent.CreateWithParser = true
	testContent.EditContent = true
	testContent.FileName = utils.GetDevFileName()
	runTest(testContent, t)
	runMultiThreadTest(testContent, t)
}

func Test_VolumeComponent(t *testing.T) {
	testContent := TestContent{}
	testContent.ComponentTypes = []schema.ComponentType{schema.VolumeComponentType}
	testContent.CreateWithParser = false
	testContent.EditContent = false
	testContent.FileName = utils.GetDevFileName()
	runTest(testContent, t)
	runMultiThreadTest(testContent, t)
}

func Test_VolumeComponentEdit(t *testing.T) {
	testContent := TestContent{}
	testContent.ComponentTypes = []schema.ComponentType{schema.VolumeComponentType}
	testContent.CreateWithParser = false
	testContent.EditContent = true
	testContent.FileName = utils.GetDevFileName()
	runTest(testContent, t)
	runMultiThreadTest(testContent, t)
}

func Test_VolumeComponentCreateWithParser(t *testing.T) {
	testContent := TestContent{}
	testContent.ComponentTypes = []schema.ComponentType{schema.VolumeComponentType}
	testContent.CreateWithParser = true
	testContent.EditContent = false
	testContent.FileName = utils.GetDevFileName()
	runTest(testContent, t)
	runMultiThreadTest(testContent, t)
}

func Test_VolumeComponentEditCreateWithParser(t *testing.T) {
	testContent := TestContent{}
	testContent.ComponentTypes = []schema.ComponentType{schema.VolumeComponentType}
	testContent.CreateWithParser = true
	testContent.EditContent = true
	testContent.FileName = utils.GetDevFileName()
	runTest(testContent, t)
	runMultiThreadTest(testContent, t)
}

func Test_MultiComponent(t *testing.T) {
	testContent := TestContent{}
	testContent.ComponentTypes = []schema.ComponentType{schema.ContainerComponentType, schema.VolumeComponentType}
	testContent.CreateWithParser = true
	testContent.EditContent = true
	testContent.FileName = utils.GetDevFileName()
	runTest(testContent, t)
	runMultiThreadTest(testContent, t)
}

func Test_GitProject(t *testing.T) {
	testContent := TestContent{}
	testContent.ProjectTypes = []schema.ProjectSourceType{schema.GitProjectSourceType}
	testContent.CreateWithParser = false
	testContent.EditContent = false
	testContent.FileName = utils.GetDevFileName()
	runTest(testContent, t)
	runMultiThreadTest(testContent, t)
}

func Test_GitProjectEdit(t *testing.T) {
	testContent := TestContent{}
	testContent.ProjectTypes = []schema.ProjectSourceType{schema.GitProjectSourceType}
	testContent.CreateWithParser = false
	testContent.EditContent = true
	testContent.FileName = utils.GetDevFileName()
	runTest(testContent, t)
	runMultiThreadTest(testContent, t)
}

func Test_GitProjectCreateWithParser(t *testing.T) {
	testContent := TestContent{}
	testContent.ProjectTypes = []schema.ProjectSourceType{schema.GitProjectSourceType}
	testContent.CreateWithParser = true
	testContent.EditContent = false
	testContent.FileName = utils.GetDevFileName()
	runTest(testContent, t)
	runMultiThreadTest(testContent, t)
}

func Test_GitProjectCreateWithParserEdit(t *testing.T) {
	testContent := TestContent{}
	testContent.ProjectTypes = []schema.ProjectSourceType{schema.GitProjectSourceType}
	testContent.CreateWithParser = true
	testContent.EditContent = true
	testContent.FileName = utils.GetDevFileName()
	runTest(testContent, t)
	runMultiThreadTest(testContent, t)
}

func Test_GithubProject(t *testing.T) {
	testContent := TestContent{}
	testContent.ProjectTypes = []schema.ProjectSourceType{schema.GitHubProjectSourceType}
	testContent.CreateWithParser = false
	testContent.EditContent = false
	testContent.FileName = utils.GetDevFileName()
	runTest(testContent, t)
	runMultiThreadTest(testContent, t)
}

func Test_GithubProjectEdit(t *testing.T) {
	testContent := TestContent{}
	testContent.ProjectTypes = []schema.ProjectSourceType{schema.GitHubProjectSourceType}
	testContent.CreateWithParser = false
	testContent.EditContent = true
	testContent.FileName = utils.GetDevFileName()
	runTest(testContent, t)
	runMultiThreadTest(testContent, t)
}

func Test_GithubProjectCreateWithParser(t *testing.T) {
	testContent := TestContent{}
	testContent.ProjectTypes = []schema.ProjectSourceType{schema.GitHubProjectSourceType}
	testContent.CreateWithParser = true
	testContent.EditContent = false
	testContent.FileName = utils.GetDevFileName()
	runTest(testContent, t)
	runMultiThreadTest(testContent, t)
}

func Test_GithubProjectCreateWithParserEdit(t *testing.T) {
	testContent := TestContent{}
	testContent.ProjectTypes = []schema.ProjectSourceType{schema.GitHubProjectSourceType}
	testContent.CreateWithParser = true
	testContent.EditContent = true
	testContent.FileName = utils.GetDevFileName()
	runTest(testContent, t)
	runMultiThreadTest(testContent, t)
}

func Test_ZipProject(t *testing.T) {
	testContent := TestContent{}
	testContent.ProjectTypes = []schema.ProjectSourceType{schema.ZipProjectSourceType}
	testContent.CreateWithParser = false
	testContent.EditContent = false
	testContent.FileName = utils.GetDevFileName()
	runTest(testContent, t)
	runMultiThreadTest(testContent, t)
}

func Test_ZipProjectEdit(t *testing.T) {
	testContent := TestContent{}
	testContent.ProjectTypes = []schema.ProjectSourceType{schema.ZipProjectSourceType}
	testContent.CreateWithParser = false
	testContent.EditContent = true
	testContent.FileName = utils.GetDevFileName()
	runTest(testContent, t)
	runMultiThreadTest(testContent, t)
}

func Test_ZipProjectCreateWithParser(t *testing.T) {
	testContent := TestContent{}
	testContent.ProjectTypes = []schema.ProjectSourceType{schema.ZipProjectSourceType}
	testContent.CreateWithParser = true
	testContent.EditContent = false
	testContent.FileName = utils.GetDevFileName()
	runTest(testContent, t)
	runMultiThreadTest(testContent, t)
}

func Test_ZipProjectCreateWithParserEdit(t *testing.T) {
	testContent := TestContent{}
	testContent.ProjectTypes = []schema.ProjectSourceType{schema.ZipProjectSourceType}
	testContent.CreateWithParser = true
	testContent.EditContent = true
	testContent.FileName = utils.GetDevFileName()
	runTest(testContent, t)
	runMultiThreadTest(testContent, t)
}

func Test_ProjectAll(t *testing.T) {
	testContent := TestContent{}
	testContent.ProjectTypes = []schema.ProjectSourceType{schema.GitProjectSourceType, schema.GitHubProjectSourceType, schema.ZipProjectSourceType}
	testContent.CreateWithParser = true
	testContent.EditContent = true
	testContent.FileName = utils.GetDevFileName()
	runTest(testContent, t)
	runMultiThreadTest(testContent, t)
}

func Test_GitStarterProject(t *testing.T) {
	testContent := TestContent{}
	testContent.StarterProjectTypes = []schema.ProjectSourceType{schema.GitProjectSourceType}
	testContent.CreateWithParser = false
	testContent.EditContent = false
	testContent.FileName = utils.GetDevFileName()
	runTest(testContent, t)
	runMultiThreadTest(testContent, t)
}

func Test_GitStarterProjectEdit(t *testing.T) {
	testContent := TestContent{}
	testContent.StarterProjectTypes = []schema.ProjectSourceType{schema.GitProjectSourceType}
	testContent.CreateWithParser = false
	testContent.EditContent = true
	testContent.FileName = utils.GetDevFileName()
	runTest(testContent, t)
	runMultiThreadTest(testContent, t)
}

func Test_GitStarterProjectCreateWithParser(t *testing.T) {
	testContent := TestContent{}
	testContent.StarterProjectTypes = []schema.ProjectSourceType{schema.GitProjectSourceType}
	testContent.CreateWithParser = true
	testContent.EditContent = false
	testContent.FileName = utils.GetDevFileName()
	runTest(testContent, t)
	runMultiThreadTest(testContent, t)
}

func Test_GitStarterProjectCreateWithParserEdit(t *testing.T) {
	testContent := TestContent{}
	testContent.StarterProjectTypes = []schema.ProjectSourceType{schema.GitProjectSourceType}
	testContent.CreateWithParser = true
	testContent.EditContent = true
	testContent.FileName = utils.GetDevFileName()
	runTest(testContent, t)
	runMultiThreadTest(testContent, t)
}

func Test_GithubStarterProject(t *testing.T) {
	testContent := TestContent{}
	testContent.StarterProjectTypes = []schema.ProjectSourceType{schema.GitHubProjectSourceType}
	testContent.CreateWithParser = false
	testContent.EditContent = false
	testContent.FileName = utils.GetDevFileName()
	runTest(testContent, t)
	runMultiThreadTest(testContent, t)
}

func Test_GithubStarterProjectEdit(t *testing.T) {
	testContent := TestContent{}
	testContent.StarterProjectTypes = []schema.ProjectSourceType{schema.GitHubProjectSourceType}
	testContent.CreateWithParser = false
	testContent.EditContent = true
	testContent.FileName = utils.GetDevFileName()
	runTest(testContent, t)
	runMultiThreadTest(testContent, t)
}

func Test_GithubStarterProjectCreateWithParser(t *testing.T) {
	testContent := TestContent{}
	testContent.StarterProjectTypes = []schema.ProjectSourceType{schema.GitHubProjectSourceType}
	testContent.CreateWithParser = true
	testContent.EditContent = false
	testContent.FileName = utils.GetDevFileName()
	runTest(testContent, t)
	runMultiThreadTest(testContent, t)
}

func Test_GithubStarterProjectCreateWithParserEdit(t *testing.T) {
	testContent := TestContent{}
	testContent.StarterProjectTypes = []schema.ProjectSourceType{schema.GitHubProjectSourceType}
	testContent.CreateWithParser = true
	testContent.EditContent = true
	testContent.FileName = utils.GetDevFileName()
	runTest(testContent, t)
	runMultiThreadTest(testContent, t)
}

func Test_ZipStarterProject(t *testing.T) {
	testContent := TestContent{}
	testContent.StarterProjectTypes = []schema.ProjectSourceType{schema.ZipProjectSourceType}
	testContent.CreateWithParser = false
	testContent.EditContent = false
	testContent.FileName = utils.GetDevFileName()
	runTest(testContent, t)
	runMultiThreadTest(testContent, t)
}

func Test_ZipStarterProjectEdit(t *testing.T) {
	testContent := TestContent{}
	testContent.StarterProjectTypes = []schema.ProjectSourceType{schema.ZipProjectSourceType}
	testContent.CreateWithParser = false
	testContent.EditContent = true
	testContent.FileName = utils.GetDevFileName()
	runTest(testContent, t)
	runMultiThreadTest(testContent, t)
}

func Test_ZipStarterProjectCreateWithParser(t *testing.T) {
	testContent := TestContent{}
	testContent.StarterProjectTypes = []schema.ProjectSourceType{schema.ZipProjectSourceType}
	testContent.CreateWithParser = true
	testContent.EditContent = false
	testContent.FileName = utils.GetDevFileName()
	runTest(testContent, t)
	runMultiThreadTest(testContent, t)
}

func Test_ZipStarterProjectCreateWithParserEdit(t *testing.T) {
	testContent := TestContent{}
	testContent.StarterProjectTypes = []schema.ProjectSourceType{schema.ZipProjectSourceType}
	testContent.CreateWithParser = true
	testContent.EditContent = true
	testContent.FileName = utils.GetDevFileName()
	runTest(testContent, t)
	runMultiThreadTest(testContent, t)
}

func Test_StarterProjectAll(t *testing.T) {
	testContent := TestContent{}
	testContent.StarterProjectTypes = []schema.ProjectSourceType{schema.GitProjectSourceType, schema.GitHubProjectSourceType, schema.ZipProjectSourceType}
	testContent.CreateWithParser = true
	testContent.EditContent = true
	testContent.FileName = utils.GetDevFileName()
	runTest(testContent, t)
	runMultiThreadTest(testContent, t)
}
func Test_Everything(t *testing.T) {
	testContent := TestContent{}
	testContent.CommandTypes = []schema.CommandType{schema.ExecCommandType, schema.CompositeCommandType}
	testContent.ComponentTypes = []schema.ComponentType{schema.ContainerComponentType, schema.VolumeComponentType}
	testContent.ProjectTypes = []schema.ProjectSourceType{schema.GitProjectSourceType, schema.GitHubProjectSourceType, schema.ZipProjectSourceType}
	testContent.StarterProjectTypes = []schema.ProjectSourceType{schema.GitProjectSourceType, schema.GitHubProjectSourceType, schema.ZipProjectSourceType}
	testContent.CreateWithParser = true
	testContent.EditContent = true
	testContent.FileName = utils.GetDevFileName()
	runTest(testContent, t)
	runMultiThreadTest(testContent, t)
}

// runMultiThreadTest : Runs the same test on multiple threads, the test is based on the content of the specified TestContent
func runMultiThreadTest(testContent TestContent, t *testing.T) {

	utils.LogMessage(fmt.Sprintf("\n\nStart Threaded test for %s", testContent.FileName))

	devfileName := testContent.FileName
	var i int
	for i = 1; i < numThreads; i++ {
		testContent.FileName = utils.AddSuffixToFileName(devfileName, strconv.Itoa(i))
		go runTest(testContent, t)
	}
	testContent.FileName = utils.AddSuffixToFileName(devfileName, strconv.Itoa(i))
	runTest(testContent, t)

	utils.LogMessage(fmt.Sprintf("Sleep 2 seconds to allow all threads to complete : %s", devfileName))
	time.Sleep(2 * time.Second)
	utils.LogMessage(fmt.Sprintf("Sleep complete : %s", devfileName))

}

// runTest : Runs a test beased on the content of the specified TestContent
func runTest(testContent TestContent, t *testing.T) {

	utils.LogMessage(fmt.Sprintf("Start test for %s", testContent.FileName))
	testDevfile := utils.GetDevfile(testContent.FileName)

	if len(testContent.CommandTypes) > 0 {
		numCommands := utils.GetRandomNumber(maxCommands)
		for i := 0; i < numCommands; i++ {
			commandIndex := utils.GetRandomNumber(len(testContent.CommandTypes))
			testDevfile.AddCommand(testContent.CommandTypes[commandIndex-1])
		}
	}

	if len(testContent.ComponentTypes) > 0 {
		numComponents := utils.GetRandomNumber(maxComponents)
		for i := 0; i < numComponents; i++ {
			componentIndex := utils.GetRandomNumber(len(testContent.ComponentTypes))
			testDevfile.AddComponent(testContent.ComponentTypes[componentIndex-1])
		}
	}

	if len(testContent.ProjectTypes) > 0 {
		numProjects := utils.GetRandomNumber(maxProjects)
		for i := 0; i < numProjects; i++ {
			projectIndex := utils.GetRandomNumber(len(testContent.ProjectTypes))
			testDevfile.AddProject(testContent.ProjectTypes[projectIndex-1])
		}
	}

	if len(testContent.StarterProjectTypes) > 0 {
		numProjects := utils.GetRandomNumber(maxStarterProjects)
		for i := 0; i < numProjects; i++ {
			projectIndex := utils.GetRandomNumber(len(testContent.StarterProjectTypes))
			testDevfile.AddStarterProject(testContent.StarterProjectTypes[projectIndex-1])
		}
	}

	err := testDevfile.CreateDevfile(testContent.CreateWithParser)
	if err != nil {
		t.Fatalf(utils.LogErrorMessage(fmt.Sprintf("ERROR creating devfile :  %s : %v", testContent.FileName, err)))
	}

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
		if len(testContent.ProjectTypes) > 0 {
			err = testDevfile.EditProjects()
			if err != nil {
				t.Fatalf(utils.LogErrorMessage(fmt.Sprintf("ERROR editing projects :  %s : %v", testContent.FileName, err)))
			}
		}
		if len(testContent.StarterProjectTypes) > 0 {
			err = testDevfile.EditStarterProjects()
			if err != nil {
				t.Fatalf(utils.LogErrorMessage(fmt.Sprintf("ERROR editing starter projects :  %s : %v", testContent.FileName, err)))
			}
		}

	}

	err = testDevfile.Verify()
	if err != nil {
		t.Fatalf(utils.LogErrorMessage(fmt.Sprintf("ERROR verifying devfile content : %s : %v", testContent.FileName, err)))
	}

}
