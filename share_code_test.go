package share_code

import (
    "github.com/stretchr/testify/assert"
    "testing"
)

func TestEncode(t *testing.T) {
    code := Encode(581171235203846144)

    assert.Equal(t, code, "LYD9811WQDI4")
}

func TestDecode(t *testing.T) {
    id := Decode("LYD9811WQDI4")

    assert.Equal(t, id, int64(581171235203846144))
}
