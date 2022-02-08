package img2emoji

import (
	"errors"
	"image"
	"image/color"
	"strings"
)

var DefaultMapping = map[color.Color]string{
	color.RGBA{49, 55, 61, 255}:    "⬛",
	color.RGBA{120, 177, 89, 255}:  "🟩",
	color.RGBA{253, 203, 88, 255}:  "🟨",
	color.RGBA{85, 172, 238, 255}:  "🟦",
	color.RGBA{193, 105, 79, 255}:  "🟫",
	color.RGBA{244, 144, 12, 255}:  "🟧",
	color.RGBA{170, 142, 214, 255}: "🟪",
	color.RGBA{221, 46, 68, 255}:   "🟥",
	color.RGBA{230, 231, 232, 255}: "⬜",
}

func Convert(img image.Image, mapping map[color.Color]string) (string, error) {
	if len(mapping) == 0 {
		return "", errors.New("img2emoji: mapping is empty")
	}
	p := make(color.Palette, 0, len(mapping))
	for c := range mapping {
		p = append(p, c)
	}

	var builder strings.Builder
	bounds := img.Bounds()
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			emoji, ok := mapping[p.Convert(img.At(x, y))]
			if !ok {
				panic("invariant violated: empty palette")
			}
			builder.WriteString(emoji)
		}
		builder.WriteByte('\n')
	}
	return builder.String(), nil
}
