package cutils

import "github.com/google/uuid"

// CreateUUIDString returns new random UUID string.
func CreateUUIDString() (string, error) {
	res, err := CreateUUID()
	if err != nil {
		return "", err
	}

	return res.String(), nil
}

// CreateUUID returns new random UUID.
func CreateUUID() (*uuid.UUID, error) {
	res, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	return &res, err
}
