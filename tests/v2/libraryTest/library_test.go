package api

import (
	"testing"

	schema "github.com/devfile/api/v2/pkg/apis/workspaces/v1alpha2"
	commonUtils "github.com/devfile/library/tests/v2/utils/common"
	libraryUtils "github.com/devfile/library/tests/v2/utils/library"
)

// TestContent - structure used by a test to configure the tests to run
type TestContent struct {
	CommandTypes   []schema.CommandType
	ComponentTypes []schema.ComponentType
	FileName       string
	EditContent    bool
}

func Best_ExecCommand(t *testing.T) {
	testContent := commonUtils.TestContent{}
	testContent.CommandTypes = []schema.CommandType{schema.ExecCommandType}
	testContent.EditContent = false
	testContent.FileName = commonUtils.GetDevFileName()
	libraryUtils.RunTest(testContent, t)
	//libraryUtils.RunMultiThreadTest(testContent, t)
}
func Best_ExecCommandEdit(t *testing.T) {
	testContent := commonUtils.TestContent{}
	testContent.CommandTypes = []schema.CommandType{schema.ExecCommandType}
	testContent.EditContent = true
	testContent.FileName = commonUtils.GetDevFileName()
	libraryUtils.RunTest(testContent, t)
	//libraryUtils.RunMultiThreadTest(testContent, t)
}

func Best_ApplyCommand(t *testing.T) {
	testContent := commonUtils.TestContent{}
	testContent.CommandTypes = []schema.CommandType{schema.ApplyCommandType}
	testContent.EditContent = false
	testContent.FileName = commonUtils.GetDevFileName()
	libraryUtils.RunTest(testContent, t)
	//libraryUtils.RunMultiThreadTest(testContent, t)
}

func Best_ApplyCommandEdit(t *testing.T) {
	testContent := commonUtils.TestContent{}
	testContent.CommandTypes = []schema.CommandType{schema.ApplyCommandType}
	testContent.EditContent = true
	testContent.FileName = commonUtils.GetDevFileName()
	libraryUtils.RunTest(testContent, t)
	//libraryUtils.RunMultiThreadTest(testContent, t)
}

func Best_CompositeCommand(t *testing.T) {
	testContent := commonUtils.TestContent{}
	testContent.CommandTypes = []schema.CommandType{schema.CompositeCommandType}
	testContent.EditContent = false
	testContent.FileName = commonUtils.GetDevFileName()
	libraryUtils.RunTest(testContent, t)
	//libraryUtils.RunMultiThreadTest(testContent, t)
}
func Best_CompositeCommandEdit(t *testing.T) {
	testContent := commonUtils.TestContent{}
	testContent.CommandTypes = []schema.CommandType{schema.CompositeCommandType}
	testContent.EditContent = true
	testContent.FileName = commonUtils.GetDevFileName()
	libraryUtils.RunTest(testContent, t)
	//libraryUtils.RunMultiThreadTest(testContent, t)
}

func Best_MultiCommand(t *testing.T) {
	testContent := commonUtils.TestContent{}
	testContent.CommandTypes = []schema.CommandType{schema.ExecCommandType,
		schema.CompositeCommandType,
		schema.ApplyCommandType}
	testContent.EditContent = true
	testContent.FileName = commonUtils.GetDevFileName()
	libraryUtils.RunTest(testContent, t)
	//libraryUtils.RunMultiThreadTest(testContent, t)
}

func Best_ContainerComponent(t *testing.T) {
	testContent := commonUtils.TestContent{}
	testContent.ComponentTypes = []schema.ComponentType{schema.ContainerComponentType}
	testContent.EditContent = false
	testContent.FileName = commonUtils.GetDevFileName()
	libraryUtils.RunTest(testContent, t)
	//libraryUtils.RunMultiThreadTest(testContent, t)
}

func Best_ContainerComponentEdit(t *testing.T) {
	testContent := commonUtils.TestContent{}
	testContent.ComponentTypes = []schema.ComponentType{schema.ContainerComponentType}
	testContent.EditContent = true
	testContent.FileName = commonUtils.GetDevFileName()
	libraryUtils.RunTest(testContent, t)
	//libraryUtils.RunMultiThreadTest(testContent, t)
}

func Best_VolumeComponent(t *testing.T) {
	testContent := commonUtils.TestContent{}
	testContent.ComponentTypes = []schema.ComponentType{schema.VolumeComponentType}
	testContent.EditContent = false
	testContent.FileName = commonUtils.GetDevFileName()
	libraryUtils.RunTest(testContent, t)
	//libraryUtils.RunMultiThreadTest(testContent, t)
}

func Best_VolumeComponentEdit(t *testing.T) {
	testContent := commonUtils.TestContent{}
	testContent.ComponentTypes = []schema.ComponentType{schema.VolumeComponentType}
	testContent.EditContent = true
	testContent.FileName = commonUtils.GetDevFileName()
	libraryUtils.RunTest(testContent, t)
	//libraryUtils.RunMultiThreadTest(testContent, t)
}

func Best_MultiComponent(t *testing.T) {
	testContent := commonUtils.TestContent{}
	testContent.ComponentTypes = []schema.ComponentType{schema.ContainerComponentType, schema.VolumeComponentType}
	testContent.EditContent = true
	testContent.FileName = commonUtils.GetDevFileName()
	libraryUtils.RunTest(testContent, t)
	//libraryUtils.RunMultiThreadTest(testContent, t)
}

func Best_Everything(t *testing.T) {
	testContent := commonUtils.TestContent{}
	testContent.CommandTypes = []schema.CommandType{schema.ExecCommandType, schema.CompositeCommandType, schema.ApplyCommandType}
	testContent.ComponentTypes = []schema.ComponentType{schema.ContainerComponentType, schema.VolumeComponentType}
	testContent.EditContent = true
	testContent.FileName = commonUtils.GetDevFileName()
	libraryUtils.RunTest(testContent, t)
	//libraryUtils.RunMultiThreadTest(testContent, t)
}

func Best_Projects(t *testing.T) {
	testContent := commonUtils.TestContent{}
	testContent.ProjectTypes = []schema.ProjectSourceType{
		schema.GitProjectSourceType,
		schema.GitHubProjectSourceType,
		schema.ZipProjectSourceType}
	testContent.FileName = commonUtils.GetDevFileName()
	libraryUtils.RunTest(testContent, t)
}

func Best_StarterProjects(t *testing.T) {
	testContent := commonUtils.TestContent{}
	testContent.StarterProjectTypes = []schema.ProjectSourceType{
		schema.GitProjectSourceType,
		schema.GitHubProjectSourceType,
		schema.ZipProjectSourceType}
	testContent.FileName = commonUtils.GetDevFileName()
	libraryUtils.RunTest(testContent, t)
}

func Test_Events(t *testing.T) {
	testContent := commonUtils.TestContent{}
	testContent.AddEvents = true
	testContent.FileName = commonUtils.GetDevFileName()
	libraryUtils.RunTest(testContent, t)
}
