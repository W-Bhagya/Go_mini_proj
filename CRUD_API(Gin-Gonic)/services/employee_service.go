package services

import (
	"context"
	"errors"

	"crud.com/api/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type EmpServices struct {
	usercollection *mongo.Collection
	ctx            context.Context
}

func NewEmpService(usercollection *mongo.Collection, ctx context.Context) EmpService {
	return &EmpServices {
		usercollection: usercollection,
		ctx:            ctx,
	}
}

func (u *EmpServices) CreateUser(emp *models.Employee) error {
	_, err := u.usercollection.InsertOne(u.ctx, emp)
	return err
}

func (u *EmpServices) GetUser(emp_name *string) (*models.Employee, error) {
	var emp *models.Employee
	query := bson.D{bson.E{Key: "emp_name", Value: emp_name}}
	err := u.usercollection.FindOne(u.ctx, query).Decode(&emp)
	return emp, err
}

func (u *EmpServices) GetAll() ([]*models.Employee, error) {
	var emps []*models.Employee
	cursor, err := u.usercollection.Find(u.ctx, bson.D{{}})
	if err != nil {
		return nil, err
	}
	for cursor.Next(u.ctx) {
		var emp models.Employee
		err := cursor.Decode(&emp)
		if err != nil {
			return nil, err
		}
		emps = append(emps, &emp)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	cursor.Close(u.ctx)

	if len(emps) == 0 {
		return nil, errors.New("Player Not Found")
	}
	return emps, nil
}

func (u *EmpServices) UpdateUser(emp *models.Employee) error {
	filter := bson.D{primitive.E{Key: "emp_name", Value: emp.Name}}
	update := bson.D{primitive.E{Key: "$set", Value: bson.D{primitive.E{Key: "emp_name", Value: emp.Name}, primitive.E{Key: "age", Value: emp.Age}, primitive.E{Key: "address", Value: emp.Address}}}}
	result, _ := u.usercollection.UpdateOne(u.ctx, filter, update)
	if result.MatchedCount != 1 {
		return errors.New("No Matched Player Found For Update")
	}
	return nil
}

func (u *EmpServices) DeleteUser(emp_name *string) error {
	filter := bson.D{primitive.E{Key: "emp_name", Value: emp_name}}
	result, _ := u.usercollection.DeleteOne(u.ctx, filter)
	if result.DeletedCount != 1 {
		return errors.New("No Matched Player Found For Delete")
	}
	return nil
}