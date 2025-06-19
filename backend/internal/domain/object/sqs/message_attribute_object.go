package object

type MessageAttribute struct {
	MessageType AttributeMessageType // メッセージタイプ
	Paymethod   *AttributePaymethod  // 決済方法
}
