package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestName(t *testing.T) {
	testCases := []struct{
		name string
		args []string
		exp  *Arguments
	}{
		{
			name: "-n 2s ls -l",
			args: []string{"-n", "2s", "ls", "-l"},
			exp: &Arguments{
				interval:  2 * time.Second,
				isPrecise: false,
				isActual:  false,
				cmd:       "ls",
				args:      []string{"-l"},
			},
		},
	}

	for _, tt := range testCases {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			argument, err := parseArguments(tt.args)
			assert.NoError(t, err)
			assert.Equal(t, tt.exp, argument)
		})
	}
}
