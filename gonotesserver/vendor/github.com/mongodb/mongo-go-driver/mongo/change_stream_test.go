// Copyright (C) MongoDB, Inc. 2017-present.
//
// Licensed under the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at http://www.apache.org/licenses/LICENSE-2.0

package mongo

import (
	"context"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/mongodb/mongo-go-driver/core/command"
	"github.com/mongodb/mongo-go-driver/x/bsonx"
	"github.com/stretchr/testify/require"
)

func isServerError(err error) bool {
	_, ok := err.(command.Error)
	return ok
}

// TODO(GODRIVER-251): Replace manual check with functionality of improved testing framework.
func skipIfBelow36(t *testing.T) {
	serverVersion, err := getServerVersion(createTestDatabase(t, nil))
	require.NoError(t, err)

	if compareVersions(t, serverVersion, "3.6") < 0 {
		t.Skip()
	}
}

func getNextChange(changes Cursor) {
	for !changes.Next(context.Background()) {
	}
}

func TestChangeStream_firstStage(t *testing.T) {
	t.Parallel()

	if testing.Short() {
		t.Skip()
	}
	skipIfBelow36(t)

	if os.Getenv("TOPOLOGY") != "replica_set" {
		t.Skip()
	}

	coll := createTestCollection(t, nil, nil)

	// Ensure the database is created.
	_, err := coll.InsertOne(context.Background(), bsonx.Doc{{"x", bsonx.Int32(1)}})
	require.NoError(t, err)

	changes, err := coll.Watch(context.Background(), nil)
	require.NoError(t, err)

	require.NotEqual(t, len(changes.(*changeStream).pipeline), 0)

	elem := changes.(*changeStream).pipeline[0]

	doc := elem.Document()
	require.Equal(t, 1, len(doc))

	_, err = doc.LookupErr("$changeStream")
	require.NoError(t, err)
}

func TestChangeStream_noCustomStandaloneError(t *testing.T) {
	t.Parallel()

	if testing.Short() {
		t.Skip()
	}
	skipIfBelow36(t)

	topology := os.Getenv("TOPOLOGY")
	if topology == "replica_set" || topology == "sharded_cluster" {
		t.Skip()
	}

	coll := createTestCollection(t, nil, nil)

	// Ensure the database is created.
	_, err := coll.InsertOne(context.Background(), bsonx.Doc{{"x", bsonx.Int32(1)}})
	require.NoError(t, err)

	_, err = coll.Watch(context.Background(), nil)
	require.Error(t, err)
	if _, ok := err.(command.Error); !ok {
		t.Errorf("Should have returned command error, but got %T", err)
	}
}

func TestChangeStream_trackResumeToken(t *testing.T) {
	t.Parallel()

	if testing.Short() {
		t.Skip()
	}
	skipIfBelow36(t)

	if os.Getenv("TOPOLOGY") != "replica_set" {
		t.Skip()
	}

	coll := createTestCollection(t, nil, nil)

	// Ensure the database is created.
	_, err := coll.InsertOne(context.Background(), bsonx.Doc{{"y", bsonx.Int32(1)}})
	require.NoError(t, err)

	changes, err := coll.Watch(context.Background(), nil)
	require.NoError(t, err)

	for i := 1; i <= 4; i++ {
		_, err = coll.InsertOne(context.Background(), bsonx.Doc{{"x", bsonx.Int32(int32(i))}})
		require.NoError(t, err)
	}

	for i := 1; i <= 4; i++ {
		getNextChange(changes)
		var doc bsonx.Doc
		err := changes.Decode(&doc)
		require.NoError(t, err)

		id, err := doc.LookupErr("_id")
		require.NoError(t, err)

		if !cmp.Equal(id.Document(), changes.(*changeStream).resumeToken) {
			t.Errorf("Resume tokens do not match. got %v; want %v", id.Document(), changes.(*changeStream).resumeToken)
		}
	}
}

