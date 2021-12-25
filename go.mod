module hellogo

go 1.14

require (
	github.com/jaecheon/greetings v0.0.0-20211225104851-f6a0fff9382b
	jclee.com/greeting2 v0.0.0-00010101000000-000000000000
)

replace jclee.com/greeting2 => ./greeting2
