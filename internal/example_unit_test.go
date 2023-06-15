package internal

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestTableTestExample(t *testing.T) {
	testCases := []struct {
		name           string
		input          int
		expectedResult int
	}{
		{name: "1 equals 1", input: 1, expectedResult: 1},
		{name: "2 equals 2", input: 2, expectedResult: 2},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			thingBeingTested := tc.input
			assert.Equal(t, tc.expectedResult, thingBeingTested)
		})
	}
}

func TestMockingExample(t *testing.T) {
	mockedDependency := &mockDependency{}
	mockedDependency.On("DependencyMethod", 1).Return(false, nil)
	thingBeingTested := exampleType{
		exampleDependency: mockedDependency,
	}

	thingBeingTested.DoSomething(1)
	mockedDependency.AssertCalled(t, "DependencyMethod", 1)
	mockedDependency.AssertNumberOfCalls(t, "DependencyMethod", 1)
}

type exampleType struct {
	exampleDependency
}

func (e *exampleType) DoSomething(value int) {
	_, _ = e.exampleDependency.DependencyMethod(value)
}

type exampleDependency interface {
	DependencyMethod(int) (bool, error)
}

type mockDependency struct {
	mock.Mock
}

func (m *mockDependency) DependencyMethod(number int) (bool, error) {
	args := m.Called(number)
	return args.Bool(0), args.Error(1)
}
