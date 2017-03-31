package keva

type kevaError struct {}
func (ke kevaError) Error() string {
	return "Keva Error: Unable to parse file or file not found: Honestly.... figure it out"
}
