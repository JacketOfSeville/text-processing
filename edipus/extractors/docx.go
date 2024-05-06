package extractors

import (
	"bytes"
	"regexp"
	"strings"

	"github.com/Gustrb/text-processing/edipus/models"
	"github.com/nguyenthenguyen/docx"
)

func extractTextFromDocxXML(contentXML string) string {
	// Define the regular expression to match XML tags
	re := regexp.MustCompile(`<[^>]+>`)

	// Replace all occurrences of XML tags with an empty string
	stripped := re.ReplaceAllString(contentXML, "")

	//trim
	return strings.TrimSpace(stripped)
}

func ExtractTextFromDocx(fileToProcess models.FileToProcess) (string, error) {
	// Create a reader from fileToProcess.Content
	// Use the reader to read the content of the docx file
	reader := bytes.NewReader(fileToProcess.Content)

	r, err := docx.ReadDocxFromMemory(reader, int64(len(fileToProcess.Content)))
	if err != nil {
		return "", err
	}

	defer r.Close()

	contentXML := r.Editable().GetContent()

	return extractTextFromDocxXML(contentXML), nil
}
