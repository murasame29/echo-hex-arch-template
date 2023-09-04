package server

import (
	"github.com/murasame29/echo-hex-arch-template/pkg/core/application"
	"github.com/murasame29/echo-hex-arch-template/pkg/core/middleware"
)

func (es *EchoServer) LoginRoute() {
	login := application.NewLoginHTTPService(es.ctx, es.db, es.maker, *es.env)

	es.POST("/login", login.Login)
}

func (es *EchoServer) UserRoute() {
	user := application.NewUserHTTPService(es.ctx, es.db)

	es.GET("/users", user.ListUser)
	es.POST("/users", user.CreateUser)
	es.DELETE("/users", user.DeleteAllUser)

	es.GET("/users/:user_id", user.GetUser)
	es.PUT("/users/:user_id", user.UpdateUser)
	es.DELETE("/users/:user_id", user.DeleteUser)
}

func (es *EchoServer) TodoRoute() {
	todo := application.NewTodoHTTPService(es.ctx, es.db)

	es.GET("/todo", todo.ListTodo)
	es.POST("/todo", todo.CreateTodo)
	es.DELETE("/todo", todo.DeleteAllTodo)

	es.GET("/todo/:todo_id", todo.GetTodo)
	es.PUT("/todo/:todo_id", todo.UpdateTodo)
	es.DELETE("/todo/:todo_id", todo.DeleteTodo)
}

func (es *EchoServer) routes() {
	es.Use(middleware.Verify(es.maker))

	es.LoginRoute()
	es.TodoRoute()
	es.UserRoute()
}
