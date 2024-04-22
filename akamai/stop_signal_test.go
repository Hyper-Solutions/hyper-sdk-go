package akamai

import "testing"

func TestIsCookieValid(t *testing.T) {
	const (
		validCookie = `0C8A2251CC04F60F59160D6AD92DA8A0~0~YAAQlivJF6o1GjGGAQAAaNihYgldsErwKa3aAlB+oRlgZYviinJa+Q29XMXmkwJUNCgQPooQUyhfjhAgavSMACfCk1doYnUa4dYmsVUWbWB+QGFEPuwcvLVQscLV8taHWIFuFxb94vEJ8MnSY9sQlhRN9i2iNgZ5QJz8h2s2mMm4ZO+i890DPaHfJPkSrYtc9ivgbjDA/jFpK6k2Pq8Pu25dCI55zOqOeSyaChgtJyF6KvlnlyVrqOa12tThX+prb52et7FRGmqhw8LU5X1E07WShiKDmJw2Rb8+odHcA28bD3EITXTy43QFb2PKR9Q3jy57KFEQYFaeW4xfAe8BxjdpkYt5vSH50nmbi1SXOWIxL4QV0b8psJgMCIq+ZMR+ZBU2opuHgxAucktvbIffGuJPWFYJu8thjxr5HGBtZBUqc6LwccFQI+DZd+hpZfsRJscvNx0yWmiPN8/gJiVLGkjHWYL5xmVaVCYceTtGL+7N~-1~-1~-1`

		invalidCookie = `3B508597CDC152514C3D85CF2749455C~-1~YAAQNcMTAgMgNEyKAQAAaxaxTwrFBKU37u+fFlXF8NMV3M3n4A2lNyvNTV14HhbYAyJWQdsaUYXmjc/7GGsPQ8EYyPjUZr6X8guTu9q9mZY2ZeF5IWdB2jRHLVzEttloRl8RMGS+dP34QSaMYx98elcgQchq+DAiRDB1XoeKDzwdZnxhfLRu8vAilIaR+i/NPf1Y1fR+n85SIp5OlpMJYK71eoL/D9wZTbpZQVHbSP4rhHG4wNaHNkuo5KkwWpgbQfLWIwhatuhn/xZ1mJEQVnZwhHg9aDF+ooxiPm+2+HK/QI5zSGZ8zz9CMZ9PI+jegXtqznth6eBXJ7ZXCr5VfgjnepKpJW3WJfKv53dHHWUtkuBEOTozjuR1RYcsFo07oNmUqWBhQ0EslRCJpMAn7ehwuyC2w7BRlrZuSvDOAcxUgbgN2LSBEA==~-1~-1~-1`
	)

	if !IsCookieValid(validCookie, 1) {
		t.Fail()
	}

	if IsCookieValid(invalidCookie, 1) {
		t.Fail()
	}
}

func BenchmarkIsCookieValid(b *testing.B) {
	const cookie = `0C8A2251CC04F60F59160D6AD92DA8A0~0~YAAQlivJF6o1GjGGAQAAaNihYgldsErwKa3aAlB+oRlgZYviinJa+Q29XMXmkwJUNCgQPooQUyhfjhAgavSMACfCk1doYnUa4dYmsVUWbWB+QGFEPuwcvLVQscLV8taHWIFuFxb94vEJ8MnSY9sQlhRN9i2iNgZ5QJz8h2s2mMm4ZO+i890DPaHfJPkSrYtc9ivgbjDA/jFpK6k2Pq8Pu25dCI55zOqOeSyaChgtJyF6KvlnlyVrqOa12tThX+prb52et7FRGmqhw8LU5X1E07WShiKDmJw2Rb8+odHcA28bD3EITXTy43QFb2PKR9Q3jy57KFEQYFaeW4xfAe8BxjdpkYt5vSH50nmbi1SXOWIxL4QV0b8psJgMCIq+ZMR+ZBU2opuHgxAucktvbIffGuJPWFYJu8thjxr5HGBtZBUqc6LwccFQI+DZd+hpZfsRJscvNx0yWmiPN8/gJiVLGkjHWYL5xmVaVCYceTtGL+7N~-1~-1~-1`

	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		IsCookieValid(cookie, 1)
	}
}

func TestIsCookieInvalidated(t *testing.T) {
	const (
		invalidatedCookie = `7D715603CF98EF4593FB1BABFA0BA525~-1~YAAQ5oFlX6qoLeeGAQAAi7B99gl3E4KSKtdW8AEn7RB8D1CpzSqsnYR5E24Q66mWSu8yXMOVdmjYPVFVad6QNKZ/w6xs2sU9sX/t4GNgLFqNLn3Qcags4msWUL6Mdlmh/MKPZWiBnU6pGnAec9cdYW9gAWiZ3kSvCxJHYD536EBIJKKkZ/EcCCKauQbnn+TUuSp4D2jSQfUEOaXMTiSREKRnqLpc9lmgG8hkFBeeyWlu7vv+iussTelN6o5zCHwH16ztaLQTDVRclRGaUo2jqN7dpDd8V0WvZ7NnbNsiU2Ac52TBM7Kjl5/l2ltAAlYr+vgnc3QRhbCo8trn2RrEP7nkCRF1RzQ3HvG097nul3hcRPXitfIslgVG9ur67LTwpRt58DqjgjNz4qHR5R77VzyTUQPt8ZQzMeh4s9TOr/E=~0~-1~-1`
	)

	if !IsCookieInvalidated(invalidatedCookie) {
		t.Fail()
	}
}

func BenchmarkIsCookieInvalidated(b *testing.B) {
	const cookie = `0C8A2251CC04F60F59160D6AD92DA8A0~0~YAAQlivJF6o1GjGGAQAAaNihYgldsErwKa3aAlB+oRlgZYviinJa+Q29XMXmkwJUNCgQPooQUyhfjhAgavSMACfCk1doYnUa4dYmsVUWbWB+QGFEPuwcvLVQscLV8taHWIFuFxb94vEJ8MnSY9sQlhRN9i2iNgZ5QJz8h2s2mMm4ZO+i890DPaHfJPkSrYtc9ivgbjDA/jFpK6k2Pq8Pu25dCI55zOqOeSyaChgtJyF6KvlnlyVrqOa12tThX+prb52et7FRGmqhw8LU5X1E07WShiKDmJw2Rb8+odHcA28bD3EITXTy43QFb2PKR9Q3jy57KFEQYFaeW4xfAe8BxjdpkYt5vSH50nmbi1SXOWIxL4QV0b8psJgMCIq+ZMR+ZBU2opuHgxAucktvbIffGuJPWFYJu8thjxr5HGBtZBUqc6LwccFQI+DZd+hpZfsRJscvNx0yWmiPN8/gJiVLGkjHWYL5xmVaVCYceTtGL+7N~-1~-1~-1`

	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		IsCookieInvalidated(cookie)
	}
}
