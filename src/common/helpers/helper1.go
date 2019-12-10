package helpers

import (
	"encoding/gob"
	log "github.com/sirupsen/logrus"
	"bytes"
)

func ConvertStructToBytes(val interface{}) (bytes.Buffer, error) {
	var network bytes.Buffer        // Stand-in for a network connection
	enc := gob.NewEncoder(&network) // Will write to network.
	//dec := gob.NewDecoder(&network) // Will read from network.
	// Encode (send) the value.
	err := enc.Encode(val)
	if err != nil {
		log.Error("encode error:", err)

		return bytes.Buffer{}, err
	}
	return network, nil
}



