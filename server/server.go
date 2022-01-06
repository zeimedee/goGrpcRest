package server

import (
	"context"

	student "github.com/zeimedee/goGrpcRest/api/v1"
)

type StudentServer struct {
	student.UnimplementedStudentServiceServer
}

func (*StudentServer) Create(ctx context.Context, request *student.CreateStudentRequest) (*student.CreateStudentResponse, error) {
	fname := request.Student.Firstname
	lname := request.Student.Lastname
	age := request.Student.Age
	id := request.Student.Id
	response := &student.CreateStudentResponse{
		Student: &student.Student{
			Firstname: fname,
			Lastname:  lname,
			Age:       age,
			Id:        id,
		},
	}

	return response, nil
}

func (*StudentServer) Read(ctx context.Context, request *student.ReadStudentRequest) (*student.ReadStudentResponse, error) {
	// id := request.Student.Id

	response := &student.ReadStudentResponse{
		Student: &student.Student{
			Firstname: "abu",
			Lastname:  "yousef",
			Age:       20,
			Id:        "ay1999",
		},
	}
	return response, nil
}

func (*StudentServer) Update(ctx context.Context, request *student.UpdateStudentRequest) (*student.UpdateStudentResponse, error) {
	response := &student.UpdateStudentResponse{
		Student: &student.Student{
			Firstname: "steve",
			Lastname:  "arthur",
			Age:       25,
			Id:        "aa0099",
		},
	}

	return response, nil
}

func (*StudentServer) Delete(ctx context.Context, req *student.DeleteStudentRequest) (*student.DeleteStudentResponse, error) {
	res := &student.DeleteStudentResponse{
		Student: &student.Student{
			Firstname: "eddy",
			Lastname:  "dowoedoe",
			Age:       26,
			Id:        "ed9999",
		},
	}
	return res, nil
}
