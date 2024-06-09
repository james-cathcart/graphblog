package user

import (
	"github.com/james-cathcart/golog"
	"go.uber.org/mock/gomock"
	"graphblog/graph/model"
	"reflect"
	"testing"
)

func TestGetAll(t *testing.T) {

	ctrl := gomock.NewController(t)
	mockUserDAO := NewMockDAO(ctrl)
	svc := NewDefaultService(mockUserDAO)

	golog.SetLoggingLevel(golog.Disabled)

	type data struct {
		err     error
		records []*model.User
	}

	tests := []struct {
		name    string
		input   data
		mock    data
		expect  data
		prepare func(userDaoMock *MockDAO, input data, mock data)
	}{
		{
			name: `happy path`,
			mock: data{
				records: []*model.User{
					{
						ID:   `1`,
						Name: `User 1`,
					},
					{
						ID:   `2`,
						Name: `User 2`,
					},
					{
						ID:   `3`,
						Name: `User 3`,
					},
				},
			},
			expect: data{
				records: []*model.User{
					{
						ID:   `1`,
						Name: `User 1`,
					},
					{
						ID:   `2`,
						Name: `User 2`,
					},
					{
						ID:   `3`,
						Name: `User 3`,
					},
				},
			},
			prepare: func(userDaoMock *MockDAO, input data, mock data) {
				userDaoMock.EXPECT().GetAll().Return(mock.records, nil)
			},
		},
	}

	for _, test := range tests {

		t.Run(test.name, func(t *testing.T) {

			// prepare mocks
			if test.prepare != nil {
				test.prepare(mockUserDAO, test.input, test.mock)
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

				t.Log(`Result should be as expected`)
				{
					noErrors := true
					for i := range test.expect.records {

						if !reflect.DeepEqual(*test.expect.records[i], *result[i]) {
							noErrors = false
							t.Errorf("\tFAIL -> expected: %v, actual: %v", *test.expect.records[i], *result[i])
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
	mockUserDao := NewMockDAO(ctrl)
	svc := NewDefaultService(mockUserDao)

	golog.SetLoggingLevel(golog.Disabled)

	type data struct {
		id     int64
		record model.User
		err    error
	}

	tests := []struct {
		name    string
		input   data
		mock    data
		expect  data
		prepare func(userDaoMock *MockDAO, input data, mock data)
	}{
		{
			name: `happy path`,
			input: data{
				record: model.User{
					Name: `User 1`,
				},
			},
			mock: data{
				id: 1,
			},
			expect: data{
				record: model.User{
					ID:   `1`,
					Name: `User 1`,
				},
			},
			prepare: func(userDaoMock *MockDAO, input data, mock data) {
				userDaoMock.EXPECT().Create(input.record).Return(mock.id, nil)
			},
		},
	}

	for _, test := range tests {

		t.Run(test.name, func(t *testing.T) {

			// prepare mocks
			if test.prepare != nil {
				test.prepare(mockUserDao, test.input, test.mock)
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
