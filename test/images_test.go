package test

import "testing"

func TestUploadImage(t *testing.T) {
	var body struct {
		Name string `json:"title"`
		Path string `json:"url"`
	}
	body.Name = "test"
	body.Path = "test"
	if body.Name == "" {
		t.Errorf("Name is empty")
	}
	if body.Path == "" {
		t.Errorf("Path is empty")
	}

}
