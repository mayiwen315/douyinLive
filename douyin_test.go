package douyinlive

import (
	"DouyinLive/generated/douyin"
	"log"
	"testing"

	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

func TestNewDouyinLive(t *testing.T) {
	d, _ := NewDouyinLive("96306655034")
	d.Subscribe(func(eventData *douyin.Message) {
		// if eventData.Method == WebcastChatMessage {
		// 	msg := &douyin.ChatMessage{}

		// 	proto.Unmarshal(eventData.Payload, msg)
		// 	marshal, _ := protojson.Marshal(msg)
		// 	//msg.String()
		// 	log.Println("聊天msg", msg.User.Id, msg.User.NickName, msg.Content, string(marshal))
		// }

		if eventData.Method == WebcastGiftMessage {
			msg := &douyin.GiftMessage{}

			proto.Unmarshal(eventData.Payload, msg)
			marshal, _ := protojson.Marshal(msg)
			//msg.String()
			log.Println("礼物msg", msg.User.Id, msg.User.NickName, string(marshal))
		}
		//log.Println(eventData.Method, string(eventData.Payload))
		//msg := &douyin.ChatMessage{}
		//
		//proto.Unmarshal(eventData.Payload, msg)
		//marshal, _ := protojson.Marshal(msg)
		////msg.String()
		//log.Println("聊天msg", msg.User.Id, msg.User.NickName, msg.Content, string(marshal), msg.String())
	})

	d.Start()

}
