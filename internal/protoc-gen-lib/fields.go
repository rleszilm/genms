package protocgenlib

// Fields is a struct that contains data about the messages fields.
type Fields struct {
	fields       []*Field
	fieldNames   []string
	fieldsByName map[string]*Field
}

// Fields returns the field names.
func (f *Fields) Fields() []string {
	return f.fieldNames
}

// ByName returns the specified field.
func (f *Fields) ByName(n string) *Field {
	return f.fieldsByName["f:-"+n]
}

// NewFields returns a new Fields.
func NewFields(msg *Message) *Fields {
	fields := []*Field{}
	fieldNames := []string{}
	fieldsByName := map[string]*Field{}

	for _, f := range msg.Message.Fields {
		field := NewField(msg.Outfile, msg.Message, f)
		fields = append(fields, field)
		fieldNames = append(fieldNames, field.Name())
		fieldsByName[field.Name()] = field
	}

	return &Fields{
		fields:       fields,
		fieldNames:   fieldNames,
		fieldsByName: fieldsByName,
	}
}
