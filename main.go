package go_example_lib

type ConfigType struct{
	Brand string
	BrandName string
}

func NewConfigType(brand string, brandName string) ConfigType {
	return ConfigType{
		Brand:     brand,
		BrandName: brandName,
	}
}

func MultiplyIt(i int, j int) int {
	return i * j
}
