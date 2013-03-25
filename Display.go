package gowl

/*

   The core global object.  This is a special singleton object.  It
   is used for internal Wayland protocol features.

*/

type Display struct{}

/*

	These errors are global and can be emitted in response to any
	server request.

*/
var Error int

const (
	InvalidObject = 0

	InvalidMethod = 1

	NoMemory = 2
)

func (*Display) Sync(Callback new_id) {
}

func (*Display) GetRegistry(Callback new_id) {
}

func (*Display) Error(ObjectId object, Code uint, Message string) {
}

func (*Display) DeleteId(Id uint) {
}