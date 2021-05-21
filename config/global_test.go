package config

import (
	"bytes"
	"testing"
)

func TestWrite(t *testing.T) {
	tests := []struct {
		Global   *Global
		Expected string
	}{
		{
			&Global{},
			"",
		},
	}

	for _, test := range tests {
		var buf bytes.Buffer

		if err := test.Global.Write(&buf); err != nil {
			t.Fatal(err)
		}

		if test.Expected != buf.String() {
			t.Fatal(test.Expected)
		}
	}
}
