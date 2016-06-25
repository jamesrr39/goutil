package user

import (
        "github.com/stretchr/testify/assert"
        "strings"
        "testing"
)

func TestExpandUser(t *testing.T) {

        dir, err := ExpandUser("~/Documents")

        assert.Nil(t, err)
        assert.Equal(t, "/", string(dir[0]))
        assert.True(t, strings.HasSuffix(dir, "/Documents"))

}

