package channels

func OK(chan bool) bool {
	select {
	case ok := <- done :
		if ok {
			return true
		}
		return false
	}
}