package o

func Try(f func(), c func(Throwable)) {
	defer func() {
		if e := recover(); e != nil {
			c(e.(Throwable))
		}
	}()
}

func Throw(t Throwable) {
	panic(t)
}
