package ibm

import (
	"fmt"
	"time"

	"github.com/IBM/vpc-go-sdk/vpcclassicv1"
	"github.com/IBM/vpc-go-sdk/vpcv1"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

const (
	isImages = "images"
)

func dataSourceIBMISImages() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceIBMISImagesRead,

		Schema: map[string]*schema.Schema{

			isImages: {
				Type:        schema.TypeList,
				Description: "List of images",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"status": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"visibility": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"os": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"architecture": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"crn": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceIBMISImagesRead(d *schema.ResourceData, meta interface{}) error {
	userDetails, err := meta.(ClientSession).BluemixUserDetails()
	if err != nil {
		return err
	}
	if userDetails.generation == 1 {
		err := classicImageList(d, meta)
		if err != nil {
			return err
		}
	} else {
		err := imageList(d, meta)
		if err != nil {
			return err
		}
	}
	return nil
}

func classicImageList(d *schema.ResourceData, meta interface{}) error {
	sess, err := classicVpcClient(meta)
	if err != nil {
		return err
	}
	start := ""
	allrecs := []vpcclassicv1.Image{}
	for {
		listImagesOptions := &vpcclassicv1.ListImagesOptions{}
		if start != "" {
			listImagesOptions.Start = &start
		}
		availableImages, response, err := sess.ListImages(listImagesOptions)
		if err != nil {
			return fmt.Errorf("Error Fetching Images %s\n%s", err, response)
		}
		start = GetNext(availableImages.Next)
		allrecs = append(allrecs, availableImages.Images...)
		if start == "" {
			break
		}
	}
	imagesInfo := make([]map[string]interface{}, 0)
	for _, image := range allrecs {

		l := map[string]interface{}{
			"name":         *image.Name,
			"id":           *image.ID,
			"status":       *image.Status,
			"crn":          *image.CRN,
			"visibility":   *image.Visibility,
			"os":           *image.OperatingSystem.Name,
			"architecture": *image.OperatingSystem.Architecture,
		}
		imagesInfo = append(imagesInfo, l)
	}
	d.SetId(dataSourceIBMISSubnetsID(d))
	d.Set(isImages, imagesInfo)
	return nil
}

func imageList(d *schema.ResourceData, meta interface{}) error {
	sess, err := vpcClient(meta)
	if err != nil {
		return err
	}
	start := ""
	allrecs := []vpcv1.Image{}
	for {
		listImagesOptions := &vpcv1.ListImagesOptions{}
		if start != "" {
			listImagesOptions.Start = &start
		}
		availableImages, response, err := sess.ListImages(listImagesOptions)
		if err != nil {
			return fmt.Errorf("Error Fetching Images %s\n%s", err, response)
		}
		start = GetNext(availableImages.Next)
		allrecs = append(allrecs, availableImages.Images...)
		if start == "" {
			break
		}
	}
	imagesInfo := make([]map[string]interface{}, 0)
	for _, image := range allrecs {

		l := map[string]interface{}{
			"name":         *image.Name,
			"id":           *image.ID,
			"status":       *image.Status,
			"crn":          *image.CRN,
			"visibility":   *image.Visibility,
			"os":           *image.OperatingSystem.Name,
			"architecture": *image.OperatingSystem.Architecture,
		}
		imagesInfo = append(imagesInfo, l)
	}
	d.SetId(dataSourceIBMISSubnetsID(d))
	d.Set(isImages, imagesInfo)
	return nil
}

// dataSourceIBMISImagesId returns a reasonable ID for a image list.
func dataSourceIBMISImagesID(d *schema.ResourceData) string {
	return time.Now().UTC().String()
}
