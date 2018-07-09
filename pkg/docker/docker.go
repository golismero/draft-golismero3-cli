package docker

func docker(image string, params []string) func([]byte) ([]byte, error) {
	return func(inp []byte) ([]byte, error) {
		return []byte{}, nil
	}
}
