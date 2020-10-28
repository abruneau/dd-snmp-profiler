package parsers

import "os"

// WriteToFile writes a byte array to a file
func WriteToFile(file string, content *[]byte) error {

	f, err := os.Create(file)
	if err != nil {
		return err
	}

	defer f.Close()

	f.Write(*content)
	f.Sync()

	return nil
}
