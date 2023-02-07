package accessdata

import (
	"fmt"
)

func _ExampleOperators() {
	op := Operators[DurationOperator]
	fmt.Printf("test: Operator() -> %v\n", op)

	op = Operators[StartTimeOperator]
	fmt.Printf("test: Operator() -> %v\n", op)

	//Output:
	//fail
}

func Example_IsDirectOperator() {
	op := Operator{Name: "test", Value: "   %"}
	fmt.Printf("test: IsDirectOperator() -> %v [value:%v]\n", IsDirectOperator(op), op.Value)

	op = Operator{Name: "test", Value: "%"}
	fmt.Printf("test: IsDirectOperator() -> %v [value:%v]\n", IsDirectOperator(op), op.Value)

	//Output:
	//test: IsDirectOperator() -> true [value:   %]
	//test: IsDirectOperator() -> false [value:%]
}

func Example_IsRequestOperator() {
	op := Operator{}
	ok := IsRequestOperator(op)
	fmt.Printf("test: IsRequestOperator(<empty>) -> %v\n", ok)

	op = Operator{Name: " ", Value: " "}
	ok = IsRequestOperator(op)
	fmt.Printf("test: IsRequestOperator(<empty>) -> %v\n", ok)

	op = Operator{Name: "", Value: "REQ "}
	ok = IsRequestOperator(op)
	fmt.Printf("test: IsRequestOperator(%v) -> %v\n", op, ok)

	op = Operator{Name: "", Value: "%REQ(header"}
	ok = IsRequestOperator(op)
	fmt.Printf("test: IsRequestOperator(%v) -> %v\n", op, ok)

	op = Operator{Name: "", Value: "%REQ(header)"}
	ok = IsRequestOperator(op)
	fmt.Printf("test: IsRequestOperator(%v) -> %v\n", op, ok)

	op = Operator{Name: "", Value: "%REQ()"}
	ok = IsRequestOperator(op)
	fmt.Printf("test: IsRequestOperator(%v) -> %v\n", op, ok)

	op = Operator{Name: "", Value: "%REQ(1)%"}
	ok = IsRequestOperator(op)
	fmt.Printf("test: IsRequestOperator(%v) -> %v\n", op, ok)

	op = Operator{Name: "", Value: "%REQ(header-name)%"}
	ok = IsRequestOperator(op)
	fmt.Printf("test: IsRequestOperator(%v) -> %v\n", op, ok)

	//Output:
	//test: IsRequestOperator(<empty>) -> false
	//test: IsRequestOperator(<empty>) -> false
	//test: IsRequestOperator({ REQ }) -> false
	//test: IsRequestOperator({ %REQ(header}) -> false
	//test: IsRequestOperator({ %REQ(header)}) -> false
	//test: IsRequestOperator({ %REQ()}) -> false
	//test: IsRequestOperator({ %REQ(1)%}) -> true
	//test: IsRequestOperator({ %REQ(header-name)%}) -> true

}

func Example_RequestOperatorHeaderName() {
	op := Operator{}
	name := RequestOperatorHeaderName(op)
	fmt.Printf("test: RequestOperatorHeaderName() -> %v [op:%v]\n", name, op.Value)

	op = Operator{Name: "", Value: "%REQ("}
	name = RequestOperatorHeaderName(op)
	fmt.Printf("test: RequestOperatorHeaderName() -> %v [op:%v]\n", name, op.Value)

	op = Operator{Name: "", Value: "%REQ()"}
	name = RequestOperatorHeaderName(op)
	fmt.Printf("test: RequestOperatorHeaderName() -> %v [op:%v]\n", name, op.Value)

	op = Operator{Name: "", Value: "%REQ()%"}
	name = RequestOperatorHeaderName(op)
	fmt.Printf("test: RequestOperatorHeaderName() -> %v [op:%v]\n", name, op.Value)

	op = Operator{Name: "", Value: "%REQ(1)%"}
	name = RequestOperatorHeaderName(op)
	fmt.Printf("test: RequestOperatorHeaderName() -> %v [op:%v]\n", name, op.Value)

	op = Operator{Name: "", Value: "%REQ(name)%"}
	name = RequestOperatorHeaderName(op)
	fmt.Printf("test: RequestOperatorHeaderName() -> %v [op:%v]\n", name, op.Value)

	//Output:
	//test: RequestOperatorHeaderName() ->  [op:]
	//test: RequestOperatorHeaderName() ->  [op:%REQ(]
	//test: RequestOperatorHeaderName() ->  [op:%REQ()]
	//test: RequestOperatorHeaderName() ->  [op:%REQ()%]
	//test: RequestOperatorHeaderName() -> 1 [op:%REQ(1)%]
	//test: RequestOperatorHeaderName() -> name [op:%REQ(name)%]

}

