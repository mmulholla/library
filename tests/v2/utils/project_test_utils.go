package utils

import (
	"errors"
	"fmt"
	schema "github.com/devfile/api/v2/pkg/apis/workspaces/v1alpha2"
	"io/ioutil"

	"github.com/google/go-cmp/cmp"
	"sigs.k8s.io/yaml"
)

// getRemotes creates and returns a map of remotes
func getRemotes() map[string]string {
	remotes := make(map[string]string)
	numRemotes := GetRandomNumber(5)
	for i := 0; i < numRemotes; i++ {
		key := GetRandomUniqueString(8,false)
		remotes[key] = GetRandomUniqueString(8,false)
		LogInfoMessage(fmt.Sprintf("Set remote key= %s, value= %s",key,remotes[key]))
	}
	return remotes
}

// AddProject adds a project of the specified type, with random attributes, to the devfile schema
func (devfile *TestDevfile) AddProject(projectType schema.ProjectSourceType) string {
	project := generateProject(projectType)
	devfile.SchemaDevFile.Projects = append(devfile.SchemaDevFile.Projects, project)
	return project.Name
}

// AddStarterProject adds a starter project of the specified type, with random attributes, to the devfile schema
func (devfile *TestDevfile) AddStarterProject(projectType schema.ProjectSourceType) string {
	project := generateStarterProject(projectType)
	devfile.SchemaDevFile.StarterProjects = append(devfile.SchemaDevFile.StarterProjects, project)
	return project.Name
}


// generateProject creates a project of a specified type in a schema project structure
func generateProject(projectType schema.ProjectSourceType) schema.Project {
	project := schema.Project{}
	project.Name = GetRandomUniqueString(GetRandomNumber(63), true)
	LogInfoMessage(fmt.Sprintf("Create Project Name: %s", project.Name))

	setProjectCommonValues(&project)

	if projectType == schema.GitProjectSourceType {
		project.Git = createGitProject()
	} else if projectType == schema.GitHubProjectSourceType {
		project.Github = createGithubProject()
	} else if projectType == schema.ZipProjectSourceType {
		project.Zip = createZipProject()
	}
	return project
}


// generateStarterProject creates a starter project of a specified type in a schema starter project structure
func generateStarterProject(projectType schema.ProjectSourceType) schema.StarterProject {
	starterProject := schema.StarterProject{}
	starterProject.Name = GetRandomUniqueString(GetRandomNumber(63), true)
	LogInfoMessage(fmt.Sprintf("Create StarterProject Name: %s", starterProject.Name))

	setStarterProjectCommonValues(&starterProject)

	if projectType == schema.GitProjectSourceType {
		starterProject.Git = createGitProject()
	} else if projectType == schema.GitHubProjectSourceType {
		starterProject.Github = createGithubProject()
	} else if projectType == schema.ZipProjectSourceType {
		starterProject.Zip = createZipProject()
	}
	return starterProject

}

// createGitProject creates a git project structure a loads it with random attributes
func createGitProject() *schema.GitProjectSource {
	project :=  schema.GitProjectSource{}
	setGitProjectValues(&project)
	return &project
}

// createGithubProject creates a github project structure a loads it with random attributes
func createGithubProject() *schema.GithubProjectSource {
	project :=  schema.GithubProjectSource{}
	setGithubProjectValues(&project)
	return &project
}

// createZipProject creates a zip project structure a loads it with random attributes
func createZipProject() *schema.ZipProjectSource {
	project :=  schema.ZipProjectSource{}
	setZipProjectValues(&project)
	return &project
}

// setProjectCommonValues sets project attributes, common to all projects, to random values.
func setProjectCommonValues(project *schema.Project) {
	if GetBinaryDecision() {
		project.ClonePath = "./" + GetRandomString(GetRandomNumber(12),false)
		LogInfoMessage(fmt.Sprintf("Set ClonePath : %s",project.ClonePath))
	}

	if GetBinaryDecision() {
		var sparseCheckoutDirs []string
		numDirs := GetRandomNumber(6)
		for i := 0; i < numDirs; i++ {
			sparseCheckoutDirs = append(sparseCheckoutDirs,GetRandomString(8,false))
			LogInfoMessage(fmt.Sprintf("Set sparseCheckoutDir : %s",sparseCheckoutDirs[i]))
		}
		project.SparseCheckoutDirs = sparseCheckoutDirs
	}
}

