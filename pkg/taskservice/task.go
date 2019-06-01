package taskservice

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/bcmmbaga/microsvc/pkg/taskpb"
	"github.com/golang/protobuf/ptypes/empty"
	"go.mongodb.org/mongo-driver/mongo"
)

// taskService implements taskpb.TaskServiceServer
type taskService struct {
	client *mongo.Client
}

// NewTaskService create a new service for task management
func NewTaskService(ctx context.Context, hosts ...string) taskpb.TaskServiceServer {
	applyURI := dbHost(27017, hosts...)
	c, err := mongo.Connect(ctx, options.Client().ApplyURI(applyURI))
	if err != nil {
		return nil
	}

	return &taskService{client: c}
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
	return new(taskpb.CreateResponse), http.ErrHandlerTimeout
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
