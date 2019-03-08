package mondia

import "github.com/astaxie/beego/orm"

func init() {
	orm.RegisterModel(new(Mo), new(Notification), new(AffTrack), new(Postback),new(MdSubscribe),
		new(UnsubPin),  new(AlreadySub),)
}

func MoTBName() string {
	return "mo"
}

func AffTrackTBName() string{
	return "aff_track"
}

func NotificationTBName() string {
	return "notification"
}

func PostbackTBName() string {
	return "postback"
}

func WapResponseTBName()string{
	return "wap_response"
}
