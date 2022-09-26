package apiFunWithGo

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"gopkg.in/resty.v1"
	"testing"
)

type ZippopotamUsTestSuite struct {
	suite.Suite
	ApiClient *resty.Client
}

const zipcode = "90210"

func (suite *ZippopotamUsTestSuite) SetupTest() {
	suite.ApiClient = resty.New()
}

func (suite *ZippopotamUsTestSuite) Test_GetUs90210_StatusCodeShouldEqual200() {
	resp, _ := suite.ApiClient.R().Get(BASE_ENDPOINT + "/" + zipcode)

	assert.Equal(suite.T(), 200, resp.StatusCode())
}

func (suite *ZippopotamUsTestSuite) Test_GetUs90210_ContentTypeShouldEqualApplicationJson() {

	resp, _ := suite.ApiClient.R().Get(BASE_ENDPOINT + "/" + zipcode)

	assert.Equal(suite.T(), "application/json", resp.Header().Get("Content-Type"))
}

func (suite *ZippopotamUsTestSuite) Test_GetUs90210_CountryShouldEqualUnitedStates() {

	resp, _ := suite.ApiClient.R().Get(BASE_ENDPOINT + "/" + zipcode)

	myResponse := LocationResponse{}

	err := json.Unmarshal(resp.Body(), &myResponse)

	if err != nil {
		fmt.Println(err)
		return
	}

	assert.Equal(suite.T(), "United States", myResponse.Country)
}

func TestZippopotamUsSuite(t *testing.T) {
	suite.Run(t, new(ZippopotamUsTestSuite))
}