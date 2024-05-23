package tg_a_c_i_su

import (
	"context"
	"fmt"
	"testing"
)

func TestClient(t *testing.T) {
	ch := make(chan ResponseData)
	ctx := context.Background()
	client := CreateACIClient(ch, ctx)
	client.Run("")
	m := <-ch
	fmt.Print(m.Count)
}
