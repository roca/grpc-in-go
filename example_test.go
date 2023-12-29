package backoffice

import (
	"fmt"
	"pb"
	"testing"
	"time"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func TestInvoice(t *testing.T) {
	// TODO: Change to protobuf invoice
	time := time.Date(2023, time.January, 7, 13, 45, 0, 0, time.UTC)
	inv := pb.Invoice{
		ID:       "2023-0123",
		Time:     timestamppb.New(time),
		Customer: "Wile E. Coyote",
		Items: []*pb.LineItem{
			{SKU: "hammer-20", Amount: 1, Price: 249},
			{SKU: "nail-9", Amount: 100, Price: 1},
			{SKU: "glue-5", Amount: 1, Price: 799},
		},
	}
	// fmt.Printf("%v\n", &inv) // Make compiler happy
	// TODO: Encode to []byte using protobuf
	data, err := proto.Marshal(&inv)
	if err == nil {
		fmt.Println("size:", len(data))
	} else {
		fmt.Println("ERROR:", err)
	}

	if inv.ID != "2023-0123" {
		t.Errorf("Expected ID to be 2023-0123, got %q", inv.ID)
	}

	if len(data) != 82 {
		t.Errorf("Expected size to be 82, got %d", len(data))
	}
}
