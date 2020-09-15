package smartconf

type Profile interface {
	BeforeParse()
	AfterParse()
}
