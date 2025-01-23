package main

import (
	"3-1/internal"
	"3-1/internal/contract"
	"log"
)

func main() {
	log.SetFlags(0)

	// given
	pewDiePie := internal.NewChannel("PewDiePie")
	waterballAcademy := internal.NewChannel("水球軟體學院")
	var waterball contract.ChannelSubscriber
	waterball = internal.NewChannelSubscriber("水球", func(video contract.Video) {
		if video.Length >= 3*60 {
			log.Printf("%s 對影片 \"%s\" 按讚。", waterball.GetName(), video.Title)
		}
	})
	var fireball contract.ChannelSubscriber
	fireball = internal.NewChannelSubscriber("火球", func(video contract.Video) {
		if video.Length <= 60 {
			video.Channel.UnSubscribe(fireball)
		}
	})

	// when
	waterballAcademy.Subscribe(waterball)
	pewDiePie.Subscribe(waterball)
	waterballAcademy.Subscribe(fireball)
	pewDiePie.Subscribe(fireball)
	waterballAcademy.Upload(internal.NewVideo("C1M1S2", "這個世界正是物件導向的呢！", 4*60))
	pewDiePie.Upload(internal.NewVideo("Hello guys", "Clickbait", 30))
	waterballAcademy.Upload(internal.NewVideo("C1M1S3", "物件 vs. 類別", 60))
	pewDiePie.Upload(internal.NewVideo("Minecraft", "Let’s play Minecraft", 30*60))
	/**
	水球 訂閱了 水球軟體學院。
	水球 訂閱了 PewDiePie。
	火球 訂閱了 水球軟體學院。
	火球 訂閱了 PewDiePie。
	頻道 水球軟體學院 上架了一則新影片 "C1M1S2"。
	水球 對影片 "C1M1S2" 按讚。
	頻道 PewDiePie 上架了一則新影片 "Hello guys"。
	火球 解除訂閱了 PewDiePie。
	頻道 水球軟體學院 上架了一則新影片 "C1M1S3"。
	火球 解除訂閱了 水球軟體學院。
	頻道 PewDiePie 上架了一則新影片 "Minecraft"。
	水球 對影片 "Minecraft" 按讚。
	*/
}
