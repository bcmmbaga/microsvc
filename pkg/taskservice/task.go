package taskservice

import (
	"context"
	"fmt"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/bson"

	"github.com/golang/protobuf/ptypes"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/bcmmbaga/microsvc/pkg/taskpb"
	"github.com/golang/protobuf/ptypes/empty"
	"go.mongodb.org/mongo-driver/mongo"
)

// taskService implements taskpb.TaskServiceServer
type taskService struct {
	coll *mongo.Collection
}

// Task todo task
type Task struct {
	ID       primitive.ObjectID `bson:"_id, omitempty"`
	Title    string             `bson:"title"`
	Desc     string             `bson:"desc"`
	Reminder time.Time          `bson:"remainder"`
}

// NewTaskService create a new service for task management
func NewTaskService(ctx context.Context, db, coll string, hosts ...string) taskpb.TaskServiceServer {
	applyURI := dbHost(27017, hosts...)
	c, err := mongo.Connect(ctx, options.Client().ApplyURI(applyURI))
	if err != nil {
		return nil
	}

	return &taskService{coll: c.Database(db).Collection(coll)}
}

// dbHost joins multiple host names into single address for connection
// Used for connecting to multiple instance of databasesrunning in k8s
func dbHost(port int, hosts ...string) string {
	URI, joinHosts := "", ""
	if len(hosts) > 1 {
		joinHosts = strings.Join(hosts, ",")
	} else {
		URI = hosts[0]
	}
	URI = fmt.Sprintf("mongodb://%s:%d", joinHosts, port)
	return URI
}

// Create creates new todo task
func (s *taskService) Create(ctx context.Context, req *taskpb.CreateRequest) (*taskpb.CreateResponse, error) {

	remainder, err := ptypes.Timestamp(req.GetRemainder())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "reminder has unknown field"+err.Error())
	}

	task := Task{
		Title:    req.GetTitle(),
		Desc:     req.GetDesc(),
		Reminder: remainder,
	}

	res, err := s.coll.InsertOne(ctx, task)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "%s", err.Error())
	}

	objID, ok := res.InsertedID.(primitive.ObjectID)
	if !ok {
		return nil, status.Errorf(codes.Internal, "Failed to retrieve taskId from db: "+err.Error())
	}

	return &taskpb.CreateResponse{
		Id: objID.Hex(),
	}, nil
}

// Read todo task
func (s *taskService) Read(ctx context.Context, req *taskpb.ReadRequest) (*taskpb.ReadResponse, error) {
	objID, err := primitive.ObjectIDFromHex(req.GetId())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "%s"+err.Error())
	}

	res := &Task{}
	err = s.coll.FindOne(ctx, bson.D{{Key: "_id", Value: objID}}).Decode(&res)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "%s"+err.Error())
	}

	remainder, err := ptypes.TimestampProto(res.Reminder)
	if err != nil {
		return nil, status.Error(codes.Internal, "Failed to convert time.time to timestamp"+err.Error())
	}

	task := &taskpb.Task{
		Id:       res.ID.Hex(),
		Title:    res.Title,
		Desc:     res.Desc,
		Reminder: remainder,
	}
	return &taskpb.ReadResponse{Task: task}, nil
}

// Update todo task
func (s *taskService) Update(ctx context.Context, req *taskpb.UpdateRequest) (*taskpb.UpdateResponse, error) {
	task := req.GetTask()

	remainder, err := ptypes.Timestamp(task.Reminder)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "reminder has unknown field"+err.Error())
	}

	objID, err := primitive.ObjectIDFromHex(task.Id)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "%s"+err.Error())
	}

	updateTask := Task{
		ID:       objID,
		Title:    task.GetTitle(),
		Desc:     task.GetDesc(),
		Reminder: remainder,
	}

	res := new(Task)
	err = s.coll.FindOneAndReplace(ctx, bson.D{{Key: "_id", Value: updateTask.ID}}, bson.D{{Key: "$set", Value: updateTask}}).Decode(&res)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to update todo task "+err.Error())
	}

	return &taskpb.UpdateResponse{Id: res.ID.Hex()}, nil
}

// Delete todo task
func (s *taskService) Delete(ctx context.Context, req *taskpb.DeleteRequest) (*taskpb.DeleteResponse, error) {
	objID, err := primitive.ObjectIDFromHex(req.GetId())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "%s"+err.Error())
	}

	res := new(Task)
	err = s.coll.FindOneAndDelete(ctx, bson.D{{Key: "_id", Value: objID}}).Decode(&res)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "%s"+err.Error())
	}

	return &taskpb.DeleteResponse{Id: res.ID.Hex()}, nil
}

// Read all todo tasks
func (s *taskService) ReadAll(ctx context.Context, req *empty.Empty) (*taskpb.ReadAllResponse, error) {
	results := make([]*Task, 0)

	cur, err := s.coll.Find(ctx, nil)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "%s"+err.Error())
	}
	defer cur.Close(ctx)

	res := new(Task)
	for cur.Next(ctx) {
		err = cur.Decode(&res)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "%s"+err.Error())
		}
		results = append(results, res)
	}
	if err := cur.Err(); err != nil {
		return nil, status.Errorf(codes.Internal, "%s"+err.Error())
	}

	var tasks []*taskpb.Task
	for _, task := range results {
		remainder, _ := ptypes.TimestampProto(task.Reminder)
		tasks = append(tasks, &taskpb.Task{
			Id:       task.ID.Hex(),
			Title:    task.Title,
			Desc:     task.Desc,
			Reminder: remainder,
		})
	}

	return &taskpb.ReadAllResponse{Tasks: tasks}, nil
}