func Example_IsStringValue() {
	op := Operator{Name: "test", Value: "   %"}
	fmt.Printf("test: IsStringValue() -> %v [value:%v]\n", IsStringValue(op), op.Value)

	op = Operator{Name: "test", Value: "%"}
	fmt.Printf("test: IsStringValue() -> %v [value:%v]\n", IsStringValue(op), op.Value)

	op = Operator{Name: "test", Value: DurationOperator}
	fmt.Printf("test: IsStringValue() -> %v [value:%v]\n", IsStringValue(op), op.Value)

	op = Operator{Name: "test", Value: TimeoutDurationOperator}
	fmt.Printf("test: IsStringValue() -> %v [value:%v]\n", IsStringValue(op), op.Value)

	op = Operator{Name: "test", Value: RateLimitOperator}
	fmt.Printf("test: IsStringValue() -> %v [value:%v]\n", IsStringValue(op), op.Value)

	op = Operator{Name: "test", Value: RateBurstOperator}
	fmt.Printf("test: IsStringValue() -> %v [value:%v]\n", IsStringValue(op), op.Value)

	op = Operator{Name: "test", Value: RetryOperator}
	fmt.Printf("test: IsStringValue() -> %v [value:%v]\n", IsStringValue(op), op.Value)

	op = Operator{Name: "test", Value: RetryRateLimitOperator}
	fmt.Printf("test: IsStringValue() -> %v [value:%v]\n", IsStringValue(op), op.Value)

	op = Operator{Name: "test", Value: RetryRateBurstOperator}
	fmt.Printf("test: IsStringValue() -> %v [value:%v]\n", IsStringValue(op), op.Value)

	op = Operator{Name: "test", Value: FailoverOperator}
	fmt.Printf("test: IsStringValue() -> %v [value:%v]\n", IsStringValue(op), op.Value)

	//Output:
	//test: IsStringValue() -> true [value:   %]
	//test: IsStringValue() -> true [value:%]
	//test: IsStringValue() -> false [value:%DURATION%]
	//test: IsStringValue() -> false [value:%TIMEOUT_DURATION%]
	//test: IsStringValue() -> false [value:%RATE_LIMIT%]
	//test: IsStringValue() -> false [value:%RATE_BURST%]
	//test: IsStringValue() -> false [value:%RETRY]
	//test: IsStringValue() -> false [value:%RETRY_RATE_LIMIT%]
	//test: IsStringValue() -> false [value:%RETRY_RATE_BURST%]
	//test: IsStringValue() -> false [value:%FAILOVER%]
}

func translateOperator(op Operator) Operator {
	newOp := Operator{Name: op.Name, Value: op.Value}
	if newOp.Name == "" {
		newOp.Name = "<empty>"
	}
	if newOp.Value == "" {
		newOp.Value = "<empty>"
	}
	//if newE.Name == "" {
	//	newE.Name = "<empty>"
	//}
	return newOp
}

