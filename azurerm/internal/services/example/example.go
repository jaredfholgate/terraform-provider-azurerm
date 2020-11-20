package example

import (
	"context"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/sdk"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/services/network/parse"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/services/network/validate"
)

type ExampleResource struct {
}

func (r ExampleResource) Arguments() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"name": {
			Type:     schema.TypeString,
			Required: true,
		},
		"number": {
			Type:     schema.TypeInt,
			Optional: true,
			Computed: true,
		},
		"enabled": {
			Type:     schema.TypeBool,
			Optional: true,
		},
		"networks": {
			Type:     schema.TypeList,
			Optional: true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
		},
		"networks_set": {
			Type:     schema.TypeSet,
			Optional: true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
		},
		"int_list": {
			Type:     schema.TypeList,
			Optional: true,
			Elem: &schema.Schema{
				Type: schema.TypeInt,
			},
		},
		"int_set": {
			Type:     schema.TypeSet,
			Optional: true,
			Elem: &schema.Schema{
				Type: schema.TypeInt,
			},
		},
		"float": {
			Type:     schema.TypeFloat,
			Optional: true,
		},
		"float_list": {
			Type:     schema.TypeList,
			Optional: true,
			Elem: &schema.Schema{
				Type: schema.TypeFloat,
			},
		},
		"float_set": {
			Type:     schema.TypeSet,
			Optional: true,
			Elem: &schema.Schema{
				Type: schema.TypeFloat,
			},
		},
		"bool_list": {
			Type:     schema.TypeList,
			Optional: true,
			Elem: &schema.Schema{
				Type: schema.TypeBool,
			},
		},
		"bool_set": {
			Type:     schema.TypeSet,
			Optional: true,
			Elem: &schema.Schema{
				Type: schema.TypeBool,
			},
		},
		"map": {
			Type:     schema.TypeMap,
			Optional: true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
		},
		"list": {
			Type:     schema.TypeList,
			Optional: true,
			MaxItems: 1,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"name": {
						Type:     schema.TypeString,
						Required: true,
					},
					"inner": {
						Type:     schema.TypeList,
						Optional: true,
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								"name": {
									Type:     schema.TypeString,
									Required: true,
								},
								"inner": {
									Type:     schema.TypeList,
									Optional: true,
									Elem: &schema.Resource{
										Schema: map[string]*schema.Schema{
											"name": {
												Type:     schema.TypeString,
												Required: true,
											},
											"should_be_fine": {
												Type:     schema.TypeBool,
												Required: true,
											},
										},
									},
								},
								"set": {
									Type:     schema.TypeSet,
									Optional: true,
									MaxItems: 1,
									Elem: &schema.Resource{
										Schema: map[string]*schema.Schema{
											"name": {
												Type:     schema.TypeString,
												Required: true,
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
		"set": {
			Type:     schema.TypeSet,
			Optional: true,
			MaxItems: 1,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"name": {
						Type:     schema.TypeString,
						Required: true,
					},
					"inner": {
						Type:     schema.TypeList,
						Optional: true,
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								"name": {
									Type:     schema.TypeString,
									Required: true,
								},
								"should_be_fine": {
									Type:     schema.TypeBool,
									Required: true,
								},
							},
						},
					},
				},
			},
		},
	}
}

// Computed Only
func (r ExampleResource) Attributes() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"output": {
			Type: schema.TypeString,
		},
	}
}

func (r ExampleResource) ResourceType() string {
	return "azurerm_example"
}

// NOTE: i guess we could return schema object to ensure everything is mapped and valid idk
type ExampleObj struct {
	Name        string            `hcl:"name"`
	Number      int               `hcl:"number"`
	Output      string            `hcl:"output" computed:"true"`
	Enabled     bool              `hcl:"enabled"`
	Networks    []string          `hcl:"networks"`
	NetworksSet []string          `hcl:"networks_set"`
	IntList     []int             `hcl:"int_list"`
	IntSet      []int             `hcl:"int_set"`
	FloatList   []float64         `hcl:"float_list"`
	FloatSet    []float64         `hcl:"float_set"`
	BoolList    []bool            `hcl:"bool_list"`
	BoolSet     []bool            `hcl:"bool_set"`
	List        []NetworkList     `hcl:"list"`
	Set         []NetworkSet      `hcl:"set"`
	Float       float64           `hcl:"float"`
	Map         map[string]string `hcl:"map"`
}

