package helpers

import (
	"github.com/google/uuid"
)

func NewUUID()(string,error){
	val,err:= uuid.NewUUID()
	if err!=nil{
		return "", err
	}
	return val.String(),nil
}
