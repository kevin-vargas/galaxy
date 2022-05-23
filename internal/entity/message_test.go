package entity

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Message(t *testing.T) {
	testCases := []struct {
		desc      string
		messages  [][]string
		expect    string
		errExpect error
	}{
		{
			"Simple",
			[][]string{{"este,", "", "", "mensaje", ""}, {"", "es", "", "", "secreto"}, {"este,", "", "un", "", ""}},
			"este, es un mensaje secreto",
			nil,
		},
		{
			"Simple error",
			[][]string{{"este,", "", "", "mensaje", "", ""}, {"", "es", "", "", "secreto"}, {"este,", "", "un", "", ""}},
			"",
			errors.New("invalid message"),
		},
	}

	for _, tt := range testCases {
		t.Run(tt.desc, func(t *testing.T) {
			// arrange
			msg := NewMessage()

			// act
			for _, message := range tt.messages {
				_, err := msg.Add(message)
				if err != nil {
					t.FailNow()
				}
			}
			result, errResult := msg.Build()

			// assert
			assert.Equal(t, tt.expect, result)
			assert.Equal(t, tt.errExpect, errResult)
		})
	}
}
