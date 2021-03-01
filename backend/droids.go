package main

type DroidModel string

const (
	DroidModelC3PO DroidModel = "C3-PO"
	DroidModelR2D2 DroidModel = "R2-D2"
)

type droid struct {
	Name string `json:"name"`
	Model DroidModel `json:"model"`
}

// makeshift database
var droids []*droid

func getDroidByName(name string) *droid {
	for _, d := range droids {
		if d.Name == name {
			return d
		}
	}
	return nil
}

func getDroidByModel(model DroidModel) *droid {
	for _, d := range droids {
		if d.Model == model {
			return d
		}
	}
	return nil
}