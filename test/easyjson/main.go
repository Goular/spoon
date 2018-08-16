package main

import "spoon/test/easyjson/model"

func main() {
	Test()
}

func Test() {
	u := model.User{Name: "yan1", Age: 12}
	user1, err := u.MarshalJSON()
	if err != nil {
		println("error")
	}
	println("after marshal user = ", user1)

	//json:=`{"id":11,"name":"yan","age":11}`
	user2 := model.User{}
	user2.UnmarshalJSON([]byte(user1))
	println("name = ", user2.Name)
}


