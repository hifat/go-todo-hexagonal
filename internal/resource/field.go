package resource

/*
	It is used for validation because some display field names and field names are different.
	Found an issue in the validator inconsistent with json fields.
*/
var Fields = map[string]string{
	"Username": "username",
	"Password": "password",
}
