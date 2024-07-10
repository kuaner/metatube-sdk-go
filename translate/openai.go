package translate

import (
	"os"
	"strconv"

	openai "github.com/zijiren233/openai-translator"
)

func OpenaiTranslate(q, source, target, key string) (result string, err error) {
	opt := []openai.Option{
		openai.WithFrom(source),
	}
	if withURL := os.Getenv("OPENAI_BASE_URL"); withURL != "" {
		opt = append(opt, openai.WithUrl(os.Getenv("OPENAI_BASE_URL")))
	}
	if withModel := os.Getenv("OPENAI_MODEL"); withModel != "" {
		opt = append(opt, openai.WithModel(os.Getenv("OPENAI_MODEL")))
	}
	if withTemperature := os.Getenv("OPENAI_TEMPERATURE"); withTemperature != "" {
		// convert string to float32
		temperature, err := strconv.ParseFloat(withTemperature, 32)
		if err == nil {
			opt = append(opt, openai.WithTemperature(float32(temperature)))
		}
	}
	return openai.Translate(q, target, key, opt...)
}