/*
func Example_createHeaderOperator() {
	op := createHeaderOperator(Operator{Name: "", Value: ""})
	fmt.Printf("test: Operator(\"\") -> [%v]\n", translateOperator(op))

	op = createHeaderOperator(Operator{Value: "test", Name: ""})
	fmt.Printf("test: Operator(\"test\") -> [%v]\n", translateOperator(op))

	op = createHeaderOperator(Operator{Value: "%REQ(", Name: ""})
	fmt.Printf("test: Operator(\"REQ(\") -> [%v]\n", translateOperator(op))

	op = createHeaderOperator(Operator{Value: "%REQ(t", Name: ""})
	fmt.Printf("test: Operator(\"REQ(t\") -> [%v]\n", translateOperator(op))

	op = createHeaderOperator(Operator{Value: "%REQ()", Name: ""})
	fmt.Printf("test: Operator(\"REQ()\") -> [%v]\n", translateOperator(op))

	op = createHeaderOperator(Operator{Value: "%REQ(member)", Name: ""})
	fmt.Printf("test: Operator(\"REQ(member)\") -> [%v]\n", translateOperator(op))

	op = createHeaderOperator(Operator{Value: "%REQ(member)", Name: "alias-member"})
	fmt.Printf("test: Operator(\"REQ(member)\") -> [%v]\n", translateOperator(op))

	//Output:
	//test: Operator("") -> [{<empty> <empty>}]
	//test: Operator("test") -> [{<empty> <empty>}]
	//test: Operator("REQ(") -> [{<empty> <empty>}]
	//test: Operator("REQ(t") -> [{<empty> <empty>}]
	//test: Operator("REQ()") -> [{<empty> <empty>}]
	//test: Operator("REQ(member)") -> [{member header:member}]
	//test: Operator("REQ(member)") -> [{alias-member header:member}]

}


*/
func Example_createOperator() {
	op, err := createOperator(Operator{})
	fmt.Printf("test: createOperator({}) -> [%v] [err:%v]\n", translateOperator(op), err)

	op, err = createOperator(Operator{Name: " ", Value: "static"})
	fmt.Printf("test: createOperator(\"static\") -> [%v] [err:%v]\n", translateOperator(op), err)

	op, err = createOperator(Operator{Name: "static", Value: "value"})
	fmt.Printf("test: createOperator(\"static\") -> [%v] [err:%v]\n", translateOperator(op), err)

	op, err = createOperator(Operator{Name: "", Value: "%TRAFFIC__%"})
	fmt.Printf("test: createOperator(\"TRAFFIC__\") -> [%v] [err:%v]\n", translateOperator(op), err)

	op, err = createOperator(Operator{Name: "", Value: "%REQ("})
	fmt.Printf("test: createOperator(\"REQ(static)\") -> [%v] [err:%v]\n", translateOperator(op), err)

	op, err = createOperator(Operator{Name: "", Value: "%REQ(test"})
	fmt.Printf("test: createOperator(\"REQ(static)\") -> [%v] [err:%v]\n", translateOperator(op), err)

	//op, err = createOperator(Operator{Name: "", Value: "%REQ()%"})
	//fmt.Printf("test: createOperator(\"REQ(static)\") -> [%v] [err:%v]\n", translateOperator(op), err)

	op, err = createOperator(Operator{Name: "", Value: "%REQ(static)%"})
	fmt.Printf("test: createOperator(\"REQ(static)\") -> [%v] [err:%v]\n", translateOperator(op), err)

	op, err = createOperator(Operator{Name: "new-name", Value: "%REQ(static)%"})
	fmt.Printf("test: createOperator(\"REQ(static)\") -> [%v] [err:%v]\n", translateOperator(op), err)

	op, err = createOperator(Operator{Name: "", Value: "%TRAFFIC%"})
	fmt.Printf("test: createOperator(\"TRAFFIC\") -> [%v] [err:%v]\n", translateOperator(op), err)

	op, err = createOperator(Operator{Name: "new-name", Value: "%TRAFFIC%"})
	fmt.Printf("test: createOperator(\"TRAFFIC\") -> [%v] [err:%v]\n", translateOperator(op), err)

	//Output:
	//test: createOperator({}) -> [{<empty> <empty>}] [err:invalid operator: value is empty ]
	//test: createOperator("static") -> [{<empty> <empty>}] [err:invalid operator: name is empty [static]]
	//test: createOperator("static") -> [{static value}] [err:<nil>]
	//test: createOperator("TRAFFIC__") -> [{<empty> <empty>}] [err:invalid operator: value not found or invalid %TRAFFIC__%]
	//test: createOperator("REQ(static)") -> [{<empty> <empty>}] [err:invalid operator: value not found or invalid %REQ(]
	//test: createOperator("REQ(static)") -> [{<empty> <empty>}] [err:invalid operator: value not found or invalid %REQ(test]
	//test: createOperator("REQ(static)") -> [{static %REQ(static)%}] [err:<nil>]
	//test: createOperator("REQ(static)") -> [{new-name %REQ(static)%}] [err:<nil>]
	//test: createOperator("TRAFFIC") -> [{traffic %TRAFFIC%}] [err:<nil>]
	//test: createOperator("TRAFFIC") -> [{new-name %TRAFFIC%}] [err:<nil>]

}

