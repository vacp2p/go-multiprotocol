package multiprotocol

import (
	"encoding/csv"
	"io"
	"os"
	"strconv"
	"strings"

	"github.com/multiformats/go-multiaddr"
)

// Init initializes multiprotocol with a CSV file
func Init(path string) error {

	file, err := os.Open(path)
	if err != nil {
		return err
	}

	r := csv.NewReader(file)

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}

		if err != nil {
			return err
		}

		if record[0] == "code" {
			continue
		}

		p, err := protocol(record)
		if err != nil {
			return err
		}

		err = AddProtocol(*p)
		if err != nil {
			return err
		}

	}

	return nil
}

func protocol(vals []string) (*Protocol, error) {
	code, err := strconv.Atoi(strings.TrimSpace(vals[0]))
	if err != nil {
		return nil, err
	}

	size, err := size(vals[1])
	if err != nil {
		return nil, err
	}

	return &Protocol{
		Name: strings.TrimSpace(vals[2]),
		Code: code,
		VCode: multiaddr.CodeToVarint(code),
		Size: size,
		Transcoder: multiaddr.TranscoderUnix,
	}, nil
}

func size(u string) (int, error) {
	u = strings.TrimSpace(u)

	if u == "0" {
		return 0, nil
	}

	if u == "V" {
		return LengthPrefixedVarSize, nil
	}

	return strconv.Atoi(u)
}
