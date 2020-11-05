package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"../proto"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

func main()  {
	conn, err := grpc.Dial("localhost:4040", grpc.WithInsecure())
	if err!=nil {
		panic(err)
	}

	client := proto.NewAddServiceClient(conn)


	gin_server := gin.Default()

	gin_server.GET("/add/:a/:b", func(ctx *gin.Context) {
		a, err := strconv.ParseUint(ctx.Param("a"),10,64)
		if err !=nil{
			ctx.JSON(http.StatusBadRequest,gin.H{"error" : "Invalid Parameter a"})
		}

		b, err := strconv.ParseUint(ctx.Param("b"),10,64)
		if err!=nil{
			ctx.JSON(http.StatusBadRequest, gin.H{"error" : "Invalid Parameter b"})
		}

		req := &proto.Request{A : int64(a), B: int64(b)}
		
		if response, err := client.Add(ctx, req); err == nil{
			ctx.JSON(http.StatusOK, gin.H{
				"result" : fmt.Sprint(response.Result),
			})
		} else{
			ctx.JSON(http.StatusInternalServerError, gin.H{"error" : err.Error()})
		}
	})

	gin_server.GET("/multi/:a/:b", func(ctx *gin.Context){
		a, err := strconv.ParseUint(ctx.Param("a"),10,64)
		if err!=nil {
			ctx.JSON(http.StatusBadRequest,gin.H{"error": "Invalid parameter a"})
		}

		b, err := strconv.ParseUint(ctx.Param("b"),10,64)
		if err!=nil{
			ctx.JSON(http.StatusBadRequest, gin.H{"error" : "Invalid parameter b"})
		}

		req := &proto.Request{A : int64(a), B : int64(b)}

		if response, err := client.Multiply(ctx, req); err == nil{
			ctx.JSON(http.StatusOK, gin.H{
				"result" : fmt.Sprint(response.Result),
			})
		}else{
			ctx.JSON(http.StatusInternalServerError,gin.H{"error": err.Error()})
		}
	})


	func insertdetails(cli proto.UserRequestClient,id int,email string, password string) {
		req := &proto.UserRequest{
			adduser := &proto.AddUser{
				Id : id,
				email : email,
				password : password,

			},
		}
		resp, err := cli.AddUser(context.Background(), req)

		if err != nil {
			log.Printf("insert response %+v\n", resp)
		}

	func readdetails(cli proto.UserRequestClient, email string) {
		req := &proto.ReadRequest {
			email : email,
		}
		resp, err := cli.Read(context.Background9), req)
		if err != nil {
			log.Printf("Error while calling the read function %v\n", err)
			return
		}
		log.Printf("read response is : %+v\n", resp.GetUser)
	}

	func updatedetails(cli proto.UserRequestClient, updateUser *proto.Update) {
		req := &proto.UpdateRequest {
			NewContact : updateUser,
		}
		resp, err := cli.Update(context.Background(), req)
		if err != nil {
			log.Printf("Error while calling the update function %v\n", err)
			return
		}
		log.Printf("update response %+v\n", resp.GetUpdateUser())
	}
	func deleteUser(cli proto.UserRequestClient, email string) {
		req := &proto.DeleteRequest{
			email : email,
		}
		resp, err := cli.Delete(context.Background(), req)
		if err != nil {
			log.Printf("Error while deleting %v\n", err)
			return
		}
		log.Printf("delete response %+v\n", resp)
	}


	}


	if err := gin_server.Run(":8080"); err!=nil {
		log.Fatal("Failed to run server: %v", err)
	}
}
