package appendtest

func appendStringsToInterfaceSlice1(ifs []interface{}, ss []string) []interface{} {
	if cap(ifs) < len(ifs)+len(ss) {
		out := make([]interface{}, len(ifs)+len(ss))
		n := copy(out, ifs)
		for i, s := range ss {
			out[n+i] = s
		}
		return out
	}
	for _, s := range ss {
		ifs = append(ifs, s)
	}
	return ifs
}

func appendStringsToInterfaceSlice2(ifs []interface{}, ss []string) []interface{} {
	ss2 := make([]interface{}, len(ss))
	for i, s := range ss {
		ss2[i] = s
	}

	return append(ifs, ss2...)
}

func appendStringsToInterfaceSlice3(ifs []interface{}, ss []string) []interface{} {
	for _, s := range ss {
		ifs = append(ifs, s)
	}

	return ifs
}
