package kata19

import "testing"
import "github.com/stretchr/testify/assert"

//  test case
func TestKata19(t *testing.T) {
	assert.Equal(t, []string{"cat", "cot", "cog", "dog"}, findWordChains("cat", "dog"))
	assert.Equal(t, []string{"dog", "cog", "cot", "cat"}, findWordChains("dog", "cat"))

	assert.Equal(t, []string{"lead", "load", "goad", "gold"}, findWordChains("lead", "gold"))
	assert.Equal(t, []string{"gold", "goad", "load", "lead"}, findWordChains("gold", "lead"))

	// That supposed to be the result, but my code found a shorter ways :p
	//assert.Equal(t, []string{"ruby", "rubs", "robs", "rods", "rode", "code"}, findWordChains("ruby", "code"))
	assert.Equal(t, []string{"ruby", "rube", "robe", "rode", "code"}, findWordChains("ruby", "code"))
	assert.Equal(t, []string{"code", "rode", "robe", "rube", "ruby"}, findWordChains("code", "ruby"))

	assert.Equal(t, []string{}, findWordChains("zenith", "faiths"))
}