// setStarterProjectCommonValues sets starter project attributes, common to all starter projects, to random values.
func setStarterProjectCommonValues(project *schema.StarterProject) {
	if GetBinaryDecision() {
		numWords := GetRandomNumber(6)
		for i := 0 ; i < numWords ; i++ {
			if i > 0 {
				project.Description += " "
			}
			project.Description +=  GetRandomString(8, false)
		}
		LogInfoMessage(fmt.Sprintf("Set Description : %s",project.Description))
	}

	if GetBinaryDecision() {
		project.SubDir = GetRandomString(12,false)
		LogInfoMessage(fmt.Sprintf("Set SubDir : %s",project.SubDir))
	}

}

// setGitProjectValues randomly sets attributes for a Git project
func setGitProjectValues(gitProject *schema.GitProjectSource) {

	gitProject.Remotes = getRemotes()
	if len(gitProject.Remotes) > 1 {
		numKey := GetRandomNumber(len(gitProject.Remotes))
		for key,_ := range gitProject.Remotes {
			numKey--
			if numKey <= 0 {
				gitProject.CheckoutFrom = &schema.CheckoutFrom{}
				gitProject.CheckoutFrom.Remote = key
				gitProject.CheckoutFrom.Revision = GetRandomString(8,false)
				LogInfoMessage(fmt.Sprintf("set CheckoutFrom remote = %s, and revision = %s",gitProject.CheckoutFrom.Remote,gitProject.CheckoutFrom.Revision))
				break
			}
		}
	}
}

// setGithubProjectValues randomly sets attributes for a Github project
func setGithubProjectValues(githubProject *schema.GithubProjectSource) {

	githubProject.Remotes = getRemotes()
	if len(githubProject.Remotes) > 1 {
		numKey := GetRandomNumber(len(githubProject.Remotes))
		for key,_ := range githubProject.Remotes {
			numKey--
			if numKey <= 0 {
				githubProject.CheckoutFrom = &schema.CheckoutFrom{}
				githubProject.CheckoutFrom.Remote = key
				githubProject.CheckoutFrom.Revision = GetRandomString(8,false)
				LogInfoMessage(fmt.Sprintf("set CheckoutFrom remote = %s, and revision = %s",githubProject.CheckoutFrom.Remote,githubProject.CheckoutFrom.Revision))
				break
			}
		}
	}
}

// ssetZipProjectValues randomly sets attributes for a Zip Project
func setZipProjectValues(zipProject *schema.ZipProjectSource) {
	zipProject.Location = GetRandomString(12,false)
}

// getSchemaProject get a Project from the saved devfile schema structure
func getSchemaProject(projects []schema.Project, name string) (*schema.Project, bool) {
	found := false
	var schemaProject schema.Project
	for _, project := range projects {
		if project.Name == name {
			schemaProject = project
			found = true
			break
		}
	}
	return &schemaProject, found
}

// replaceSchemaProject replace a Project in the saved devfile schema structure
func (devfile TestDevfile) replaceSchemaProject(project schema.Project) {
	for i := 0; i < len(devfile.SchemaDevFile.Projects); i++ {
		if devfile.SchemaDevFile.Projects[i].Name == project.Name {
			devfile.SchemaDevFile.Projects[i] = project
			break
		}
	}
}

// getSchemaStarterProject gets a Starter Project from the saved devfile schema structure
func getSchemaStarterProject(starterProjects []schema.StarterProject, name string) (*schema.StarterProject, bool) {
	found := false
	var schemaStarterProject schema.StarterProject
	for _, starterProject := range starterProjects {
		if starterProject.Name == name {
			schemaStarterProject = starterProject
			found = true
			break
		}
	}
	return &schemaStarterProject, found
}

