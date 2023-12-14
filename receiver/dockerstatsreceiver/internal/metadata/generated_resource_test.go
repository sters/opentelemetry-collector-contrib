// Code generated by mdatagen. DO NOT EDIT.

package metadata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestResourceBuilder(t *testing.T) {
	for _, test := range []string{"default", "all_set", "none_set"} {
		t.Run(test, func(t *testing.T) {
			cfg := loadResourceAttributesConfig(t, test)
			rb := NewResourceBuilder(cfg)
			rb.SetContainerCommandLine("container.command_line-val")
			rb.SetContainerHostname("container.hostname-val")
			rb.SetContainerID("container.id-val")
			rb.SetContainerImageID("container.image.id-val")
			rb.SetContainerImageName("container.image.name-val")
			rb.SetContainerName("container.name-val")
			rb.SetContainerRuntime("container.runtime-val")

			res := rb.Emit()
			assert.Equal(t, 0, rb.Emit().Attributes().Len()) // Second call should return empty Resource

			switch test {
			case "default":
				assert.Equal(t, 5, res.Attributes().Len())
			case "all_set":
				assert.Equal(t, 7, res.Attributes().Len())
			case "none_set":
				assert.Equal(t, 0, res.Attributes().Len())
				return
			default:
				assert.Failf(t, "unexpected test case: %s", test)
			}

			val, ok := res.Attributes().Get("container.command_line")
			assert.Equal(t, test == "all_set", ok)
			if ok {
				assert.EqualValues(t, "container.command_line-val", val.Str())
			}
			val, ok = res.Attributes().Get("container.hostname")
			assert.True(t, ok)
			if ok {
				assert.EqualValues(t, "container.hostname-val", val.Str())
			}
			val, ok = res.Attributes().Get("container.id")
			assert.True(t, ok)
			if ok {
				assert.EqualValues(t, "container.id-val", val.Str())
			}
			val, ok = res.Attributes().Get("container.image.id")
			assert.Equal(t, test == "all_set", ok)
			if ok {
				assert.EqualValues(t, "container.image.id-val", val.Str())
			}
			val, ok = res.Attributes().Get("container.image.name")
			assert.True(t, ok)
			if ok {
				assert.EqualValues(t, "container.image.name-val", val.Str())
			}
			val, ok = res.Attributes().Get("container.name")
			assert.True(t, ok)
			if ok {
				assert.EqualValues(t, "container.name-val", val.Str())
			}
			val, ok = res.Attributes().Get("container.runtime")
			assert.True(t, ok)
			if ok {
				assert.EqualValues(t, "container.runtime-val", val.Str())
			}
		})
	}
}