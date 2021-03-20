package digitalocean

import "io"

type digitalOceanSpace struct {
	BasePath string
}

func NewDOSpace(basepath string) digitalOceanSpace {
	return digitalOceanSpace{BasePath: basepath}
}

func (s *digitalOceanSpace) Save(filename string, contents io.Reader) error {

	return nil
}
