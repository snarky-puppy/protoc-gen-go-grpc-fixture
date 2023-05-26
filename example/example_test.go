package example

import "testing"

func Test_Petstore(t *testing.T) {
	// create petstore fixture client
	petstore := NewPetStoreFixtures("fixtures")
	// inputs don't matter
	pet, _ := petstore.CreatePet(nil, nil)
	// outputs are loaded from fixtures
	if pet == nil {
		t.Fail()
	}
	if pet.Name != "Fido" {
		t.Fail()
	}
}
