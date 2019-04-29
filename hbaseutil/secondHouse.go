package hbaseutil

import(
	"github.com/tsuna/gohbase/hrpc"
	"context"
	"fmt"
)

type HbaseResult struct {
	Row string
	Val string
	Family string
	Qualifier string
}

func ScanHousePrice() (result []HbaseResult, err error) {
	scanReq, err := hrpc.NewScanStr(context.Background(), "second_house")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	scanRsp := Client.Scan(scanReq)
	scanResults := make([]HbaseResult, 0)
	for {
		result, err := scanRsp.Next()
		if err != nil {
			fmt.Println(err)
			break
		}
		cells := result.Cells
		for _, cell := range cells {
			itemResult := HbaseResult{ Row: string(cell.Row), Val: string(cell.Value), Family: string(cell.Family), Qualifier: string(cell.Qualifier) }
			scanResults = append(scanResults, itemResult)
		}
	}
	return scanResults, nil
}