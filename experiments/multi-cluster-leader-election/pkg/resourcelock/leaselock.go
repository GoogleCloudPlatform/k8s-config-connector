package resourcelock

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"cloud.google.com/go/storage"
	leaderelectionv1 "github.com/600lyy/multi-cluster-leader-election/api/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"k8s.io/klog/v2"
)

// LeaseLock implements the resourcelock.Interface
// in the controller manager
type LeaseLock struct {
	*storage.Client
	LockConfig ResourceLockConfig
	ProjectId  string
	BucketName string
	LeaseFile  string
}

// Get returns the lease object from the storage bucket
func (ll *LeaseLock) Get(ctx context.Context) (*LeaderElectionRecord, []byte, error) {
	var rc *storage.Reader
	var gcsErr error
	var len int
	var recordByte = make([]byte, 200)
	var record = LeaderElectionRecord{}

	if ll.Client == nil {
		return nil, nil, errors.New("storage client is not empty, initiate it first")
	}

	if rc, gcsErr = ll.Client.Bucket(ll.BucketName).Object(ll.LeaseFile).NewReader(ctx); gcsErr != nil {
		return nil, nil, ll.convertGcsErrToLeaseErr(gcsErr)
	}
	defer rc.Close()

	if len, gcsErr = rc.Read(recordByte); gcsErr != nil {
		return nil, nil, ll.convertGcsErrToLeaseErr(gcsErr)
	}

	if gcsErr = json.Unmarshal(recordByte[:len], &record); gcsErr != nil {
		return nil, nil, ll.convertGcsErrToLeaseErr(gcsErr)
	}
	klog.Infof("The candidate %v reads %d bytes from [%v/%v] in Google storage", ll.LockConfig.Identity, len, ll.BucketName, ll.LeaseFile)
	return &record, recordByte, nil
}

// Create attempts to create a lease
func (ll *LeaseLock) Create(ctx context.Context, ler LeaderElectionRecord) error {
	var bkt *storage.BucketHandle
	var attr *storage.BucketAttrs
	var err error
	var len int

	if ll.Client == nil {
		return errors.New("storage client is not empty, initiate it first")
	}

	bkt = ll.Client.Bucket(ll.BucketName)
	attr, err = bkt.Attrs(ctx)
	if err != nil {
		return err
	}
	klog.Infof("Succesfully connected to the bucket [%v] in Google storage", attr.Name)

	wc := bkt.Object(ll.LeaseFile).NewWriter(ctx)
	writeByte, err := json.Marshal(ler)
	if err != nil {
		return err
	}
	if len, err = wc.Write(writeByte); err != nil {
		return err
	}
	if err = wc.Close(); err != nil {
		return err
	}
	klog.Infof("Create lease: Succesfully write [%d] bytes to the lease [%v/%v] in Google storage", len, attr.Name, wc.ObjectAttrs.Name)
	return nil
}

// Update will update an exising lease
// For GCS, leader needs to update the time stamps in the lease file
func (ll *LeaseLock) Update(ctx context.Context, ler LeaderElectionRecord) error {
	var err error
	var len int

	if ll.Client == nil {
		return errors.New("storage client is not empty, initiate it first")
	}
	wc := ll.Client.Bucket(ll.BucketName).Object(ll.LeaseFile).NewWriter(ctx)
	writeByte, err := json.Marshal(ler)
	if err != nil {
		return err
	}
	if len, err = wc.Write(writeByte); err != nil {
		return err
	}
	if err = wc.Close(); err != nil {
		return err
	}
	klog.Infof("Update lease: Succesfully write [%d] bytes to the lease [%v/%v] in Google storage", len, ll.BucketName, wc.ObjectAttrs.Name)
	return nil
}

// Describe is used to convert details on current resource lock
// into a string
func (ll *LeaseLock) Describe() string {
	return fmt.Sprintf("%v/%v", ll.BucketName, ll.LeaseFile)
}

// Identity will return the locks Identity
func (ll *LeaseLock) Identity() string {
	return ll.LockConfig.Identity
}

func (ll *LeaseLock) convertGcsErrToLeaseErr(gcsErr error) *apierrors.StatusError {
	var statusReason metav1.StatusReason
	var statusErr *apierrors.StatusError

	if errors.Is(gcsErr, storage.ErrObjectNotExist) {
		statusReason = metav1.StatusReasonNotFound
	} else {
		statusReason = metav1.StatusReasonInternalError
	}
	statusErr = &apierrors.StatusError{
		ErrStatus: metav1.Status{
			Status:  metav1.StatusFailure,
			Code:    http.StatusNotFound,
			Reason:  statusReason,
			Message: gcsErr.Error(),
		},
	}
	return statusErr
}

func LeaderElectionRecordToLeaseSpec(ler *LeaderElectionRecord) leaderelectionv1.LeaseSpec {
	leaseDurationSeconds := int32(ler.LeaseDurationSeconds)
	return leaderelectionv1.LeaseSpec{
		HolderIdentity:       ler.HolderIdentity,
		LeaseDurationSeconds: leaseDurationSeconds,
	}
}

func LeaderElectionRecordToLeaseStatus(ler *LeaderElectionRecord) leaderelectionv1.LeaseStatus {
	return leaderelectionv1.LeaseStatus{
		ObservedHolderIdentity: ler.HolderIdentity,
		ObservedAcquireTime:    metav1.MicroTime(ler.AcquireTime),
		ObservedRenewTime:      metav1.MicroTime(ler.RenewTime),
		LeaseTransitions:       int32(ler.LeaderTransitions),
	}
}
