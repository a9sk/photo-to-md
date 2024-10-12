# Photo to .md
Go code to translate a photo to a markdown file.

## Usage
```
   	Usage: go run main.go [options] <image-file>

Options:
  -s                         Use hOCR mode (specific mode), which includes advanced styling and formatting like bold and italic.
  -mode (path/capture)       Specify the mode to get the input file. If not provided, path has to be passed (in this version it is default).
  -path                      Specify the path to the input image.

Environment Variables:
  ISSPECIFIC=true            Set this environment variable to enable hOCR mode automatically (without needing -s (--specific)).

Example Usage:
  1. Run OCR on an image with default settings:
     go run main.go your_image.png

  2. Run OCR in hOCR mode for advanced styling:
     go run main.go -mode path -path your_image.png -s

  3. Display this help message:
     go run main.go --help

```

## Dependencies
The initial version of the code depends on:
- tesseract (https://github.com/tesseract-ocr) for the ocr

NOTE: Further improvements might remove this dependencies or add others.

## License
```
Copyright (c) 2024 a9sk

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
```