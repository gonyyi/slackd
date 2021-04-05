module github.com/gonyyi/slackd

go 1.16

require (
	github.com/aws/aws-sdk-go v1.38.12 // indirect
	github.com/gonyyi/afmt v1.0.0 // indirect
	github.com/gonyyi/agraceful v0.1.0 // indirect
	github.com/gonyyi/alog v0.7.7
	github.com/gonyyi/areq v0.1.0
	github.com/gonyyi/atype v0.0.0-20210331015131-bacc33c0d2cb
	github.com/gonyyi/gointf v0.0.0-20210325184721-c115720eadd2
	github.com/gonyyi/gointf/log_alog v0.0.0-20210325184721-c115720eadd2
	github.com/gonyyi/gointf/store_bbolt v0.0.0-20210325184721-c115720eadd2
	github.com/gonyyi/mutt v0.0.0-20210331023810-9333a661a62b
	golang.org/x/crypto v0.0.0-20210322153248-0c34fe9e7dc2 // indirect
	golang.org/x/sys v0.0.0-20210324051608-47abb6519492 // indirect
)

replace (
	github.com/gonyyi/atype => /Users/gonyi/go/src/github.com/gonyyi/atype
	github.com/gonyyi/gointf => /Users/gonyi/go/src/github.com/gonyyi/gointf
	github.com/gonyyi/mutt => /Users/gonyi/go/src/github.com/gonyyi/mutt
	github.com/gonyyi/slackd/... => /Users/gonyi/go/src/github.com/gonyyi/slackd/...
)
