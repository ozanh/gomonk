package help_test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/urfave/cli/v2"

	"gomonk/cmd/help"
)

func TestHelp(t *testing.T) {
	execName, err := os.Executable()
	require.NoError(t, err)

	app := &cli.App{
		Commands: []*cli.Command{
			help.Command(),
		},
	}

	testArgs := []string{execName, "help"}
	require.NoError(t, app.Run(testArgs))
}
