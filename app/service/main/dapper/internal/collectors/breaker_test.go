package collectors

import (
	"angrymiao-go/app/service/main/dapper/internal/model"
	"fmt"
	"testing"
)

func TestServiceBreaker(t *testing.T) {
	breaker := NewServiceBreakerProcess(10)
	for i := 0; i < 20; i++ {
		err := breaker.Process(&model.Span{ServiceName: "test", OperationName: fmt.Sprintf("opt_%d", i)})
		if i < 10 {
			if err != nil {
				t.Error(err)
			}
		} else {
			if err == nil {
				t.Error("expect breaked")
			}
		}
	}
}
