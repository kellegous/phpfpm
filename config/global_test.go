package config

import (
	"bytes"
	"strings"
	"testing"
)

func fromLines(lines ...string) string {
	return strings.Join(lines, "\n")
}

func TestGlobalWrite(t *testing.T) {
	tests := []struct {
		Input    func() *Global
		Expected string
	}{
		{
			func() *Global {
				g := New()
				g.WithPool("www", "127.0.0.1:8080", "root", TypeStatic, 3)
				return g
			},
			fromLines(
				"daemon = no",
				"[www]",
				"listen = 127.0.0.1:8080",
				"user = root",
				"pm = static",
				"pm.max_children = 3",
				"",
			),
		},
	}

	for _, test := range tests {
		var buf bytes.Buffer

		g := test.Input()
		if err := g.Write(&buf); err != nil {
			t.Fatal(err)
		}

		if test.Expected != buf.String() {
			t.Fatalf("results do not match\nEXPECTED:\n%s\nGOT:\n%s\n", test.Expected, buf.String())
		}
	}
}
