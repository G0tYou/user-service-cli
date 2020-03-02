package helper

import (
	"encoding/json"
	pb "github.com/G0tYou/user-service/proto"
	"io/ioutil"
)

func ParseFile(file string)(*pb.User, error){
	var user *pb.User
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