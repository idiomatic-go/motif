package resource

import (
	"errors"
	"fmt"
)

func ExampleValidateConfig() {
	m := map[string]string{"database-url": "postgres://{user}:{pswd}@{sub-domain}.{db-name}.cloud.timescale.com:31770/tsdb?sslmode=require", "ping-path": "", "postgres-urn": "urn:postgres", "postgres-pgxsql-uri": "github.com/idiomatic-go/postgresql/pgxsql"}
	errs := ValidateConfig(nil, nil)
	fmt.Printf("test: ValidateConfig(nil,nil) -> %v\n", errs)

	errs = ValidateConfig(m, errors.New("file I/O error"))
	fmt.Printf("test: ValidateConfig(m,err) -> %v\n", errs)

	errs = ValidateConfig(m, nil, "not-found")
	fmt.Printf("test: Validate(m,nil,not-found) -> %v\n", errs)

	errs = ValidateConfig(m, nil, "database-url", "ping-path", "postgres-pgxsql-uri")
	fmt.Printf("test: Validate(m,nil,...) -> %v\n", errs)

	//Output:
	//test: ValidateConfig(nil,nil) -> [config map is nil]
	//test: ValidateConfig(m,err) -> [config map read error: file I/O error]
	//test: Validate(m,nil,not-found) -> [[config map error: key does not exist [not-found]]
	//test: Validate(m,nil,...) -> [config map error: value for key does not exist [ping-path]]

}