func Example_InitOperators() {
	var items []Operator

	items, err := InitOperators([]Operator{})
	fmt.Printf("test: InitOperators({}) -> [err:%v] [%v]\n", err, items)

	items, err = InitOperators([]Operator{{Name: "name", Value: ""}})
	fmt.Printf("test: InitOperators(\"items: nil\") -> [err:%v] [%v]\n", err, items)

	items, err = InitOperators([]Operator{{Name: "", Value: "%INVALID"}})
	fmt.Printf("test: InitOperators(\"Value: INVALID\") -> [err:%v] [%v]\n", err, items)

	items, err = InitOperators([]Operator{{Name: "name", Value: "static"}})
	fmt.Printf("test: InitOperators(\"Value: static\") -> [err:%v] [%v]\n", err, items)

	items, err = InitOperators([]Operator{{Name: "", Value: "%START_TIME%"}})
	fmt.Printf("test: InitOperators(\"Value: START_TIME\") -> [err:%v] [%v]\n", err, items)

	items, err = InitOperators([]Operator{{Name: "duration", Value: "%DURATION%"}})
	fmt.Printf("test: InitOperators(\"Value: START_TIME\") -> [err:%v] [%v]\n", err, items)

	var newItems []Operator

	newItems, err = InitOperators([]Operator{{Name: "duration", Value: "%DURATION%"}, {Name: "duration", Value: "%DURATION%"}})
	fmt.Printf("test: InitOperators(\"Value: START_TIME\") -> [err:%v] [%v]\n", err, newItems)

	//Output:
	//test: InitOperators({}) -> [err:invalid configuration: configuration slice is empty] [[]]
	//test: InitOperators("items: nil") -> [err:invalid operator: value is empty name] [[]]
	//test: InitOperators("Value: INVALID") -> [err:invalid operator: value not found or invalid %INVALID] [[]]
	//test: InitOperators("Value: static") -> [err:<nil>] [[{name static}]]
	//test: InitOperators("Value: START_TIME") -> [err:<nil>] [[{start_time %START_TIME%}]]
	//test: InitOperators("Value: START_TIME") -> [err:<nil>] [[{duration %DURATION%}]]
	//test: InitOperators("Value: START_TIME") -> [err:invalid operator: name is a duplicate [duration]] [[]]

}

func Example_CreateOperators() {
	var items []Operator

	items, err := CreateOperators([]string{TrafficOperator,
		StartTimeOperator,
		DurationOperator,
		OriginRegionOperator})
	fmt.Printf("test: CreateOperators({}) -> [err:%v] [%v]\n", err, items)

	//Output:
	//test: CreateOperators({}) -> [err:<nil>] [[{traffic %TRAFFIC%} {start_time %START_TIME%} {duration_ms %DURATION%} {region %REGION%}]]

}
