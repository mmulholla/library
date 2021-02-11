package common

import (
	"errors"
	"fmt"
	"io/ioutil"

	schema "github.com/devfile/api/v2/pkg/apis/workspaces/v1alpha2"
)

// commandAdded adds a new command to the test schema data and to the parser data
func (devfile *TestDevfile) commandAdded(command schema.Command) {
	LogInfoMessage(fmt.Sprintf("command added Id: %s", command.Id))
	devfile.SchemaDevFile.Commands = append(devfile.SchemaDevFile.Commands, command)
	if devfile.Writer != nil {
		devfile.Writer.AddCommand(command)
	}
}

// commandUpdated updates a command in the parser data
func (devfile *TestDevfile) commandUpdated(command schema.Command) {
	LogInfoMessage(fmt.Sprintf("command updated Id: %s", command.Id))
	if devfile.Writer != nil {
		devfile.Writer.UpdateCommand(command)
	}
}

// addEnv creates and returns a specifed number of env attributes in a schema structure
func addEnv(numEnv int) []schema.EnvVar {
	commandEnvs := make([]schema.EnvVar, numEnv)
	for i := 0; i < numEnv; i++ {
		commandEnvs[i].Name = "Name_" + GetRandomString(5, false)
		commandEnvs[i].Value = "Value_" + GetRandomString(5, false)
		LogInfoMessage(fmt.Sprintf("Add Env: %s", commandEnvs[i]))
	}
	return commandEnvs
}

// addAttributes creates returns a specifed number of attributes in a schema structure
func addAttributes(numAtrributes int) map[string]string {
	attributes := make(map[string]string)
	for i := 0; i < numAtrributes; i++ {
		AttributeName := "Name_" + GetRandomString(6, false)
		attributes[AttributeName] = "Value_" + GetRandomString(6, false)
		LogInfoMessage(fmt.Sprintf("Add attribute : %s = %s", AttributeName, attributes[AttributeName]))
	}
	return attributes
}

// addGroup creates and returns a group in a schema structure
func (devfile *TestDevfile) addGroup() *schema.CommandGroup {

	commandGroup := schema.CommandGroup{}
	commandGroup.Kind = GetRandomGroupKind()
	LogInfoMessage(fmt.Sprintf("group Kind: %s, default already set %t", commandGroup.Kind, devfile.GroupDefaults[commandGroup.Kind]))
	// Ensure only one and at least one of each type are labelled as default
	if !devfile.GroupDefaults[commandGroup.Kind] {
		devfile.GroupDefaults[commandGroup.Kind] = true
		commandGroup.IsDefault = true
	} else {
		commandGroup.IsDefault = false
	}
	LogInfoMessage(fmt.Sprintf("group isDefault: %t", commandGroup.IsDefault))
	return &commandGroup
}

// AddCommand creates a command of a specified type in a schema structure and pupulates it with random attributes
func (devfile *TestDevfile) AddCommand(commandType schema.CommandType) schema.Command {

	var command *schema.Command
	if commandType == schema.ExecCommandType {
		command = devfile.createExecCommand()
		devfile.setExecCommandValues(command)
	} else if commandType == schema.CompositeCommandType {
		command = devfile.createCompositeCommand()
		devfile.setCompositeCommandValues(command)
	} else if commandType == schema.ApplyCommandType {
		command = devfile.createApplyCommand()
		devfile.setApplyCommandValues(command)
	} else if commandType == schema.VscodeTaskCommandType {
		command = devfile.createVscodeTaskCommand()
		devfile.setVscodeTaskCommandValues(command)
	} else if commandType == schema.VscodeLaunchCommandType {
		command = devfile.createVscodeLaunchCommand()
		devfile.setVscodeLaunchCommandValues(command)
	}
	return *command
}


// createExecCommand creates and returns an empty exec command in a schema structure
func (devfile *TestDevfile) createExecCommand() *schema.Command {

	LogInfoMessage("Create an exec command :")
	command := schema.Command{}
	command.Id = GetRandomUniqueString(8, true)
	LogInfoMessage(fmt.Sprintf("command Id: %s", command.Id))
	command.Exec = &schema.ExecCommand{}
	devfile.commandAdded(command)
	return &command

}

// setExecCommandValues randomly sets exec command attribute to random values
func (devfile *TestDevfile) setExecCommandValues(command *schema.Command) {

	execCommand := command.Exec

	// exec command must be mentioned by a container component
	execCommand.Component = devfile.GetContainerName()

	execCommand.CommandLine = GetRandomString(4, false) + " " + GetRandomString(4, false)
	LogInfoMessage(fmt.Sprintf("....... commandLine: %s", execCommand.CommandLine))

	// If group already leave it to make sure defaults are not deleted or added
	if execCommand.Group == nil {
		if GetRandomDecision(2, 1) {
			execCommand.Group = devfile.addGroup()
		}
	}

	if GetBinaryDecision() {
		execCommand.Label = GetRandomString(12, false)
		LogInfoMessage(fmt.Sprintf("....... label: %s", execCommand.Label))
	} else {
		execCommand.Label = ""
	}

	if GetBinaryDecision() {
		execCommand.WorkingDir = "./tmp"
		LogInfoMessage(fmt.Sprintf("....... WorkingDir: %s", execCommand.WorkingDir))
	} else {
		execCommand.WorkingDir = ""
	}

	execCommand.HotReloadCapable = GetBinaryDecision()
	LogInfoMessage(fmt.Sprintf("....... HotReloadCapable: %t", execCommand.HotReloadCapable))

	if GetBinaryDecision() {
		execCommand.Env = addEnv(GetRandomNumber(4))
	} else {
		execCommand.Env = nil
	}
	devfile.commandUpdated(*command)

}

