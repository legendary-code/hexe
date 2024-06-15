package unsafe

import "io"

func CloseIgnoreError(closer io.Closer) {
	_ = closer.Close()
}
