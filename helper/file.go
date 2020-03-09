package helper

import (
	"encoding/json"
	userPB "github.com/G0tYou/user-service/proto"
	"io/ioutil"
)

func ParseFile(file string)(*userPB.User, error){
	var user *userPB.User
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(data, &user)
	if err!=nil{
		return nil,err
	}
	return user, err
}