package main

import (
	"os"
	"strings"

	"github.com/envoyproxy/protoc-gen-validate/validate"
	pgs "github.com/lyft/protoc-gen-star"
	pgsgo "github.com/lyft/protoc-gen-star/lang/go"
	"github.com/intrinsec/protoc-gen-checker/checker"
)

// CheckerModule adds Checker methods on PB
type CheckerModule struct {
	*pgs.ModuleBase
	ctx               pgsgo.Context
	strict            bool
	missingValidation bool
}

// Checker returns an initialized CheckerPlugin
func Checker() *CheckerModule {
	return &CheckerModule{
		ModuleBase:        &pgs.ModuleBase{},
		missingValidation: false,
	}
}

// InitContext populates the module with needed context and fields
func (c *CheckerModule) InitContext(b pgs.BuildContext) {
	b.Debug("InitContext")
	c.ModuleBase.InitContext(b)
	c.ctx = pgsgo.InitContext(b.Parameters())
}

// Name satisfies the generator.Plugin interface.
func (c *CheckerModule) Name() string { return "Checker" }

// Execute generates checking code for files
func (c *CheckerModule) Execute(targets map[string]pgs.File, pkgs map[string]pgs.Package) []pgs.Artifact {
	c.Debug("Execute")

	if ok, _ := c.Parameters().Bool("strict"); ok {
		c.strict = true
	}

	for _, t := range targets {
		v := initCheckerVisitor(c)
		c.CheckErr(pgs.Walk(v, t), "unable to check validation")
		if v.missingValidation {
			c.missingValidation = true
		}
	}

	return nil
}

func (c *CheckerModule) ExitCheck() {
	if c.missingValidation && c.strict {
		c.Log("Checker strict mode enabled. Encountered errors or file/message/field not respecting the rules. Check messages above.")
		os.Exit(1)
	}
}

type checkVisitor struct {
	pgs.Visitor
	pgs.DebuggerCommon
	missingValidation bool
}

func initCheckerVisitor(d pgs.DebuggerCommon) *checkVisitor {
	v := &checkVisitor{DebuggerCommon: d}
	v.Visitor = pgs.PassThroughVisitor(v)
	return v
}

func (v *checkVisitor) VisitFile(f pgs.File) (pgs.Visitor, error) {
	var disableValidateOnFile bool

	ok, err := f.Extension(checker.E_DisableFileValidate, &disableValidateOnFile)
	disabled, reason := v.isValidationDisabled(ok, disableValidateOnFile, f, err)
	if disabled {
		v.Debugf("%v:%d: validation disabled on whole file. Reason: %v",
			f.File().Name(),
			f.SourceCodeInfo().Location().Span[0]+1,
			reason,
		)
		return nil, nil
	}

	return v, nil
}

func (v *checkVisitor) VisitMessage(m pgs.Message) (pgs.Visitor, error) {
	// Protoc-gen-validate propose two ways to indicates a message should not be validated.
	// See https://github.com/envoyproxy/protoc-gen-validate#message-global
	var disableValidateOnMessage bool
	ok, err := m.Extension(validate.E_Disabled, &disableValidateOnMessage)
	disabled, disabledReason := v.isValidationDisabled(ok, disableValidateOnMessage, m, err)
	ok, err = m.Extension(validate.E_Ignored, &disableValidateOnMessage)
	ignored, ignoredReason := v.isValidationDisabled(ok, disableValidateOnMessage, m, err)

	var reason string

	if disabled {
		reason = disabledReason
	} else if ignored {
		reason = ignoredReason
	}

	if disabled || ignored {
		v.Debugf("%v:%d: validation disabled for message '%v'. Reason: %v",
			m.File().Name(),
			m.SourceCodeInfo().Location().Span[0]+1,
			m.Name(),
			reason,
		)
		return nil, nil
	}

	return v, nil
}

func (v *checkVisitor) VisitField(f pgs.Field) (pgs.Visitor, error) {
	var validateRulesOnField validate.FieldRules
	ok, err := f.Extension(validate.E_Rules, &validateRulesOnField)
	if ok && err == nil {
		v.Debugf("%v:%d: validation rules defined for '%v' in message %v",
			f.File().Name(),
			f.SourceCodeInfo().Location().Span[0]+1,
			f.Name(),
			f.Message().Name(),
		)

		return nil, nil
	}

	var disableValidateOnField bool
	ok, err = f.Extension(checker.E_DisableFieldValidate, &disableValidateOnField)
	disabled, reason := v.isValidationDisabled(ok, disableValidateOnField, f, err)
	if disabled {
		v.Debugf("%v:%d: validation disabled for field '%v' in message '%v'. Reason: %v",
			f.File().Name(),
			f.SourceCodeInfo().Location().Span[0]+1,
			f.Name(),
			f.Message().Name(),
			reason,
		)

		return nil, nil
	}

	v.Logf("%v:%d: no validation on field '%v' in message %v",
		f.File().Name(),
		f.SourceCodeInfo().Location().Span[0]+1,
		f.Name(),
		f.Message().Name(),
	)
	v.missingValidation = true

	return nil, nil
}

// isValidationDisabled check if the validation is correctly disabled on given entity
// Disabled validation without a proper reason is not considered valid.
func (v *checkVisitor) isValidationDisabled(ok bool, disabled bool, e pgs.Entity, err error) (bool, string) {
	if !ok || err != nil || !disabled {
		return false, ""
	}

	reason := getNoValidationReason(e.SourceCodeInfo().LeadingComments())
	if len(reason) == 0 {
		v.Logf("%v:%d: no reason given for disabled validation on %v",
			e.File().Name(),
			e.SourceCodeInfo().Location().Span[0]+1,
			e.Name(),
		)
		v.missingValidation = true

		return false, reason
	}

	return true, reason
}

const noValidationMarker = " No Validation Reason: "

// getNoValidationReason extract the reason provided for the lack of validation
func getNoValidationReason(comments string) string {
	for _, comment := range strings.Split(comments, "\n") {
		if strings.HasPrefix(comment, noValidationMarker) {
			return strings.TrimPrefix(comment, noValidationMarker)
		}
	}

	return ""
}
