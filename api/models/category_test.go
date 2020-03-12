package models

func TestDescriptionModel(t *testing.T) {
	c := &Category{}
	c.Description = "Smart Phone"
	if err := c.Validate(); err != nil {
		t.Error(err)
	}
}