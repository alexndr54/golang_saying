package golang_saying

func Filter(s string) interface{} {
	if s == "anjing" {
		return "heh ga boleh gitu"
	}

	if s == "babi" {
		return "dibilangin jangan gitu"
	}

	return nil
}
