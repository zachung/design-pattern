package internal

import "3-1/internal/contract"

func NewVideo(title, description string, length int) *interface{} {
	video := &contract.Video{
		Title:       title,
		Description: description,
		Length:      length,
	}
	i := interface{}(video)

	return &i
}
