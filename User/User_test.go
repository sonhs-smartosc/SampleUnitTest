package User

import (
	"github.com/golang/mock/gomock"
	"mocktest/mocks"
	"testing"
)

func TestUser(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockUser := mocks.NewMockIUserInterface(mockCtrl)
	testUser := &User{IUser: mockUser}

	mockUser.EXPECT().AddUser(1, "new").Return(nil).Times(1)
	testUser.Use()
}
