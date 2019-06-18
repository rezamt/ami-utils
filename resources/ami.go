package resources

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ec2"
	"sort"
	"time"
)

type By func(p1, p2 *ec2.Image) bool

func (by By) Sort(images []*ec2.Image) {
	ps := &sortedImageList{
		images: images,
		by:     by,
	}
	sort.Sort(ps)
}

type sortedImageList struct {
	images []*ec2.Image
	by     func(im2, im1 *ec2.Image) bool
}

func (s *sortedImageList) Len() int {
	return len(s.images)
}

func (s *sortedImageList) Swap(i, j int) {
	s.images[i], s.images[j] = s.images[j], s.images[i]
}

func (s *sortedImageList) Less(i, j int) bool {
	return s.by(s.images[i], s.images[j])
}

func ListAMIs(name string) (*ec2.DescribeImagesOutput, error) {

	input := &ec2.DescribeImagesInput{
		Filters: []*ec2.Filter{
			{
				Name:   aws.String("name"),
				Values: []*string{&name},
			},
		},
	}

	result, err := Service.DescribeImages(input)

	// we are sorting based on creation date
	if err == nil {
		images := result.Images

		layout := "2006-01-02T15:04:05.000Z"

		By(func(im1, im2 *ec2.Image) bool {

			t1, err1 := time.Parse(layout, *im1.CreationDate)
			t2, err2 := time.Parse(layout, *im2.CreationDate)

			if err1 != nil && err2 != nil {
				fmt.Print("Data conversion error")

				return false
			}

			return t1.After(t2)
		}).Sort(images)
	}

	return result, err
}
