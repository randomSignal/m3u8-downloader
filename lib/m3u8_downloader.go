package lib

import (
	"context"
	"fmt"
	m3u8_decoder "github.com/changxiliu/m3u8-decoder"
)

func WriteQueue(m3u8Url string, filePath string) error {
	var err error

	fn := func() (string, error) {
		return m3u8Url, nil
	}

	callback := func(ts m3u8_decoder.M3u8Ts) error {
		var job Job
		job.FilePath = filePath
		job.TsUrl = ts.Url
		fmt.Println("WriteQueue job.TsUrl:", job.TsUrl)
		Queue <- job

		return nil
	}

	fmt.Println("before lib.WriteQueue")
	ctx, _ := context.WithCancel(context.Background())
	err = m3u8_decoder.NewM3u8Decoder(fn).WithContext(ctx).StartDecode(callback)
	if err != nil {
		return err
	}
	fmt.Println("end lib.WriteQueue")
	return nil
}
