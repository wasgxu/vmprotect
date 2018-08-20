package vmprotect

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTooShort(t *testing.T) {
	serial := "aGVsbG8sIHdvcmxk"
	_, err := ParseLicense(serial, vmpPublic, vmpModulus, vmpProductCode, vmpBits)

	require.NotNil(t, err, "The input is too short, it must be invalid")
}

func TestTooLong(t *testing.T) {
	serial := "aGVsbG8sIHdvcmxkIGhlbGxvLCB3b3JsZCBoZWxsbywgd29ybGQgaGVsbG8sIHdvcmxkIGhlbGxvLCB3b3JsZCBoZWxsbywgd29ybGQgaGVsbG8sIHdvcmxkIGhlbGxvLCB3b3JsZCBoZWxsbywgd29ybGQgaGVsbG8sIHdvcmxkIGhlbGxvLCB3b3JsZCBoZWxsbywgd29ybGQgaGVsbG8sIHdvcmxkIGhlbGxvLCB3b3JsZCBoZWxsbywgd29ybGQgaGVsbG8sIHdvcmxkIGhlbGxvLCB3b3JsZCBoZWxsbywgd29ybGQgaGVsbG8sIHdvcmxkIGhlbGxvLCB3b3JsZCBoZWxsbywgd29ybGQgaGVsbG8sIHdvcmxkIGhlbGxvLCB3b3JsZCBoZWxsbywgd29ybGQgaGVsbG8sIHdvcmxkIGhlbGxvLCB3b3JsZCBoZWxsbywgd29ybGQgaGVsbG8sIHdvcmxkIGhlbGxvLCB3b3JsZCBoZWxsbywgd29ybGQgaGVsbG8sIHdvcmxkIGhlbGxvLCB3b3JsZCBoZWxsbywgd29ybGQgaGVsbG8sIHdvcmxkIGhlbGxvLCB3b3JsZCBoZWxsbywgd29ybGQgaGVsbG8sIHdvcmxkIGhlbGxvLCB3b3JsZCBoZWxsbywgd29ybGQgaGVsbG8sIHdvcmxkIGhlbGxvLCB3b3JsZCBoZWxsbywgd29ybGQgaGVsbG8sIHdvcmxkIGhlbGxvLCB3b3JsZCBoZWxsbywgd29ybGQgaGVsbG8sIHdvcmxkIGhlbGxvLCB3b3JsZCBoZWxsbywgd29ybGQgaGVsbG8sIHdvcmxkIGhlbGxvLCB3b3JsZCBoZWxsbywgd29ybGQgaGVsbG8sIHdvcmxkIGhlbGxvLCB3b3JsZCBoZWxsbywgd29ybGQgaGVsbG8sIHdvcmxkIGhlbGxvLCB3b3JsZCBoZWxsbywgd29ybGQgaGVsbG8sIHdvcmxkIGhlbGxvLCB3b3JsZCBoZWxsbywgd29ybGQgaGVsbG8sIHdvcmxkIGhlbGxvLCB3b3JsZCBoZWxsbywgd29ybGQgaGVsbG8sIHdvcmxkIGhlbGxvLCB3b3JsZCBoZWxsbywgd29ybGQgaGVsbG8sIHdvcmxkIGhlbGxvLCB3b3JsZCBoZWxsbywgd29ybGQgaGVsbG8sIHdvcmxkIGhlbGxvLCB3b3JsZCBoZWxsbywgd29ybGQgaGVsbG8sIHdvcmxkIGhlbGxvLCB3b3JsZCBoZWxsbywgd29ybGQgaGVsbG8sIHdvcmxkIGhlbGxvLCB3b3JsZCBoZWxsbywgd29ybGQgaGVsbG8sIHdvcmxkIGhlbGxvLCB3b3JsZCBoZWxsbywgd29ybGQgaGVsbG8sIHdvcmxkIGhlbGxvLCB3b3JsZCBoZWxsbywgd29ybGQgaGVsbG8sIHdvcmxkIGhlbGxvLCB3b3JsZCBoZWxsbywgd29ybGQgaGVsbG8sIHdvcmxkIGhlbGxvLCB3b3JsZCBoZWxsbywgd29ybGQgaGVsbG8sIHdvcmxkIGhlbGxvLCB3b3JsZCBoZWxsbywgd29ybGQgaGVsbG8sIHdvcmxkIGhlbGxvLCB3b3JsZCBoZWxsbywgd29ybGQgaGVsbG8sIHdvcmxkIGhlbGxvLCB3b3JsZCBoZWxsbywgd29ybGQgaGVsbG8sIHdvcmxkIGhlbGxvLCB3b3JsZCBoZWxsbywgd29ybGQgaGVsbG8sIHdvcmxkIGhlbGxvLCB3b3JsZCBoZWxsbywgd29ybGQgaGVsbG8sIHdvcmxkIGhlbGxvLCB3b3JsZCBoZWxsbywgd29ybGQgaGVsbG8sIHdvcmxkIGhlbGxvLCB3b3JsZCBoZWxsbywgd29ybGQgaGVsbG8sIHdvcmxkIGhlbGxvLCB3b3JsZCBoZWxsbywgd29ybGQgaGVsbG8sIHdvcmxkIGhlbGxvLCB3b3JsZCBoZWxsbywgd29ybGQgaGVsbG8sIHdvcmxkIGhlbGxvLCB3b3JsZCBoZWxsbywgd29ybGQgaGVsbG8sIHdvcmxkIGhlbGxvLCB3b3JsZCBoZWxsbywgd29ybGQgaGVsbG8sIHdvcmxkIGhlbGxvLCB3b3JsZCBoZWxsbywgd29ybGQgaGVsbG8sIHdvcmxkIGhlbGxvLCB3b3JsZCBoZWxsbywgd29ybGQgaGVsbG8sIHdvcmxkIGhlbGxvLCB3b3JsZCBoZWxsbywgd29ybGQgaGVsbG8sIHdvcmxkIGhlbGxvLCB3b3JsZCBoZWxsbywgd29ybGQgaGVsbG8sIHdvcmxkIGhlbGxvLCB3b3JsZCBoZWxsbywgd29ybGQgaGVsbG8sIHdvcmxkIGhlbGxvLCB3b3JsZCBoZWxsbywgd29ybGQgaGVsbG8sIHdvcmxkIGhlbGxvLCB3b3JsZCBoZWxsbywgd29ybGQgaGVsbG8sIHdvcmxkIGhlbGxvLCB3b3JsZCBoZWxsbywgd29ybGQgaGVsbG8sIHdvcmxkIGhlbGxvLCB3b3JsZCBoZWxsbywgd29ybGQgaGVsbG8sIHdvcmxkIGhlbGxvLCB3b3JsZCBoZWxsbywgd29ybGQgaGVsbG8sIHdvcmxkIGhlbGxvLCB3b3JsZCBoZWxsbywgd29ybGQgaGVsbG8sIHdvcmxkIGhlbGxvLCB3b3JsZCBoZWxsbywgd29ybGQgaGVsbG8sIHdvcmxkIGhlbGxvLCB3b3JsZCBoZWxsbywgd29ybGQgaGVsbG8sIHdvcmxkIGhlbGxvLCB3b3JsZCBoZWxsbywgd29ybGQgaGVsbG8sIHdvcmxkIGhlbGxvLCB3b3JsZCBoZWxsbywgd29ybGQgaGVsbG8sIHdvcmxkIGhlbGxvLCB3b3JsZCBoZWxsbywgd29ybGQgaGVsbG8sIHdvcmxkIGhlbGxvLCB3b3JsZCBoZWxsbywgd29ybGQgaGVsbG8sIHdvcmxkIGhlbGxvLCB3b3JsZCBoZWxsbywgd29ybGQgaGVsbG8sIHdvcmxkIGhlbGxvLCB3b3JsZCBoZWxsbywgd29ybGQgaGVsbG8sIHdvcmxkIGhlbGxvLCB3b3JsZCBoZWxsbywgd29ybGQgaGVsbG8sIHdvcmxkIGhlbGxvLCB3b3JsZCBoZWxsbywgd29ybGQgaGVsbG8sIHdvcmxkIGhlbGxvLCB3b3JsZCBoZWxsbywgd29ybGQgaGVsbG8sIHdvcmxkIGhlbGxvLCB3b3JsZCA="
	_, err := ParseLicense(serial, vmpPublic, vmpModulus, vmpProductCode, vmpBits)

	require.NotNil(t, err, "The input is too long, it must be invalid")
}

