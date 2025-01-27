package run

import (
	"context"
	"github.com/kurtosis-tech/kurtosis/cli/cli/command_framework/lowlevel/args"
	"github.com/kurtosis-tech/kurtosis/cli/cli/command_framework/lowlevel/flags"
	"github.com/stretchr/testify/require"
	"testing"
)

const testScriptArg = "."

var testCtx context.Context = nil
var testParsedFlags *flags.ParsedFlags = nil

func TestValidateArgsJson_valid(t *testing.T) {
	inputArgs := `{"hello": "world!"}`
	parsedArgs, err := args.ParseArgsForValidation(StarlarkRunCmd.Args, []string{testScriptArg, inputArgs})
	require.Nil(t, err)
	require.NotNil(t, parsedArgs)

	err = validatePackageArgs(testCtx, testParsedFlags, parsedArgs)
	require.Nil(t, err)
}

func TestValidateArgsYaml_valid(t *testing.T) {
	inputArgs := `hello: world`
	parsedArgs, err := args.ParseArgsForValidation(StarlarkRunCmd.Args, []string{testScriptArg, inputArgs})
	require.Nil(t, err)
	require.NotNil(t, parsedArgs)

	err = validatePackageArgs(testCtx, testParsedFlags, parsedArgs)
	require.Nil(t, err)
}

func TestValidateArgs_invalid(t *testing.T) {
	inputArgs := `"hello" - "world!"` // missing { }
	parsedArgs, err := args.ParseArgsForValidation(StarlarkRunCmd.Args, []string{testScriptArg, inputArgs})
	require.Nil(t, err)
	require.NotNil(t, parsedArgs)

	err = validatePackageArgs(testCtx, testParsedFlags, parsedArgs)
	require.NotNil(t, err)
}
