package shelter

type PetRepository interface {
	SelectAllPets() ([]Pet, error)
	SelectOnePet(id int) (*Pet, error)
	UpdatePetName(id int, newName string) error
}
