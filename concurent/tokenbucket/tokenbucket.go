package tokenbucket

const (
	tcount = 30
	freq   = 10
)

type TokenBucket interface {
	GetToken() bool
}
