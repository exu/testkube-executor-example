package runner

import (
	"testing"

	"github.com/kubeshop/testkube/pkg/api/v1/testkube"
	"github.com/stretchr/testify/assert"
)

func TestRun(t *testing.T) {

	t.Run("successful result", func(t *testing.T) {
		runner := NewRunner()
		res, err := runner.Run(testkube.Execution{
			ScriptContent: "https://testkube.io",
		})

		assert.NoError(t, err)
		assert.Equal(t, testkube.ExecutionStatusSuccess, res.Status)
	})

	t.Run("failed 404 result", func(t *testing.T) {
		runner := NewRunner()
		res, err := runner.Run(testkube.Execution{
			ScriptContent: "https://testkube.io/some-non-existing-uri-blablablabl",
		})

		assert.NoError(t, err)
		assert.Equal(t, testkube.ExecutionStatusError, res.Status)

	})

	t.Run("network connection issues returns errors", func(t *testing.T) {
		runner := NewRunner()
		_, err := runner.Run(testkube.Execution{
			ScriptContent: "blabla://non-existing-uri",
		})

		assert.Error(t, err)
	})

}