type NetworkList struct {
	Name  string         `hcl:"name"`
	Inner []NetworkInner `hcl:"inner"`
}

type NetworkListSet struct {
	Name string `hcl:"name"`
}

type NetworkSet struct {
	Name  string       `hcl:"name"`
	Inner []InnerInner `hcl:"inner"`
}

type NetworkInner struct {
	Name  string           `hcl:"name"`
	Inner []InnerInner     `hcl:"inner"`
	Set   []NetworkListSet `hcl:"set"`
}

type InnerInner struct {
	Name         string `hcl:"name"`
	ShouldBeFine bool   `hcl:"should_be_fine"`
}

func (r ExampleResource) Create() sdk.ResourceFunc {
	return CreateUpdate()
}

func (r ExampleResource) Read() sdk.ResourceFunc {
	return sdk.ResourceFunc{
		Func: func(ctx context.Context, metadata sdk.ResourceMetaData) error {
			return metadata.Encode(&ExampleObj{
				Name:        "updated",
				Number:      123,
				Enabled:     true,
				Networks:    []string{"123", "124"},
				NetworksSet: []string{"asdf", "qwer"},
				IntList:     []int{1, 2, 3},
				IntSet:      []int{3, 4, 5},
				List: []NetworkList{{
					Name: "test1232",
					Inner: []NetworkInner{
						{
							Name: "oiadsjfgoijs",
							Inner: []InnerInner{
								{
									Name:         "sure why not",
									ShouldBeFine: true,
								},
								{
									Name:         "sure why not",
									ShouldBeFine: true,
								},
								{
									Name:         "sure why not",
									ShouldBeFine: true,
								},
							},
						},
						{
							Name: "second",
							Set: []NetworkListSet{
								{
									Name: "bingo bango",
								},
							},
							Inner: []InnerInner{
								{
									Name:         "sure why not",
									ShouldBeFine: true,
								},
							},
						},
					},
				}},
				Set: []NetworkSet{{
					Name: "set1232",
					Inner: []InnerInner{
						{
							Name:         "do a thing",
							ShouldBeFine: true,
						},
					},
				}},
				Float: float64(123),
			})
		},
		Timeout: 5 * time.Minute,
	}
}

// copy pasta create
func (r ExampleResource) Update() sdk.ResourceFunc {
	return CreateUpdate()
}

func (r ExampleResource) Delete() sdk.ResourceFunc {
	return sdk.ResourceFunc{
		Func: func(ctx context.Context, metadata sdk.ResourceMetaData) error {
			return nil
		},
		Timeout: 5 * time.Minute,
	}
}

func (r ExampleResource) IDValidationFunc() schema.SchemaValidateFunc {
	return validate.SubnetID
}

func CreateUpdate() sdk.ResourceFunc {
	return sdk.ResourceFunc{
		Func: func(ctx context.Context, metadata sdk.ResourceMetaData) error {
			//metadata.ResourceData
			//metadata.Logger.WarnF("OHHAI %d", 3)
			//metadata.Client.Account.SubscriptionId
			metadata.Logger.Info("HEYO")

			var obj ExampleObj
			if err := metadata.Decode(&obj); err != nil {
				return err
			}

			id := parse.SubnetId{
				ResourceGroup:      "production-resources",
				VirtualNetworkName: "production-network",
				Name:               obj.Name,
			}

			metadata.Logger.Infof("Name is %s", obj.Name)
			metadata.Logger.Infof("Number is %d", obj.Number)
			metadata.Logger.Infof("Float is %d", obj.Float)
			metadata.Logger.Infof("Networks are %+v", obj.Networks)
			metadata.Logger.Infof("Networks Set is %+v", obj.NetworksSet)
			metadata.Logger.Infof("List is %+v", obj.List)
			metadata.Logger.Infof("Set is %+v", obj.Set)

			metadata.SetID(id)
			return nil
		},
		Timeout: 5 * time.Minute,
	}
}
