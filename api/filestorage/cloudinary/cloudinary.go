package cloudinary

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
)

const (
	baseUploadUrl   = "http://api.cloudinary.com/v1_1"
	baseResource    = "res.cloudinary.com"
	baseResourceUrl = "http://res.cloudinary.com"
	imageType       = "image"
)

type ResourceType int

const (
	ImageType ResourceType = iota
)

type Service struct {
	cloudName     string
	apiKey        string
	apiSecret     string
	uploadURI     *url.URL
	uploadResType ResourceType
}

type Resource struct {
	PublicId     string `json:"public_id"`
	Version      string `json:"version"`
	ResourceType string `json:"resource_type"`
	Format       string `json:"format"`
	Size         string `json:"bytes"`
	Width        string `json:"width"`
	Height       string `json:"height"`
	Url          string `json:"url"`
	SecureUrl    string `json:"secure_url"`
}

type uploadResponse struct {
	PublicId     string `json:"public_id"`
	Version      uint   `json:"version"`
	ResourceType string `json:"resource_type"`
	Format       string `json:"format"`
	Size         int    `json:"bytes"`
}

// cloudinary://api_key:api_secret@cloud_name
func Dial(uri string) (*Service, error) {
	u, err := url.Parse(uri)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	if u.Scheme != "cloudinary" {
		return nil, errors.New("Missing cloudinary:// scheme in URI")
	}

	secret, exists := u.User.Password()

	if !exists {
		return nil, errors.New("No API secret provided in URI")
	}

	s := &Service{
		cloudName:     u.Host,
		apiKey:        u.User.Username(),
		apiSecret:     secret,
		uploadResType: ImageType,
	}

	up, err := url.Parse(fmt.Sprintf("%s/%s/image/upload", baseUploadUrl, s.cloudName))
	if err != nil {
		return nil, err
	}

	s.uploadURI = up

	return s, nil
}

func (s *Service) CloudName() string {
	return s.cloudName
}

func (s *Service) ApiKey() string {
	return s.apiKey
}

func (s *Service) DefaultUploadURI() *url.URL {
	return s.uploadURI
}

func (s *Service) UploadImage(fullPath string, data io.Reader) (string, error) {

	upURI := s.uploadURI.String()

	req, err := http.NewRequest("POST", upURI, data)

	if err != nil {
		return fullPath, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return fullPath, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		dec := json.NewDecoder(resp.Body)
		upInfo := new(uploadResponse)
		if err := dec.Decode(upInfo); err != nil {
			log.Println(err)
			return fullPath, err
		}
		fmt.Println("fajar")
		fmt.Println(dec)
		return upInfo.PublicId, nil
	} else {
		fmt.Println("Error")
		fmt.Println(resp)
		return fullPath, errors.New("Request error:" + resp.Status)
	}
}

func (s *Service) Url(publicId string, namedTransformation string) string {
	return fmt.Sprintf("http://%s-%s/image/upload/%s/%s.jpg", s.cloudName, baseResource, namedTransformation, publicId)
}
