package main 	//porque estoy en el mismo nivel

import "testing"
import "tudai2021.com/model"
import "github.com/stretchr/testify/assert"

func TestchangeName(t *testing.T){

	p := model.NewPerson(1,"Maria")
	changeName(&p,"Agustin")

	assert.Equal(t,p.Name,"Agustin","Los valores no coinciden") //en vez del if
}