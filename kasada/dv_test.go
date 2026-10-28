package kasada

import "testing"

// A user-agent and dv built from one set of inputs, in the four-field form
// the i-2.x iOS SDK sends.
func TestGenerateDV(t *testing.T) {
	const (
		osVersion  = "18.0.1"
		hardwareID = "iPhone13,2"
		appID      = "com.example.app"
		appVersion = "9.15.0"
	)
	wantDV := "QkZWEmcDRUBEDloaAg8FDhpTAwhEGxFFWRFcQAVCVHFXFAlBUDRRDl8DBgABUxYUGkxERn0HDwNEVFxAFgZAVglAXV0dV0MTG0gbAEEWYVZfEl1XVg8OFA1NHAxIBxFQ"
	if got := GenerateDV(osVersion, hardwareID, appID, appVersion); got != wantDV {
		t.Errorf("GenerateDV = %q, want %q", got, wantDV)
	}
	wantUA := "7016adc4ba8d433533954ebad33be424"
	if got := GenerateUserAgent(osVersion, appID, appVersion); got != wantUA {
		t.Errorf("GenerateUserAgent = %q, want %q", got, wantUA)
	}
}
