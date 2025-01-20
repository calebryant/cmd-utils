package stdin

import (
	"io"
	"os"
)

func ReadFromStdin() ([]byte, error) {
	file := os.Stdin
	fileStat, err := file.Stat()
	if err != nil {
		return nil, err
	}
	// checks if there is a pipe
	if (fileStat.Mode() & os.ModeCharDevice) == 0 {
		return []byte{}, nil
	}
	// check if the pipe has data in it
	if fileStat.Size() <= 0 {
		return []byte{}, nil
	}
	return io.ReadAll(file)
}
