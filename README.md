![build](https://github.com/mathisve/GoRandomNoiseImage/actions/workflows/go.yaml/badge.svg)
# GoRandomNoiseImage
Generate random noise RGB images with Golang.

```go
img := image.NewRGBA(image.Rectangle{Min: min, Max: max})
for x := 0; x < width; x++ {
	for y := 0; y < height; y++ {
		img.Set(x, y, GetRandomColor())
	}
}
```


Example:

![img1](https://github.com/mathisve/GoRandomNoiseImage/blob/master/images/image.png)
