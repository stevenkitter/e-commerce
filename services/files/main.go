package main

import (
	"github.com/minio/minio-go"
	"log"
)

func main() {
	endpoint := "localhost:9000"
	accessKeyID := "AKIAIOSFOANN7EXAMPLE"
	secretAccessKey := "wJalrXUtkFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY"
	minioClient, err := minio.New(endpoint, accessKeyID, secretAccessKey, false)
	if err != nil {
		log.Fatalln(err)
	}
	bucketName := "mymusic"
	location := "china"
	err = minioClient.MakeBucket(bucketName, location)
	if err != nil {
		// 检查存储桶是否已经存在。
		exists, err := minioClient.BucketExists(bucketName)
		if err == nil && exists {
			log.Printf("We already own %s\n", bucketName)
		} else {
			log.Fatalln(err)
		}
	}
	log.Printf("Successfully created %s\n", bucketName)

	// 上传一个zip文件。
	objectName := "123.txt"
	filePath := "./123.txt"
	contentType := "text/plain"

	// 使用FPutObject上传一个zip文件。
	n, err := minioClient.FPutObject(bucketName, objectName, filePath, minio.PutObjectOptions{ContentType: contentType})
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf("Successfully uploaded %s of size %d\n", objectName, n)
}
