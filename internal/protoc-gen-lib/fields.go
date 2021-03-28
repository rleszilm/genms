package protocgenlib

// Fields is a struct that contains data about the messages fields.
type Fields struct {
	fieldNames   []string
	fieldsByName map[string]*Field
}

// Fields returns the field names.
func (f *Fields) Names() []string {
	return f.fieldNames
}

// ByName returns the specified field.
func (f *Fields) ByName(n string) *Field {
	return f.fieldsByName[n]
}

// NewFields returns a new Fields.
func NewFields(msg *Message) *Fields {
	fieldNames := []string{}
	fieldsByName := map[string]*Field{}

	for _, f := range msg.message.Fields {
		field := NewField(msg, f)
		fieldNames = append(fieldNames, field.Name())
		fieldsByName[field.Name()] = field
	}

	return &Fields{
		fieldNames:   fieldNames,
		fieldsByName: fieldsByName,
	}
}
