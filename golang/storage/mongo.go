package storage

import (
	"context"
	"sync"

	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MongoStorage represents mongo storage model.
type MongoStorage struct {
	db *mongo.Database
}

// MongoCollection represents mongo collection model.
type MongoCollection struct {
	*mongo.Collection
}

// MongoTranscation represents mongo transaction model.
type MongoTranscation struct {
	ctx        context.Context
	client     *mongo.Client
	operations []func(*mongo.SessionContext) error
}

// NewMongo returns new mongo database connection.
func NewMongo(ctx context.Context, connString, database string) (*MongoStorage, error) {
	client, err := mongo.NewClient(options.Client().ApplyURI(connString))
	if err != nil {
		return nil, errors.Wrap(err, "could not create mongo client")
	}

	err = client.Connect(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "could not connect")
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, errors.Wrap(err, "could not ping connection")
	}

	return &MongoStorage{
		db: client.Database(database),
	}, nil
}

// Collection creates new or updates existing collection.
func (s *MongoStorage) Collection(ctx context.Context, name string, indexes ...mongo.IndexModel) (*MongoCollection, error) {
	coll := s.db.Collection(name)
	if len(indexes) != 0 {
		_, err := coll.Indexes().CreateMany(ctx, indexes)
		if err != nil {
			return nil, errors.Wrap(err, "could not create indexes")
		}
	}
	return &MongoCollection{coll}, nil
}

// Disconnect closes database connection.
func (s *MongoStorage) Disconnect(ctx context.Context) error {
	return s.db.Client().Disconnect(ctx)
}

// NewTranscation returns new mongo transcation instance.
func (s *MongoStorage) NewTranscation(ctx context.Context) *MongoTranscation {
	return &MongoTranscation{
		ctx:        ctx,
		client:     s.db.Client(),
		operations: make([]func(*mongo.SessionContext) error, 0),
	}
}

// NewOperation creates new operation in the transaction.
func (t *MongoTranscation) NewOperation(operation func(*mongo.SessionContext) error) *MongoTranscation {
	t.operations = append(t.operations, operation)
	return t
}

// Execute runs transaction and all related operations.
func (t *MongoTranscation) Execute() error {
	return t.execTransaction(false)
}

// ExecuteAsync runs transaction and all related operations.
// Each operation will be executed in async mode.
func (t *MongoTranscation) ExecuteAsync() error {
	return t.execTransaction(true)
}

func (t *MongoTranscation) execTransaction(async bool) error {
	return t.client.UseSession(t.ctx, func(sc mongo.SessionContext) error {
		if err := sc.StartTransaction(); err != nil {
			return err
		}

		errChan := make(chan error, 1)

		switch async {
		case true:
			var wg sync.WaitGroup
			wg.Add(len(t.operations))
			for i := range t.operations {
				go func(operation func(*mongo.SessionContext) error) {
					err := operation(&sc)
					if err != nil {
						errChan <- err
					}
					wg.Done()
				}(t.operations[i])
			}
			wg.Wait()
			close(errChan)
		default:
			var err error
			for i := range t.operations {
				if err = t.operations[i](&sc); err != nil {
					errChan <- err
					break
				}
			}
			if err == nil {
				close(errChan)
			}
		}

		if err, ok := <-errChan; err != nil && ok {
			return sc.AbortTransaction(t.ctx)
		}

		return sc.CommitTransaction(t.ctx)
	})
}
