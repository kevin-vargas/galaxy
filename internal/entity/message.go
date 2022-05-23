package entity

import (
	"errors"
	"strings"
)

const (
	DEFAULT_JOIN_VALUE = " "
)

type Message struct {
	acum      map[int]string
	joinValue string
}

type MessageOptions func(*Message)

func NewMessage(options ...MessageOptions) *Message {
	msg := &Message{
		acum:      make(map[int]string),
		joinValue: DEFAULT_JOIN_VALUE,
	}
	for _, option := range options {
		option(msg)
	}
	return msg
}

func (m *Message) Build() (string, error) {
	lastOne := len(m.acum) - 1
	var sb strings.Builder
	for i := 0; i <= lastOne; i++ {
		acum := m.acum[i]
		if acum == "" {
			return "", errors.New("invalid message")
		}
		_, err := sb.WriteString(acum)
		if i != lastOne {
			_, err = sb.WriteString(m.joinValue)
		}
		if err != nil {
			return "", err
		}
	}
	return sb.String(), nil
}

func (m *Message) Add(strs []string) (*Message, error) {
	for i, str := range strs {
		value := m.acum[i]
		if value == "" {
			m.acum[i] = str
			continue
		}
		if str != "" && value != str {
			return nil, errors.New("two different values")
		}
	}
	return m, nil
}
