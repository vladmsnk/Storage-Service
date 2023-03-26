package service

import (
	"bytes"
	"git.miem.hse.ru/1206/material-library/internal/service/mapper"
	"git.miem.hse.ru/1206/proto/pb"
	"io"
)

func (s *LibraryServer) UploadMaterial(stream pb.Library_UploadMaterialServer) error {
	metaDataRequest, err := stream.Recv()
	if err != nil {
		return err
	}

	data := bytes.Buffer{}
	var fileSize int64
	for {
		dataReq, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		chunkData := dataReq.GetChunkData()
		fileSize += int64(len(chunkData))
		_, err = data.Write(chunkData)
		if err != nil {
			return err
		}
	}

	reader := bytes.NewReader(data.Bytes())
	uploadMaterialDTO := mapper.UploadMaterialRequest(metaDataRequest, fileSize, reader)
	uploadMaterialResponseDTO, err := s.uc.UploadMaterial(&uploadMaterialDTO)
	if err != nil {
		return err
	}
	return stream.SendAndClose(mapper.UploadMaterialResponse(uploadMaterialResponseDTO))
}
