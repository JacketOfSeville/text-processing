package extractors

import (
	"github.com/Gustrb/text-processing/edipus/models"
)

func ExtractTextFromTxt(fileToProcess models.FileToProcess) (string, error) {
	return string(fileToProcess.Content), nil
}
