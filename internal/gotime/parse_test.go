package gotime

import (
	"bytes"
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewParseCmd(t *testing.T) {
	epochNow := "1632438665"
	cmd := NewParseCmd()

	t.Run("Error: failed to parse input", func(t *testing.T) {
		cmd.SetArgs([]string{"asdf"})
		err := cmd.Execute()
		assert.EqualError(t, err, "failed to parse input: strconv.ParseInt: parsing \"asdf\": invalid syntax")
	})
	t.Run("Error: invalid time zone", func(t *testing.T) {
		cmd.SetArgs([]string{epochNow, "PST"})
		err := cmd.Execute()
		assert.EqualError(t, err, "PST is not a supported time zone")
	})
	t.Run("Success: Parse with time zone", func(t *testing.T) {
		buf := new(bytes.Buffer)
		cmd.SetOut(buf)
		cmd.SetArgs([]string{epochNow, "America/New_York"})
		err := cmd.Execute()
		require.NoError(t, err)
		assert.Equal(t, "2021-09-23 19:11:05 -0400 EDT\n", buf.String())
	})
	t.Run("Success: Parse into local time", func(t *testing.T) {
		i, _ := strconv.ParseInt(epochNow, 10, 64)
		localTime := time.Unix(i, 0)

		buf := new(bytes.Buffer)
		cmd.SetOut(buf)
		cmd.SetArgs([]string{epochNow})
		err := cmd.Execute()
		require.NoError(t, err)
		assert.Equal(t, localTime.String()+"\n", buf.String())
	})
}
