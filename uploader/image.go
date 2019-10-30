// Package uploader contains the implementation needed
// to create an image data depending on the type png, jpg or gif
// it creates the image and appends it to a tmp directory
package uploader

// gopherPNG creates an io.Reader by decoding the base64 encoded image data string in the gopher constant.
func gopherPNG() io.Reader { return base64.NewDecoder(base64.StdEncoding, strings.NewReader(gopher)) }


