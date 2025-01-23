package internal

import "3-1/internal/contract"

func NewVideo(title, description string, length int) contract.Video {
	return contract.Video{
		Title:       title,
		Description: description,
		Length:      length,
	}
}
