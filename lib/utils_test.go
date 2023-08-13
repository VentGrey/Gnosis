package lib

import (
	"testing"
)

func TestInferType(t *testing.T) {
    tests := []struct {
        input  interface{}
        output string
    }{
        {input: "string", output: "string"},
        {input: 123, output: "number"},
        {input: 12.3, output: "number"},
        {input: true, output: "boolean"},
        {input: []string{"a", "b"}, output: "Array<string>"},
    }

    for _, tt := range tests {
        result := inferType(tt.input)
        if result != tt.output {
            t.Errorf("For input %v, expected %v but got %v", tt.input, tt.output, result)
        }
    }
}

func TestGenerateTypeScriptInterface(t *testing.T) {
    input := map[string]interface{}{
        "@collectionName": "bands",
        "name": "Opeth",
        "genre": "Progressive Metal",
        "origin": "Sweden",
        "formed": 1990,
    }

    collectionName := "bands"

    expected := `interface Bands {
    name: string;
    genre: string;
    origin: string;
    formed: number;
}`
    result := GenerateTypeScriptInterface(input, collectionName)

    if result != expected {
        t.Errorf("Expected %s but got %s", expected, result)
    }
}
