package cleaner_test

import (
	"fmt"
	"os"

	"github.com/IGZgustavomarin/dir-cleaner/pkg/cleaner"
)

func ExampleApp_Clean() {
	app := &cleaner.App{}
	dir := "/tmp/dir-cleaner-example"

	_ = os.MkdirAll(dir, 0777)

	stats, _ := app.Clean(&cleaner.Options{
		DryRun: true,
		Path:   dir,
	})

	fmt.Println(stats)

	// Output:
	// No match found on [/tmp/dir-cleaner-example]
}
