package oss

import (
	"context"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

var MinioClient *minio.Client

const VIDEO string = "video"
const IMAGE string = "image"

// Minio 连接oss数据库minio
func Minio(){
	endpoint := os.Getenv("MINIO_ENDPOINT")
	accessKeyID := os.Getenv("MINIO_ACCESS_KEY_ID")
	secretAccessKey := os.Getenv("MINIO_SECRET_ACCESS_KEY")
	connectTimeout,err := strconv.Atoi(os.Getenv("MINIO_CONNECT_TIMEOUT"))
	if err != nil {
		log.Panicln("MINIO_CONNECT_TIMEOUT"+"配置错误")
	}

	ctx,cancel := context.WithTimeout(context.Background(),time.Second*time.Duration(connectTimeout))
	defer cancel()
	successFlag := make(chan struct{})

	// Initialize minio client object.
	go func() {
		minioClient, err := minio.New(endpoint, &minio.Options{
			Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		})
		if err != nil {
			log.Panicln(err)
		}
		MinioClient = minioClient

		//存放视频内容的桶
		BucketInit(VIDEO,"us-east-1")
		//存放视频封面的桶
		BucketInit(IMAGE,"us-east-1")

		successFlag <- struct{}{}
	}()

	select {
	case <- ctx.Done():
		log.Panicln("连接 "+endpoint+" minio数据库超时")
	case <- successFlag:
		log.Println("连接 "+endpoint+" minio数据库成功")
	}

}

func BucketInit(bucketName string,location string){
	ctx := context.Background()
	err := MinioClient.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{Region: location})
	if err != nil {
		// Check to see if we already own this bucket (which happens if you run this twice)
		exists, errBucketExists := MinioClient.BucketExists(ctx, bucketName)
		if errBucketExists == nil && exists {
			log.Printf("We already own %s\n", bucketName)
		} else {
			log.Fatalln(err)
		}
	} else {
		log.Printf("Successfully created %s\n", bucketName)
	}
}
