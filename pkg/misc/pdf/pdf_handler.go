package pdf

import (
	"fmt"
	"os"

	"github.com/dslipak/pdf"
)

func ReadPdf(path string) (string, error) {
	r, err := openFileEOF(path)
	if err != nil {
		fmt.Println("Error trying to open file:", err)
		return "", err
	}
	totalPage := r.NumPage()

	fonts := make(map[string]*pdf.Font)
	for _, name := range r.Page(totalPage).Fonts() { // cache fonts so we don't continually parse charmap
		if _, ok := fonts[name]; !ok {
			f := r.Page(totalPage).Font(name)
			fonts[name] = &f
		}
	}
	p := r.Page(totalPage)
	if p.V.IsNull() {
		return "", err
	}
	bpage, _ := p.GetPlainText(fonts)

	return bpage, nil
}

func openFileEOF(path string) (*pdf.Reader, error) {
	f, er4 := os.Open(path)
	fmt.Println(er4)
	fl, _ := f.Stat()
	return pdf.NewReader(f, fl.Size())
	//r, er2 := rscpdf.NewReaderEncrypted(f, fl.Size(), func() string { return "" })
}
