// Package all conveniently loads all the inbuilt/supported host drivers.
package all

import (
	_ "github.com/alittlebrighter/embd/host/bbb"
	_ "github.com/alittlebrighter/embd/host/rpi"
)
