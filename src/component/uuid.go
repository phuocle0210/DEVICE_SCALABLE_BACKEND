package component

import "github.com/google/uuid"

type TUUID struct {
	UUID     uuid.UUID
	UUIDText string
}

func UUID() *TUUID {
	id := uuid.New()
	return &TUUID{
		UUID:     id,
		UUIDText: id.String(),
	}
}
