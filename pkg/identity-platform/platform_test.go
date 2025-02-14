package platform

import (
	"io/ioutil"
	"secure-banking-uk-initializer/pkg/httprest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	mocks "secure-banking-uk-initializer/pkg/mocks/am"
)

func TestFindExistingAlphaClient(t *testing.T) {
	mockRestReaderWriter := &mocks.RestReaderWriter{}
	httprest.Client = mockRestReaderWriter
	buffer, _ := ioutil.ReadFile("client-check-test.json")
	mockRestReaderWriter.On("Get", mock.Anything, mock.Anything).
		Return(buffer)

	b := httprest.AlphaClientsExist("Doesnt existy")
	assert.False(t, b)
	mockRestReaderWriter.AssertCalled(t, "Get", mock.Anything, mock.Anything)

	b = httprest.AlphaClientsExist("ig-client")
	assert.True(t, b)
}
