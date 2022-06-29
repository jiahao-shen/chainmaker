package tests

import (
	"fmt"
	"reflect"
	"testing"
	"time"
)

func Test_1(t *testing.T) {
	now := time.Now().Unix()
	fmt.Println(now)
	fmt.Println(reflect.TypeOf(now))
	fmt.Println(time.Unix(now, 0))
}
