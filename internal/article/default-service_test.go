package article

import (
	"github.com/james-cathcart/golog"
	"go.uber.org/mock/gomock"
	"graphblog/graph/model"
	"reflect"
	"testing"
)

func TestGetAll(t *testing.T) {

	ctrl := gomock.NewController(t)
	mockArticleDAO := NewMockDAO(ctrl)
	svc := NewDefaultService(mockArticleDAO)

	golog.SetLoggingLevel(golog.Disabled)

	type data struct {
		err     error
		records []*model.Article
	}

	tests := []struct {
		name    string
		input   data
		mock    data
		expect  data
		prepare func(articleDaoMock *MockDAO, input data, mock data)
	}{
		{
			name: `happy path`,
			mock: data{
				records: []*model.Article{
					{
						ID:      `1`,
						Title:   `Title 1`,
						Content: `Content 1`,
						Status:  `published`,
						User: &model.User{
							ID:   `1`,
							Name: `UserName`,
						},
					},
					{
						ID:      `2`,
						Title:   `Title 2`,
						Content: `Content 2`,
						Status:  `published`,
						User: &model.User{
							ID:   `2`,
							Name: `UserName`,
						},
					},
					{
						ID:      `3`,
						Title:   `Title 3`,
						Content: `Content 3`,
						Status:  `published`,
						User: &model.User{
							ID:   `3`,
							Name: `UserName`,
						},
					},
				},
			},
			expect: data{
				records: []*model.Article{
					{
						ID:      `1`,
						Title:   `Title 1`,
						Content: `Content 1`,
						Status:  `published`,
						User: &model.User{
							ID:   `1`,
							Name: `UserName`,
						},
					},
					{
						ID:      `2`,
						Title:   `Title 2`,
						Content: `Content 2`,
						Status:  `published`,
						User: &model.User{
							ID:   `2`,
							Name: `UserName`,
						},
					},
					{
						ID:      `3`,
						Title:   `Title 3`,
						Content: `Content 3`,
						Status:  `published`,
						User: &model.User{
							ID:   `3`,
							Name: `UserName`,
						},
					},
				},
			},
			prepare: func(articleDaoMock *MockDAO, input data, mock data) {
				articleDaoMock.EXPECT().GetAll().Return(mock.records, nil)
			},
		},
	}

	for _, test := range tests {

		t.Run(test.name, func(t *testing.T) {

			// prepare mocks
			if test.prepare != nil {
				test.prepare(mockArticleDAO, test.input, test.mock)
			}

			// execute logic
			result, err := svc.GetAll()

			if test.expect.err != nil {

				// handle negative cases

				t.Log(`Error should be as expected`)
				{
					if !reflect.DeepEqual(test.expect.err, err) {
						t.Errorf("\tFAIL -> expected: (%T|%v), actual: (%T|%v)",
							test.expect.err, test.expect.err, err, err)
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

				t.Log("Result should be as expected")
				{
					noErrors := true

					for i := range test.expect.records {

						if !reflect.DeepEqual(*result[i], *test.expect.records[i]) {
							noErrors = false
							t.Errorf("\tFAIL -> expected: %v, actual: %v", *result[i], *test.expect.records[i])
						}

						if !reflect.DeepEqual(*result[i].User, *test.expect.records[i].User) {
							noErrors = false
							t.Errorf("\tFAIL -> expected: %v, actual: %v", *result[i].User, *test.expect.records[i].User)
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

func TestCreate(t *testing.T) {

	ctrl := gomock.NewController(t)
	mockArticleDAO := NewMockDAO(ctrl)
	svc := NewDefaultService(mockArticleDAO)

	golog.SetLoggingLevel(golog.Disabled)

	type data struct {
		err    error
		id     int64
		record model.Article
	}

	tests := []struct {
		name    string
		input   data
		mock    data
		expect  data
		prepare func(articleDaoMock *MockDAO, input data, mock data)
	}{
		{
			name: `happy path`,
			input: data{
				record: model.Article{
					Title:   `Title`,
					Content: `Content`,
					Status:  `published`,
					User: &model.User{
						ID:   `1`,
						Name: `UserName`,
					},
				},
			},
			mock: data{
				id: 1,
			},
			expect: data{
				record: model.Article{
					ID:      `1`,
					Title:   `Title`,
					Content: `Content`,
					Status:  `published`,
					User: &model.User{
						ID:   `1`,
						Name: `UserName`,
					},
				},
			},
			prepare: func(articleDaoMock *MockDAO, input data, mock data) {
				articleDaoMock.EXPECT().Create(input.record).Return(mock.id, nil)
			},
		},
	}

	for _, test := range tests {

		t.Run(test.name, func(t *testing.T) {

			// prepare mocks
			if test.prepare != nil {
				test.prepare(mockArticleDAO, test.input, test.mock)
			}

			// execute logic
			result, err := svc.Create(test.input.record)

			if test.expect.err != nil {

				// handle negative cases

				t.Log(`Error should be as expected`)
				{
					if !reflect.DeepEqual(test.expect.err, err) {
						t.Errorf("\tFAIL -> expected: (%T|%v), actual: (%T|%v)",
							test.expect.err, test.expect.err, err, err)
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

				t.Log(`Result should be as expected`)
				{
					if !reflect.DeepEqual(test.expect.record, result) {
						t.Errorf("\tFAIL -> expected: %v, actual: %v", test.expect.record, result)
					} else {
						t.Log("\tSuccess")
					}
				}
			}

		})
	}
}
