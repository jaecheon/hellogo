module hellogo

go 1.14

require (
	github.com/jaecheon/greetings v0.0.0-20211225104851-f6a0fff9382b
	jclee.com/gochannel v0.0.0-00010101000000-000000000000
	jclee.com/gocrawler v0.0.0-00010101000000-000000000000
	jclee.com/gogeneric v0.0.0-00010101000000-000000000000
	jclee.com/goroutine v0.0.0-00010101000000-000000000000
	jclee.com/greeting2 v0.0.0-00010101000000-000000000000
)

replace jclee.com/greeting2 => ./greeting2

replace jclee.com/goroutine => ./goroutine

replace jclee.com/gochannel => ./channel

replace jclee.com/gogeneric => ./generic

replace jclee.com/gocrawler => ./crawler
