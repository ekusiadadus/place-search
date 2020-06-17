package interactor

import (
	"os"
	"testing"

	"github.com/golang/mock/gomock"
	mock_gateway "github.com/yagi-eng/place-search/mock/gateway"
	mock_ipresenter "github.com/yagi-eng/place-search/mock/presenter"
	"github.com/yagi-eng/place-search/usecases/dto/googlemapdto"
	"github.com/yagi-eng/place-search/usecases/dto/searchdto"
)

// TestHundle1
// 正常系のみ
func TestHundle1(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	q := "東京"
	query := os.Getenv("QUERY")
	gmo := googlemapdto.Output{}
	gmos := []googlemapdto.Output{gmo}

	mockGoogleMapGW := mock_gateway.NewMockIGoogleMapGateway(ctrl)
	mockGoogleMapGW.EXPECT().
		GetPlaceDetailsAndPhotoURLsWithQuery(q + query).
		Return(gmos)

	expected := searchdto.Output{
		ReplyToken:       replyToken,
		Q:                q,
		GoogleMapOutputs: gmos,
	}
	mockLinePrst := mock_ipresenter.NewMockILinePresenter(ctrl)
	mockLinePrst.EXPECT().
		Search(expected).
		Return()

	interactor := NewSearchInteractor(mockGoogleMapGW, mockLinePrst)

	in := searchdto.Input{
		ReplyToken: replyToken,
		Q:          q,
	}
	interactor.Hundle(in)
}
