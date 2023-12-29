package backoffice

import (
	"fmt"
	"pb"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func ExampleInvoice() {
	// TODO: Change to protobuf invoice
	inv := pb.Invoice{
		ID:       "2023-0123",
		Time:     &timestamppb.Timestamp{Seconds: 1671678300},
		Customer: "Wile E. Coyote",
		Items: []*pb.LineItem{
			{SKU: "hammer-20", Amount: 1, Price: 249},
			{SKU: "nail-9", Amount: 100, Price: 1},
			{SKU: "glue-5", Amount: 1, Price: 799},
		},
	}
	fmt.Printf("%v\n", &inv) // Make compiler happy
	// TODO: Encode to []byte using protobuf
	data, err := proto.Marshal(&inv)
	if err == nil {
		fmt.Println("size:", len(data))
	} else {
		fmt.Println("ERROR:", err)
	}

	// Output:
	// &{2023-0123 2023-01-07 13:45:00 +0000 UTC Wile E. Coyote [{hammer-20 1 249} {nail-9 100 1} {glue-5 1 799}]}
	// size: 0
}
