package executortest

import (
	"os"
	"path"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
	logger "github.com/subrahamanyam341/andes-logger-123"
	scenexec "github.com/subrahamanyam341/andes-scenario-1234/scenario/executor"
	scenio "github.com/subrahamanyam341/andes-scenario-1234/scenario/io"
)

func init() {
	_ = logger.SetLogLevel("*:NONE")
}

func getTestRoot() string {
	exePath, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	vmTestRoot := filepath.Join(exePath, "")
	return vmTestRoot
}

// ScenariosTestBuilder defines the Scenarios builder component
type ScenariosTestBuilder struct {
	t            *testing.T
	folder       string
	singleFile   string
	exclusions   []string
	currentError error
}

// ScenariosTest will create a new ScenariosTestBuilder instance
func ScenariosTest(t *testing.T) *ScenariosTestBuilder {
	return &ScenariosTestBuilder{
		t:          t,
		folder:     "",
		singleFile: "",
	}
}

// Folder sets the folder
func (mtb *ScenariosTestBuilder) Folder(folder string) *ScenariosTestBuilder {
	mtb.folder = folder
	return mtb
}

// File sets the file
func (mtb *ScenariosTestBuilder) File(fileName string) *ScenariosTestBuilder {
	mtb.singleFile = fileName
	return mtb
}

// Exclude sets the exclusion path
func (mtb *ScenariosTestBuilder) Exclude(path string) *ScenariosTestBuilder {
	mtb.exclusions = append(mtb.exclusions, path)
	return mtb
}

// Run will start the testing process
func (mtb *ScenariosTestBuilder) Run() *ScenariosTestBuilder {
	vmBuilder := &DummyVMBuilder{}
	executor := scenexec.NewScenarioExecutor(vmBuilder)
	defer executor.Close()

	runner := scenio.NewScenarioController(
		executor,
		scenio.NewDefaultFileResolver(),
		vmBuilder.GetVMType(),
	)

	if len(mtb.singleFile) > 0 {
		fullPath := path.Join(getTestRoot(), mtb.folder)
		fullPath = path.Join(fullPath, mtb.singleFile)

		mtb.currentError = runner.RunSingleJSONScenario(
			fullPath,
			scenio.DefaultRunScenarioOptions())
	} else {
		mtb.currentError = runner.RunAllJSONScenariosInDirectory(
			getTestRoot(),
			mtb.folder,
			".scen.json",
			mtb.exclusions,
			scenio.DefaultRunScenarioOptions())
	}

	return mtb
}

// CheckNoError does an assert for the containing error
func (mtb *ScenariosTestBuilder) CheckNoError() *ScenariosTestBuilder {
	if mtb.currentError != nil {
		mtb.t.Error(mtb.currentError)
	}
	return mtb
}

// RequireError does an assert for the containing error
func (mtb *ScenariosTestBuilder) RequireError(expectedErrorMsg string) *ScenariosTestBuilder {
	require.EqualError(mtb.t, mtb.currentError, expectedErrorMsg)
	return mtb
}
