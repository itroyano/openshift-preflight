package container

import (
	"time"
)

var (
	checkContainerTimeout time.Duration = 10 * time.Second
	waitContainer         time.Duration = 2 * time.Second
)