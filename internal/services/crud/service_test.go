package crud_test

import (
	"testing"

	"github.com/stretchr/testify/suite"
	"go.uber.org/mock/gomock"

	"github.com/DoktorGhost/mock_test/internal/entity"
	"github.com/DoktorGhost/mock_test/internal/services/crud"
)

type CrudServiceTestSuite struct {
	suite.Suite

	repo *crud.MockRepository
}

func (suite *CrudServiceTestSuite) SetupTest() {
	ctrl := gomock.NewController(suite.T())
	defer ctrl.Finish()

	suite.repo = crud.NewMockRepository(ctrl)
}

func (suite *CrudServiceTestSuite) TestCreate() {
	suite.repo.EXPECT().Save(gomock.Any()).Return(nil)

	service := crud.NewService(suite.repo)
	err := service.Create(&entity.WeatherInfo{})
	suite.NoError(err)
}

func TestCrudServiceTestSuite(t *testing.T) {
	suite.Run(t, new(CrudServiceTestSuite))
}
