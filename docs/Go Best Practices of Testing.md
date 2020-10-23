# Go Best Practices — Testing

## 1. Basics

1) `Setup` and `TearDown`:

```go
func TestMain(m *testing.M) {
	fmt.Println("write setup code here...") // setup before test
	retCode := m.Run()                      // run test
	fmt.Printf("result code: %d\n", retCode)
	fmt.Println("write teardown code here...") // tear down after test
}
```

2) `Setup` and `TearDown` for sub tests:

```go
// Setup and Teardown for test group
func setupTestCase(t *testing.T) func(t *testing.T) {
	t.Log("before testing, setup")
	return func(t *testing.T) {
		t.Log("after testing, teardown")
	}
}

// Setup and Teardown for sub group
func setupSubTest(t *testing.T) func(t *testing.T) {
	t.Log("before testing, setup")
	return func(t *testing.T) {
		t.Log("after testing, teardown")
	}
}
```



## 2. Table based tests

In almost all cases when you test a method you’d like to test a couple of scenarios on the function. Usually with different input parameters or different mock responses. I like to group all these tests into one `Test*` function and then have a loop running through all the test cases. Here’s a basic example:

**Function:**

```go
func Trim(s string) string {
	return strings.Trim(s, " ")
}
```

Goland can auto generate the test code frame:

**Tests:**

```go
func TestTrim(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "normal",
			args: args{
				s: "hello",
			},
			want: "hello",
		},
		{
			name: "tailing blank",
			args: args{
				s: "hello   ",
			},
			want: "hello",
		},
		{
			name: "heading blank",
			args: args{
				s: "   hello",
			},
			want: "hello",
		},
		{
			name: "heading and tailing blank",
			args: args{
				s: "   hello ",
			},
			want: "hello",
		},
		{
			name: "empty",
			args: args{
				s: "",
			},
			want: "",
		},
	}
  
  teardownTestCase := setupTestCase(t) // setup before test
	defer teardownTestCase(t)            // tear down after test
  
	for _, tt := range tests {
    teardownSubTest := setupSubTest(t) // setup for sub test
		defer teardownSubTest(t)           // teardown for sub test
		t.Run(tt.name, func(t *testing.T) {
			if got := Trim(tt.args.s); got != tt.want {
				t.Errorf("trim() = %v, want %v", got, tt.want)
			}
		})
	}
}
```



## 3. Mock external resource:

### Case:

Test downloader function, which depends on `getPage` method:

```go
package demo

import (
	"io/ioutil"
	"mvdan.cc/xurls/v2"
	"net/http"
)

func getPage(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		return ""
	}
	defer resp.Body.Close()

	contents, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return ""
	}
	return string(contents)
}

func downloader() {
	content := getPage("https://golang.org/")
	rxStrict := xurls.Strict()
	matches := rxStrict.FindAllString(content, -1)

	for _, match := range matches {
		go getPage(match)
	}
}

```





### Solutions:

#### 1) gomock

##### Basic Usage:

Usage of *GoMock* follows four basic steps:

1. Use `mockgen` to generate a mock for the interface you wish to mock.
2. In your test, create an instance of `gomock.Controller` and pass it to your mock object’s constructor to obtain a mock object.
3. Call `EXPECT()` on your mocks to set up their expectations and return values
4. Call `Finish()` on the mock controller to assert the mock’s expectations

##### Example:

Refactor: create interface `PageGetter` in file `page.go`:

```go
package demo

import (
	"io/ioutil"
	"net/http"
)

type PageGetter interface {
	getPage(url string) string
}

type PageGetterImpl struct {
}

func (g PageGetterImpl) getPage(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		return ""
	}
	defer resp.Body.Close()

	contents, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return ""
	}
	return string(contents)
}

```



then, generate mock code:

```bash
mockgen -source=page.go -destination=page_mock.go -package=demo
```

Or use short command, if under interface file directory:

```bash
mockgen . PageGetter
```

**Main:**

```go
package demo

import (
	"mvdan.cc/xurls/v2"
)

type Downloader struct {
	pageGetter PageGetter
}

func (d Downloader) download(baseUrl string) {
	content := d.pageGetter.getPage(baseUrl)
	rxStrict := xurls.Strict()
	matches := rxStrict.FindAllString(content, -1)

	for _, match := range matches {
		go d.pageGetter.getPage(match)
	}
}

```

**Test:**

```go
package demo

import (
	"github.com/golang/mock/gomock"
	"testing"
)

func Test_download(t *testing.T) {
	t.Run("download", func(t *testing.T) {
		controller := gomock.NewController(t)
		defer controller.Finish()
		pageGetter := NewMockPageGetter(controller)

		pageGetter.EXPECT().getPage(gomock.Any()).Return("").Times(100)
		downloader := &Downloader{
			pageGetter: pageGetter,
		}
		downloader.download("https://golang.org/")
	})
}
```



**Third part mock:**

for example, mock `ILogger` in `git.garena.com/shopee/platform/splog`

```bash
mockgen git.garena.com/shopee/platform/splog ILogger >> mock_splog.go
```



#### 2) Other mocking technicals

##### Method 1: method as paramter

Pass `getPage()` as a parameter of `downloader()`

```go
type PageGetter func(url string) string

func downloader(pageGetterFunc PageGetter) {
    // ...
    content := pageGetterFunc(BASE_URL)
    // ...
}
```

**Main:**

```go
func getPage(url string) string { /* ... */ }

func main() {
    downloader(getPage)
}
```

**Test:**

```go
func mockGetPage(url string) string {
    // mock your 'getPage()' function here
}

func TestDownloader(t *testing.T) {
    downloader(mockGetPage)
}
```

##### Method 2: define type

Make `download()` a method of a type `Downloader`:

```go
type PageGetter func(url string) string

type Downloader struct {
    getPage PageGetter
}

func NewDownloader(pg PageGetter) *Downloader {
    return &Downloader{getPage: pg}
}

func (d *Downloader) download() {
    //...
    content := d.getPage(BASE_URL)
    //...
}
```

**Main:**

```go
func getPage(url string) string { /* ... */ }

func main() {
    d := NewDownloader(getPage)
    d.download()
}
```

**Test:**

```go
func mockGetPage(url string) string {
    // mock your 'getPage()' function here
}

func TestDownloader() {
    d := NewDownloader(mockGetPage)
    d.download()
}
```

##### Method 3: define function

Change your function definition to use a variable instead:

```go
var getPage = func(url string) string {
    //...
}
```

You can override it in your tests:

```go
func TestDownloader(t *testing.T) {
    getPage = func(url string) string {
        if url != "expected" {
            t.Fatal("good message")
        }
        return "something"
    }
    downloader()
}
```





## Refers:

1. Refers to: [mock in go](https://stackoverflow.com/questions/19167970/mock-functions-in-go)
2. https://github.com/golang/mock
3. https://www.liwenzhou.com/posts/Go/16_test/