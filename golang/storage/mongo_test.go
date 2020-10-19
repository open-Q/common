package storage

import (
	"context"
	"testing"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func Test_NewMongo(t *testing.T) {
	t.Run("create client error", func(t *testing.T) {
		_, err := NewMongo(context.Background(), "url://", "test-db")
		require.Error(t, err)
		require.EqualError(t, err, errors.Wrap(errors.New("error parsing uri: scheme must be \"mongodb\" or \"mongodb+srv\""), "could not create mongo client").Error())
	})
	t.Run("ping error", func(t *testing.T) {
		ctx, cancel := context.WithCancel(context.Background())
		go func() {
			cancel()
		}()
		_, err := NewMongo(ctx, "mongodb://127.0.0.1:27019", "test-db")
		require.Error(t, err)
		require.EqualError(t, err, errors.Wrap(errors.New("context canceled"), "could not ping connection").Error())
	})
	t.Run("all ok", func(t *testing.T) {
		db, err := NewMongo(context.Background(), "mongodb://127.0.0.1:27017", "test-db")
		require.NoError(t, err)
		require.NotNil(t, db)
	})
}

func TestMongoStorage_Collection(t *testing.T) {
	t.Run("create indexes error", func(t *testing.T) {
		db := newTestConnection(t)
		defer closeTestConnection(t, db)
		_, err := db.Collection(context.Background(), "test-data", mongo.IndexModel{
			Keys: nil,
		})
		require.Error(t, err)
		require.EqualError(t, err, errors.Wrap(errors.New("index model keys cannot be nil"), "could not create indexes").Error())
	})
	t.Run("all ok", func(t *testing.T) {
		db := newTestConnection(t)
		defer closeTestConnection(t, db)
		_, err := db.Collection(context.Background(), "test-data", mongo.IndexModel{
			Keys: bson.M{
				"some field": 1,
			},
		})
		require.NoError(t, err)
		collections, err := db.db.ListCollectionNames(context.Background(), bson.D{})
		require.NoError(t, err)
		require.Equal(t, 1, len(collections))
		require.Equal(t, "test-data", collections[0])
	})
}

func TestMongoStorage_Disconnect(t *testing.T) {
	t.Run("disconnection error", func(t *testing.T) {
		db := newTestConnection(t)
		closeTestConnection(t, db)
		err := db.Disconnect(context.Background())
		require.Error(t, err)
	})

	t.Run("all ok", func(t *testing.T) {
		db := newTestConnection(t)
		err := db.Disconnect(context.Background())
		require.NoError(t, err)
	})
}

func TestMongoStorage_NewTranscation(t *testing.T) {
	db := newTestConnection(t)
	defer closeTestConnection(t, db)
	tr := db.NewTranscation(context.Background())
	require.NotNil(t, tr)
	require.NotNil(t, tr.ctx)
	require.NotNil(t, tr.client)
	require.NotNil(t, tr.operations)
}

func TestMongoStorage_NewOperation(t *testing.T) {
	db := newTestConnection(t)
	defer closeTestConnection(t, db)
	tr := db.NewTranscation(context.Background())
	require.NotNil(t, tr)
	for i := 0; i < 3; i++ {
		tr.NewOperation(func(m *mongo.SessionContext) error {
			return nil
		})
	}
	require.Equal(t, 3, len(tr.operations))
}

func TestMongoTranscation_ExecuteAsync(t *testing.T) {
	t.Run("duplicate value error", func(t *testing.T) {
		ctx := context.Background()
		db := newTestConnection(t)
		defer closeTestConnection(t, db)
		coll, err := db.Collection(ctx, "test-data")
		require.NoError(t, err)

		id := primitive.NewObjectID()
		err = db.NewTranscation(ctx).
			NewOperation(func(m *mongo.SessionContext) error {
				_, err := coll.InsertOne(*m, bson.M{"_id": id})
				return err
			}).
			NewOperation(func(m *mongo.SessionContext) error {
				_, err := coll.InsertOne(*m, bson.M{"_id": id})
				return err
			}).
			ExecuteAsync()
		require.NoError(t, err)
		count, err := coll.CountDocuments(ctx, bson.D{})
		require.NoError(t, err)
		require.Equal(t, int64(0), count)
	})
	t.Run("all ok", func(t *testing.T) {
		ctx := context.Background()
		db := newTestConnection(t)
		defer closeTestConnection(t, db)
		coll, err := db.Collection(ctx, "test-data")
		require.NoError(t, err)

		id1 := primitive.NewObjectID()
		id2 := primitive.NewObjectID()
		err = db.NewTranscation(ctx).
			NewOperation(func(m *mongo.SessionContext) error {
				id, err := coll.InsertOne(*m, bson.M{"_id": id1})
				require.Equal(t, id1, id.InsertedID.(primitive.ObjectID))
				return err
			}).
			NewOperation(func(m *mongo.SessionContext) error {
				id, err := coll.InsertOne(*m, bson.M{"_id": id2})
				require.Equal(t, id2, id.InsertedID.(primitive.ObjectID))
				return err
			}).
			ExecuteAsync()
		require.NoError(t, err)
		count, err := coll.CountDocuments(ctx, bson.D{})
		require.NoError(t, err)
		require.Equal(t, int64(2), count)
	})
}

func TestMongoTranscation_Execute(t *testing.T) {
	t.Run("duplicate value error", func(t *testing.T) {
		ctx := context.Background()
		db := newTestConnection(t)
		defer closeTestConnection(t, db)
		coll, err := db.Collection(ctx, "test-data")
		require.NoError(t, err)

		id := primitive.NewObjectID()
		err = db.NewTranscation(ctx).
			NewOperation(func(m *mongo.SessionContext) error {
				_, err := coll.InsertOne(*m, bson.M{"_id": id})
				return err
			}).
			NewOperation(func(m *mongo.SessionContext) error {
				_, err := coll.InsertOne(*m, bson.M{"_id": id})
				return err
			}).
			Execute()
		require.NoError(t, err)
		count, err := coll.CountDocuments(ctx, bson.D{})
		require.NoError(t, err)
		require.Equal(t, int64(0), count)
	})
	t.Run("all ok", func(t *testing.T) {
		ctx := context.Background()
		db := newTestConnection(t)
		defer closeTestConnection(t, db)
		coll, err := db.Collection(ctx, "test-data")
		require.NoError(t, err)

		id1 := primitive.NewObjectID()
		id2 := primitive.NewObjectID()
		err = db.NewTranscation(ctx).
			NewOperation(func(m *mongo.SessionContext) error {
				id, err := coll.InsertOne(*m, bson.M{"_id": id1})
				require.Equal(t, id1, id.InsertedID.(primitive.ObjectID))
				return err
			}).
			NewOperation(func(m *mongo.SessionContext) error {
				id, err := coll.InsertOne(*m, bson.M{"_id": id2})
				require.Equal(t, id2, id.InsertedID.(primitive.ObjectID))
				return err
			}).
			Execute()
		require.NoError(t, err)
		count, err := coll.CountDocuments(ctx, bson.D{})
		require.NoError(t, err)
		require.Equal(t, int64(2), count)
	})
}

func newTestConnection(t *testing.T) *MongoStorage {
	db, err := NewMongo(context.Background(), "mongodb://127.0.0.1:27017", "test-db")
	require.NoError(t, err)
	return db
}

func closeTestConnection(t *testing.T, db *MongoStorage) {
	ctx := context.Background()
	collections, err := db.db.ListCollectionNames(ctx, bson.D{})
	require.NoError(t, err)
	for i := range collections {
		_, err := db.db.Collection(collections[i]).DeleteMany(ctx, bson.D{})
		require.NoError(t, err)
	}
	err = db.db.Client().Disconnect(ctx)
	require.NoError(t, err)
}
