package models

import (
	"github.com/RichardKnop/jsonhal"
)

type GetLabelsResponse struct {
	jsonhal.Hal
	Labels []Label `json:"labels"`
}

type PostCreateLabelResponse struct {
	jsonhal.Hal
	ID LabelID `json:"id"`
}
