package runtime

import "fmt"

func ExampleExpand() {
	t := "file_{env}.txt"
	fmt.Printf("test: EnvExpansion(%v) -> %v\n", t, EnvExpansion(t))

	SetRuntimeEnv("prod")
	fmt.Printf("test: EnvExpansion(%v) -> %v\n", t, EnvExpansion(t))

	t = "file_{env_var_invalid}.txt"
	fmt.Printf("test: EnvExpansion(%v) -> %v\n", t, EnvExpansion(t))

	// Output:
	//test: EnvExpansion(file_{env}.txt) -> file_dev.txt
	//test: EnvExpansion(file_{env}.txt) -> file_prod.txt
	//test: EnvExpansion(file_{env_var_invalid}.txt) -> invalid or missing environment variable reference: {env}

}
