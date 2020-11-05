package main

import (
	"context"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"grpc-go-demo/database"
	"grpc-go-demo/proto"
	"io/ioutil"
	"log"
	"net/http"
)
type server struct {

}

func (s *server) GetUsers(ctx context.Context,e interface{}) (*proto.UserListResponse, error) {
	db, err := database.GetDatabase()
	if err!=nil{
		panic(err.Error())
	}

	query, err := db.Query("select * from user")
	if err!=nil{
		panic(err.Error())
	}

	res := []*proto.UserRequest{}
	for query.Next(){
		var u *proto.UserRequest
		err := query.Scan(&u.Email,&u.Password,&u.Id)
		if err!=nil{
			panic(err.Error())
		}
		res = append(res, u)
	}
	return &proto.UserListResponse
}


func (s *server) GetUser(ctx context.Context,request *proto.UserRequest) (response *proto.UserResponse,err error) {
	db, err := database.GetDatabase()
	user_id := request.GetId()
	fmt.Println(user_id)
	if err!=nil{
		panic(err.Error())
	}
	query, err := db.Query("SELECT * FROM user WHERE id=?",user_id)
	if err!=nil{
		panic(err.Error())
	}
	fmt.Println("Query",query)
	res := User{}
	for query.Next(){
		var user User
		err := query.Scan(&user.Email,&user.Password,&user.Id)
		if err!=nil{
			panic(err.Error())
		}
		fmt.Println(user.Email,user.Password)
		res = user
	}
	fmt.Println(query.Next())
	return &proto.UserResponse
}

func (s *server) AddUser(ctx context.Context,request *proto.UserRequest) (*proto.UserStringResponse, error) {
	db, err := database.GetDatabase()
	user_email := request.GetEmail()
	user_password := request.GetPassword()
	user_id := request.GetId()
	if err!=nil{
		panic(err.Error())
	}
	result, err := db.Query("insert into user values(?,?,?)",user_email,user_password,user_id)
	if err!=nil{
		panic(err.Error())
	}
	fmt.Println("Result", result.Next())

	
}

func (s *server) DeleteUser(ctx context.Context,request *proto.UserRequest) (*proto.UserStringResponse, error) {
	db, err := database.GetDatabase()
	user_id := request.GetId()
	if err!=nil{
		panic(err.Error())
	}
	result, err := db.Query("DELETE FROM user WHERE id=?",user_id)
	if err!=nil{
		panic(err.Error())
	}


}

func (s *server) UpdateUser(context.Context, *proto.UserRequest) (*proto.UserStringResponse, error) {
	panic("implement me")
}

type User struct{
	Email string `json:"email"`
	Password int64 `json:"password"`
	Id int64 `json:"id"`
}

func index_handle(w http.ResponseWriter, r *http.Request) {
	db, err := database.GetDatabase()
	if err!=nil{
		panic(err.Error())
	}
	result, err := db.Query("Select * from user where id=200")
	if err!=nil{
		panic(err.Error())
	}
	res := []User{}
	for result.Next() {
		var user User
		err := result.Scan(&user.Email,&user.Password,&user.Id)
		if err!=nil{
			panic(err.Error())
		}
		fmt.Println(user.Password)
		res = append(res,user)
	}
	ress, err := json.Marshal(res)
	if err!=nil{
		panic(err.Error())
	}
	w.Write(ress)
}

	

func main() {
	listener, err := net.Listen("tcp", ":4040")
	if err != nil {
		panic(err)
	}

	srv := grpc.NewServer()

	proto.RegisterAddServiceServer(srv, &server{})



	reflection.Register(srv)

	if e := srv.Serve(listener); e!= nil {
		panic(e)
	}
}
*/
func (s *server) Add(ctx context.Context, request *proto.Request) (*proto.Response, error){
	fmt.Println("Service hitted")
	a, b := request.GetA(), request.GetB()
	result := a+b
	fmt.Println("Service hitted")
	return &proto.Response{Result: result}, nil
}

func (s *server) Multiply(ctx context.Context, request *proto.Request) (*proto.Response, error){
	a, b := request.GetA(), request.GetB()
	result := a*b
	return &proto.Response{Result: result}, nil
}
