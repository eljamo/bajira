package toml

import (
	"log"
	"os"

	"github.com/BurntSushi/toml"
	"github.com/eljamo/bajira/internal/errorconc"
)

// encodeToFile encodes the given struct into a .toml file
func EncodeToFile[T any](data T, path string) error {
	file, err := os.Create(path)
	if err != nil {
		return errorconc.LocalizedError(err, "failed to create file")
	}
	defer func() {
		if err := file.Close(); err != nil {
			log.Printf(`failed to close file "%v": %v`, file, err)
		}
	}()

	encoder := toml.NewEncoder(file)
	if err := encoder.Encode(data); err != nil {
		return errorconc.LocalizedError(err, "failed to encode data")
	}

	return nil
}

// DecodeFromFile decodes the .toml file into the given struct
func DecodeFromFile[T any](path string, result *T) error {
	if _, err := toml.DecodeFile(path, result); err != nil {
		return errorconc.LocalizedError(err, "failed to decode file")
	}

	return nil
}
