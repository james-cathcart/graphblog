package graph

import (
	"errors"
	"github.com/99designs/gqlgen/client"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/james-cathcart/golog"
	"go.uber.org/mock/gomock"
	"graphblog/graph/model"
	"graphblog/internal/article"
	"graphblog/internal/user"
	"reflect"
	"testing"
)

func TestResolver_GetUsers(t *testing.T) {

	// build mocks
	ctrl := gomock.NewController(t)
	mockArticleSvc := article.NewMockService(ctrl)
	mockUserSvc := user.NewMockService(ctrl)
	resolver := NewResolver(mockArticleSvc, mockUserSvc)
	golog.SetLoggingLevel(golog.Disabled)
	golog.SetLoggingLevel(golog.Disabled)

	c := client.New(handler.NewDefaultServer(NewExecutableSchema(Config{Resolvers: resolver})))

	// create graphql response struct
	type response struct {
		Users []*model.User `json:"users"`
	}

	type errResponse struct {
		Message string   `json:"message"`
		Path    []string `json:"path"`
	}

	// test data struct
	type data struct {
		articleRecords []*model.Article
		records        []*model.User
		query          string
		errResp        *errResponse
		resp           response
		err            error
	}

	// test cases
	tests := []struct {
		name    string
		input   data
		mock    data
		expect  data
		prepare func(articleSvcMock *article.MockService, userSvcMock *user.MockService, input data, mock data)
	}{
		{
			name: `happy path`,
			input: data{
				query: `query { users{ id name } }`,
			},
			mock: data{
				articleRecords: []*model.Article{
					{
						ID:      `1`,
						Title:   `Title`,
						Content: `Content`,
						User: &model.User{
							ID:   `1`,
							Name: `One`,
						},
					},
				},
				records: []*model.User{
					{
						ID:   `1`,
						Name: `One`,
					},
					{
						ID:   `2`,
						Name: `Two`,
					},
					{
						ID:   `3`,
						Name: `three`,
					},
				},
			},
			expect: data{
				resp: response{
					Users: []*model.User{
						{
							ID:   `1`,
							Name: `One`,
						},
						{
							ID:   `2`,
							Name: `Two`,
						},
						{
							ID:   `3`,
							Name: `three`,
						},
					},
				},
			},
			prepare: func(articleSvcMock *article.MockService, userSvcMock *user.MockService, input data, mock data) {
				userSvcMock.EXPECT().GetAll().Return(mock.records, nil)
			},
		},
		{
			name: `user service returns error`,
			input: data{
				query: `query { users{ id name } }`,
			},
			mock: data{
				err: errors.New(`some error`),
			},
			expect: data{
				err: client.RawJsonError{
					RawMessage: []byte(`[{"message":"some error","path":["users"]}]`),
				},
				errResp: &errResponse{
					Message: `some error`,
					Path: []string{
						`users`,
					},
				},
			},
			prepare: func(articleSvcMock *article.MockService, userSvcMock *user.MockService, input data, mock data) {
				userSvcMock.EXPECT().GetAll().Return(nil, mock.err)
			},
		},
	}

	// execute test cases
	for _, test := range tests {

		t.Run(test.name, func(t *testing.T) {

			// prepare mocks
			if test.prepare != nil {
				test.prepare(mockArticleSvc, mockUserSvc, test.input, test.mock)
			}

			// execute logic
			var resp response
			err := c.Post(test.input.query, &resp)

			if test.expect.errResp != nil {

				// handle negative cases
				t.Log(`Error should be as expected`)
				{
					if !reflect.DeepEqual(test.expect.err, err) {
						t.Errorf("\tFAIL -> expected: (%T|%v), actual: (%T|%v)", test.expect.err, test.expect.err, err, err)
					} else {
						t.Log("\tSuccess")
					}
				}

				t.Log(`Users response should be nil`)
				{
					if resp.Users != nil {
						t.Errorf("\tFAIL -> expected: nil, actual: %v", resp.Users)
					} else {
						t.Log("\tSuccess")
					}
				}

			} else {

				// handle positive cases

				t.Log(`Error should be nil`)
				{
					if err != nil {
						t.Errorf("\tFAIL -> expected: nil, actual: (%T|%v)", err, err)
					} else {
						t.Log("\tSuccess")
					}
				}

				t.Log(`Response should be as expected`)
				{
					noErrors := true
					for i := range test.expect.resp.Users {
						if !reflect.DeepEqual(*test.expect.resp.Users[i], *resp.Users[i]) {
							noErrors = false
							t.Errorf("\tFAIL -> expected: %v, actual: %v", *test.expect.resp.Users[i], *resp.Users[i])
						}
					}
					if noErrors {
						t.Log("\tSuccess")
					}
				}
			}
		})
	}
}
