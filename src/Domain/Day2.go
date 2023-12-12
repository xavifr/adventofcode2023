package Domain

type D2Set struct {
	Red, Green, Blue int
}

type D2Game struct {
	ID   int
	Sets []D2Set
}