func TestChangeStream_errorMissingResponseToken(t *testing.T) {
	t.Parallel()

	if testing.Short() {
		t.Skip()
	}
	skipIfBelow36(t)

	if os.Getenv("TOPOLOGY") != "replica_set" {
		t.Skip()
	}

	coll := createTestCollection(t, nil, nil)

	// Ensure the database is created.
	_, err := coll.InsertOne(context.Background(), bsonx.Doc{{"y", bsonx.Int32(1)}})
	require.NoError(t, err)

	// Project out the response token
	changes, err := coll.Watch(context.Background(), []bsonx.Doc{
		{{"$project", bsonx.Document(bsonx.Doc{{"_id", bsonx.Int32(0)}})}},
	})
	require.NoError(t, err)

	_, err = coll.InsertOne(context.Background(), bsonx.Doc{{"x", bsonx.Int32(1)}})
	require.NoError(t, err)

	getNextChange(changes)
	require.Error(t, changes.Decode(&bsonx.Doc{}))
}

func TestChangeStream_resumableError(t *testing.T) {
	// Skipping this test due to flakiness - test sometimes has resume set, sometimes does not.  Not investigating
	// because this is being superseded by new changestream code
	t.Skip()
	t.Parallel()

	if testing.Short() {
		t.Skip()
	}
	skipIfBelow36(t)

	if os.Getenv("TOPOLOGY") != "replica_set" {
		t.Skip()
	}

	coll := createTestCollection(t, nil, nil)

	// Ensure the database is created.
	_, err := coll.InsertOne(context.Background(), bsonx.Doc{{"y", bsonx.Int32(1)}})
	require.NoError(t, err)

	changes, err := coll.Watch(context.Background(), nil)
	require.NoError(t, err)

	// Create a context that will expire before the operation can finish.
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Nanosecond)

	// "Use" the cancel function, which go vet complains if we throw away.
	func(context.CancelFunc) {}(cancel)

	require.False(t, changes.Next(ctx))

	err = changes.Err()
	require.Error(t, err)
	require.True(t, strings.Contains(err.Error(), "context deadline exceeded"))

	// If the ResumeAfter option is present, the the operation attempted to resume.
	hasResume := false

	for _, opt := range changes.(*changeStream).options {
		if opt.Key == "resumeAfter" {
			hasResume = true
			break
		}
	}

	require.True(t, hasResume)
}

// TODO: GODRIVER-247 Test that a change stream does not attempt to resume after a server error.

func TestChangeStream_resumeAfterKillCursors(t *testing.T) {
	t.Parallel()
	if testing.Short() {
		t.Skip()
	}
	skipIfBelow36(t)

	if os.Getenv("TOPOLOGY") != "replica_set" {
		t.Skip()
	}

	coll := createTestCollection(t, nil, nil)

	// Ensure the database is created.
	_, err := coll.InsertOne(context.Background(), bsonx.Doc{{"y", bsonx.Int32(1)}})
	require.NoError(t, err)

	changes, err := coll.Watch(context.Background(), nil)
	require.NoError(t, err)

	oldns := coll.namespace()
	killCursors := command.KillCursors{
		NS:  command.Namespace{DB: oldns.DB, Collection: oldns.Collection},
		IDs: []int64{changes.ID()},
	}

	ss, err := coll.client.topology.SelectServer(context.Background(), coll.readSelector)
	require.NoError(t, err)

	conn, err := ss.Connection(context.Background())
	require.NoError(t, err)
	defer conn.Close()

	_, err = killCursors.RoundTrip(context.Background(), ss.Description(), conn)
	require.NoError(t, err)

	// insert a document after blocking call to getNextChange below
	go func() {
		time.Sleep(time.Millisecond * 500)
		_, err = coll.InsertOne(context.Background(), bsonx.Doc{{"x", bsonx.Int32(1)}})
		require.NoError(t, err)
	}()

	getNextChange(changes)
	var doc bsonx.Doc
	require.NoError(t, changes.Decode(&doc))
}
