package main

import (
	"errors"

	"github.com/nathancastelein/go-course-api/shelter"
	"github.com/stretchr/testify/require"
)

type petRepositoryStub struct {
	require *require.Assertions
}

func (p petRepositoryStub) SelectAllPets() ([]shelter.Pet, error) {
	return []shelter.Pet{
		{
			Id:       1,
			Name:     "Daffy",
			Category: "duck",
		},
		{
			Id:       2,
			Name:     "Mickey",
			Category: "mouse",
		},
	}, nil
}

func (p petRepositoryStub) SelectOnePet(id int) (*shelter.Pet, error) {
	p.require.Equal(1, id, "expect id 1")
	return &shelter.Pet{
		Id:       1,
		Name:     "Daffy",
		Category: "duck",
	}, nil
}

func (p petRepositoryStub) UpdatePetName(id int, newName string) error {
	p.require.Equal(1, id, "expect id 1")
	p.require.Equal("Daisy", newName, "expect new name Daisy")

	return nil
}

type petRepositorySpy struct {
	petRepositoryStub
	updateCalled bool
}

func (p *petRepositorySpy) UpdatePetName(id int, newName string) error {
	p.petRepositoryStub.UpdatePetName(id, newName)
	p.updateCalled = true

	return nil
}

type petRepositoryStubError struct {
}

func (p petRepositoryStubError) SelectAllPets() ([]shelter.Pet, error) {
	return nil, errors.New("an error message")
}

func (p petRepositoryStubError) SelectOnePet(id int) (*shelter.Pet, error) {
	return nil, errors.New("an error message")
}

func (p petRepositoryStubError) UpdatePetName(id int, newName string) error {
	return errors.New("an error message")
}