// replaceSchemaProject replaces a Project in the saved devfile schema structure
func (devfile TestDevfile) replaceSchemaStarterProject(starterProject schema.StarterProject) {
	for i := 0; i < len(devfile.SchemaDevFile.StarterProjects); i++ {
		if devfile.SchemaDevFile.StarterProjects[i].Name == starterProject.Name {
			devfile.SchemaDevFile.StarterProjects[i] = starterProject
			break
		}
	}
}

// UpdateProject randomly modifies an existing project
func (devfile *TestDevfile) UpdateProject(parserProject *schema.Project) error {

	var err error
	testProject, found := getSchemaProject(devfile.SchemaDevFile.Projects, parserProject.Name)
	if found {
		LogInfoMessage(fmt.Sprintf("Updating Project id: %s", parserProject.Name))
		setProjectCommonValues(parserProject)
		if testProject.Git != nil {
			setGitProjectValues(parserProject.Git)
		} else if testProject.Github != nil {
			setGithubProjectValues(parserProject.Github)
		} else if testProject.Zip != nil {
			setZipProjectValues(parserProject.Zip)
		}
		devfile.replaceSchemaProject(*parserProject)
	} else {
		err = errors.New(LogErrorMessage(fmt.Sprintf("Project not found in test : %s", parserProject.Name)))
	}
	return err

}

// UpdateStarterProject randomly modifies an existing starter project
func (devfile *TestDevfile) UpdateStarterProject(parserStarterProject *schema.StarterProject) error {

	var err error
	testStarterProject, found := getSchemaStarterProject(devfile.SchemaDevFile.StarterProjects, parserStarterProject.Name)
	if found {
		LogInfoMessage(fmt.Sprintf("Updating Starter Project id: %s", parserStarterProject.Name))
		setStarterProjectCommonValues(parserStarterProject)
		if testStarterProject.Git != nil {
			setGitProjectValues(parserStarterProject.Git)
		} else if testStarterProject.Github != nil {
			setGithubProjectValues(parserStarterProject.Github)
		} else if testStarterProject.Zip != nil {
			setZipProjectValues(parserStarterProject.Zip)
		}
		devfile.replaceSchemaStarterProject(*parserStarterProject)
	} else {
		err = errors.New(LogErrorMessage(fmt.Sprintf("Starter Project not found in test : %s", parserStarterProject.Name)))
	}
	return err
}

// VerifyProjects verifies projects returned by the parser are the same as those saved in the devfile schema
func (devfile TestDevfile) VerifyProjects(parserProjects []schema.Project) error {

	LogInfoMessage("Enter VerifyProjects")
	var errorString []string

	// Compare entire array of projects
	if !cmp.Equal(parserProjects, devfile.SchemaDevFile.Projects) {
		// Compare failed so compare each project to find which one(s) don't compare
		errorString = append(errorString, LogErrorMessage(fmt.Sprintf("Project array compare failed.")))
		for _, project := range parserProjects {
			if testProject, found := getSchemaProject(devfile.SchemaDevFile.Projects, project.Name); found {
				if !cmp.Equal(project, *testProject) {
					parserFilename := AddSuffixToFileName(devfile.FileName, "_"+project.Name+"_Parser")
					testFilename := AddSuffixToFileName(devfile.FileName, "_"+project.Name+"_Test")
					LogInfoMessage(fmt.Sprintf(".......marshall and write devfile %s", parserFilename))
					c, err := yaml.Marshal(project)
					if err != nil {
						errorString = append(errorString, LogErrorMessage(fmt.Sprintf(".......marshall devfile %s", parserFilename)))
					} else {
						err = ioutil.WriteFile(parserFilename, c, 0644)
						if err != nil {
							errorString = append(errorString, LogErrorMessage(fmt.Sprintf(".......write devfile %s", parserFilename)))
						}
					}
					LogInfoMessage(fmt.Sprintf(".......marshall and write devfile %s", testFilename))
					c, err = yaml.Marshal(testProject)
					if err != nil {
						errorString = append(errorString, LogErrorMessage(fmt.Sprintf(".......marshall devfile %s", testFilename)))
					} else {
						err = ioutil.WriteFile(testFilename, c, 0644)
						if err != nil {
							errorString = append(errorString, LogErrorMessage(fmt.Sprintf(".......write devfile %s", testFilename)))
						}
					}
					errorString = append(errorString, LogErrorMessage(fmt.Sprintf("Project %s did not match, see files : %s and %s", project.Name, parserFilename, testFilename)))
				} else {
					LogInfoMessage(fmt.Sprintf(" --> Project matched : %s", project.Name))
				}
			} else {
				errorString = append(errorString, LogErrorMessage(fmt.Sprintf("Project from parser not known to test - id : %s ", project.Name)))
			}
		}
		// Check test does not include projects which the parser did not return
		for _, project := range devfile.SchemaDevFile.Projects {
			if _, found := getSchemaProject(parserProjects, project.Name); !found {
				errorString = append(errorString, LogErrorMessage(fmt.Sprintf("Project from test not returned by parser : %s ", project.Name)))
			}
		}
	} else {
		LogInfoMessage(fmt.Sprintf("Project structures matched"))
	}

	var err error
	if len(errorString) > 0 {
		err = errors.New(fmt.Sprint(errorString))
	}
	return err
}