func TestRightLength(t *testing.T) {
	serial := "RxtFJ+i0xRj8KNoyR9oAfQtOEHNRTKL5t/kWW7f/l0HFt9MC0uOhy0TeVC40LSjixdZ3xbON69bzloRDaAVyZMk44VPfT5jYt5PYyVEsZsWg1xkMdBNd1J2yrzI5cl8Ly2/w4JqZcJai7ubQDEz9MsTE1dfaoYT+vEnO4AjG+ynXYpZKQbIUAfDMBORR9KuVZp5wIJEFsloKZibU+W/72XLva5kSdXXGbtYIvKG8x41sO897aXBpVYbZ+WpnbGs7O3AqBhw0PNtJ2WAfN4Ce/RW0AkPDpfk4uqd5wvER5l9RNZnDjIhN1+/4dfmF1PMeFPaqhEcH3vf4SzJcfd6P4g=="
	_, err := ParseLicense(serial, vmpPublic, vmpModulus, vmpProductCode, vmpBits)

	require.NotNil(t, err, "The input is of a right length, but still invalid")
}

func TestSpacesAndNewLines(t *testing.T) {
	serial := "RxtFJ+i0xRj8KN\noyR9oAfQtOEHNRTKL5t/kWW7f/l0HFt9MC0u Ohy0TeVC40LSjixdZ3xbON69bzloRDaAVyZMk   44VPfT5jYt5PYyVEsZsWg1xkMdBNd1J2yrzI5cl8Ly2/w4JqZcJai\n\n\n7ubQDEz9MsTE1dfaoYT+vEnO4AjG+ynX\r\nYpZKQ\n\rbIUAfDMBORR9KuVZp5wIJEFsloKZibU+W/72XLva5kSdXXGbtYIvKG8x41sO897aXBpVYbZ+WpnbGs7O3AqBhw0PNtJ2WAfN4Ce/RW0AkPDpfk4uqd5wvER5l9RNZnDjIhN1+/4dfmF1PMeFPaqhEcH3vf4SzJcfd6P4g=="
	_, err := ParseLicense(serial, vmpPublic, vmpModulus, vmpProductCode, vmpBits)

	require.NotNil(t, err, "The input is right just has some newlines and spaces, but still invalid")
}

