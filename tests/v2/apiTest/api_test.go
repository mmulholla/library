package api

import (
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/devfile/library/tests/v2/utils/common"
	schema "github.com/devfile/api/v2/pkg/apis/workspaces/v1alpha2"
)

const (
	// numThreads :  Number of threads used by multi-thread tests
	numThreads = 5
	// maxCommands : The maximum number of commands to include in a generated devfile
	maxCommands = 10
	// maxComponents : The maximum number of components to include in a generated devfile
	maxComponents = 10
)

func Test_ExecCommand(t *testing.T) {
	testContent := common.TestContent{}
	testContent.CommandTypes = []schema.CommandType{schema.ExecCommandType}
	testContent.EditContent = false
	testContent.FileName = utils.GetDevFileName()
	common.runTest(testContent, t)
}
func Test_ExecCommandEdit(t *testing.T) {
	testContent := common.TestContent{}
	testContent.CommandTypes = []schema.CommandType{schema.ExecCommandType}
	testContent.EditContent = true
	testContent.FileName = utils.GetDevFileName()
	common.runTest(testContent, t)
}

func Test_ApplyCommand(t *testing.T) {
	testContent := common.TestContent{}
	testContent.CommandTypes = []schema.CommandType{schema.ApplyCommandType}
	testContent.EditContent = false
	testContent.FileName = utils.GetDevFileName()
	common.runTest(testContent, t)
}

func Test_ApplyCommandEdit(t *testing.T) {
	testContent := common.TestContent{}
	testContent.CommandTypes = []schema.CommandType{schema.ApplyCommandType}
	testContent.EditContent = true
	testContent.FileName = utils.GetDevFileName()
	common.runTest(testContent, t)
}

func Test_VscodeLaunchCommand(t *testing.T) {
	testContent := common.TestContent{}
	testContent.CommandTypes = []schema.CommandType{schema.VscodeLaunchCommandType}
	testContent.EditContent = false
	testContent.FileName = utils.GetDevFileName()
	common.runTest(testContent, t)
}

func Test_VscodeLaunchCommandEdit(t *testing.T) {
	testContent := common.TestContent{}
	testContent.CommandTypes = []schema.CommandType{schema.VscodeLaunchCommandType}
	testContent.EditContent = true
	testContent.FileName = utils.GetDevFileName()
	common.runTest(testContent, t)
}

func Test_VscodeTaskCommand(t *testing.T) {
	testContent := common.TestContent{}
	testContent.CommandTypes = []schema.CommandType{schema.VscodeTaskCommandType}
	testContent.EditContent = false
	testContent.FileName = utils.GetDevFileName()
	common.runTest(testContent, t)
}

func Test_VscodeTaskCommandEdit(t *testing.T) {
	testContent := common.TestContent{}
	testContent.CommandTypes = []schema.CommandType{schema.VscodeTaskCommandType}
	testContent.EditContent = true
	testContent.FileName = utils.GetDevFileName()
	common.runTest(testContent, t)
}

func Test_CompositeCommand(t *testing.T) {
	testContent := common.TestContent{}
	testContent.CommandTypes = []schema.CommandType{schema.CompositeCommandType}
	testContent.EditContent = false
	testContent.FileName = utils.GetDevFileName()
	common.runTest(testContent, t)
}

func Test_CompositeCommandEdit(t *testing.T) {
	testContent := common.TestContent{}
	testContent.CommandTypes = []schema.CommandType{schema.CompositeCommandType}
	testContent.EditContent = true
	testContent.FileName = utils.GetDevFileName()
	common.runTest(testContent, t)
}

func Test_MultiCommand(t *testing.T) {
	testContent := common.TestContent{}
	testContent.CommandTypes = []schema.CommandType{schema.ExecCommandType,
		schema.CompositeCommandType}
	testContent.EditContent = true
	testContent.FileName = utils.GetDevFileName()
	common.runTest(testContent, t)
}

func Test_ContainerComponent(t *testing.T) {
	testContent := common.TestContent{}
	testContent.ComponentTypes = []schema.ComponentType{schema.ContainerComponentType}
	testContent.EditContent = false
	testContent.FileName = utils.GetDevFileName()
	common.runTest(testContent, t)
}

func Test_ContainerComponentEdit(t *testing.T) {
	testContent := common.TestContent{}
	testContent.ComponentTypes = []schema.ComponentType{schema.ContainerComponentType}
	testContent.EditContent = true
	testContent.FileName = utils.GetDevFileName()
	common.runTest(testContent, t)
}

func Test_VolumeComponent(t *testing.T) {
	testContent := common.TestContent{}
	testContent.ComponentTypes = []schema.ComponentType{schema.VolumeComponentType}
	testContent.EditContent = false
	testContent.FileName = utils.GetDevFileName()
	common.runTest(testContent, t)
}

func Test_VolumeComponentEdit(t *testing.T) {
	testContent := common.TestContent{}
	testContent.ComponentTypes = []schema.ComponentType{schema.VolumeComponentType}
	testContent.EditContent = true
	testContent.FileName = utils.GetDevFileName()
	common.runTest(testContent, t)
}

func Test_MultiComponent(t *testing.T) {
	testContent := common.TestContent{}
	testContent.ComponentTypes = []schema.ComponentType{schema.ContainerComponentType, schema.VolumeComponentType}
	testContent.EditContent = true
	testContent.FileName = utils.GetDevFileName()
	common.runTest(testContent, t)
}

func Test_Everything(t *testing.T) {
	testContent := common.TestContent{}
	testContent.CommandTypes = []schema.CommandType{schema.ExecCommandType, schema.CompositeCommandType}
	testContent.ComponentTypes = []schema.ComponentType{schema.ContainerComponentType, schema.VolumeComponentType}
	testContent.EditContent = true
	testContent.FileName = utils.GetDevFileName()
	common.runTest(testContent, t)
}

