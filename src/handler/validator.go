package handler

import (
	"ful/RESTful/src/storage"
)

func IsIdExists(s *storage.Storage, id int) bool {
	al := (*s).GetAlbums()
	for _, item := range al {
		if item.ID == id {
			return true 
		}
	}
	return false
}