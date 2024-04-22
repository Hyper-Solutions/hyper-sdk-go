package akamai

import (
	"io"
	"strings"
	"testing"
)

const secCptHtml = `<iframe id="sec-cpt-if" provider="crypto" class="crypto" challenge="eyJ0b2tlbiI6IkFBUUFBQUFKX19fX184cXRFWk5MVVRlUFhyRXNIdU1BNnF0Zk53aDBhRjR5MEJYaGQ2TlJ1dlRFdGpTc1pFNW5ja2hfSDRfSFA5bm5sd0V6ODRYWUEtUGlFck80QmNDelc5SmV5N1lMckZiY041SDdPX0RicFFKWEo5czQ2THJyYlRCZDBYdVRVcENYWFJtclp3YmxXNWZGLW14UUFXSGllTmk2SUZ1Y1RTTkx6bFlnbUpScEJMWGRHUnBHOUtZQXU5Y1dVNGRTdUNqMi1JQ0ZRY2gyLUNQX3BaRzFlVWJndzVELTFKNE90S285R0JLNWNMYW1qSENHQWU1U3VpaS1tTExiYWlvWWduNks0VmdzWm1PeVpwMkc2eHFXaU43eWZHREl5NXhEbXpscXV5QmFHNVlBUExMcGFpb3NtWHp2UGtNIiwidGltZXN0YW1wIjoxNzEyMjM3NTQ5LCJub25jZSI6ImFiMzAyOGQ1MTg3MzhiZDhmYWFkIiwiZGlmZmljdWx0eSI6MjUwMCwidGltZW91dCI6MTAwMCwiY3B1IjpmYWxzZX0="  data-key="" data-duration=5 src="/_sec/cp_challenge/ak-challenge-4-3.htm"></iframe>`

const secCptJson = `{"sec-cp-challenge":"true","provider":"crypto","branding_url_content":"/_sec/cp_challenge/crypto_message-4-3.htm","chlg_duration":30,"token":"AAQAAAAJ_____9z_ZPsdHbk36hg2f6np2sGJDXmkwGmBiMBr_DDEmSWfi8Zt7BdtjWrNd9KD4DS_vim0VnK2wsa8tIC7XWsCshkvDF9J9Rf5EFwBU00c6SMXTaSNSTcDR-HVFGp3uAa67Mb3I6HeifXbjALcEomjcnwa9ZNQdDWuTAUTgNGbYw09A8AXIuP9DNv3QktUx488FV38Rm6xBXr66-MmD05hsBhucIYpLS_VCJVs9OFPnWsksPJ19ibw2K3fabfJbzIdB3Xv3J0kzLQ0gY7bpLRXK1oAcUTxNNsy-LQGe_lyV6INQ4ojPLGJpOTk","timestamp":1713283747,"nonce":"ebccdb479fcb92636fbc","difficulty":15000,"timeout":1000,"cpu":false}`

func TestParseSecCptChallenge(t *testing.T) {
	input := strings.NewReader(secCptHtml)

	challenge, err := ParseSecCptChallenge(input)
	if err != nil {
		t.Fatal(err)
	}

	if challenge.duration != 5 {
		t.Errorf("Expected duration to be 5, got %d", challenge.duration)
	}

	if challenge.challengeData.Difficulty != 2500 {
		t.Errorf("Expected difficulty to be 2500, got %d", challenge.challengeData.Difficulty)
	}

	if challenge.ChallengePath != "/_sec/cp_challenge/ak-challenge-4-3.htm" {
		t.Errorf("Expected challenge path to be '/_sec/cp_challenge/ak-challenge-4-3.htm', got %s", challenge.ChallengePath)
	}
}

