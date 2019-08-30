package wraptypes

//WrapBool retorna um ponteiro a partir do primitivo
func WrapBool(b bool) *bool {
	wrap := b
	return &wrap
}

//WrapInt32 retorna um ponteiro a partir do primitivo
func WrapInt32(i int32) *int32 {
	wrap := i
	return &wrap
}

//WrapString retorna um ponteiro a partir do primitivo
func WrapString(s string) *string {
	wrap := s
	return &wrap
}

//WrapFloat32 retorna um ponteiro a partir do primitivo
func WrapFloat32(i float32) *float32 {
	wrap := i
	return &wrap
}
