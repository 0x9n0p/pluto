package processor

import (
	"io"
	"net/http"
	"pluto"
	"strings"
)

const (
	Category_Flow          = "Flow"
	Category_InputOutpt    = "InputOutput"
	Category_Communication = "Communication"
	Category_Validator     = "Validator"
)

var Processors = []Descriptor{
	//{
	//	Name:        "Execute processor and join the result",
	//	Description: "Execute processor and join the result ..",
	//	Icon:        "https://..",
	//	Arguments: []pluto.ValueDescriptor{
	//		{
	//			Name:     "Processor",
	//			Type:     pluto.TypeProcessor,
	//			Required: true,
	//			ValueValidator: func(arg pluto.Value, _ pluto.ValueDescriptor) (err error) {
	//				m, ok := arg.Value.(map[string]any)
	//				if !ok {
	//					return &pluto.Error{
	//						HTTPCode: http.StatusBadRequest,
	//						Message:  "Value of (Processor) is not a processor",
	//					}
	//				}
	//
	//				defer func() {
	//					if v := recover(); v != nil {
	//						err = &pluto.Error{
	//							HTTPCode: http.StatusBadRequest,
	//							Message:  fmt.Sprintf("Missing fields or incorrect types: %s", v),
	//						}
	//					}
	//
	//				}()
	//
	//				_ = m["name"].(string)
	//				_ = m["arguments"].([]any)
	//
	//				return
	//			},
	//		},
	//	},
	//	Input: []pluto.ValueDescriptor{
	//		/*
	//			Input of the inner processor.
	//		*/
	//	},
	//	Output: []pluto.ValueDescriptor{
	//		/*
	//			Output of the inner processor.
	//		*/
	//	},
	//	Category: Category_Flow,
	//},
	{
		Name:        pluto.ProcessorName_Execute,
		Description: "Executes the pipeline",
		Icon:        "https://...",
		Arguments: []pluto.ValueDescriptor{
			{
				Name:     "name",
				Type:     pluto.TypeText,
				Required: true,
			},
			{
				Name:     "append_result",
				Type:     pluto.TypeBoolean,
				Required: false,
				Default:  false,
			},
		},
		Input:    []pluto.ValueDescriptor{},
		Output:   []pluto.ValueDescriptor{},
		Category: Category_Flow,
	},
	{
		Name:        pluto.ProcessorName_Fork,
		Description: "Executes the pipeline in background",
		Icon:        "https://...",
		Arguments: []pluto.ValueDescriptor{
			{
				Name:     "name",
				Type:     pluto.TypeText,
				Required: true,
			},
		},
		Input:    []pluto.ValueDescriptor{},
		Output:   []pluto.ValueDescriptor{},
		Category: Category_Flow,
	},
	{
		Name:        pluto.ProcessorName_RuntimeProcessorCreator,
		Description: "Creates the processor and adds it to the body",
		Icon:        "https://...",
		Arguments: []pluto.ValueDescriptor{
			{
				Name:     "processor_name",
				Type:     pluto.TypeText,
				Required: true,
			},
			{
				Name:     "append_name",
				Type:     pluto.TypeText,
				Required: true,
			},
		},
		Input: []pluto.ValueDescriptor{},
		Output: []pluto.ValueDescriptor{
			{
				Name:     "$append_name",
				Type:     pluto.TypeProcessor,
				Required: true,
			},
		},
		Category: Category_Flow,
	},
	{
		Name:        pluto.ProcessorName_SendResponse,
		Description: "Sends body to producer",
		Icon:        "https://...",
		Arguments: []pluto.ValueDescriptor{
			{
				Name:     "pipeline_name",
				Type:     pluto.TypeText,
				Required: true,
			},
		},
		Input:    []pluto.ValueDescriptor{},
		Output:   []pluto.ValueDescriptor{},
		Category: Category_Communication,
	},
	{
		Name:        pluto.ProcessorName_NumberValidator,
		Description: "Validates the input",
		Icon:        "https://...",
		Arguments: []pluto.ValueDescriptor{
			{
				Name:     "name",
				Type:     pluto.TypeText,
				Required: true,
			},
			{
				Name:     "minimum",
				Type:     pluto.TypeNumeric,
				Required: true,
			},
			{
				Name:     "maximum",
				Type:     pluto.TypeNumeric,
				Required: true,
			},
		},
		Input:    []pluto.ValueDescriptor{},
		Output:   []pluto.ValueDescriptor{},
		Category: Category_Validator,
	},
	{
		Name:        pluto.ProcessorName_ChannelCreator,
		Description: "Creates a channel",
		Icon:        "https://...",
		Arguments: []pluto.ValueDescriptor{
			{
				Name:     "name",
				Type:     pluto.TypeText,
				Required: true,
			},
			{
				Name:    "length",
				Type:    pluto.TypeNumeric,
				Default: 10,
			},
		},
		Input: []pluto.ValueDescriptor{},
		Output: []pluto.ValueDescriptor{
			{
				Name:     "channel",
				Type:     pluto.TypeChannel,
				Required: true,
			},
		},
		Category: Category_Communication,
	},
	{
		Name:        pluto.ProcessorName_ChannelFinder,
		Description: "Finds channel by the given name",
		Icon:        "https://...",
		Arguments: []pluto.ValueDescriptor{
			{
				Name:     "name",
				Type:     pluto.TypeText,
				Required: true,
			},
		},
		Input: []pluto.ValueDescriptor{},
		Output: []pluto.ValueDescriptor{
			{
				Name:     "channel",
				Type:     pluto.TypeChannel,
				Required: true,
			},
		},
		Category: Category_Communication,
	},
	{
		Name:        pluto.ProcessorName_JoinChannel,
		Description: "When a processable is published by the channel, the processor is executed.",
		Icon:        "https://...",
		Arguments:   []pluto.ValueDescriptor{},
		Input: []pluto.ValueDescriptor{
			{
				Name:     "channel",
				Type:     pluto.TypeChannel,
				Required: true,
			},
			{
				Name:     "identifier",
				Type:     pluto.TypeIdentifier,
				Required: true,
			},
			{
				Name:     "processor",
				Type:     pluto.TypeProcessor,
				Required: true,
			},
		},
		Output:   []pluto.ValueDescriptor{},
		Category: Category_Communication,
	},
	{
		Name:        pluto.ProcessorName_IOWriter,
		Description: "Write to Input/Output interfaces directly",
		Icon:        "https://...",
		Arguments: []pluto.ValueDescriptor{
			{
				Name:     "io_interface",
				Type:     pluto.TypeInternalInterface,
				Required: true,
				ValueValidator: func(arg pluto.Value, _ pluto.ValueDescriptor) error {
					_, ok := arg.Value.(io.Writer)
					if !ok {
						return &pluto.Error{
							HTTPCode: http.StatusBadRequest,
							Message:  "Value of (io_interface) is not an io.Writer",
						}
					}
					return nil
				},
			},
		},
		Input: []pluto.ValueDescriptor{
			/*
				The processable.body is Processable.GetBody()
			*/
			{
				Name:     "processable.body",
				Type:     pluto.TypeBytes,
				Required: true,
			},
		},
		Output:   []pluto.ValueDescriptor{},
		Category: Category_InputOutpt,
	},
}

func GetDescriptor(name string) (Descriptor, bool) {
	for _, processor := range Processors {
		if name == processor.Name {
			return processor, true
		}
	}
	return Descriptor{}, false
}

type DescriptorFinder struct {
	Name string `query:"name"`
}

func (f *DescriptorFinder) Find() []Descriptor {
	if f.Name == "" {
		return Processors
	}

	found := make([]Descriptor, 0)
	for _, descriptor := range Processors {
		if strings.Contains(strings.ToLower(descriptor.Name), strings.ToLower(f.Name)) {
			found = append(found, descriptor)
		}
	}
	return found
}
