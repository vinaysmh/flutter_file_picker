package file_picker

import (
	"strings"

	"github.com/gen2brain/dlgs"
	"github.com/pkg/errors"
)

func fileFilter(method string, extensions []string, size int) (string, error) {
	switch method {
	case "any":
		return `*.*`, nil
	case "image":
		return `*.png *.jpg *.jpeg`, nil
	case "audio":
		return `*.mp3`, nil
	case "video":
		return `*.webm *.mpeg *.mkv *.mp4 *.avi *.mov *.flv`, nil
	case "custom":
		var i int
		var filters = ""
		for i = 0 ; i<size ; i++ {
	  		filters += `*.` + extensions[i] + ` `
		}
		return filters, nil
	default:
		return "", errors.New("unknown method")
	}

}

func fileDialog(title string, filter string) (string, error) {
	filePath, _, err := dlgs.File(title, filter, false)
	if err != nil {
		return "", errors.Wrap(err, "failed to open dialog picker")
	}
	return filePath, nil
}
