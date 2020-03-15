package lib_test

import (
	"github.com/randomSignal/m3u8-downloader/lib"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestTsDownloader(t *testing.T) {
	err := lib.TsDownloader("https://bp1.dkkomo.com/stream/full/japan/3000/heyzo_2166.m3u8", "/Users/liucx/Desktop/data")
	require.NoError(t, err)

	err = lib.TsDownloader("http://bp1.dkkomo.com/stream/full/japan/3000/heyzo_2166/000.ts", "/Users/liucx/Desktop/data")
	require.NoError(t, err)
}
