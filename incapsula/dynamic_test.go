package incapsula

import (
	"os"
	"testing"
)

func TestParseDynamicReeseScript(t *testing.T) {
	file, err := os.Open("../scripts/reese.html")
	if err != nil {
		t.Fatal(err)
	}
	defer file.Close()

	sensorPath, scriptPath, err := ParseDynamicReeseScript(file, "https://www.smythstoys.com/")
	if err != nil {
		t.Fatal(err)
	}

	t.Log(sensorPath, scriptPath)
}
