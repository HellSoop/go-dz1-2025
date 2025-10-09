package task3

func ScaleSlice(slice *[]int, scaleFactor uint32) error {
	if slice == nil {
		return nil
	}

	if scaleFactor == 0 {
		*slice = nil
		return nil
	}

	max_capacity := ^uint32(0) // zero with inverted bits is the max value
	basic_length := uint32(len(*slice))

	if basic_length > max_capacity/scaleFactor {
		return ErrOverflow
	}

	for i := uint32(1); i < scaleFactor; i++ {
		*slice = append(*slice, (*slice)[:basic_length]...)
	}
	return nil
}
