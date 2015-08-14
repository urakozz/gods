package sort

type Comparable interface {
	CompareTo(interface{}) int
}

type CString string

func(s CString) CompareTo(st interface{}) int {
	var str string
	cstr, ok := st.(CString)
	if !ok {
		str, ok = st.(string)
		if !ok {return 0}
		cstr = CString(str)
	}
	str = string(cstr)
	if string(s) == str {return 0}
	if string(s) > str { return 1}
	return -1
}