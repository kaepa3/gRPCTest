package handler

import (
	"context"
	"math/rand"
	"sync"
	"time"

	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/kaepa3/gRPCTest/ch1/api/gen/api"
	"google.golang.org/protobuf/types/known/timestamp"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

type BakerHandler struct {
	report *report
}

type report struct {
	sync.Mutex
	data map[api.Pancake_Menu]int
}

func NewBakerHandler() *BakerHandler {
	return &BakerHandler{
		report: &report{
			data: make(map[api.Pancake_Menu]int),
		},
	}
}

func (h *BakerHandler) Bake(ctx context.Context, req *api.BakeRequest) (*api.BakeResponce, error) {
	if req.Menu == api.Pancake_UNKNOWN || req.Menu > api.Pancake_SPICY_CURRY {
		return nil, status.Errorf(codes.InvalidArgument, "パンケーキを選んでください！")
	}
	now := time.Now()
	h.report.Lock()
	h.report.data[req.Menu] = h.report.data[req.Menu] + 1
	h.report.Unlock()

	return &api.BakeResponce{
		Pancake: &api.Pancake{
			Menu:           req.Menu,
			ChefName:       "gami",
			TechnicalScore: rand.Float32(),
			CreateTime: &timestamp.Timestamp{
				Seconds: now.Unix(),
				Nanos:   int32(now.Nanosecond()),
			},
		},
	}, nil
}
func (h *BakerHandler) Report(ctx context.Context, req *api.ReportRequest) (*api.ReportResponce, error) {
	counts := make([]*api.Report_BakeCount, 0)
	h.report.Lock()
	for k, v := range h.report.data {
		counts = append(counts, &api.Report_BakeCount{
			Menu:  k,
			Count: int32(v),
		})
	}
	h.report.Unlock()
	return &api.ReportResponce{
		Report: &api.Report{
			BakeCounts: counts,
		},
	}, nil
}
