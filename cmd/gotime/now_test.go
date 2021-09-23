package gotime

import (
	"bytes"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewNowCmd(t *testing.T) {
	currentTime = time.Date(2021, 10, 29, 10, 10, 10, 10, time.UTC)
	cmd := NewNowCmd()

	t.Run("Success: current time", func(t *testing.T) {
		buf := new(bytes.Buffer)
		cmd.SetOut(buf)
		err := cmd.Execute()
		require.NoError(t, err)
		assert.Equal(t, "2021-10-29 10:10:10.00000001 +0000 UTC\n", buf.String())
	})

	t.Run("Success: epoch time", func(t *testing.T) {
		buf := new(bytes.Buffer)
		cmd.SetOut(buf)
		cmd.SetArgs([]string{"-e"})
		err := cmd.Execute()
		require.NoError(t, err)
		assert.Equal(t, "1635502210\n", buf.String())
	})

	t.Run("Success: current time in different time zone", func(t *testing.T) {
		buf := new(bytes.Buffer)
		cmd.SetOut(buf)
		cmd.SetArgs([]string{"America/Los_Angeles"})
		err := cmd.Execute()
		require.NoError(t, err)
		assert.Equal(t, "2021-10-29 03:10:10.00000001 -0700 PDT\n", buf.String())
	})

	t.Run("Error: Invalid time zone", func(t *testing.T) {
		cmd.SetArgs([]string{"PST"})
		err := cmd.Execute()
		assert.EqualError(t, err, "PST is not a supported time zone\n")
	})
}
