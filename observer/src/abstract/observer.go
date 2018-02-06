package abstract

type Observer interface {
	Update()
	SetSubject(s Subject)
}
