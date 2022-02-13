package handler

import (
	"bytes"
	"io"
	"net/http"
	"sync"

	"image.upload/gen/pb"

	"github.com/google/uuid"
)

type ImageUploadHandler struct {
	sync.Mutex
	files map[string][]byte
}

func NewImageUploadHandler() *ImageUploadHandler {
	return &ImageUploadHandler{
		files: make(map[string][]byte),
	}
}

func (h *ImageUploadHandler) Upload(stream pb.ImageUploadService_UploadServer) error {
	req, err := stream.Recv()
	if err != nil {
		return err
	}
	meta := req.GetFileMeta()
	filename := meta.Filename

	u, err := uuid.NewRandom()
	if err != nil {
		return err
	}
	uuid := u.String()

	buf := &bytes.Buffer{}
	for {
		r, err := stream.Recv()

		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}
		chunk := r.GetData()
		_, err = buf.Write(chunk)
		if err != nil {
			return err
		}
	}
	data := buf.Bytes()
	mimeType := http.DetectContentType(data)
	h.files[uuid] = data

	err = stream.SendAndClose(&pb.ImageUploadResponse{
		Uuid:        uuid,
		Size:        int32(len(data)),
		Filename:    filename,
		ContentType: mimeType,
	})
	return err
}