// VerifyStarterProjects verifies starter projects returned by the parser are the same as those saved in the devfile schema
func (devfile TestDevfile) VerifyStarterProjects(parserStarterProjects []schema.StarterProject) error {

	LogInfoMessage("Enter VerifyStarterProjects")
	var errorString []string

	// Compare entire array of projects
	if !cmp.Equal(parserStarterProjects, devfile.SchemaDevFile.StarterProjects) {
		// Compare failed so compare each project to find which one(s) don't compare
		errorString = append(errorString, LogErrorMessage(fmt.Sprintf("Starter Project array compare failed.")))
		for _, starterProject := range parserStarterProjects {
			if testStarterProject, found := getSchemaStarterProject(devfile.SchemaDevFile.StarterProjects, starterProject.Name); found {
				if !cmp.Equal(starterProject, *testStarterProject) {
					parserFilename := AddSuffixToFileName(devfile.FileName, "_"+starterProject.Name+"_Parser")
					testFilename := AddSuffixToFileName(devfile.FileName, "_"+starterProject.Name+"_Test")
					LogInfoMessage(fmt.Sprintf(".......marshall and write devfile %s", parserFilename))
					c, err := yaml.Marshal(starterProject)
					if err != nil {
						errorString = append(errorString, LogErrorMessage(fmt.Sprintf(".......marshall devfile %s", parserFilename)))
					} else {
						err = ioutil.WriteFile(parserFilename, c, 0644)
						if err != nil {
							errorString = append(errorString, LogErrorMessage(fmt.Sprintf(".......write devfile %s", parserFilename)))
						}
					}
					LogInfoMessage(fmt.Sprintf(".......marshall and write devfile %s", testFilename))
					c, err = yaml.Marshal(testStarterProject)
					if err != nil {
						errorString = append(errorString, LogErrorMessage(fmt.Sprintf(".......marshall devfile %s", testFilename)))
					} else {
						err = ioutil.WriteFile(testFilename, c, 0644)
						if err != nil {
							errorString = append(errorString, LogErrorMessage(fmt.Sprintf(".......write devfile %s", testFilename)))
						}
					}
					errorString = append(errorString, LogErrorMessage(fmt.Sprintf("Starter Project %s did not match, see files : %s and %s", starterProject.Name, parserFilename, testFilename)))
				} else {
					LogInfoMessage(fmt.Sprintf(" --> Starter Project matched : %s", starterProject.Name))
				}
			} else {
				errorString = append(errorString, LogErrorMessage(fmt.Sprintf("Starter Project from parser not known to test - id : %s ", starterProject.Name)))
			}
		}
		// Check test does not include projects which the parser did not return
		for _, starterProject := range devfile.SchemaDevFile.StarterProjects {
			if _, found := getSchemaStarterProject(parserStarterProjects, starterProject.Name); !found {
				errorString = append(errorString, LogErrorMessage(fmt.Sprintf("Starter Project from test not returned by parser : %s ", starterProject.Name)))
			}
		}
	} else {
		LogInfoMessage(fmt.Sprintf("Starter Project structures matched"))
	}

	var err error
	if len(errorString) > 0 {
		err = errors.New(fmt.Sprint(errorString))
	}
	return err
}