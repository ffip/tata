package text

import "C"
import "unsafe"

// https://stackoverflow.com/questions/36188649/cgo-char-to-slice-string?rq=1
func toGoStrings(argc C.int, argv **C.char) []string {
	length := int(argc)

	if length > 0 {
		tmpSlice := (*[1 << 30]*C.char)(unsafe.Pointer(argv))[:length:length]
		goStrings := make([]string, length)

		for i, s := range tmpSlice {
			goStrings[i] = C.GoString(s)
		}

		return goStrings
	}

	return make([]string, 0)
}
