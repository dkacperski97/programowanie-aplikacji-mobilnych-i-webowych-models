package models

import (
	"errors"
	"regexp"
	"strconv"
)

type (
	LabelID string
	Label   struct {
		ID             LabelID `json:"id"`
		Sender         string  `json:"sender"`
		Recipient      string  `json:"recipient"`
		Locker         string  `json:"locker"`
		Size           int     `json:"size"`
		AssignedParcel string  `json:"assignedParcel"`
	}
)

func CreateLabel(sender, recipient, locker, size string) (*Label, error, error) {
	sizeValue, err := strconv.Atoi(size)
	if err != nil {
		return nil, errors.New("Niepoprawny rozmiar paczki"), nil
	}
	validationErr, err := IsLabelValid(sender, recipient, locker, sizeValue)
	if validationErr != nil {
		return nil, validationErr, nil
	}
	if err != nil {
		return nil, nil, err
	}
	l := new(Label)
	l.Sender = sender
	l.Recipient = recipient
	l.Locker = locker
	l.Size = sizeValue
	return l, nil, nil
}

func IsLabelValid(sender, recipient, locker string, size int) (error, error) {
	matched, err := regexp.MatchString(`[a-z]{3,12}`, sender)
	if err != nil {
		return nil, err
	}
	if !matched {
		return errors.New("Niepoprawny nadawca"), nil
	}

	if len(recipient) == 0 && len(recipient) > 100 {
		return errors.New("Niepoprawny adresat"), nil
	}

	matched, err = regexp.MatchString(`[A-Z0-9]{9}`, locker)
	if err != nil {
		return nil, err
	}
	if !matched {
		return errors.New("Niepoprawny identyfikator skrytki"), nil
	}

	if size < 1 && size > 8000 {
		return errors.New("Niepoprawny rozmiar paczki"), nil
	}

	return nil, nil
}
