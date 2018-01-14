package coinmarketcap

import (
	"testing"
)

func TestClient(t *testing.T) {

}

// TestEncoding test encodeBody
func TestEncoding(t *testing.T) {
	type Mock struct {
		Foo string `json:"foo"`
	}

	input := &Mock{Foo: "bar"}
	expected := `{"foo":"bar"}`

	encoded, err := encodeBody(input)
	if err != nil {
		t.Fatal(err)
	}

	encodedstr := string(encoded)
	if encodedstr != expected {
		t.Fatalf("expected '%v', got '%v'", expected, encodedstr)
	}
}