func TestIgnoreUnsupportedCharacters(t *testing.T) {
	serial := "RxtFJ+i0xRj8KNoyR9oAfQtOEHNRTKL5t/kWW7f/l0HFt9MC0uOhy0TeVC40LSjixdZ3xbO$$$N69bzloRDaAVyZMk44VPfT5jYt5PYyVEsZs---Wg1xkMdBNd1J2yrzI5cl8Ly2/w4JqZcJai7ubQDEz9MsT###E1dfaoYT+vEnO4AjG+ynXYpZKQbIUAfDMBORR9KuVZp5wIJEFsloKZibU+W/72XLva5kSdXXGbtYIvKG8x41sO897aXBpVYbZ+WpnbGs7O3AqBh()w0PNtJ2WAfN4Ce/RW0AkP{}Dpfk4uqd5wvE][R5l9RNZnDjIhN1+/4dfmF1PMeFPaqhEcH3vf4SzJcfd6P4g=="
	_, err := ParseLicense(serial, vmpPublic, vmpModulus, vmpProductCode, vmpBits)

	require.NotNil(t, err, "The input has some unsupported characters that must be ignored. The input must be valid, but the serial number itself - doesn't")
}

func TestVmpParamsNotInitialized(t *testing.T) {
	serial := "b2HUC5SA0qqHSmJHAJe+pM9Q5sey+iqCqkW3e0cK8R3kSxlGsFrVzVJ/OZ5etJ8DeDHCKBbmismtwd3I9uzJwitfR/NJJ93u/n/5J0RFDAkklyJ+A23mEDtdwP/w/LS97jvFMfXwX0SMBtQ28948iraiu7VeruU9SZcUerlPLtXj4AKoUOzfciWYJ9xDMA+daJOFioMd7zNZ2AW7bz8PB9+X5Vrtg6fg7QPaJuuXBqkQyxKaoBm/YCcVNBST0LpP0upDV/FDAhHXJL6hjvt55RE6vdHt75othC9diQAIxREN8JhrGkZnOGEypwB5wBCGYeD43bc8s+AM3P7AtUlxxg=="
	lic, err := ParseLicense(serial, "", "", "", 2048)

	require.NotNil(t, err, "There must be an error here")
	require.Nil(t, lic, "License must not be returned")
}
