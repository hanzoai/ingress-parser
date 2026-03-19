package env

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindPrefixedEnvVars(t *testing.T) {
	testCases := []struct {
		desc     string
		environ  []string
		element  interface{}
		expected []string
	}{
		{
			desc:     "exact name",
			environ:  []string{"INGRESS_FOO"},
			element:  &Yo{},
			expected: []string{"INGRESS_FOO"},
		},
		{
			desc:     "prefixed name",
			environ:  []string{"INGRESS_FII01"},
			element:  &Yo{},
			expected: []string{"INGRESS_FII01"},
		},
		{
			desc:     "excluded env vars",
			environ:  []string{"INGRESS_NOPE", "INGRESS_NO"},
			element:  &Yo{},
			expected: nil,
		},
		{
			desc:     "filter",
			environ:  []string{"INGRESS_NOPE", "INGRESS_NO", "INGRESS_FOO", "INGRESS_FII01"},
			element:  &Yo{},
			expected: []string{"INGRESS_FOO", "INGRESS_FII01"},
		},
	}

	for _, test := range testCases {
		t.Run(test.desc, func(t *testing.T) {
			t.Parallel()

			vars := FindPrefixedEnvVars(test.environ, DefaultNamePrefix, test.element)

			assert.Equal(t, test.expected, vars)
		})
	}
}

func Test_getRootFieldNames(t *testing.T) {
	testCases := []struct {
		desc     string
		element  interface{}
		expected []string
	}{
		{
			desc:     "simple fields",
			element:  &Yo{},
			expected: []string{"INGRESS_FOO", "INGRESS_FII", "INGRESS_FUU", "INGRESS_YI", "INGRESS_YU"},
		},
		{
			desc:     "embedded struct",
			element:  &Yu{},
			expected: []string{"INGRESS_FOO", "INGRESS_FII", "INGRESS_FUU"},
		},
		{
			desc:     "embedded struct pointer",
			element:  &Ye{},
			expected: []string{"INGRESS_FOO", "INGRESS_FII", "INGRESS_FUU"},
		},
	}

	for _, test := range testCases {
		t.Run(test.desc, func(t *testing.T) {
			t.Parallel()

			names := getRootPrefixes(test.element, DefaultNamePrefix)

			assert.Equal(t, test.expected, names)
		})
	}
}