// getSchemaCommand get a specified command from the devfile schema structure
func getSchemaCommand(commands []schema.Command, id string) (*schema.Command, bool) {
	found := false
	var schemaCommand schema.Command
	for _, command := range commands {
		if command.Id == id {
			schemaCommand = command
			found = true
			break
		}
	}
	return &schemaCommand, found
}

// createCompositeCommand creates an empty composite command in a schema structure
func (devfile *TestDevfile) createCompositeCommand() *schema.Command {

	LogInfoMessage("Create a composite command :")
	command := schema.Command{}
	command.Id = GetRandomUniqueString(8, true)
	LogInfoMessage(fmt.Sprintf("command Id: %s", command.Id))
	command.Composite = &schema.CompositeCommand{}
	devfile.commandAdded(command)

	return &command
}

// setCompositeCommandValues randomly sets composite command attribute to random values
func (devfile *TestDevfile) setCompositeCommandValues(command *schema.Command) {

	compositeCommand := command.Composite
	numCommands := GetRandomNumber(3)

	for i := 0; i < numCommands; i++ {
		execCommand := devfile.AddCommand(schema.ExecCommandType)
		compositeCommand.Commands = append(compositeCommand.Commands, execCommand.Id)
		LogInfoMessage(fmt.Sprintf("....... command %d of %d : %s", i, numCommands, execCommand.Id))
	}

	// If group already exists - leave it to make sure defaults are not deleted or added
	if compositeCommand.Group == nil {
		if GetRandomDecision(2, 1) {
			compositeCommand.Group = devfile.addGroup()
		}
	}

	if GetBinaryDecision() {
		compositeCommand.Label = GetRandomString(12, false)
		LogInfoMessage(fmt.Sprintf("....... label: %s", compositeCommand.Label))
	}

	if GetBinaryDecision() {
		compositeCommand.Parallel = true
		LogInfoMessage(fmt.Sprintf("....... Parallel: %t", compositeCommand.Parallel))
	}

	devfile.commandUpdated(*command)
}

// createApplyCommand creates an apply command in a schema structure
func (devfile *TestDevfile) createApplyCommand() *schema.Command {

	LogInfoMessage("Create a apply command :")
	command := schema.Command{}
	command.Id = GetRandomUniqueString(8, true)
	LogInfoMessage(fmt.Sprintf("command Id: %s", command.Id))
	command.Apply = &schema.ApplyCommand{}
	devfile.commandAdded(command)
	return &command
}

// setApplyCommandValues randomly sets apply command attributes to random values
func (devfile *TestDevfile) setApplyCommandValues(command *schema.Command) {
	applyCommand := command.Apply

	applyCommand.Component = devfile.GetContainerName()

	if GetRandomDecision(2, 1) {
		applyCommand.Group = devfile.addGroup()
	}

	if GetBinaryDecision() {
		applyCommand.Label = GetRandomString(63, false)
		LogInfoMessage(fmt.Sprintf("....... label: %s", applyCommand.Label))
	}

	devfile.commandUpdated(*command)
}

// createVscodeLaunchCommand creates an vscodeLaunch command in a schema structure
func (devfile *TestDevfile) createVscodeLaunchCommand() *schema.Command {

	LogInfoMessage("Create a vscode command :")
	command := schema.Command{}
	command.Id = GetRandomUniqueString(8, true)
	LogInfoMessage(fmt.Sprintf("command Id: %s", command.Id))
	command.VscodeLaunch = &schema.VscodeConfigurationCommand{}
	devfile.commandAdded(command)
	return &command
}

// createVscodeTaskCommand creates an vscodeTask command in a schema structure
func (devfile *TestDevfile) createVscodeTaskCommand() *schema.Command {

	LogInfoMessage("Create a vscode command :")
	command := schema.Command{}
	command.Id = GetRandomUniqueString(8, true)
	LogInfoMessage(fmt.Sprintf("command Id: %s", command.Id))
	command.VscodeTask = &schema.VscodeConfigurationCommand{}
	devfile.commandAdded(command)
	return &command
}

// setVscodeLaunchCommandValues andomly sets VscodeLaunch command attributes to random values
func (devfile *TestDevfile) setVscodeLaunchCommandValues(command *schema.Command) {
	devfile.setVscodeCommandValues(command.VscodeLaunch)
	devfile.commandUpdated(*command)
}

// setVscodeTaskCommandValues andomly sets VscodeTask command attributes to random values
func (devfile *TestDevfile) setVscodeTaskCommandValues(command *schema.Command) {
	devfile.setVscodeCommandValues(command.VscodeTask)
	devfile.commandUpdated(*command)
}

// setVscodeCommandValues randomly sets VscodeConfigurationCommand attributes to random values
func (devfile *TestDevfile) setVscodeCommandValues(vscodeCommand *schema.VscodeConfigurationCommand) {

	if GetRandomDecision(2, 1) {
		vscodeCommand.Group = devfile.addGroup()
	}

	if GetBinaryDecision() {
		vscodeCommand.Uri = "http://" + GetRandomString(GetRandomNumber(24), false)
		LogInfoMessage(fmt.Sprintf("....... uri: %s", vscodeCommand.Uri))
		vscodeCommand.Inlined = ""
	} else {
		vscodeCommand.Inlined = GetRandomString(GetRandomNumber(12), false)
		LogInfoMessage(fmt.Sprintf("....... inlined: %s", vscodeCommand.Inlined))
		vscodeCommand.Uri = ""
	}

}
