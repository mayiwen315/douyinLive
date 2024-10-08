package douyinlive

import (
	"fmt"
	"log"
	"testing"

	"github.com/mayiwen315/douyinLive/generated/douyin"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

func TestNewDouyinLive(t *testing.T) {
	d, err := NewDouyinLive("520305235818", "./jssrc/webmssdk.js", "")
	if err != nil {
		fmt.Println("err:"+err.Error())
	}
	d.Subscribe(func(eventData *douyin.Message) {
		if eventData.Method == WebcastChatMessage {
			msg := &douyin.ChatMessage{}
			proto.Unmarshal(eventData.Payload, msg)
			marshal, _ := protojson.Marshal(msg)
			//msg.String()
			log.Println("聊天msg", msg.User.Id, msg.User.NickName, msg.Content, string(marshal))
		}
		// if eventData.Method == WebcastGiftMessage {
		// 	msg := &douyin.GiftMessage{}
		// 	proto.Unmarshal(eventData.Payload, msg)
		// 	marshal, _ := protojson.Marshal(msg)
		// 	//msg.String()
		// 	log.Println("礼物msg", msg.Gift.Name, string(marshal))
		// }
		// if eventData.Method == WebcastMemberMessage {
		// 	msg := &douyin.MemberMessage{}
		// 	proto.Unmarshal(eventData.Payload, msg)
		// 	marshal, _ := protojson.Marshal(msg)
		// 	log.Println("用户msg", string(marshal))
		// }
		//log.Println(eventData.Method, string(eventData.Payload))
		//msg := &douyin.ChatMessage{}
		//
		//proto.Unmarshal(eventData.Payload, msg)
		//marshal, _ := protojson.Marshal(msg)
		////msg.String()
		//log.Println("聊天msg", msg.User.Id, msg.User.NickName, msg.Content, string(marshal), msg.String())
	})

	d.ErrorMsgHandler(func(err error) {
		log.Println("err", err)
	})
	d.Start()

}
