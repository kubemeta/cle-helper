package unit

type TimeUnit string

const (
	Days         = TimeUnit("d")
	Hours        = TimeUnit("h")
	Minutes      = TimeUnit("m")
	Seconds      = TimeUnit("s")
	Milliseconds = TimeUnit("ms")
	Microseconds = TimeUnit("micros")
	Nanoseconds  = TimeUnit("nanos")
)

type ByteSizeUnit string

const (
	Bytes     = ByteSizeUnit("b")
	Kilobytes = ByteSizeUnit("kb")
	Megabytes = ByteSizeUnit("mb")
	Gigabytes = ByteSizeUnit("gb")
	Terabytes = ByteSizeUnit("tb")
	Petabytes = ByteSizeUnit("pb")
)
