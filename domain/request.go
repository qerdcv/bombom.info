package domain

import (
	"fmt"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type JoinRequest struct {
	UserTag      string `json:"userTag"`
	TelegramName string `json:"telegramName"`
}

func (jr JoinRequest) String() string {
	return fmt.Sprintf("user tag: %s", jr.UserTag)
}

func (jr JoinRequest) Validate() error {
	return validation.ValidateStruct(&jr,
		validation.Field(&jr.UserTag, validation.Required),
		validation.Field(&jr.TelegramName, validation.Required),
	)
}
