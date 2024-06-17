// Copyright 2024 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package bucketlease

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"cloud.google.com/go/storage"
	"github.com/go-logr/logr"
	"google.golang.org/api/googleapi"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/tools/record"

	"github.com/GoogleCloudPlatform/k8s-config-connector/experiments/multicluster/pkg/resourcelock"
)

func New(ctx context.Context, c Config) *BucketLease {
	bl := BucketLease{
		identity:      c.Identity,
		bucketName:    c.BucketName,
		leaseName:     c.LeaseName,
		log:           c.Log,
		eventRecorder: c.EventRecorder,
	}
	client, err := storage.NewClient(ctx)
	if err != nil {
		return nil
	}
	bl.client = client
	return &bl
}

type Config struct {
	Identity      string
	BucketName    string
	LeaseName     string
	Log           logr.Logger
	EventRecorder record.EventRecorder
}

type BucketLease struct {
	client        *storage.Client
	identity      string
	bucketName    string
	leaseName     string
	log           logr.Logger
	eventRecorder record.EventRecorder
}

var _ resourcelock.Interface = &BucketLease{}

func (l *BucketLease) Get(ctx context.Context) (*resourcelock.LeaderElectionRecord, []byte, error) {
	r, err := l.client.Bucket(l.bucketName).Object(l.leaseName).NewReader(ctx)
	if err != nil {
		l.log.Error(err, "error creating storage reader") // TODO: do not print this is error is not found
		return nil, nil, toAPIError(err)
	}
	defer r.Close()
	recordBytes, err := io.ReadAll(r)
	if err != nil {
		l.log.Error(err, "error reading record")
		return nil, nil, err
	}
	var record resourcelock.LeaderElectionRecord
	if err := json.Unmarshal(recordBytes, &record); err != nil {
		l.log.Error(err, "error unmarshaling record")
		return nil, nil, err
	}
	return &record, recordBytes, nil
}

func (l *BucketLease) Create(ctx context.Context, record resourcelock.LeaderElectionRecord) error {
	b := l.client.Bucket(l.bucketName)
	attrs, err := b.Attrs(ctx)
	if err != nil {
		return fmt.Errorf("error getting bucket attributes: %v", err)
	}
	w := b.If(storage.BucketConditions{MetagenerationMatch: attrs.MetaGeneration}).Object(l.leaseName).NewWriter(ctx)

	recordByte, err := json.Marshal(record)
	if err != nil {
		return fmt.Errorf("error marshaling record: %v", err)
	}
	if _, err := w.Write(recordByte); err != nil {
		return fmt.Errorf("error writing record: %v", err)
	}

	if err := w.Close(); err != nil {
		switch ee := err.(type) {
		case *googleapi.Error:
			if ee.Code == http.StatusPreconditionFailed {
				// The condition presented in the If failed.
				// TODO: add metrics
				return fmt.Errorf("error writing record because the precondition failed: %v", err)
			}
		default:
			return fmt.Errorf("error writing record: %v", err)
		}
	}
	return nil
}

func (l *BucketLease) Update(ctx context.Context, record resourcelock.LeaderElectionRecord) error {
	b := l.client.Bucket(l.bucketName)
	objAttrs, err := b.Object(l.leaseName).Attrs(ctx)
	if err != nil {
		return fmt.Errorf("error getting object attributes: %v", err)
	}
	w := b.Object(l.leaseName).If(storage.Conditions{GenerationMatch: objAttrs.Generation}).NewWriter(ctx)

	recordByte, err := json.Marshal(record)
	if err != nil {
		return fmt.Errorf("error marshaling record: %v", err)
	}
	if _, err := w.Write(recordByte); err != nil {
		return fmt.Errorf("error writing record: %v", err)
	}

	if err := w.Close(); err != nil {
		switch ee := err.(type) {
		case *googleapi.Error:
			if ee.Code == http.StatusPreconditionFailed {
				// The condition presented in the If failed.
				// TODO: add metrics
				return fmt.Errorf("error writing record due to precondition failed: %v", err)
			}
		default:
			return fmt.Errorf("error writing record: %v", err)
		}
	}
	return nil
}

func (l *BucketLease) RecordEvent(s string) {
	if l.eventRecorder == nil {
		return
	}
	// TODO: record event
}

func (l *BucketLease) Identity() string {
	return l.identity
}

func (l *BucketLease) Describe() string {
	return fmt.Sprintf("%v/%v", l.bucketName, l.leaseName)
}

func toAPIError(err error) error {
	if err == nil {
		return nil
	}
	if errors.Is(err, storage.ErrObjectNotExist) {
		return &apierrors.StatusError{
			ErrStatus: metav1.Status{
				Status:  metav1.StatusFailure,
				Code:    http.StatusNotFound,
				Reason:  metav1.StatusReasonNotFound,
				Message: err.Error(),
			},
		}
	}
	// TODO: handle more error types

	return &apierrors.StatusError{
		ErrStatus: metav1.Status{
			Status:  metav1.StatusFailure,
			Code:    http.StatusInternalServerError,
			Reason:  metav1.StatusReasonInternalError,
			Message: err.Error(),
		},
	}
}
