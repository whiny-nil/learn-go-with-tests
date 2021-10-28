package dictionary

type Dictionary map[string]string

const (
	ErrWordNotFound = DictionaryError("search term not in dictionary")
	ErrWordExists   = DictionaryError("term exists in dictionary")
)

type DictionaryError string

func (e DictionaryError) Error() string {
	return string(e)
}

func (d Dictionary) Search(searchTerm string) (string, error) {
	definition, ok := d[searchTerm]

	if ok {
		return definition, nil
	} else {
		return "", ErrWordNotFound
	}
}

func (d Dictionary) Add(term, definition string) error {
	_, err := d.Search(term)

	switch err {
	case ErrWordNotFound:
		d[term] = definition
		return nil
	case nil:
		return ErrWordExists
	default:
		return err
	}
}

func (d Dictionary) Update(term, definition string) error {
	_, err := d.Search(term)

	switch err {
	case nil:
		d[term] = definition
		return nil
	default:
		return err
	}
}

func (d Dictionary) Delete(term string) {
	delete(d, term)
}
