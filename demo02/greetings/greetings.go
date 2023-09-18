package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("name是空的")
	}
	// return fmt.Sprintf(randomFormat(), name), nil
	return fmt.Sprint(randomFormat()), nil
}

func randomFormat() string {
	formats := []string{
		"Hi, %v. Weclome!",
		"Creat to see you, %v!",
		"Hail, %v! Well met!",
	}
	return formats[rand.Intn(len(formats))]
}

func Hellos(names []string) (map[string]string, error) {
	messages := make(map[string]string)

	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}
		messages[name] = message
	}
	return messages, nil
}
