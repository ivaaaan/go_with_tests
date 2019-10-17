package main

import "errors"

type Dictionary map[string]string

var (
	ErrNotFound         = errors.New("could not find the word you were looking for")
	ErrWordExists       = errors.New("cannot add word because it already exists")
	ErrWordDoesNotExist = errors.New("cannot update word because it does not exist")
)

func (dictionary Dictionary) Search(needle string) (string, error) {
	definition, ok := dictionary[needle]
	if !ok {
		return "", ErrNotFound
	}
	return definition, nil
}

func (dictionary Dictionary) Add(key, value string) error {
	_, err := dictionary.Search(key)
	switch err {
	case ErrNotFound:
		dictionary[key] = value
	case nil:
		return ErrWordExists
	default:
		return err
	}

	return nil
}

func (dictionary Dictionary) Update(key, value string) error {
	_, err := dictionary.Search(key)
	switch err {
	case ErrNotFound:
		return ErrWordDoesNotExist
	case nil:
		dictionary[key] = value
	default:
		return err
	}
	return nil
}

func (dictionary Dictionary) Delete(key string) {
	delete(dictionary, key)
}
