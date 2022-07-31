package main

import (
	tf "SeirekiWarekiConversionApp/pkg/TimeFlag"
	wareki "SeirekiWarekiConversionApp/pkg/Wareki"
	"flag"
	"fmt"
	"log"
	"time"
)

var t = &time.Time{}

func main() {

	// Time argument. Check `pkg/TimeFlag/timeFlag.go` for details
	flag.Var(&tf.TimeValue{t}, "seireki", "Time to parse 20060102")
	flag.Var(&tf.TimeValue{t}, "s", "Time to parse 20060102")

	var (
		i = flag.Int("int", 0, "int flag")
		b = flag.Bool("bool", false, "bool flag")
	)
	flag.Parse()

	// 初期値の場合はフラグ設定されていないとみなす。
	if t.String() == "0001-01-01 00:00:00 +0000 UTC" {
		log.Fatal("time引数が利用されていません。")
	} else {
		result, err := wareki.ConvertToWareki(t)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(result)
	}

	if *i != 0 {
		fmt.Println(*i)
	}

	if *b {
		fmt.Println(*b)
	}
}
