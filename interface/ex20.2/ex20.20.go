package main
import "LEARNGOLANG/interface/koreaPost"
type Sender interface {
	Send(parcel string)
}
func SendBook(name string,sender Sender) {
	sender.Send(name)
}
func main() {
	var sender Sender = &koreaPost.PostSender{}
	SendBook("어린왕자",sender)
	SendBook("으으응ㅇ", sender)
}