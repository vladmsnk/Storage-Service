package main

import (
	"bufio"
	"context"
	"git.miem.hse.ru/1206/proto/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"io"
	"log"
	"os"
	"time"
)

const filePath = "/Users/vladislavmoiseenkov/Downloads/video1248215149.mp4"

type TestMaterial struct {
	author    string
	fileTitle string
	fileType  string
	fileLink  string
}

func uploadMaterial(client pb.LibraryClient) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}

	testMaterial := TestMaterial{author: "vlad", fileTitle: "lecture", fileType: "mp4", fileLink: filePath}

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	stream, err := client.UploadMaterial(ctx)
	if err != nil {
		return err
	}

	infoReq := &pb.UploadMaterialRequest{
		Data: &pb.UploadMaterialRequest_Info{
			Info: &pb.MaterialInfo{
				Author: testMaterial.author,
				Title:  testMaterial.fileTitle,
				Type:   testMaterial.fileType,
				Link:   testMaterial.fileLink,
			},
		},
	}
	err = stream.Send(infoReq)
	if err != nil {
		return err
	}
	reader := bufio.NewReader(file)
	buffer := make([]byte, 1024)

	for {
		n, err := reader.Read(buffer)
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		req := &pb.UploadMaterialRequest{
			Data: &pb.UploadMaterialRequest_ChunkData{
				ChunkData: buffer[:n],
			},
		}
		err = stream.Send(req)
		if err != nil {
			return err
		}
	}
	res, err := stream.CloseAndRecv()
	if err != nil {
		return err
	}
	log.Printf("Video uploaded with id: %s, size: %d", res.GetId(), res.GetSize())
	return nil
}

func main() {
	//creds, err := credentials.NewServerTLSFromFile("server.crt", "")

	conn, err := grpc.Dial("localhost:9002", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("Could not establish connection!", err)
	}
	defer conn.Close()

	client := pb.NewLibraryClient(conn)
	err = uploadMaterial(client)
	log.Fatal(err)
}