func TestParseSecCptChallengeData(t *testing.T) {
	input := strings.NewReader(secCptHtml)
	src, _ := io.ReadAll(input)

	challengeData, err := parseSecCptChallengeData(src)
	if err != nil {
		t.Fatal(err)
	}

	if challengeData.Token != "AAQAAAAJ_____8qtEZNLUTePXrEsHuMA6qtfNwh0aF4y0BXhd6NRuvTEtjSsZE5nckh_H4_HP9nnlwEz84XYA-PiErO4BcCzW9Jey7YLrFbcN5H7O_DbpQJXJ9s46LrrbTBd0XuTUpCXXRmrZwblW5fF-mxQAWHieNi6IFucTSNLzlYgmJRpBLXdGRpG9KYAu9cWU4dSuCj2-ICFQch2-CP_pZG1eUbgw5D-1J4OtKo9GBK5cLamjHCGAe5Suii-mLLbaioYgn6K4VgsZmOyZp2G6xqWiN7yfGDIy5xDmzlquyBaG5YAPLLpaiosmXzvPkM" {
		t.Errorf("Expected token to be 'AAQAAAAJ_____8qtEZNLUTePXrEsHuMA6qtfNwh0aF4y0BXhd6NRuvTEtjSsZE5nckh_H4_HP9nnlwEz84XYA-PiErO4BcCzW9Jey7YLrFbcN5H7O_DbpQJXJ9s46LrrbTBd0XuTUpCXXRmrZwblW5fF-mxQAWHieNi6IFucTSNLzlYgmJRpBLXdGRpG9KYAu9cWU4dSuCj2-ICFQch2-CP_pZG1eUbgw5D-1J4OtKo9GBK5cLamjHCGAe5Suii-mLLbaioYgn6K4VgsZmOyZp2G6xqWiN7yfGDIy5xDmzlquyBaG5YAPLLpaiosmXzvPkM', got %s", challengeData.Token)
	}

	if challengeData.Timestamp != 1712237549 {
		t.Errorf("Expected timestamp to be 1712237549, got %d", challengeData.Timestamp)
	}

	if challengeData.Nonce != "ab3028d518738bd8faad" {
		t.Errorf("Expected nonce to be 'ab3028d518738bd8faad', got %s", challengeData.Nonce)
	}

	if challengeData.Difficulty != 2500 {
		t.Errorf("Expected difficulty to be 2500, got %d", challengeData.Difficulty)
	}

	if challengeData.Timeout != 1000 {
		t.Errorf("Expected timeout to be 1000, got %d", challengeData.Timeout)
	}
}

func TestParseSecCptChallengeFromJson(t *testing.T) {
	challenge, err := ParseSecCptChallengeFromJson(strings.NewReader(secCptJson))
	if err != nil {
		t.Fatal(err)
	}

	if challenge.challengeData.Token != "AAQAAAAJ_____9z_ZPsdHbk36hg2f6np2sGJDXmkwGmBiMBr_DDEmSWfi8Zt7BdtjWrNd9KD4DS_vim0VnK2wsa8tIC7XWsCshkvDF9J9Rf5EFwBU00c6SMXTaSNSTcDR-HVFGp3uAa67Mb3I6HeifXbjALcEomjcnwa9ZNQdDWuTAUTgNGbYw09A8AXIuP9DNv3QktUx488FV38Rm6xBXr66-MmD05hsBhucIYpLS_VCJVs9OFPnWsksPJ19ibw2K3fabfJbzIdB3Xv3J0kzLQ0gY7bpLRXK1oAcUTxNNsy-LQGe_lyV6INQ4ojPLGJpOTk" {
		t.Fail()
	}

	if challenge.challengeData.Timestamp != 1713283747 {
		t.Fail()
	}

	if challenge.challengeData.Nonce != "ebccdb479fcb92636fbc" {
		t.Fail()
	}

	if challenge.challengeData.Difficulty != 15000 {
		t.Fail()
	}

	if challenge.challengeData.Timeout != 1000 {
		t.Fail()
	}
}

func BenchmarkParseSecCptChallenge(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		if _, err := ParseSecCptChallenge(strings.NewReader(secCptHtml)); err != nil {
			b.Error(err)
		}
	}
}

func BenchmarkSecCptChallenge_GenerateSecCptPayload(b *testing.B) {
	input := strings.NewReader(secCptHtml)

	challenge, err := ParseSecCptChallenge(input)
	if err != nil {
		b.Fatal(err)
	}
	const cookie = `3F3B2D3E2ABE67693EE8134E57C501C8~...`

	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		if _, err := challenge.GenerateSecCptPayload(cookie); err != nil {
			b.Error(err)
		}
	}
}
