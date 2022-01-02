package apitest

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestApiTestcaseCode(t *testing.T) {
	s := &ApiTestcaseService{}
	s.ApiTestcaseCode()

}

func TestAAA(t *testing.T) {

	s1 := make([]string, 0)
	require.Equal(t, 0, len(s1))
	require.NotEqual(t, nil, s1)

	var s2 []string
	require.Equal(t, 0, len(s1))
	require.NotEqual(t, nil, s2)
}
