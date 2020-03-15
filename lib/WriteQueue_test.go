package lib_test

import (
	"github.com/randomSignal/m3u8-downloader/lib"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestDownload(t *testing.T) {
	err := lib.WriteQueue("https://bp1.dkkomo.com/stream/full/japan/3000/heyzo_2166.m3u8", "/Users/liucx/Desktop/data")
	require.NoError(t, err)
}
