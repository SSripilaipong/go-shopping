package testing

import (
	gfun "go-shopping/go/gfunc"
	"go-shopping/lambler"
)

func DynamodbEvent(records ...lambler.Json) lambler.Json {
	return lambler.Json{
		"Records": gfun.Map(records, func(record lambler.Json) any {
			return record
		}),
	}
}

func DynamodbEventRecordWithNewImage(keys lambler.Json, newImage lambler.Json) lambler.Json {
	return lambler.Json{
		"Records": lambler.Array{
			lambler.Json{
				"awsRegion": "ap-southeast-1",
				"dynamodb": lambler.Json{
					"ApproximateCreationDateTime": 1.68399118e+09,
					"Keys":                        keys,
					"NewImage":                    newImage,
					"SequenceNumber":              "38323900000000005876538041",
					"SizeBytes":                   123,
					"StreamViewType":              "NEW_IMAGE",
				},
				"eventID":        "d40f6e9cc8fb3274b28310ed81c5e66d",
				"eventName":      "INSERT",
				"eventSource":    "aws:dynamodb",
				"eventSourceARN": "arn:aws:dynamodb:ap-southeast-1:123456789012:table/inventory-item-repo/stream/2023-05-13T12:19:32.902",
				"eventVersion":   "1.1",
			},
		},
	}
}
