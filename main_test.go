package main

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	s3Ctr := CreateS3Controller(Region, Endpoint, Profile)
	buckets := []string{
		Bucket,
	}

	s3Ctr.CreateBuckets(buckets)
	result := m.Run()
	os.Exit(result)
}

func TestUpload(t *testing.T) {
	s3Ctr := CreateS3Controller(Region, Endpoint, Profile)

	tests := []struct {
		bucket     string
		bucketPath string
		filePath   string
	}{
		{"go-localstack-github-actions-sample", "hello.txt", "sample/hello.txt"},
		{"go-localstack-github-actions-sample", "workflow.txt", "sample/workflow.txt"},
	}

	for _, tt := range tests {
		err := s3Ctr.UploadFile(tt.bucket, tt.bucketPath, tt.filePath)
		if err != nil {
			t.Fatalf("Failed to Upload: %s\n", err.Error())
		}
	}
}

func TestUploadError(t *testing.T) {
	s3Ctr := CreateS3Controller(Region, Endpoint, Profile)

	tests := []struct {
		bucket     string
		bucketPath string
		filePath   string
	}{
		{"xxxxxx", "hello.txt", "sample/hello.txt"},                              //non-existent bucket
		{"go-localstack-github-actions-sample", "workflow.txt", "sample/xxxxxx"}, //non-existent file
	}

	for _, tt := range tests {
		err := s3Ctr.UploadFile(tt.bucket, tt.bucketPath, tt.filePath)
		if err == nil {
			t.Fatal("No error was returned")
		}
	}
}
