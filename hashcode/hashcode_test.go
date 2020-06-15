package hashcode_test

import (
	"phraseservice/hashcode"
	"testing"
)

type mockHash struct {
	data []byte
}

func (m *mockHash) Write(p []byte) (n int, err error) {
	m.data = p
	return len(p), nil
}

func (m *mockHash) Sum(_ []byte) []byte {
	return m.data
}

func (m *mockHash) Reset() {
}

func (m *mockHash) Size() int {
	panic("implement me")
}

func (m *mockHash) BlockSize() int {
	panic("implement me")
}

var getCodeTests = map[string]struct {
	input    string
	expected int64
}{
	"empty string": {input: "", expected: 0},
	"small hash":   {input: string([]byte{0xFF}), expected: 255},
	"8-byte long hash, but it is less than maxInt": {
		input:    string([]byte{0x11, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11}),
		expected: 1229782938247303441,
	},
	"8-byte long hash is greater than maxInt": {
		input:    string([]byte{0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF}),
		expected: 1<<63 - 1,
	},
	"hash length is more than 8 byte": {
		input:    string([]byte{0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF}),
		expected: 1<<63 - 1,
	},
}

func TestHashCode_GetCode(t *testing.T) {
	for testName, tt := range getCodeTests {
		hc := hashcode.New(&mockHash{})
		hc.LoadString(tt.input)

		actual := hc.GetCode()
		if tt.expected != actual {
			t.Errorf("Test %q failed: \n\t expected: %d \n\t      get: %d", testName, tt.expected, actual)
		}
	}
}
