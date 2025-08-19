package davka16

import (
	"errors"
	"fmt"
	"io"
	"strings"

	"golang.org/x/text/encoding/charmap"
)

// Validate checks the batch for semantic rules per spec (II-5.16).
func Validate(h Header, docs []Document) error {
	if h.Dtyp != "16" {
		return errors.New("header DTYP must be '16'")
	}
	if len(docs) == 0 {
		return errors.New("at least one document required")
	}
	if h.Poc != 0 && h.Poc != len(docs) {
		return errors.New("header POC must match number of documents or be 0")
	}
	for i, doc := range docs {
		if len(doc.Naklads) < 1 {
			return fmt.Errorf("document %d: at least one 'U' line required", i+1)
		}
		if len(doc.Naklads) > 50 {
			return fmt.Errorf("document %d: max 50 'U' lines", i+1)
		}
		if len(doc.Sdelenis) > 20 {
			return fmt.Errorf("document %d: max 20 'S' lines", i+1)
		}
	}
	return nil
}

// Generate produces the full dávka content as []byte, encoded in CP852 (PC Latin-2).
// Auto-sets Header.Poc if it's 0.
// Lines end with CRLF.
func Generate(h Header, docs []Document) ([]byte, error) {
	if err := Validate(h, docs); err != nil {
		return nil, err
	}
	if h.Poc == 0 {
		h.Poc = len(docs)
	}

	var content strings.Builder
	content.WriteString(generateHeader(h) + "\r\n")

	for _, doc := range docs {
		content.WriteString(generateL(doc.L) + "\r\n")
		for _, u := range doc.Naklads {
			content.WriteString(generateU(u) + "\r\n")
		}
		for _, s := range doc.Sdelenis {
			content.WriteString(generateS(s) + "\r\n")
		}
	}

	// Convert UTF-8 to CP852
	encoder := charmap.CodePage852.NewEncoder()
	encoded, err := encoder.String(content.String())
	if err != nil {
		return nil, fmt.Errorf("failed to encode to CP852: %w", err)
	}
	return []byte(encoded), nil
}

// WriteTo writes the generated CP852-encoded content to an io.Writer.
func WriteTo(w io.Writer, h Header, docs []Document) error {
	data, err := Generate(h, docs)
	if err != nil {
		return err
	}
	_, err = w.Write(data)
	return err
}
