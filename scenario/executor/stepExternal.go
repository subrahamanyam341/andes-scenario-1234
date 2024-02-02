package scenexec

import (
	scenio "github.com/subrahamanyam341/andes-scenario-1234/scenario/io"
	scenmodel "github.com/subrahamanyam341/andes-scenario-1234/scenario/model"
)

// ExecuteExternalStep executes an external step referenced by the scenario.
func (ae *ScenarioExecutor) ExecuteExternalStep(step *scenmodel.ExternalStepsStep) error {
	log.Trace("ExternalStepsStep", "path", step.Path)
	if len(step.Comment) > 0 {
		log.Trace("ExternalStepsStep", "comment", step.Comment)
	}

	fileResolverBackup := ae.fileResolver
	clonedFileResolver := ae.fileResolver.Clone()
	externalStepsRunner := scenio.NewScenarioController(ae, clonedFileResolver, ae.vmBuilder.GetVMType())

	extAbsPth := ae.fileResolver.ResolveAbsolutePath(step.Path)
	setExternalStepGasTracing(ae, step)

	err := externalStepsRunner.RunSingleJSONScenario(extAbsPth, scenio.DefaultRunScenarioOptions())
	if err != nil {
		return err
	}

	ae.fileResolver = fileResolverBackup

	return nil
}
