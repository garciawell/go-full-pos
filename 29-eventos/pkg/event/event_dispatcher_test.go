package event

import (
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"gotest.tools/v3/assert"
)

type TestEvent struct {
	Name    string
	Payload interface{}
}

func (e *TestEvent) GetName() string {
	return e.Name
}

func (e *TestEvent) GetPayload() interface{} {
	return e.Payload
}

func (e *TestEvent) GetDateType() time.Time {
	return time.Now()
}

type TestEventHandler struct {
	ID int
}

func (h *TestEventHandler) Handle(event EventInterface, wg *sync.WaitGroup) {

}

type EventDispatcherTestSuite struct {
	suite.Suite
	event           TestEvent
	event2          TestEvent
	handler         TestEventHandler
	handler2        TestEventHandler
	handler3        TestEventHandler
	EventDispatcher *EventDispatcher
}

func (suite *EventDispatcherTestSuite) SetupTest() {
	suite.EventDispatcher = NewEventDispatcher()
	suite.handler = TestEventHandler{
		ID: 1,
	}
	suite.handler2 = TestEventHandler{
		ID: 2,
	}
	suite.handler3 = TestEventHandler{
		ID: 3,
	}
	suite.event = TestEvent{Name: "test", Payload: "test"}
	suite.event2 = TestEvent{Name: "test2", Payload: "test2"}
}

func (suite *EventDispatcherTestSuite) TestEventDispatcher_Register() {
	err := suite.EventDispatcher.Register(suite.event.GetName(), &suite.handler)
	suite.Nil(err)
	suite.Equal(1, len(suite.EventDispatcher.handlers[suite.event.GetName()]))

	err = suite.EventDispatcher.Register(suite.event.GetName(), &suite.handler2)
	suite.Nil(err)
	suite.Equal(2, len(suite.EventDispatcher.handlers[suite.event.GetName()]))

	assert.Equal(suite.T(), &suite.handler, suite.EventDispatcher.handlers[suite.event.GetName()][0])
	assert.Equal(suite.T(), &suite.handler2, suite.EventDispatcher.handlers[suite.event.GetName()][1])
}

func (suite *EventDispatcherTestSuite) TestEventDispatcher_RegisterErr() {
	err := suite.EventDispatcher.Register(suite.event.GetName(), &suite.handler)
	suite.Nil(err)
	suite.Equal(1, len(suite.EventDispatcher.handlers[suite.event.GetName()]))

	err = suite.EventDispatcher.Register(suite.event.GetName(), &suite.handler)
	suite.Equal(ErrHandlerNotFound, err)
	suite.Equal(1, len(suite.EventDispatcher.handlers[suite.event.GetName()]))
}
func (suite *EventDispatcherTestSuite) TestEventDispatcher_Clear() {
	// Event 1
	err := suite.EventDispatcher.Register(suite.event.GetName(), &suite.handler)
	suite.Nil(err)
	suite.Equal(1, len(suite.EventDispatcher.handlers[suite.event.GetName()]))

	err = suite.EventDispatcher.Register(suite.event.GetName(), &suite.handler2)
	suite.Nil(err)
	suite.Equal(2, len(suite.EventDispatcher.handlers[suite.event.GetName()]))

	// Event 2
	err = suite.EventDispatcher.Register(suite.event2.GetName(), &suite.handler3)
	suite.Nil(err)
	suite.Equal(1, len(suite.EventDispatcher.handlers[suite.event2.GetName()]))

	suite.EventDispatcher.Clear()
	suite.Equal(0, len(suite.EventDispatcher.handlers))
}

func (suite *EventDispatcherTestSuite) TestEventDispatcher_Has() {
	err := suite.EventDispatcher.Register(suite.event.GetName(), &suite.handler)
	suite.Nil(err)
	suite.Equal(1, len(suite.EventDispatcher.handlers[suite.event.GetName()]))

	err = suite.EventDispatcher.Register(suite.event.GetName(), &suite.handler2)
	suite.Nil(err)
	suite.Equal(2, len(suite.EventDispatcher.handlers[suite.event.GetName()]))

	assert.Assert(suite.T(), suite.EventDispatcher.Has(suite.event.GetName(), &suite.handler))
	assert.Assert(suite.T(), suite.EventDispatcher.Has(suite.event.GetName(), &suite.handler2))
	assert.Assert(suite.T(), !suite.EventDispatcher.Has(suite.event.GetName(), &suite.handler3))
}

type MockHandler struct {
	mock.Mock
}

func (m *MockHandler) Handle(event EventInterface, wg *sync.WaitGroup) {
	m.Called(event)
	wg.Done()
}

func (suite *EventDispatcherTestSuite) TestEventDispatch_Dispatch() {
	eh := &MockHandler{}
	eh.On("Handle", &suite.event)
	suite.EventDispatcher.Register(suite.event.GetName(), eh)
	suite.EventDispatcher.Dispatch(&suite.event)
	eh.AssertExpectations(suite.T())
	eh.AssertNumberOfCalls(suite.T(), "Handle", 1)
}

func (suite *EventDispatcherTestSuite) TestEventDispatch_Remove() {
	// Event 1
	err := suite.EventDispatcher.Register(suite.event.GetName(), &suite.handler)
	suite.Nil(err)
	suite.Equal(1, len(suite.EventDispatcher.handlers[suite.event.GetName()]))

	err = suite.EventDispatcher.Register(suite.event.GetName(), &suite.handler2)
	suite.Nil(err)
	suite.Equal(2, len(suite.EventDispatcher.handlers[suite.event.GetName()]))

	// Event 2
	err = suite.EventDispatcher.Register(suite.event2.GetName(), &suite.handler3)
	suite.Nil(err)
	suite.Equal(1, len(suite.EventDispatcher.handlers[suite.event2.GetName()]))

	suite.EventDispatcher.Remove(suite.event.GetName(), &suite.handler)
	suite.Equal(1, len(suite.EventDispatcher.handlers[suite.event.GetName()]))
	assert.Equal(suite.T(), &suite.handler2, suite.EventDispatcher.handlers[suite.event.GetName()][0])

	suite.EventDispatcher.Remove(suite.event.GetName(), &suite.handler2)
	suite.Equal(0, len(suite.EventDispatcher.handlers[suite.event.GetName()]))

	suite.EventDispatcher.Remove(suite.event2.GetName(), &suite.handler3)
	suite.Equal(0, len(suite.EventDispatcher.handlers[suite.event2.GetName()]))
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(EventDispatcherTestSuite))
}
