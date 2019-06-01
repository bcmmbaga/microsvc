package taskservice

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"time"

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
func NewTaskService(ctx context.Context, hosts ...string) taskpb.TaskServiceServer {
	applyURI := dbHost(27017, hosts...)
	c, err := mongo.Connect(ctx, options.Client().ApplyURI(applyURI))
	if err != nil {
		return nil
	}

	return &taskService{coll: c.Database("todosdb").Collection("todos")}
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
	return new(taskpb.ReadResponse), http.ErrHandlerTimeout
}

// Update todo task
func (s *taskService) Update(ctx context.Context, req *taskpb.UpdateRequest) (*taskpb.UpdateResponse, error) {
	return new(taskpb.UpdateResponse), http.ErrHandlerTimeout
}

// Delete todo task
func (s *taskService) Delete(ctx context.Context, req *taskpb.DeleteRequest) (*taskpb.DeleteResponse, error) {
	return new(taskpb.DeleteResponse), http.ErrHandlerTimeout
}

// Read all todo tasks
func (s *taskService) ReadAll(ctx context.Context, req *empty.Empty) (*taskpb.ReadAllResponse, error) {
	return new(taskpb.ReadAllResponse), http.ErrHandlerTimeout
}
