package incapsula

import (
	"os"
	"testing"
)

func TestParseUtmvcScriptPath(t *testing.T) {
	file, err := os.Open("../scripts/utmvc.html")
	if err != nil {
		t.Fatal(err)
	}

	scriptPath, err := ParseUtmvcScriptPath(file)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(scriptPath)
}
